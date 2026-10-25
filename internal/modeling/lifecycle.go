package modeling

import "github.com/stephen-mcelhose/testable-spec-pattern/internal/modeling/types"

// ExecuteSoftwareDevelopmentWorkflow is the top-level function.
// It composes the four phases in sequence.
func ExecuteSoftwareDevelopmentWorkflow(idea types.Idea) (types.CompletedProject, error) {
	requirements, err := GatherRequirements(idea)
	if err != nil {
		return types.CompletedProject{}, err
	}
	contracts, err := BuildContracts(requirements)
	if err != nil {
		return types.CompletedProject{}, err
	}
	project, err := Implement(requirements, contracts)
	if err != nil {
		return types.CompletedProject{}, err
	}
	return types.CompletedProject{Requirements: requirements, Contracts: contracts, Project: project}, nil
}
