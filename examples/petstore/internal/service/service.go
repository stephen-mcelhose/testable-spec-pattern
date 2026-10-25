package service

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"buf.build/go/protovalidate"
	"google.golang.org/protobuf/proto"
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
//
// This example keeps Store on generated petstorev1 types (thin wire). Introduce
// domain objects + mappers when domain language diverges from the proto.
type Store interface {
	Create(ctx context.Context, p *petstorev1.Pet) (*petstorev1.Pet, error)
	Get(ctx context.Context, name string) (*petstorev1.Pet, error)
	Update(ctx context.Context, p *petstorev1.Pet, etag string) (*petstorev1.Pet, error)
	Delete(ctx context.Context, name string, etag string) error
	List(ctx context.Context, pageSize int32, pageToken string) (pets []*petstorev1.Pet, nextPageToken string, err error)
}

// PetService is a thin AIP-flavoured Pets API over generated petstorev1 types:
// (1) protovalidate.Validate shape → (2) validate*BusinessRules → (3) Store + error map.
// Embeds UnimplementedPetServiceServer so UploadPhoto stays unimplemented
// until a later issue; a separate transport adapter (KnownError → gRPC status)
// waits on server wiring.
type PetService struct {
	petstorev1.UnimplementedPetServiceServer
	Store Store
}

var _ petstorev1.PetServiceServer = (*PetService)(nil)

// CreatePet creates a pet (AIP-133): pet_id builds the resource name; Pet.Name on
// the body is ignored; returns the fully populated stored Pet.
// Duplicate id → AlreadyExists.
func (s *PetService) CreatePet(ctx context.Context, req *petstorev1.CreatePetRequest) (*petstorev1.Pet, error) {
	// 1. Proto / shape — stereotypical check: protovalidate.Validate (buf.validate).
	if err := validateShape(req); err != nil {
		return nil, err
	}

	// 2. Business rules.
	name, display, err := validateCreatePetBusinessRules(req)
	if err != nil {
		return nil, err
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
	// 1. Proto / shape — stereotypical check: protovalidate.Validate (buf.validate).
	if err := validateShape(req); err != nil {
		return nil, err
	}

	// 2. Business rules.
	canonical, err := validateGetPetBusinessRules(req)
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
	// 1. Proto / shape — stereotypical check: protovalidate.Validate (buf.validate).
	if err := validateShape(req); err != nil {
		return nil, err
	}

	// 2. Business rules.
	name, display, etag, err := validateUpdatePetBusinessRules(req)
	if err != nil {
		return nil, err
	}

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
	// 1. Proto / shape — stereotypical check: protovalidate.Validate (buf.validate).
	if err := validateShape(req); err != nil {
		return nil, err
	}

	// 2. Business rules.
	canonical, etag, err := validateDeletePetBusinessRules(req)
	if err != nil {
		return nil, err
	}

	// 3. Implementation — delegate to storage.
	if err := s.Store.Delete(ctx, canonical, etag); err != nil {
		return nil, mapStoreError(err)
	}
	return &emptypb.Empty{}, nil
}

// ListPets lists pets (AIP-132 / AIP-158): default and coerce page_size; return a
// page. Empty collection is success. filter / order_by / show_deleted are out of scope.
func (s *PetService) ListPets(ctx context.Context, req *petstorev1.ListPetsRequest) (*petstorev1.ListPetsResponse, error) {
	// 1. Proto / shape — stereotypical check: protovalidate.Validate (buf.validate).
	if err := validateShape(req); err != nil {
		return nil, err
	}

	// 2. Business rules.
	pageSize, err := validateListPetsBusinessRules(req)
	if err != nil {
		return nil, err
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

// validateShape is the stereotypical proto shape check: nil guard, then
// protovalidate.Validate against buf.validate constraints on the request.
// Shape violations MUST map to InvalidArgument; storage MUST NOT be called.
func validateShape[T proto.Message](msg T) error {
	var zero T
	if any(msg) == any(zero) || isNilProto(msg) {
		return &pet.InvalidArgumentError{Message: "request is required"}
	}
	if err := protovalidate.Validate(msg); err != nil {
		return &pet.InvalidArgumentError{Message: err.Error()}
	}
	return nil
}

// isNilProto reports typed-nil pointers boxed as proto.Message (iface != nil, value nil).
func isNilProto(msg proto.Message) bool {
	v := reflect.ValueOf(msg)
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Map, reflect.Slice, reflect.Chan, reflect.Func:
		return v.IsNil()
	default:
		return false
	}
}

// validateCreatePetBusinessRules applies CreatePet rules beyond buf.validate.
// User-specified pet_id is required in this example (system-generated ids out of scope);
// display_name is required after trim.
func validateCreatePetBusinessRules(req *petstorev1.CreatePetRequest) (name, display string, err error) {
	name, err = petNameFromID(req.GetPetId())
	if err != nil {
		return "", "", err
	}
	display = strings.TrimSpace(req.GetPet().GetDisplayName())
	if display == "" {
		return "", "", &pet.InvalidArgumentError{Message: "display_name is required"}
	}
	return name, display, nil
}

// validateGetPetBusinessRules applies GetPet rules beyond buf.validate:
// name must be a canonical pets/{pet} id (trim + format).
func validateGetPetBusinessRules(req *petstorev1.GetPetRequest) (canonical string, err error) {
	return pet.NormalizePetName(req.GetName())
}

// validateUpdatePetBusinessRules applies UpdatePet rules beyond buf.validate:
// name shape; display_name required after trim; etag optional (trimmed).
func validateUpdatePetBusinessRules(req *petstorev1.UpdatePetRequest) (name, display, etag string, err error) {
	name, err = pet.NormalizePetName(req.GetPet().GetName())
	if err != nil {
		return "", "", "", err
	}
	display = strings.TrimSpace(req.GetPet().GetDisplayName())
	if display == "" {
		return "", "", "", &pet.InvalidArgumentError{Message: "display_name is required"}
	}
	etag = strings.TrimSpace(req.GetPet().GetEtag())
	return name, display, etag, nil
}

// validateDeletePetBusinessRules applies DeletePet rules beyond buf.validate:
// name shape; etag optional (trimmed).
func validateDeletePetBusinessRules(req *petstorev1.DeletePetRequest) (canonical, etag string, err error) {
	canonical, err = pet.NormalizePetName(req.GetName())
	if err != nil {
		return "", "", err
	}
	return canonical, strings.TrimSpace(req.GetEtag()), nil
}

// validateListPetsBusinessRules applies ListPets rules beyond buf.validate:
// default page_size 50; coerce above max 100 (AIP-132). Negative page_size is shape.
func validateListPetsBusinessRules(req *petstorev1.ListPetsRequest) (pageSize int32, err error) {
	pageSize = req.GetPageSize()
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	if pageSize > maxPageSize {
		pageSize = maxPageSize
	}
	return pageSize, nil
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
