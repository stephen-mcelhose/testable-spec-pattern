package modeling

import "github.com/stephen-mcelhose/testable-spec-pattern/internal/modeling/types"

// The workflow does not have to start at requirements gathering.
// Each entry point below represents a valid starting phase given its prerequisites.
// Use the earliest entry point for which you have the inputs.

// ExecuteSoftwareDevelopmentWorkflow runs all phases from an Idea.
func ExecuteSoftwareDevelopmentWorkflow(idea types.Idea) (types.CompletedProject, error) {
	requirements, err := GatherRequirements(idea)
	if err != nil {
		return types.CompletedProject{}, err
	}
	return ExecuteSoftwareDevelopmentWorkflowFromRequirements(requirements)
}

// ExecuteSoftwareDevelopmentWorkflowFromRequirements runs from existing requirements.
// Use when requirements have been gathered externally or in a prior session.
func ExecuteSoftwareDevelopmentWorkflowFromRequirements(requirements types.Requirements) (types.CompletedProject, error) {
	contracts, err := BuildContracts(requirements)
	if err != nil {
		return types.CompletedProject{}, err
	}
	return ExecuteSoftwareDevelopmentWorkflowFromContracts(requirements, contracts)
}

// ExecuteSoftwareDevelopmentWorkflowFromContracts runs from existing contracts.
// Use when contracts have been designed externally (e.g. given an existing proto or OpenAPI spec).
func ExecuteSoftwareDevelopmentWorkflowFromContracts(requirements types.Requirements, contracts types.ProtoContracts) (types.CompletedProject, error) {
	project, err := Implement(requirements, contracts)
	if err != nil {
		return types.CompletedProject{}, err
	}
	return types.CompletedProject{Requirements: requirements, Contracts: contracts, Project: project}, nil
}
