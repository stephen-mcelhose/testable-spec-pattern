package modeling

import "github.com/stephen-mcelhose/testable-spec-pattern/internal/modeling/types"

// WriteSpecification produces a prose behavioral specification —
// RFC 2119-style descriptions of how a system must behave,
// grounded in domain names from the contracts.
// Leaf: backed by the rfc-2119-spec-writer skill.
func WriteSpecification(requirements types.Requirements, contracts types.Contracts) (types.Specification, error) {
	return nil, types.ErrUnimplemented
}
