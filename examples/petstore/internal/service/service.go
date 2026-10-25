package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore/internal/pet"
)

const (
	defaultPageSize int32 = 50
	maxPageSize     int32 = 100
)

// Store is the persistence leaf used by PetService. Defined at the usage site
// (Go idiom: small consumer-owned interfaces). Speaks domain pet.Pet — not
// generated proto types (those stay at the transport boundary).
type Store interface {
	Create(ctx context.Context, p pet.Pet) (pet.Pet, error)
	Get(ctx context.Context, name string) (pet.Pet, error)
	Update(ctx context.Context, p pet.Pet, etag string) (pet.Pet, error)
	Delete(ctx context.Context, name string, etag string) error
	List(ctx context.Context, pageSize int32, pageToken string) (pets []pet.Pet, nextPageToken string, err error)
}

// PetService is the application service: business rules → Store.
// Behavioural contract: docs/specification.md
// It has no dependency on generated gRPC/proto types. The transport adapter
// (internal/transport/grpc) owns request shape checks and proto mapping.
type PetService struct {
	Store Store
}

// CreatePet creates a pet (AIP-133): pet_id builds the resource name; body name
// is ignored at the adapter. Duplicate id → AlreadyExists.
func (s *PetService) CreatePet(ctx context.Context, petID, displayName string) (pet.Pet, error) {
	// TODO: revisit, this is really transforming from ExternalInterface to DomainInterface
	// QUESTION: What is the use case for this? Why isn't it captured?
	// This speaks to this possibly being a transport concern?  All of shape is really.
	// Either way I think this is better expressed pet.New() where New can return a pet.Pet and an error.
	name, display, err := validateCreatePetBusinessRules(petID, displayName)
	if err != nil {
		return pet.Pet{}, err
	}

	// Here we are using domain object pet.Pet as our transport shape.
	created, err := s.Store.Create(ctx, pet.Pet{Name: name, DisplayName: display})
	if err != nil {
		return pet.Pet{}, mapStoreError(err)
	}
	return created, nil
}

// GetPet returns the stored Pet for a resource name.
func (s *PetService) GetPet(ctx context.Context, name string) (pet.Pet, error) {
	canonical, err := validateGetPetBusinessRules(name)
	if err != nil {
		return pet.Pet{}, err
	}

	got, err := s.Store.Get(ctx, canonical)
	if err != nil {
		return pet.Pet{}, mapStoreError(err)
	}
	return got, nil
}

// UpdatePet updates a pet (AIP-134). Optional etag → Aborted on mismatch (AIP-154).
func (s *PetService) UpdatePet(ctx context.Context, name, displayName, etag string) (pet.Pet, error) {
	canonical, display, freshEtag, err := validateUpdatePetBusinessRules(name, displayName, etag)
	if err != nil {
		return pet.Pet{}, err
	}

	updated, err := s.Store.Update(ctx, pet.Pet{Name: canonical, DisplayName: display}, freshEtag)
	if err != nil {
		return pet.Pet{}, mapStoreError(err)
	}
	return updated, nil
}

// DeletePet deletes the pet (AIP-135). Optional etag → Aborted on mismatch.
func (s *PetService) DeletePet(ctx context.Context, name, etag string) error {
	canonical, freshEtag, err := validateDeletePetBusinessRules(name, etag)
	if err != nil {
		return err
	}

	if err := s.Store.Delete(ctx, canonical, freshEtag); err != nil {
		return mapStoreError(err)
	}
	return nil
}

// ListPets lists pets (AIP-132 / AIP-158): default and coerce page_size.
func (s *PetService) ListPets(ctx context.Context, pageSize int32, pageToken string) (pets []pet.Pet, nextPageToken string, err error) {
	size, err := validateListPetsBusinessRules(pageSize)
	if err != nil {
		return nil, "", err
	}

	pets, next, err := s.Store.List(ctx, size, pageToken)
	if err != nil {
		return nil, "", mapStoreError(err)
	}
	if pets == nil {
		pets = []pet.Pet{}
	}
	return pets, next, nil
}

// validateCreatePetBusinessRules: user-specified pet_id required in this example;
// display_name required after trim. (No-op default would return success with passthrough.)
func validateCreatePetBusinessRules(petID, displayName string) (name, display string, err error) {
	name, err = petNameFromID(petID)
	if err != nil {
		return "", "", err
	}
	display = strings.TrimSpace(displayName)
	if display == "" {
		return "", "", &pet.InvalidArgumentError{Message: "display_name is required"}
	}
	return name, display, nil
}

func validateGetPetBusinessRules(name string) (canonical string, err error) {
	return pet.NormalizePetName(name)
}

func validateUpdatePetBusinessRules(name, displayName, etag string) (canonical, display, freshEtag string, err error) {
	canonical, err = pet.NormalizePetName(name)
	if err != nil {
		return "", "", "", err
	}
	display = strings.TrimSpace(displayName)
	if display == "" {
		return "", "", "", &pet.InvalidArgumentError{Message: "display_name is required"}
	}
	return canonical, display, strings.TrimSpace(etag), nil
}

func validateDeletePetBusinessRules(name, etag string) (canonical, freshEtag string, err error) {
	canonical, err = pet.NormalizePetName(name)
	if err != nil {
		return "", "", err
	}
	return canonical, strings.TrimSpace(etag), nil
}

// validateListPetsBusinessRules: default 50; coerce above max 100 (AIP-132).
// Negative page_size is a transport shape concern (buf.validate), not handled here.
func validateListPetsBusinessRules(pageSize int32) (int32, error) {
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	if pageSize > maxPageSize {
		pageSize = maxPageSize
	}
	return pageSize, nil
}

func petNameFromID(petID string) (string, error) {
	return pet.NormalizePetName("pets/" + strings.TrimSpace(petID))
}

func mapStoreError(err error) error {
	var known pet.KnownError
	if errors.As(err, &known) {
		return err
	}
	return &pet.InternalError{Message: fmt.Sprintf("store: %v", err)}
}
