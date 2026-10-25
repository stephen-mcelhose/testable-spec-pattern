package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"google.golang.org/protobuf/types/known/emptypb"

	petstorev1 "github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore/internal/gen/petstore/v1"
	"github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore/internal/pet"
)

const (
	defaultPageSize int32 = 50
	maxPageSize     int32 = 100
)

// Store is the persistence leaf used by PetService. Defined at the usage site
// (Go idiom: small consumer-owned interfaces). Tests inject a hand stub; prefer
// counterfeiter when call-count or argument capture is needed.
type Store interface {
	Create(ctx context.Context, p *petstorev1.Pet) (*petstorev1.Pet, error)
	Get(ctx context.Context, name string) (*petstorev1.Pet, error)
	Update(ctx context.Context, p *petstorev1.Pet, etag string) (*petstorev1.Pet, error)
	Delete(ctx context.Context, name string, etag string) error
	List(ctx context.Context, pageSize int32, pageToken string) (pets []*petstorev1.Pet, nextPageToken string, err error)
}

// PetService is a thin AIP-flavoured Pets API over generated petstorev1 types:
// (1) proto/shape validation → (2) business rules → (3) Store + error map.
// Embeds UnimplementedPetServiceServer so UploadPhoto stays unimplemented
// until a later issue; gRPC server registration is out of scope here.
type PetService struct {
	petstorev1.UnimplementedPetServiceServer
	Store Store
}

var _ petstorev1.PetServiceServer = (*PetService)(nil)

// CreatePet creates a pet (AIP-133): pet_id builds the resource name; Pet.Name on
// the body is ignored; returns the fully populated stored Pet.
// Duplicate id → AlreadyExists.
func (s *PetService) CreatePet(ctx context.Context, req *petstorev1.CreatePetRequest) (*petstorev1.Pet, error) {
	// 1. Proto / shape validation — must be a valid CreatePetRequest.
	if req == nil {
		return nil, &pet.InvalidArgumentError{Message: "request is required"}
	}
	if req.GetPet() == nil {
		return nil, &pet.InvalidArgumentError{Message: "pet is required"}
	}

	// 2. Business rules — user-specified pet_id required in this example
	// (system-generated ids when pet_id omitted are out of scope); display_name required.
	name, err := petNameFromID(req.GetPetId())
	if err != nil {
		return nil, err
	}
	display := strings.TrimSpace(req.GetPet().GetDisplayName())
	if display == "" {
		return nil, &pet.InvalidArgumentError{Message: "display_name is required"}
	}

	// 3. Implementation — delegate to storage.
	created, err := s.Store.Create(ctx, &petstorev1.Pet{Name: name, DisplayName: display})
	if err != nil {
		return nil, mapStoreError(err)
	}
	return created, nil
}

// GetPet validates the resource name and returns the stored Pet.
func (s *PetService) GetPet(ctx context.Context, req *petstorev1.GetPetRequest) (*petstorev1.Pet, error) {
	// 1. Proto / shape validation — must be a valid GetPetRequest.
	if req == nil {
		return nil, &pet.InvalidArgumentError{Message: "request is required"}
	}

	// 2. Business rules — name must be a canonical pets/{pet} id.
	canonical, err := pet.NormalizePetName(req.GetName())
	if err != nil {
		return nil, err
	}

	// 3. Implementation — delegate to storage.
	got, err := s.Store.Get(ctx, canonical)
	if err != nil {
		return nil, mapStoreError(err)
	}
	return got, nil
}

// UpdatePet updates a pet (AIP-134): Pet.Name identifies the resource; Pet.Etag is
// optional freshness (AIP-154) — when set and mismatched, Store returns Aborted.
// Returns the fully populated stored Pet. update_mask / allow_missing are out of scope.
func (s *PetService) UpdatePet(ctx context.Context, req *petstorev1.UpdatePetRequest) (*petstorev1.Pet, error) {
	// 1. Proto / shape validation — must be a valid UpdatePetRequest.
	if req == nil {
		return nil, &pet.InvalidArgumentError{Message: "request is required"}
	}
	if req.GetPet() == nil {
		return nil, &pet.InvalidArgumentError{Message: "pet is required"}
	}

	// 2. Business rules — name shape; display_name required; etag optional.
	name, err := pet.NormalizePetName(req.GetPet().GetName())
	if err != nil {
		return nil, err
	}
	display := strings.TrimSpace(req.GetPet().GetDisplayName())
	if display == "" {
		return nil, &pet.InvalidArgumentError{Message: "display_name is required"}
	}
	etag := strings.TrimSpace(req.GetPet().GetEtag())

	// 3. Implementation — delegate to storage.
	updated, err := s.Store.Update(ctx, &petstorev1.Pet{Name: name, DisplayName: display}, etag)
	if err != nil {
		return nil, mapStoreError(err)
	}
	return updated, nil
}

// DeletePet deletes the pet (AIP-135). Etag is optional protected-delete freshness
// (AIP-154): when provided and mismatched, Store returns Aborted.
func (s *PetService) DeletePet(ctx context.Context, req *petstorev1.DeletePetRequest) (*emptypb.Empty, error) {
	// 1. Proto / shape validation — must be a valid DeletePetRequest.
	if req == nil {
		return nil, &pet.InvalidArgumentError{Message: "request is required"}
	}

	// 2. Business rules — name shape; etag optional.
	canonical, err := pet.NormalizePetName(req.GetName())
	if err != nil {
		return nil, err
	}

	// 3. Implementation — delegate to storage.
	if err := s.Store.Delete(ctx, canonical, strings.TrimSpace(req.GetEtag())); err != nil {
		return nil, mapStoreError(err)
	}
	return &emptypb.Empty{}, nil
}

// ListPets lists pets (AIP-132 / AIP-158): default and coerce page_size; return a
// page. Empty collection is success. filter / order_by / show_deleted are out of scope.
func (s *PetService) ListPets(ctx context.Context, req *petstorev1.ListPetsRequest) (*petstorev1.ListPetsResponse, error) {
	// 1. Proto / shape validation — must be a valid ListPetsRequest.
	if req == nil {
		return nil, &pet.InvalidArgumentError{Message: "request is required"}
	}

	// 2. Business rules — page_size bounds (default + coerce above max).
	pageSize := req.GetPageSize()
	if pageSize < 0 {
		return nil, &pet.InvalidArgumentError{Message: "page_size must not be negative"}
	}
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	if pageSize > maxPageSize {
		pageSize = maxPageSize // AIP-132: coerce above max; do not error
	}

	// 3. Implementation — delegate to storage.
	pets, next, err := s.Store.List(ctx, pageSize, req.GetPageToken())
	if err != nil {
		return nil, mapStoreError(err)
	}
	if pets == nil {
		pets = []*petstorev1.Pet{}
	}
	return &petstorev1.ListPetsResponse{Pets: pets, NextPageToken: next}, nil
}

// petNameFromID builds pets/{pet_id} and validates via NormalizePetName (AIP-133).
func petNameFromID(petID string) (string, error) {
	return pet.NormalizePetName("pets/" + strings.TrimSpace(petID))
}

// mapStoreError passes through KnownError and wraps anything else as Internal.
func mapStoreError(err error) error {
	var known pet.KnownError
	if errors.As(err, &known) {
		return err
	}
	return &pet.InternalError{Message: fmt.Sprintf("store: %v", err)}
}
