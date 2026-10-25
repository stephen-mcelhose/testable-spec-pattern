package modeling

import "github.com/stephen-mcelhose/testable-spec-pattern/internal/modeling/types"

// DefaultModelDomainStrategy is the default leaf implementation for ModelDomain.
// Leaf: backed by the model-as-code skill.
func DefaultModelDomainStrategy(existing types.CodeContracts, idea types.Idea) (types.CodeContracts, error) {
	return nil, types.ErrUnimplemented
}

// ModelDomain takes an idea and produces CodeContracts —
// Go interfaces and strategy function signatures at layer boundaries.
// The output is a partial code artifact: structure and intent without implementation.
func ModelDomain(idea types.Idea, strategies ...types.ModelDomainStrategy) (types.CodeContracts, error) {
	if strategies == nil {
		strategies = []types.ModelDomainStrategy{DefaultModelDomainStrategy}
	}
	var code types.CodeContracts
	var err error
	for _, strategy := range strategies {
		code, err = strategy(code, idea)
		if err != nil {
			return nil, err
		}
	}
	return code, nil
}

// DefaultBuildContractStrategy composes BuildWireContracts, preserving any
// existing CodeContracts (e.g. from a prior ModelDomain call).
func DefaultBuildContractStrategy(existing types.Contracts, requirements types.Requirements) (types.Contracts, error) {
	wire, err := BuildWireContracts(requirements)
	if err != nil {
		return existing, err
	}
	return types.Contracts{Code: existing.Code, Wire: wire}, nil
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

// DefaultBuildWireContractsStrategy composes DesignResourceModel, DefineServices,
// and GenerateProtoFiles to produce wire contracts.
func DefaultBuildWireContractsStrategy(existing types.ProtoContracts, requirements types.Requirements) (types.ProtoContracts, error) {
	resources, err := DesignResourceModel(requirements)
	if err != nil {
		return existing, err
	}
	services, err := DefineServices(resources)
	if err != nil {
		return existing, err
	}
	files, err := GenerateProtoFiles(services)
	if err != nil {
		return existing, err
	}
	return types.ProtoContracts{Files: append(existing.Files, files...)}, nil
}

// BuildWireContracts takes requirements and produces ProtoContracts —
// proto definitions, OpenAPI specs, and ROD resource definitions.
func BuildWireContracts(requirements types.Requirements, strategies ...types.BuildWireContractsStrategy) (types.ProtoContracts, error) {
	if strategies == nil {
		strategies = []types.BuildWireContractsStrategy{DefaultBuildWireContractsStrategy}
	}
	var wire types.ProtoContracts
	var err error
	for _, strategy := range strategies {
		wire, err = strategy(wire, requirements)
		if err != nil {
			return types.ProtoContracts{}, err
		}
	}
	return wire, nil
}

// DefaultDesignResourceModelStrategy is the default leaf implementation for DesignResourceModel.
// Leaf: backed by the resource-oriented-design skill.
func DefaultDesignResourceModelStrategy(existing []types.Resource, requirements types.Requirements) ([]types.Resource, error) {
	return nil, types.ErrUnimplemented
}

// DesignResourceModel derives AIP resources from requirements —
// named, revisioned domain entities with standard lifecycle methods.
func DesignResourceModel(requirements types.Requirements, strategies ...types.DesignResourceModelStrategy) ([]types.Resource, error) {
	if strategies == nil {
		strategies = []types.DesignResourceModelStrategy{DefaultDesignResourceModelStrategy}
	}
	var resources []types.Resource
	var err error
	for _, strategy := range strategies {
		resources, err = strategy(resources, requirements)
		if err != nil {
			return nil, err
		}
	}
	return resources, nil
}

// DefaultDefineServicesStrategy is the default leaf implementation for DefineServices.
// Leaf: backed by service definition tooling.
func DefaultDefineServicesStrategy(existing []types.Service, resources []types.Resource) ([]types.Service, error) {
	return nil, types.ErrUnimplemented
}

// DefineServices groups resources into gRPC service definitions,
// each with the standard methods appropriate for its resource type.
func DefineServices(resources []types.Resource, strategies ...types.DefineServicesStrategy) ([]types.Service, error) {
	if strategies == nil {
		strategies = []types.DefineServicesStrategy{DefaultDefineServicesStrategy}
	}
	var services []types.Service
	var err error
	for _, strategy := range strategies {
		services, err = strategy(services, resources)
		if err != nil {
			return nil, err
		}
	}
	return services, nil
}

// DefaultGenerateProtoFilesStrategy is the default leaf implementation for GenerateProtoFiles.
// Leaf: side effect — writes files to disk.
func DefaultGenerateProtoFilesStrategy(existing []types.ProtoFile, services []types.Service) ([]types.ProtoFile, error) {
	return nil, types.ErrUnimplemented
}

// GenerateProtoFiles produces .proto files from service definitions,
// including googleapis annotations for AIP compliance.
func GenerateProtoFiles(services []types.Service, strategies ...types.GenerateProtoFilesStrategy) ([]types.ProtoFile, error) {
	if strategies == nil {
		strategies = []types.GenerateProtoFilesStrategy{DefaultGenerateProtoFilesStrategy}
	}
	var files []types.ProtoFile
	var err error
	for _, strategy := range strategies {
		files, err = strategy(files, services)
		if err != nil {
			return nil, err
		}
	}
	return files, nil
}
