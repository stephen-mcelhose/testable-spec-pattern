package pet

// KnownError is a typed API error with a machine-readable code.
// It is a teaching stand-in for gRPC status errors — not production gRPC.
type KnownError interface {
	error
	Code() string
}

// InvalidArgumentError is returned when a request fails shape or format validation.
type InvalidArgumentError struct {
	Message string
}

func (e *InvalidArgumentError) Error() string { return e.Message }
func (e *InvalidArgumentError) Code() string  { return "INVALID_ARGUMENT" }

// NotFoundError is returned when a resource does not exist.
type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string { return e.Message }
func (e *NotFoundError) Code() string  { return "NOT_FOUND" }

// AlreadyExistsError is returned when creating a resource that already exists.
type AlreadyExistsError struct {
	Message string
}

func (e *AlreadyExistsError) Error() string { return e.Message }
func (e *AlreadyExistsError) Code() string  { return "ALREADY_EXISTS" }

// AbortedError is returned when an etag does not match (AIP-154 freshness).
type AbortedError struct {
	Message string
}

func (e *AbortedError) Error() string { return e.Message }
func (e *AbortedError) Code() string  { return "ABORTED" }

// InternalError is returned when a dependency fails for a non-domain reason.
type InternalError struct {
	Message string
}

func (e *InternalError) Error() string { return e.Message }
func (e *InternalError) Code() string  { return "INTERNAL" }
