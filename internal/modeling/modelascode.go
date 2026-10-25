package modeling

import "github.com/stephen-mcelhose/testable-spec-pattern/internal/modeling/types"

// ModelAsCode takes an idea and produces a CodeArtifact —
// a Go package modeling the idea using types, interfaces, and strategy functions.
// Leaf: backed by the model-as-code skill.
func ModelAsCode(idea types.Idea) (types.CodeArtifact, error) {
	return nil, types.ErrUnimplemented
}
