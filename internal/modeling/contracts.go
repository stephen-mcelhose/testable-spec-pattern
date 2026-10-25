package modeling

import "github.com/stephen-mcelhose/testable-spec-pattern/internal/modeling/types"

// ModelDomain takes an idea and produces CodeContracts —
// Go interfaces and strategy function signatures at layer boundaries.
// The output is a partial code artifact: structure and intent without implementation.
// Leaf: backed by the model-as-code skill.
func ModelDomain(idea types.Idea) (types.CodeContracts, error) {
	return nil, types.ErrUnimplemented
}

func DefaultBuildContractStrategy(existing types.Contracts, requirements types.Requirements) (types.Contracts, error) {
	return BuildContracts(requirements)
}

// BuildContracts takes requirements and produces the full Contracts —
// code contracts derived from the domain model and wire contracts from resource design.
func BuildContracts(requirements types.Requirements, strategies ...types.BuildContractStrategy) (types.Contracts, error) {
	if strategies == nil {
		strategies = []types.BuildContractStrategy{DefaultBuildContractStrategy}
	}
	var contracts types.Contracts
	var err error
	for _, strategy := range strategies {
		contracts, err = strategy(contracts, requirements)
		if err != nil {
			return types.Contracts{}, err
		}
	}
	return contracts, nil
}

// BuildWireContracts takes requirements and produces ProtoContracts —
// proto definitions, OpenAPI specs, and ROD resource definitions.
// It composes resource design, service definition, and proto file generation.
func BuildWireContracts(requirements types.Requirements) (types.ProtoContracts, error) {
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
