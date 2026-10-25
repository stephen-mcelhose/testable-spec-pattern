package grpc_test

import (
	"context"

	"github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore/internal/pet"
)

// stubPets records PetAPI calls so transport specs assert mapping, not business rules.
type stubPets struct {
	createPet pet.Pet
	createErr error
	getPet    pet.Pet
	getErr    error
	updatePet pet.Pet
	updateErr error
	deleteErr error
	listPets  []pet.Pet
	listNext  string
	listErr   error

	createCalls []createCall
	getCalls    []string
	updateCalls []updateCall
	deleteCalls []deleteCall
	listCalls   []listCall
}

type createCall struct {
	petID       string
	displayName string
}

type updateCall struct {
	name        string
	displayName string
	etag        string
}

type deleteCall struct {
	name string
	etag string
}

type listCall struct {
	pageSize  int32
	pageToken string
}

func (s *stubPets) CreatePet(_ context.Context, petID, displayName string) (pet.Pet, error) {
	s.createCalls = append(s.createCalls, createCall{petID: petID, displayName: displayName})
	return s.createPet, s.createErr
}

func (s *stubPets) GetPet(_ context.Context, name string) (pet.Pet, error) {
	s.getCalls = append(s.getCalls, name)
	return s.getPet, s.getErr
}

func (s *stubPets) UpdatePet(_ context.Context, name, displayName, etag string) (pet.Pet, error) {
	s.updateCalls = append(s.updateCalls, updateCall{name: name, displayName: displayName, etag: etag})
	return s.updatePet, s.updateErr
}

func (s *stubPets) DeletePet(_ context.Context, name, etag string) error {
	s.deleteCalls = append(s.deleteCalls, deleteCall{name: name, etag: etag})
	return s.deleteErr
}

func (s *stubPets) ListPets(_ context.Context, pageSize int32, pageToken string) ([]pet.Pet, string, error) {
	s.listCalls = append(s.listCalls, listCall{pageSize: pageSize, pageToken: pageToken})
	return s.listPets, s.listNext, s.listErr
}
