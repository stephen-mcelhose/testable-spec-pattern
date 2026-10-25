package cicdpipeline

import "errors"

// Undefined is the explicit placeholder for anything not yet modeled.
type Undefined = any

// ErrUnimplemented is returned by every leaf. Never panic in a model.
var ErrUnimplemented = errors.New("unimplemented")
