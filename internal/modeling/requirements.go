package modeling

import "github.com/stephen-mcelhose/testable-spec-pattern/internal/modeling/types"

// DefaultGatherRequirementStrategy composes WriteUseCases, SketchDomainTypes,
// and WriteGherkinScenarios. Assembly into a concrete Requirements value
// requires Requirements to be a defined type.
func DefaultGatherRequirementStrategy(existing types.Requirements, idea types.Idea) (types.Requirements, error) {
	useCases, err := WriteUseCases(idea)
	if err != nil {
		return existing, err
	}
	sketch, err := SketchDomainTypes(idea, useCases)
	if err != nil {
		return existing, err
	}
	scenarios, err := WriteGherkinScenarios(useCases, sketch)
	if err != nil {
		return existing, err
	}
	_ = scenarios // assembled into Requirements once the type is concrete
	return existing, types.ErrUnimplemented
}

// GatherRequirements takes an idea and produces requirements.
// The composition of views (use cases, domain sketches, Gherkin scenarios)
// is determined by the strategy, not fixed here.
func GatherRequirements(idea types.Idea, strategies ...types.GatherRequirementStrategy) (types.Requirements, error) {
	if strategies == nil {
		strategies = []types.GatherRequirementStrategy{DefaultGatherRequirementStrategy}
	}
	var requirements types.Requirements
	var err error
	for _, strategy := range strategies {
		requirements, err = strategy(requirements, idea)
		if err != nil {
			return nil, err
		}
	}
	return requirements, nil
}

// DefaultWriteUseCasesStrategy is the default leaf implementation for WriteUseCases.
// Leaf: backed by the domain-use-cases skill.
func DefaultWriteUseCasesStrategy(existing []types.UseCase, idea types.Idea) ([]types.UseCase, error) {
	return nil, types.ErrUnimplemented
}

// WriteUseCases interrogates an idea and produces a set of use cases —
// each describing an actor, a goal, and the interactions needed to reach it.
func WriteUseCases(idea types.Idea, strategies ...types.WriteUseCasesStrategy) ([]types.UseCase, error) {
	if strategies == nil {
		strategies = []types.WriteUseCasesStrategy{DefaultWriteUseCasesStrategy}
	}
	var useCases []types.UseCase
	var err error
	for _, strategy := range strategies {
		useCases, err = strategy(useCases, idea)
		if err != nil {
			return nil, err
		}
	}
	return useCases, nil
}

// DefaultSketchDomainTypesStrategy is the default leaf implementation for SketchDomainTypes.
// Leaf: backed by the domain-modeling skill.
func DefaultSketchDomainTypesStrategy(existing types.DomainSketch, idea types.Idea, useCases []types.UseCase) (types.DomainSketch, error) {
	return nil, types.ErrUnimplemented
}

// SketchDomainTypes derives a domain sketch from an idea and its use cases —
// named entities, their key attributes, and the relations between them.
func SketchDomainTypes(idea types.Idea, useCases []types.UseCase, strategies ...types.SketchDomainTypesStrategy) (types.DomainSketch, error) {
	if strategies == nil {
		strategies = []types.SketchDomainTypesStrategy{DefaultSketchDomainTypesStrategy}
	}
	var sketch types.DomainSketch
	var err error
	for _, strategy := range strategies {
		sketch, err = strategy(sketch, idea, useCases)
		if err != nil {
			return nil, err
		}
	}
	return sketch, nil
}

// DefaultWriteGherkinScenariosStrategy is the default leaf implementation for WriteGherkinScenarios.
// Leaf: backed by the gherkin-scenarios skill.
func DefaultWriteGherkinScenariosStrategy(existing []types.GherkinScenario, useCases []types.UseCase, sketch types.DomainSketch) ([]types.GherkinScenario, error) {
	return nil, types.ErrUnimplemented
}

// WriteGherkinScenarios derives Given/When/Then scenarios from use cases,
// grounded in the domain names established by the sketch.
func WriteGherkinScenarios(useCases []types.UseCase, sketch types.DomainSketch, strategies ...types.WriteGherkinScenariosStrategy) ([]types.GherkinScenario, error) {
	if strategies == nil {
		strategies = []types.WriteGherkinScenariosStrategy{DefaultWriteGherkinScenariosStrategy}
	}
	var scenarios []types.GherkinScenario
	var err error
	for _, strategy := range strategies {
		scenarios, err = strategy(scenarios, useCases, sketch)
		if err != nil {
			return nil, err
		}
	}
	return scenarios, nil
}
