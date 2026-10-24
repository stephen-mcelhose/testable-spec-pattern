package service_test

import (
	"context"

	"github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore/internal/pet"
)

// stubStore records Store calls so specs can assert the leaf boundary.
// Prefer counterfeiter when the interface is package-owned and you want generate-fakes.
type stubStore struct {
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

	createCalls []pet.Pet
	getCalls    []string
	updateCalls []updateCall
	deleteCalls []deleteCall
	listCalls   []listCall
}

type updateCall struct {
	pet  pet.Pet
	etag string
}

type deleteCall struct {
	name string
	etag string
}

type listCall struct {
	pageSize  int32
	pageToken string
}

func (s *stubStore) Create(_ context.Context, p pet.Pet) (pet.Pet, error) {
	s.createCalls = append(s.createCalls, p)
	return s.createPet, s.createErr
}

func (s *stubStore) Get(_ context.Context, name string) (pet.Pet, error) {
	s.getCalls = append(s.getCalls, name)
	return s.getPet, s.getErr
}

func (s *stubStore) Update(_ context.Context, p pet.Pet, etag string) (pet.Pet, error) {
	s.updateCalls = append(s.updateCalls, updateCall{pet: p, etag: etag})
	return s.updatePet, s.updateErr
}

func (s *stubStore) Delete(_ context.Context, name, etag string) error {
	s.deleteCalls = append(s.deleteCalls, deleteCall{name: name, etag: etag})
	return s.deleteErr
}

func (s *stubStore) List(_ context.Context, pageSize int32, pageToken string) ([]pet.Pet, string, error) {
	s.listCalls = append(s.listCalls, listCall{pageSize: pageSize, pageToken: pageToken})
	return s.listPets, s.listNext, s.listErr
}
