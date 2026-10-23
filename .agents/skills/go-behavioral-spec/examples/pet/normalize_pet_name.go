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

func invalidArgument(msg string) error {
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
		return "", invalidArgument("pet name must be pets/{pet}")
	}
	if !petIDRE.MatchString(id) {
		return "", invalidArgument(fmt.Sprintf("pet id %q must be [a-z0-9-]{1,63} with no leading or trailing hyphen", id))
	}
	return "pets/" + id, nil
}
