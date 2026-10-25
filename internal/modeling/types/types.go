package types

import "errors"

// Undefined is a placeholder for any type not yet modeled.
type Undefined = any

// ErrUnimplemented is returned by leaf functions that represent side effects
// or implementation details not yet modeled.
var ErrUnimplemented = errors.New("unimplemented")

// ---------------------------------------------------------------------------
// Core lifecycle types
// ---------------------------------------------------------------------------

// Idea is the raw starting material — a product idea, a user story,
// a feature request, a domain description. Not yet structured.
type Idea interface{ Undefined }

// CompletedProject is the full output of the workflow.
type CompletedProject struct {
	Requirements  Requirements
	Contracts     ProtoContracts
	Artifact      Artifact
	Specification Specification
}

// ---------------------------------------------------------------------------
// Requirements Gathering
// ---------------------------------------------------------------------------

// Requirements is the structured output of the requirements-gathering phase.
type Requirements struct {
	UseCases  []UseCase
	Sketch    DomainSketch
	Scenarios []GherkinScenario
}

// UseCase describes a single actor–goal interaction with the system.
type UseCase interface{ Undefined }

// Actor is an entity that interacts with the system to achieve a goal.
type Actor interface{ Undefined }

// Goal is what an Actor is trying to accomplish.
type Goal interface{ Undefined }

// DomainSketch is a lightweight model of the domain entities and their
// relations, produced before any wire format is chosen.
type DomainSketch struct {
	Types     []DomainType
	Relations []DomainRelation
}

// DomainType is a named entity in the domain — not a Go type, not a proto
// message, just a named concept with a role.
type DomainType interface{ Undefined }

// DomainRelation is a directional connection between two DomainTypes —
// ownership, association, or dependency.
type DomainRelation interface{ Undefined }

// GherkinScenario is a Given/When/Then specification of a single behaviour,
// grounded in domain names.
type GherkinScenario interface{ Undefined }

// ---------------------------------------------------------------------------
// Contract Design
// ---------------------------------------------------------------------------

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

// ---------------------------------------------------------------------------
// Implementation
// ---------------------------------------------------------------------------

// Artifact is the output of the implementation phase —
// a scaffolded, wired Go project satisfying the contracts.
type Artifact interface {
	Undefined
}

// ArtifactWithRunnableSpecification is the concrete artifact produced when
// Implement succeeds and the output includes a runnable specification.
// This is the normal case: a completed implementation always includes its spec.
type ArtifactWithRunnableSpecification struct {
	Artifact
}

// GoProject is a scaffolded Go module with idiomatic layout.
type GoProject interface{ Undefined }

// ProtoGeneration is the output of running buf generate — generated Go types
// and gRPC stubs derived from the proto files.
type ProtoGeneration interface{ Undefined }

// ---------------------------------------------------------------------------
// Model as Code
// ---------------------------------------------------------------------------

// CodeArtifact is the output of the model-as-code phase —
// a Go package expressing a domain model using types, interfaces,
// and strategy functions as the modeling primitives.
type CodeArtifact interface{ Undefined }

// Subject is a domain entity extracted from an Idea —
// the raw material for a DomainType or Go interface.
type Subject interface{ Undefined }

// Interaction is a relationship, operation, or behaviour between Subjects —
// the raw material for a Go function.
type Interaction interface{ Undefined }

// GoType is a modeled Subject expressed as a Go type definition.
type GoType interface{ Undefined }

// GoFunction is a modeled Interaction expressed as a Go function signature
// and, where it composes sub-functions, its body.
type GoFunction interface{ Undefined }

// ---------------------------------------------------------------------------
// Specification
// ---------------------------------------------------------------------------

// Specification is the runnable specification — the artifact that is
// executed to verify that the contracts are satisfied.
type Specification interface{ Undefined }

// RunnableSpecification is the executable form of the spec — Go tests named as
// RFC 2119 requirements. Running it is the definition of a completed project:
// a passing run means the implementation satisfies its contracts.
type RunnableSpecification struct {
	Artifact
}

func NewRunnableSpecification(artifact ArtifactWithRunnableSpecification) RunnableSpecification {
	return RunnableSpecification{Artifact: artifact}
}

// Run executes the specification. Returns an error if any requirement is violated.
// In practice this invokes go-behavioral-spec — the test suite is the spec.
func (r RunnableSpecification) Run() error {
	return ErrUnimplemented
}

// ---------------------------------------------------------------------------
// Strategies — pluggable phase implementations, each backed by a skill
// ---------------------------------------------------------------------------

// GatherRequirementsStrategy is the pluggable behaviour for the requirements-gathering phase.
// In practice each implementation is an agentic skill:
//   - write-use-cases
//   - sketch-domain-types
//   - write-gherkin-scenarios
//
// A composite strategy may chain these in sequence.
type GatherRequirementsStrategy func(idea Idea) (Requirements, error)

// DefaultGatherRequirementsStrategy is a no-op placeholder.
// Replace with a real strategy backed by the requirements-gathering skills.
func DefaultGatherRequirementsStrategy(idea Idea) (Requirements, error) {
	return Requirements{}, ErrUnimplemented
}

// BuildContractsStrategy is the pluggable behaviour for the contract-design phase.
// In practice the implementation is the resource-oriented-design skill,
// which produces AIP-compliant proto/OpenAPI/ROD definitions from requirements.
type BuildContractsStrategy func(requirements Requirements) (ProtoContracts, error)

// ImplementStrategy is the pluggable behaviour for the implementation phase.
// In practice the implementation is the go-project-scaffold skill plus
// any proto-wiring steps. Returns ArtifactWithRunnableSpecification so the
// caller can extract and run the specification without type-asserting.
type ImplementStrategy func(requirements Requirements, contracts ProtoContracts) (Artifact, error)

// DefaultImplementStrategy is a no-op placeholder.
// Replace with a real strategy backed by the implementation skills.
func DefaultImplementStrategy(requirements Requirements, contracts ProtoContracts) (ArtifactWithRunnableSpecification, error) {
	return ArtifactWithRunnableSpecification{}, ErrUnimplemented
}

// ModelAsCodeStrategy is the pluggable behaviour for the model-as-code phase.
// In practice the implementation is the model-as-code skill.
type ModelAsCodeStrategy func(idea Idea) (CodeArtifact, error)
