package modeling

import "github.com/stephen-mcelhose/testable-spec-pattern/internal/modeling/types"

// DefaultImplementStrategy composes ScaffoldGoProject, GenerateFromContracts,
// and WireServices to produce a BuiltProject.
func DefaultImplementStrategy(existing types.BuiltProject, requirements types.Requirements, contracts types.Contracts) (types.BuiltProject, error) {
	project, err := ScaffoldGoProject(requirements)
	if err != nil {
		return existing, err
	}
	generated, err := GenerateFromContracts(contracts.Wire)
	if err != nil {
		return existing, err
	}
	return WireServices(project, generated, contracts)
}

// Implement takes requirements and contracts and produces a BuiltProject —
// scaffolded Go project, wired services, generated proto types, and a runnable specification.
func Implement(requirements types.Requirements, contracts types.Contracts, strategies ...types.ImplementStrategy) (types.BuiltProject, error) {
	if strategies == nil {
		strategies = []types.ImplementStrategy{DefaultImplementStrategy}
	}
	var project types.BuiltProject
	var err error
	for _, strategy := range strategies {
		project, err = strategy(project, requirements, contracts)
		if err != nil {
			return nil, err
		}
	}
	return project, nil
}

// DefaultScaffoldGoProjectStrategy is the default leaf implementation for ScaffoldGoProject.
// Leaf: backed by the go-project-scaffold skill. Side effect — writes directory structure to disk.
func DefaultScaffoldGoProjectStrategy(existing types.GoProject, requirements types.Requirements) (types.GoProject, error) {
	return nil, types.ErrUnimplemented
}

// ScaffoldGoProject produces a Go module with idiomatic layout.
func ScaffoldGoProject(requirements types.Requirements, strategies ...types.ScaffoldGoProjectStrategy) (types.GoProject, error) {
	if strategies == nil {
		strategies = []types.ScaffoldGoProjectStrategy{DefaultScaffoldGoProjectStrategy}
	}
	var project types.GoProject
	var err error
	for _, strategy := range strategies {
		project, err = strategy(project, requirements)
		if err != nil {
			return nil, err
		}
	}
	return project, nil
}

// DefaultGenerateFromContractsStrategy is the default leaf implementation for GenerateFromContracts.
// Leaf: side effect — invokes codegen tooling, writes generated files to disk.
func DefaultGenerateFromContractsStrategy(existing types.GeneratedCode, contracts types.ProtoContracts) (types.GeneratedCode, error) {
	return nil, types.ErrUnimplemented
}

// GenerateFromContracts runs codegen against the wire contracts,
// producing Go types and service stubs.
func GenerateFromContracts(contracts types.ProtoContracts, strategies ...types.GenerateFromContractsStrategy) (types.GeneratedCode, error) {
	if strategies == nil {
		strategies = []types.GenerateFromContractsStrategy{DefaultGenerateFromContractsStrategy}
	}
	var generated types.GeneratedCode
	var err error
	for _, strategy := range strategies {
		generated, err = strategy(generated, contracts)
		if err != nil {
			return nil, err
		}
	}
	return generated, nil
}

// DefaultWireServicesStrategy is the default leaf implementation for WireServices.
// Leaf: side effect — writes Go source files to disk.
func DefaultWireServicesStrategy(existing types.BuiltProject, project types.GoProject, generated types.GeneratedCode, contracts types.Contracts) (types.BuiltProject, error) {
	return nil, types.ErrUnimplemented
}

// WireServices connects service implementations to generated code,
// producing a BuiltProject whose specification can be run.
func WireServices(project types.GoProject, generated types.GeneratedCode, contracts types.Contracts, strategies ...types.WireServicesStrategy) (types.BuiltProject, error) {
	if strategies == nil {
		strategies = []types.WireServicesStrategy{DefaultWireServicesStrategy}
	}
	var built types.BuiltProject
	var err error
	for _, strategy := range strategies {
		built, err = strategy(built, project, generated, contracts)
		if err != nil {
			return nil, err
		}
	}
	return built, nil
}
