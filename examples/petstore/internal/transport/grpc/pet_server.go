package grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	petstorev1 "github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore/internal/gen/petstore/v1"
	"github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore/internal/service"
)

// PetServer is the gRPC transport adapter for PetService.
// It owns: (1) protovalidate shape, (2) proto ↔ domain mapping, (3) forward to app.
// KnownError → google.rpc.Status mapping waits on server wiring.
type PetServer struct {
	petstorev1.UnimplementedPetServiceServer
	Pets *service.PetService
}

var _ petstorev1.PetServiceServer = (*PetServer)(nil)

// CreatePet — shape → map → app CreatePet → map response.
func (s *PetServer) CreatePet(ctx context.Context, req *petstorev1.CreatePetRequest) (*petstorev1.Pet, error) {
	if err := validateShape(req); err != nil {
		return nil, err
	}

	created, err := s.Pets.CreatePet(ctx, req.GetPetId(), req.GetPet().GetDisplayName())
	if err != nil {
		return nil, err
	}
	return petToProto(created), nil
}

// GetPet — shape → map → app GetPet → map response.
func (s *PetServer) GetPet(ctx context.Context, req *petstorev1.GetPetRequest) (*petstorev1.Pet, error) {
	if err := validateShape(req); err != nil {
		return nil, err
	}

	got, err := s.Pets.GetPet(ctx, req.GetName())
	if err != nil {
		return nil, err
	}
	return petToProto(got), nil
}

// UpdatePet — shape → map → app UpdatePet → map response.
func (s *PetServer) UpdatePet(ctx context.Context, req *petstorev1.UpdatePetRequest) (*petstorev1.Pet, error) {
	if err := validateShape(req); err != nil {
		return nil, err
	}

	p := req.GetPet()
	updated, err := s.Pets.UpdatePet(ctx, p.GetName(), p.GetDisplayName(), p.GetEtag())
	if err != nil {
		return nil, err
	}
	return petToProto(updated), nil
}

// DeletePet — shape → map → app DeletePet.
func (s *PetServer) DeletePet(ctx context.Context, req *petstorev1.DeletePetRequest) (*emptypb.Empty, error) {
	if err := validateShape(req); err != nil {
		return nil, err
	}

	if err := s.Pets.DeletePet(ctx, req.GetName(), req.GetEtag()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// ListPets — shape → map → app ListPets → map response.
func (s *PetServer) ListPets(ctx context.Context, req *petstorev1.ListPetsRequest) (*petstorev1.ListPetsResponse, error) {
	if err := validateShape(req); err != nil {
		return nil, err
	}

	pets, next, err := s.Pets.ListPets(ctx, req.GetPageSize(), req.GetPageToken())
	if err != nil {
		return nil, err
	}
	return &petstorev1.ListPetsResponse{Pets: petsToProto(pets), NextPageToken: next}, nil
}
