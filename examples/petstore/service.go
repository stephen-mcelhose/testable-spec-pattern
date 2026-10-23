package petstore

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

const (
	defaultPageSize int32 = 50
	maxPageSize     int32 = 100
)

// Store is the persistence leaf used by PetService. Defined at the usage site
// (Go idiom: small consumer-owned interfaces). Tests inject a hand stub; prefer
// counterfeiter when call-count or argument capture is needed.
type Store interface {
	Create(ctx context.Context, pet Pet) (Pet, error)
	Get(ctx context.Context, name string) (Pet, error)
	Update(ctx context.Context, pet Pet, etag string) (Pet, error)
	Delete(ctx context.Context, name string, etag string) error
	List(ctx context.Context, pageSize int32, pageToken string) (pets []Pet, nextPageToken string, err error)
}

// PetService is a thin AIP-flavoured Pets API: validate shape, call Store, map errors.
type PetService struct {
	Store Store
}

// CreatePet creates a pet (AIP-133): pet_id builds the resource name; Pet.Name on
// the body is ignored; returns the fully populated stored Pet.
// Duplicate id → AlreadyExists.
func (s *PetService) CreatePet(ctx context.Context, req CreatePetRequest) (Pet, error) {
	name, err := petNameFromID(req.PetID)
	if err != nil {
		return Pet{}, err
	}
	display := strings.TrimSpace(req.Pet.DisplayName)
	if display == "" {
		return Pet{}, &InvalidArgumentError{Message: "display_name is required"}
	}

	created, err := s.Store.Create(ctx, Pet{Name: name, DisplayName: display})
	if err != nil {
		return Pet{}, mapStoreError(err)
	}
	return created, nil
}

// GetPet validates the resource name and returns the stored Pet.
func (s *PetService) GetPet(ctx context.Context, name string) (Pet, error) {
	canonical, err := NormalizePetName(name)
	if err != nil {
		return Pet{}, err
	}

	got, err := s.Store.Get(ctx, canonical)
	if err != nil {
		return Pet{}, mapStoreError(err)
	}
	return got, nil
}

// UpdatePet updates a pet (AIP-134): Pet.Name identifies the resource; Pet.Etag is
// optional freshness (AIP-154) — when set and mismatched, Store returns Aborted.
// Returns the fully populated stored Pet. update_mask / allow_missing are out of scope.
func (s *PetService) UpdatePet(ctx context.Context, req UpdatePetRequest) (Pet, error) {
	name, err := NormalizePetName(req.Pet.Name)
	if err != nil {
		return Pet{}, err
	}
	display := strings.TrimSpace(req.Pet.DisplayName)
	if display == "" {
		return Pet{}, &InvalidArgumentError{Message: "display_name is required"}
	}
	etag := strings.TrimSpace(req.Pet.Etag)

	updated, err := s.Store.Update(ctx, Pet{Name: name, DisplayName: display}, etag)
	if err != nil {
		return Pet{}, mapStoreError(err)
	}
	return updated, nil
}

// DeletePet deletes the pet (AIP-135). Etag is optional protected-delete freshness
// (AIP-154): when provided and mismatched, Store returns Aborted.
func (s *PetService) DeletePet(ctx context.Context, name, etag string) error {
	canonical, err := NormalizePetName(name)
	if err != nil {
		return err
	}

	if err := s.Store.Delete(ctx, canonical, strings.TrimSpace(etag)); err != nil {
		return mapStoreError(err)
	}
	return nil
}

// ListPets lists pets (AIP-132 / AIP-158): default and coerce page_size; return a
// page. Empty collection is success. filter / order_by / show_deleted are out of scope.
func (s *PetService) ListPets(ctx context.Context, req ListPetsRequest) (ListPetsResponse, error) {
	pageSize := req.PageSize
	if pageSize < 0 {
		return ListPetsResponse{}, &InvalidArgumentError{Message: "page_size must not be negative"}
	}
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	if pageSize > maxPageSize {
		pageSize = maxPageSize // AIP-132: coerce above max; do not error
	}

	pets, next, err := s.Store.List(ctx, pageSize, req.PageToken)
	if err != nil {
		return ListPetsResponse{}, mapStoreError(err)
	}
	if pets == nil {
		pets = []Pet{}
	}
	return ListPetsResponse{Pets: pets, NextPageToken: next}, nil
}

// petNameFromID builds pets/{pet_id} and validates via NormalizePetName (AIP-133).
func petNameFromID(petID string) (string, error) {
	return NormalizePetName("pets/" + strings.TrimSpace(petID))
}

// mapStoreError passes through KnownError and wraps anything else as Internal.
func mapStoreError(err error) error {
	var known KnownError
	if errors.As(err, &known) {
		return err
	}
	return &InternalError{Message: fmt.Sprintf("store: %v", err)}
}
