package main

import (
	"log"

	"github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore/internal/service"
	petgrpc "github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore/internal/transport/grpc"
)

func main() {
	// Wire application service + gRPC adapter. A real server would inject a
	// Store implementation and RegisterPetServiceServer on a grpc.Server.
	pets := &service.PetService{}
	_ = petgrpc.PetServer{Pets: pets} // PetAPI satisfied by *service.PetService
	log.Println("petstore: ready")
}
