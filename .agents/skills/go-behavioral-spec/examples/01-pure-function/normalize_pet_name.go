package pet

import (
	"fmt"
	"regexp"
	"strings"
)

// KnownError is a small teaching stand-in for typed API errors (not gRPC status).
type KnownError interface {
	error
	Code() string
}

// InvalidArgumentError is returned when a resource name fails format rules.
type InvalidArgumentError struct {
	Message string
}

func (e *InvalidArgumentError) Error() string { return e.Message }
func (e *InvalidArgumentError) Code() string  { return "INVALID_ARGUMENT" }

// NotFoundError is returned when a pet resource does not exist.
type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string { return e.Message }
func (e *NotFoundError) Code() string  { return "NOT_FOUND" }

// InternalError is returned when a dependency fails for a non-domain reason.
type InternalError struct {
	Message string
}

func (e *InternalError) Error() string { return e.Message }
func (e *InternalError) Code() string  { return "INTERNAL" }

// AlreadyExistsError is returned when creating a resource that already exists.
type AlreadyExistsError struct {
	Message string
}

func (e *AlreadyExistsError) Error() string { return e.Message }
func (e *AlreadyExistsError) Code() string  { return "ALREADY_EXISTS" }

// AbortedError is returned when an etag does not match (AIP-135 / AIP-154 freshness).
type AbortedError struct {
	Message string
}

func (e *AbortedError) Error() string { return e.Message }
func (e *AbortedError) Code() string  { return "ABORTED" }

// FailedPreconditionError is returned when a delete/update cannot proceed because
// child resources exist and force was not set (AIP-135 cascading delete). Not
// exercised in the 03 example.
type FailedPreconditionError struct {
	Message string
}

func (e *FailedPreconditionError) Error() string { return e.Message }
func (e *FailedPreconditionError) Code() string  { return "FAILED_PRECONDITION" }

func NewInvalidArgumentError(msg string) error {
	return &InvalidArgumentError{Message: msg}
}

// petIDRE matches {pet}: [a-z0-9-]{1,63}, no leading/trailing hyphen.
var petIDRE = regexp.MustCompile(`^[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`)

// NormalizePetName validates and returns a canonical pet resource name.
//
// Input and output MUST be of the form pets/{pet}. Surrounding whitespace on
// the full name is trimmed. Illegal ids are rejected (not rewritten).
func NormalizePetName(name string) (string, error) {
	name = strings.TrimSpace(name)

	collection, id, ok := strings.Cut(name, "/")
	if !ok || collection != "pets" || id == "" || strings.Contains(id, "/") {
		return "", NewInvalidArgumentError("pet name must be pets/{pet}")
	}
	if !petIDRE.MatchString(id) {
		return "", NewInvalidArgumentError(fmt.Sprintf("pet id %q must be [a-z0-9-]{1,63} with no leading or trailing hyphen", id))
	}
	return "pets/" + id, nil
}
