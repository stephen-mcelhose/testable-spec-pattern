package petstore_test

import (
	"context"

	"example.com/go-behavioral-spec/examples/03-petstore-api"
)

// stubStore records Store calls so specs can assert the leaf boundary.
// Prefer counterfeiter when the interface is package-owned and you want generate-fakes.
type stubStore struct {
	createPet petstore.Pet
	createErr error
	getPet    petstore.Pet
	getErr    error
	updatePet petstore.Pet
	updateErr error
	deleteErr error
	listPets  []petstore.Pet
	listNext  string
	listErr   error

	createCalls []petstore.Pet
	getCalls    []string
	updateCalls []updateCall
	deleteCalls []deleteCall
	listCalls   []listCall
}

type updateCall struct {
	pet  petstore.Pet
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

func (s *stubStore) Create(_ context.Context, p petstore.Pet) (petstore.Pet, error) {
	s.createCalls = append(s.createCalls, p)
	return s.createPet, s.createErr
}

func (s *stubStore) Get(_ context.Context, name string) (petstore.Pet, error) {
	s.getCalls = append(s.getCalls, name)
	return s.getPet, s.getErr
}

func (s *stubStore) Update(_ context.Context, p petstore.Pet, etag string) (petstore.Pet, error) {
	s.updateCalls = append(s.updateCalls, updateCall{pet: p, etag: etag})
	return s.updatePet, s.updateErr
}

func (s *stubStore) Delete(_ context.Context, name, etag string) error {
	s.deleteCalls = append(s.deleteCalls, deleteCall{name: name, etag: etag})
	return s.deleteErr
}

func (s *stubStore) List(_ context.Context, pageSize int32, pageToken string) ([]petstore.Pet, string, error) {
	s.listCalls = append(s.listCalls, listCall{pageSize: pageSize, pageToken: pageToken})
	return s.listPets, s.listNext, s.listErr
}
