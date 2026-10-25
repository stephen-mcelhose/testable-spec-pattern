package modeling

import "github.com/stephen-mcelhose/testable-spec-pattern/internal/modeling/types"

// Implement takes requirements and contracts and produces a BuiltProject —
// scaffolded Go project, wired services, generated proto types, and a runnable specification.
func Implement(requirements types.Requirements, contracts types.ProtoContracts) (types.BuiltProject, error) {
	project, err := ScaffoldGoProject(requirements)
	if err != nil {
		return nil, err
	}
	generated, err := GenerateFromContracts(contracts)
	if err != nil {
		return nil, err
	}
	return WireServices(project, generated, contracts)
}

// ScaffoldGoProject produces a Go module with idiomatic layout.
// Backed by the go-project-scaffold skill.
// Leaf: side effect — writes directory structure to disk.
func ScaffoldGoProject(requirements types.Requirements) (types.GoProject, error) {
	return nil, types.ErrUnimplemented
}

// GenerateFromContracts runs codegen against the contracts,
// producing Go types and service stubs.
// Leaf: side effect — invokes codegen tooling, writes generated files to disk.
func GenerateFromContracts(contracts types.ProtoContracts) (types.GeneratedCode, error) {
	return nil, types.ErrUnimplemented
}

// WireServices connects service implementations to generated code,
// producing a BuiltProject whose specification can be run.
// Leaf: side effect — writes Go source files to disk.
func WireServices(project types.GoProject, generated types.GeneratedCode, contracts types.ProtoContracts) (types.BuiltProject, error) {
	return nil, types.ErrUnimplemented
}

// ModelAsCode takes an idea and produces a CodeArtifact —
// a Go package modeling the idea using types, interfaces, and strategy functions.
func ModelAsCode(idea types.Idea) (types.CodeArtifact, error) {
	return nil, types.ErrUnimplemented
}
