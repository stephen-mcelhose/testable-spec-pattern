package modeling

import "github.com/stephen-mcelhose/testable-spec-pattern/internal/modeling/types"

// ExecuteSoftwareDevelopmentWorkflow is the top-level function.
// It composes the four phases in sequence.
func ExecuteSoftwareDevelopmentWorkflow(idea types.Idea) (types.CompletedProject, error) {
	requirements, err := GatherRequirements(idea, types.DefaultGatherRequirementsStrategy)
	if err != nil {
		return types.CompletedProject{}, err
	}
	contracts, err := BuildContracts(requirements)
	if err != nil {
		return types.CompletedProject{}, err
	}
	artifact, err := Implement(requirements, contracts)
	if err != nil {
		return types.CompletedProject{}, err
	}

	// TODO: extract the RunnableSpecification from Artifact and run it.
	// Blocked on: Artifact interface needs a RunnableSpecification() method,
	// and ImplementStrategy must guarantee it returns ArtifactWithRunnableSpecification.
	// When unblocked:
	//   specification := artifact.RunnableSpecification()
	//   if err = specification.Run(); err != nil {
	//       return types.CompletedProject{}, err
	//   }
	return types.CompletedProject{Requirements: requirements, Contracts: contracts, Artifact: artifact}, nil
}
