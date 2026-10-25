package modeling

import "github.com/stephen-mcelhose/testable-spec-pattern/internal/modeling/types"

// Implement takes requirements and contracts and produces an Artifact —
// scaffolded Go project, wired services, generated proto types.
func Implement(requirements types.Requirements, contracts types.ProtoContracts) (types.Artifact, error) {
	project, err := ScaffoldGoProject(requirements)
	if err != nil {
		return nil, err
	}
	generated, err := GenerateGoFromProto(contracts)
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

// GenerateGoFromProto runs buf generate against the proto files,
// producing Go types and gRPC stubs.
// Leaf: side effect — invokes buf, writes generated files to disk.
func GenerateGoFromProto(contracts types.ProtoContracts) (types.ProtoGeneration, error) {
	return nil, types.ErrUnimplemented
}

// WireServices connects service implementations to generated proto types.
// Leaf: side effect — writes Go source files to disk.
func WireServices(project types.GoProject, generated types.ProtoGeneration, contracts types.ProtoContracts) (types.ArtifactWithRunnableSpecification, error) {
	return types.ArtifactWithRunnableSpecification{}, types.ErrUnimplemented
}

// RunRunnableSpecification executes the specification.
// Returns an error if any requirement is violated.
// Leaf: side effect — invokes go test.
func RunRunnableSpecification(specification types.Specification) error {
	return types.ErrUnimplemented
}

// ModelAsCode takes an idea and produces a CodeArtifact —
// a Go package modeling the idea using types, interfaces, and strategy functions.
func ModelAsCode(idea types.Idea) (types.CodeArtifact, error) {
	return nil, types.ErrUnimplemented
}
