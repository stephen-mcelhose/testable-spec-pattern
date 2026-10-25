package types

import "errors"

// Undefined is a placeholder for any type not yet modeled.
type Undefined = any

// ErrUnimplemented is returned by leaf functions that represent side effects
// or implementation details not yet modeled.
var ErrUnimplemented = errors.New("unimplemented")

// Idea is the raw starting material — a product idea, a user story,
// a feature request, a domain description. Not yet structured.
type Idea interface{ Undefined }

// CompletedProject is the full output of the workflow — provenance plus the
// built project carrying its runnable specification.
type CompletedProject struct {
	Requirements Requirements
	Contracts    ProtoContracts
	Project      BuiltProject
}

// Requirements is the structured output of the requirements-gathering phase.
type Requirements struct {
	UseCases  []UseCase
	Sketch    DomainSketch
	Scenarios []GherkinScenario
}

// UseCase describes a single actor–goal interaction with the system.
type UseCase interface{ Undefined }

// DomainSketch is a lightweight model of the domain entities and their
// relations, produced before any wire format is chosen.
type DomainSketch interface{ Undefined }

// GherkinScenario is a Given/When/Then specification of a single behaviour,
// grounded in domain names.
type GherkinScenario interface{ Undefined }

// ProtoContracts is the output of the contract-design phase.
type ProtoContracts struct {
	Files []ProtoFile
}

// ProtoFile is a single .proto file defining resources, services, and methods.
type ProtoFile interface{ Undefined }

// Resource is a domain entity modeled as an AIP resource —
// named, revisioned, with standard lifecycle methods.
type Resource interface{ Undefined }

// Service is a gRPC service definition grouping methods for a resource.
type Service interface{ Undefined }

// Method is a single RPC method on a Service.
type Method interface{ Undefined }

// BuiltProject is the output of the implementation phase — a scaffolded,
// wired Go project satisfying its contracts. Every project built by this
// workflow carries a runnable specification: that is the domain invariant.
type BuiltProject interface {
	Specification() RunnableSpecification
}

// GoProject is a scaffolded Go module with idiomatic layout.
type GoProject interface{ Undefined }

// GeneratedCode is the output of contract codegen —
// Go types and service stubs derived from the contracts.
type GeneratedCode interface{ Undefined }

// CodeArtifact is the output of the model-as-code phase —
// a Go package expressing a domain model using types, interfaces,
// and strategy functions as the modeling primitives.
type CodeArtifact interface{ Undefined }

// RunnableSpecification is the executable form of the spec — Go tests named as
// RFC 2119 requirements. A passing run means the implementation satisfies its contracts.
type RunnableSpecification interface {
	// Run executes the specification. Returns an error if any requirement is violated.
	// In practice this invokes go-behavioral-spec — the test suite is the spec.
	Run() error
}

// IsVerifiedProject reports whether a BuiltProject's specification passes.
func IsVerifiedProject(project BuiltProject) bool {
	return project.Specification().Run() == nil
}
