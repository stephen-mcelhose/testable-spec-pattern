package petstore

import (
	"fmt"
	"regexp"
	"strings"
)

// petIDRE matches a valid pet id: [a-z0-9-]{1,63}, no leading/trailing hyphen.
var petIDRE = regexp.MustCompile(`^[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`)

// NormalizePetName validates and returns the canonical resource name for a pet.
//
// Input and output MUST be of the form pets/{pet}. Surrounding whitespace on
// the full name is trimmed. Illegal ids are rejected with InvalidArgumentError.
func NormalizePetName(name string) (string, error) {
	name = strings.TrimSpace(name)

	collection, id, ok := strings.Cut(name, "/")
	if !ok || collection != "pets" || id == "" || strings.Contains(id, "/") {
		return "", &InvalidArgumentError{Message: "pet name must be pets/{pet}"}
	}
	if !petIDRE.MatchString(id) {
		return "", &InvalidArgumentError{
			Message: fmt.Sprintf("pet id %q must be [a-z0-9-]{1,63} with no leading or trailing hyphen", id),
		}
	}
	return "pets/" + id, nil
}
