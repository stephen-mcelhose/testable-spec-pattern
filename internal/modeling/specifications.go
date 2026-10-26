package modeling

import "github.com/stephen-mcelhose/testable-spec-pattern/internal/modeling/types"

// DefaultWriteSpecificationStrategy is the default leaf implementation for WriteSpecification.
// Leaf: backed by the rfc-2119-spec-writer skill.
func DefaultWriteSpecificationStrategy(existing types.Specification, requirements types.Requirements, contracts types.Contracts) (types.Specification, error) {
	return nil, types.ErrUnimplemented
}

// WriteSpecification produces a prose behavioral specification —
// RFC 2119-style descriptions of how a system must behave,
// grounded in domain names from the contracts.
func WriteSpecification(requirements types.Requirements, contracts types.Contracts, strategies ...types.WriteSpecificationStrategy) (types.Specification, error) {
	if strategies == nil {
		strategies = []types.WriteSpecificationStrategy{DefaultWriteSpecificationStrategy}
	}
	var spec types.Specification
	var err error
	for _, strategy := range strategies {
		spec, err = strategy(spec, requirements, contracts)
		if err != nil {
			return nil, err
		}
	}
	return spec, nil
}
