package modeling

import "github.com/stephen-mcelhose/testable-spec-pattern/internal/modeling/types"

// GatherRequirements takes an idea and produces structured requirements by
// composing use-case writing, domain sketching, and Gherkin scenario derivation.
func GatherRequirements(idea types.Idea) (types.Requirements, error) {
	useCases, err := WriteUseCases(idea)
	if err != nil {
		return types.Requirements{}, err
	}
	sketch, err := SketchDomainTypes(idea, useCases)
	if err != nil {
		return types.Requirements{}, err
	}
	scenarios, err := WriteGherkinScenarios(useCases, sketch)
	if err != nil {
		return types.Requirements{}, err
	}
	return types.Requirements{UseCases: useCases, Sketch: sketch, Scenarios: scenarios}, nil
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
