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
	Requirements  Requirements
	Contracts     Contracts
	Specification Specification
	Project       BuiltProject
}

// Requirements is the output of the requirements-gathering phase.
// Its internal structure depends on the strategy used — use cases,
// domain sketches, and Gherkin scenarios are views of requirements,
// not a fixed composition.
type Requirements interface{ Undefined }

// UseCase describes a single actor–goal interaction with the system.
type UseCase interface{ Undefined }

// DomainSketch is a lightweight model of the domain entities and their
// relations, produced before any wire format is chosen.
type DomainSketch interface{ Undefined }

// GherkinScenario is a Given/When/Then specification of a single behaviour,
// grounded in domain names.
type GherkinScenario interface{ Undefined }

// Contracts is the full set of contracts for a system —
// code contracts (layer boundaries, compile-time enforced) and
// wire contracts (process boundaries, runtime enforced).
type Contracts struct {
	Code CodeContracts
	Wire ProtoContracts
}

// CodeContracts is the output of the domain modeling phase —
// Go interfaces and strategy function signatures at layer boundaries.
// A domain model is a set of code contracts: it names the domain's
// interfaces and stub functions, leaving implementations unresolved.
type CodeContracts interface{ Undefined }

// ProtoContracts is the output of the wire contract design phase.
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

// Specification is the prose behavioral specification —
// RFC 2119-style descriptions of how a system must behave,
// grounded in domain names from the contracts.
type Specification interface{ Undefined }

// BuiltProject is the output of the implementation phase — a scaffolded,
// wired Go project satisfying its contracts. Every project built by this
// workflow carries a runnable specification: that is the domain invariant.
type BuiltProject interface {
	Specification() RunnableSpecification
}

// GoProject is a scaffolded Go module with idiomatic layout.
type GoProject interface{ Undefined }

// GeneratedCode is the output of contract codegen —
// Go types and service stubs derived from the wire contracts.
type GeneratedCode interface{ Undefined }

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

// Strategy types.
// Each strategy takes the current accumulated output and the phase inputs,
// returning the updated output. Strategies are applied in sequence,
// allowing composition of independent skills into a single phase.

// Requirements gathering strategies.
type GatherRequirementStrategy func(existing Requirements, idea Idea) (Requirements, error)
type WriteUseCasesStrategy func(existing []UseCase, idea Idea) ([]UseCase, error)
type SketchDomainTypesStrategy func(existing DomainSketch, idea Idea, useCases []UseCase) (DomainSketch, error)
type WriteGherkinScenariosStrategy func(existing []GherkinScenario, useCases []UseCase, sketch DomainSketch) ([]GherkinScenario, error)

// Contracts & domain modeling strategies.
type ModelDomainStrategy func(existing CodeContracts, idea Idea) (CodeContracts, error)
type BuildContractStrategy func(existing Contracts, requirements Requirements) (Contracts, error)
type BuildWireContractsStrategy func(existing ProtoContracts, requirements Requirements) (ProtoContracts, error)
type DesignResourceModelStrategy func(existing []Resource, requirements Requirements) ([]Resource, error)
type DefineServicesStrategy func(existing []Service, resources []Resource) ([]Service, error)
type GenerateProtoFilesStrategy func(existing []ProtoFile, services []Service) ([]ProtoFile, error)

// Specification strategies.
type WriteSpecificationStrategy func(existing Specification, requirements Requirements, contracts Contracts) (Specification, error)

// Implementation strategies.
type ImplementStrategy func(existing BuiltProject, requirements Requirements, contracts Contracts) (BuiltProject, error)
type ScaffoldGoProjectStrategy func(existing GoProject, requirements Requirements) (GoProject, error)
type GenerateFromContractsStrategy func(existing GeneratedCode, contracts ProtoContracts) (GeneratedCode, error)
type WireServicesStrategy func(existing BuiltProject, project GoProject, generated GeneratedCode, contracts Contracts) (BuiltProject, error)
