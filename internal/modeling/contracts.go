package modeling

import "github.com/stephen-mcelhose/testable-spec-pattern/internal/modeling/types"

// BuildContracts takes requirements and produces machine-readable contracts —
// proto definitions, OpenAPI specs, ROD resource definitions.
// It composes resource design, service definition, and proto file generation.
func BuildContracts(requirements types.Requirements) (types.ProtoContracts, error) {
	resources, err := DesignResourceModel(requirements)
	if err != nil {
		return types.ProtoContracts{}, err
	}
	services, err := DefineServices(resources)
	if err != nil {
		return types.ProtoContracts{}, err
	}
	files, err := GenerateProtoFiles(services)
	if err != nil {
		return types.ProtoContracts{}, err
	}
	return types.ProtoContracts{Files: files}, nil
}

// DesignResourceModel derives AIP resources from requirements —
// named, revisioned domain entities with standard lifecycle methods.
// Backed by the resource-oriented-design skill.
func DesignResourceModel(requirements types.Requirements) ([]types.Resource, error) {
	return nil, types.ErrUnimplemented
}

// DefineServices groups resources into gRPC service definitions,
// each with the standard methods appropriate for its resource type.
func DefineServices(resources []types.Resource) ([]types.Service, error) {
	return nil, types.ErrUnimplemented
}

// GenerateProtoFiles produces .proto files from service definitions,
// including googleapis annotations for AIP compliance.
// Leaf: side effect — writes files to disk.
func GenerateProtoFiles(services []types.Service) ([]types.ProtoFile, error) {
	return nil, types.ErrUnimplemented
}
