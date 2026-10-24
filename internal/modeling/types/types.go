package types

// Undefined is a placeholder for any type not yet modeled.
type Undefined = any

// Idea is the raw starting material — a product idea, a user story,
// a feature request, a domain description. Not yet structured.
type Idea = Undefined

// Requirements is the structured output of the requirements-gathering phase.
type Requirements = Undefined

// ProtoContracts is the output of the contract-design phase —
// proto definitions, OpenAPI specs, ROD resource definitions.
type ProtoContracts = Undefined

// Repository is the output of the implementation phase —
// a scaffolded, wired Go project satisfying the contracts.
type Repository = Undefined

// Specification is the runnable specification — the artifact that is
// executed to verify that the contracts are satisfied.
type Specification = Undefined

// CompletedProject is the full output of the workflow.
type CompletedProject struct {
	Requirements  Requirements
	Contracts     ProtoContracts
	Repository    Repository
	Specification Specification
}
