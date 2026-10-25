package modeling

import "github.com/stephen-mcelhose/testable-spec-pattern/internal/modeling/types"

func DefaultGatherRequirementStrategy(existing types.Requirements, idea types.Idea) (types.Requirements, error) {
	return nil, types.ErrUnimplemented
}

// GatherRequirements takes an idea and produces requirements.
// The composition of views (use cases, domain sketches, Gherkin scenarios)
// is determined by the strategy, not fixed here.
// Leaf: backed by requirements-gathering skills.
func GatherRequirements(idea types.Idea, strategies ...types.GatherRequirementStrategy) (types.Requirements, error) {
	if strategies == nil {
		strategies = []types.GatherRequirementStrategy{DefaultGatherRequirementStrategy}
	}
	var requirements types.Requirements
	var err error
	for _, strategy := range strategies {
		requirements, err = strategy(requirements, idea)
		if err != nil {
			return types.Requirements(nil), err
		}
	}
	return requirements, nil
}

// WriteUseCases interrogates an idea and produces a set of use cases —
// each describing an actor, a goal, and the interactions needed to reach it.
// Leaf: backed by the write-use-cases skill.
func WriteUseCases(idea types.Idea) ([]types.UseCase, error) {
	return nil, types.ErrUnimplemented
}

// SketchDomainTypes derives a domain sketch from an idea and its use cases —
// named entities, their key attributes, and the relations between them.
// Leaf: backed by the sketch-domain-types skill.
func SketchDomainTypes(idea types.Idea, useCases []types.UseCase) (types.DomainSketch, error) {
	return nil, types.ErrUnimplemented
}

// WriteGherkinScenarios derives Given/When/Then scenarios from use cases,
// grounded in the domain names established by the sketch.
// Leaf: backed by the write-gherkin-scenarios skill.
func WriteGherkinScenarios(useCases []types.UseCase, sketch types.DomainSketch) ([]types.GherkinScenario, error) {
	return nil, types.ErrUnimplemented
}
