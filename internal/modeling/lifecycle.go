package modeling

import "github.com/stephen-mcelhose/testable-spec-pattern/internal/modeling/types"

// ExecuteSoftwareDevelopmentWorkflow is the top-level function.
// It composes the four phases in sequence.
func ExecuteSoftwareDevelopmentWorkflow(idea types.Idea) (types.CompletedProject, error) {
	panic("not implemented")
}

// GatherRequirements takes an idea and produces structured requirements —
// use cases, domain type sketches, Gherkin scenarios.
func GatherRequirements(idea types.Idea) (types.Requirements, error) {
	panic("not implemented")
}

// BuildContracts takes requirements and produces machine-readable contracts —
// proto definitions, OpenAPI specs, ROD resource definitions.
func BuildContracts(requirements types.Requirements) (types.ProtoContracts, error) {
	panic("not implemented")
}

// Implement takes requirements and contracts and produces a repository —
// scaffolded Go project, wired services, generated proto types.
// This function has significant internal structure (many steps).
func Implement(requirements types.Requirements, contracts types.ProtoContracts) (types.Repository, error) {
	panic("not implemented")
}

// RunRunnableSpecification takes a specification and runs it.
// Returns an error if any requirement is violated.
// A passing run is the definition of a completed project.
func RunRunnableSpecification(specification types.Specification) error {
	panic("not implemented")
}
