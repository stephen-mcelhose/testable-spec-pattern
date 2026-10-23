package pet_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pet "example.com/go-behavioral-spec/examples/01-pure-function"
)

// TestNormalizePetName covers resource shape: accept/trim prefixed names and reject
// names that are not pets/{pet}.
// Order: valid → whitespace → shape errors.
func TestNormalizePetName(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string

		// inputs — set in for a single value, or examples for several fixtures of one rule
		in       string
		examples []string

		// expected outputs
		want           string
		wantInvalidArg bool
	}

	testCases := []testCase{
		{
			name: "MUST accept a valid pets/{pet} name and return the canonical prefixed form",
			in:   "pets/fido",
			want: "pets/fido",
		},
		{
			name: "MUST trim surrounding whitespace and return a prefixed name",
			in:   "  pets/fido  ",
			want: "pets/fido",
		},
		{
			name: "MUST return InvalidArgument when the name is not pets/{pet}",
			examples: []string{
				"fido",       // no collection
				"pet/fido",   // wrong collection
				"pets/",      // empty id
				"pets/fi/do", // extra segment
			},
			wantInvalidArg: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			runNormalizeCases(t, tc.in, tc.examples, tc.want, tc.wantInvalidArg)
		})
	}
}

// TestNormalizePetName_PetIDFormat covers the {pet} id format invariant.
// Normative cases use MUST; fixtures under examples/ use plain nested names only.
// Order: accept when format matches → reject when format fails.
func TestNormalizePetName_PetIDFormat(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string

		// inputs
		examples []string

		// expected outputs
		wantInvalidArg bool
	}

	testCases := []testCase{
		{
			name: "MUST accept a pets/{pet} name when {pet} matches the id format",
			examples: []string{
				"pets/a",      // single character
				"pets/pet-42", // internal hyphen
			},
		},
		{
			name: "MUST return InvalidArgument when {pet} does not match the id format",
			examples: []string{
				"pets/Fido",  // uppercase
				"pets/-fido", // leading hyphen
				"pets/fido-", // trailing hyphen
			},
			wantInvalidArg: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			for _, in := range tc.examples {
				t.Run(in, func(t *testing.T) {
					t.Parallel()
					got, err := pet.NormalizePetName(in)
					if tc.wantInvalidArg {
						require.Error(t, err)
						var inv *pet.InvalidArgumentError
						require.True(t, errors.As(err, &inv), "error MUST be InvalidArgument, got %T: %v", err, err)
						assert.Equal(t, "INVALID_ARGUMENT", inv.Code())
						assert.Empty(t, got)
						return
					}
					require.NoError(t, err)
					assert.Equal(t, in, got)
					assert.True(t, strings.HasPrefix(got, "pets/"), "result MUST remain prefixed")
				})
			}
		})
	}
}

func runNormalizeCases(t *testing.T, in string, examples []string, want string, wantInvalidArg bool) {
	t.Helper()
	inputs := examples
	if len(inputs) == 0 {
		inputs = []string{in}
	}
	for _, input := range inputs {
		t.Run(input, func(t *testing.T) {
			t.Parallel()
			got, err := pet.NormalizePetName(input)
			if wantInvalidArg {
				require.Error(t, err)
				var inv *pet.InvalidArgumentError
				require.True(t, errors.As(err, &inv), "error MUST be InvalidArgument, got %T: %v", err, err)
				assert.Equal(t, "INVALID_ARGUMENT", inv.Code())
				assert.Empty(t, got)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, want, got)
			assert.True(t, strings.HasPrefix(got, "pets/"), "result MUST remain prefixed")
		})
	}
}
