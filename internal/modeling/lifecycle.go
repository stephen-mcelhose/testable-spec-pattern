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

// ExecuteSoftwareDevelopmentWorkflowFromDomainModel runs from an existing domain model (CodeContracts).
// Use when a domain model has been produced via ModelDomain and wire contracts still need to be built.
func ExecuteSoftwareDevelopmentWorkflowFromDomainModel(requirements types.Requirements, code types.CodeContracts) (types.CompletedProject, error) {
	wire, err := BuildWireContracts(requirements)
	if err != nil {
		return types.CompletedProject{}, err
	}
	contracts := types.Contracts{Code: code, Wire: wire}
	return ExecuteSoftwareDevelopmentWorkflowFromContracts(requirements, contracts)
}

// ExecuteSoftwareDevelopmentWorkflowFromContracts runs from existing contracts.
// Use when contracts have been designed externally (e.g. given an existing proto or OpenAPI spec).
func ExecuteSoftwareDevelopmentWorkflowFromContracts(requirements types.Requirements, contracts types.Contracts) (types.CompletedProject, error) {
	spec, err := WriteSpecification(requirements, contracts)
	if err != nil {
		return types.CompletedProject{}, err
	}
	return ExecuteSoftwareDevelopmentWorkflowFromSpecification(requirements, contracts, spec)
}

// ExecuteSoftwareDevelopmentWorkflowFromSpecification runs from an existing specification.
// Use when contracts and a prose spec already exist and only implementation remains.
func ExecuteSoftwareDevelopmentWorkflowFromSpecification(requirements types.Requirements, contracts types.Contracts, spec types.Specification) (types.CompletedProject, error) {
	project, err := Implement(requirements, contracts)
	if err != nil {
		return types.CompletedProject{}, err
	}
	return types.CompletedProject{
		Requirements:  requirements,
		Contracts:     contracts,
		Specification: spec,
		Project:       project,
	}, nil
}
