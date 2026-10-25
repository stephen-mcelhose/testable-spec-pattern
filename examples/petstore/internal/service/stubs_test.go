package service_test

import (
	"context"

	petstorev1 "github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore/internal/gen/petstore/v1"
)

// stubStore records Store calls so specs can assert the leaf boundary.
// Prefer counterfeiter when the interface is package-owned and you want generate-fakes.
type stubStore struct {
	createPet *petstorev1.Pet
	createErr error
	getPet    *petstorev1.Pet
	getErr    error
	updatePet *petstorev1.Pet
	updateErr error
	deleteErr error
	listPets  []*petstorev1.Pet
	listNext  string
	listErr   error

	createCalls []*petstorev1.Pet
	getCalls    []string
	updateCalls []updateCall
	deleteCalls []deleteCall
	listCalls   []listCall
}

type updateCall struct {
	pet  *petstorev1.Pet
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

func (s *stubStore) Create(_ context.Context, p *petstorev1.Pet) (*petstorev1.Pet, error) {
	s.createCalls = append(s.createCalls, p)
	return s.createPet, s.createErr
}

func (s *stubStore) Get(_ context.Context, name string) (*petstorev1.Pet, error) {
	s.getCalls = append(s.getCalls, name)
	return s.getPet, s.getErr
}

func (s *stubStore) Update(_ context.Context, p *petstorev1.Pet, etag string) (*petstorev1.Pet, error) {
	s.updateCalls = append(s.updateCalls, updateCall{pet: p, etag: etag})
	return s.updatePet, s.updateErr
}

func (s *stubStore) Delete(_ context.Context, name, etag string) error {
	s.deleteCalls = append(s.deleteCalls, deleteCall{name: name, etag: etag})
	return s.deleteErr
}

func (s *stubStore) List(_ context.Context, pageSize int32, pageToken string) ([]*petstorev1.Pet, string, error) {
	s.listCalls = append(s.listCalls, listCall{pageSize: pageSize, pageToken: pageToken})
	return s.listPets, s.listNext, s.listErr
}
