package main

import (
	"log"

	"github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore/internal/service"
)

func main() {
	// Wire up the service. A real server would inject a Store implementation
	// (e.g. a database-backed store) and start an HTTP/gRPC listener here.
	_ = service.PetService{}
	log.Println("petstore: ready")
}
