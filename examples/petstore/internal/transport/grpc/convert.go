package grpc

import (
	petstorev1 "github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore/internal/gen/petstore/v1"
	"github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore/internal/pet"
)

func petToProto(p pet.Pet) *petstorev1.Pet {
	return &petstorev1.Pet{
		Name:        p.Name,
		DisplayName: p.DisplayName,
		Etag:        p.Etag,
	}
}

func petsToProto(pets []pet.Pet) []*petstorev1.Pet {
	out := make([]*petstorev1.Pet, 0, len(pets))
	for _, p := range pets {
		out = append(out, petToProto(p))
	}
	return out
}
