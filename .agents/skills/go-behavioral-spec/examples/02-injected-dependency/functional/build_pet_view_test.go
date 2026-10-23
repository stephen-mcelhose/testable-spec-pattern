package functional_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pet "example.com/go-behavioral-spec/examples/01-pure-function"
	"example.com/go-behavioral-spec/examples/02-injected-dependency/functional"
)

// stubPetGetter records Get calls so specs can assert the injected leaf boundary.
// Prefer counterfeiter when the interface is package-owned and you want generate-fakes.
type stubPetGetter struct {
	pet   functional.Pet
	err   error
	calls []string
}

func (s *stubPetGetter) Get(_ context.Context, name string) (functional.Pet, error) {
	s.calls = append(s.calls, name)
	return s.pet, s.err
}

// TestBuildPetView covers BuildPetView(ctx, name, getter) — functional injection.
// Order: Get with canonical name + view from result → invalid name skips Get →
// leaf NotFound → leaf other error mapped to Internal.
func TestBuildPetView(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string

		// inputs
		petName  string
		examples []string
		getter   stubPetGetter

		// expected outputs
		want         functional.PetView
		wantCode     string
		wantGetCalls []string
	}

	testCases := []testCase{
		{
			name:    "MUST call Get with the canonical pet name and return a PetView from the Get result",
			petName: "  pets/fido  ",
			getter: stubPetGetter{
				pet: functional.Pet{Name: "pets/fido", DisplayName: "FromGetter"},
			},
			want:         functional.PetView{Name: "pets/fido", DisplayName: "FromGetter"},
			wantGetCalls: []string{"pets/fido"},
		},
		{
			name: "MUST return InvalidArgument without calling Get when the pet name is invalid",
			examples: []string{
				"fido",
				"pets/",
				"pets/Fido",
			},
			getter:       stubPetGetter{pet: functional.Pet{DisplayName: "must-not-appear"}},
			wantCode:     "INVALID_ARGUMENT",
			wantGetCalls: nil,
		},
		{
			name:    "MUST return NotFound when Get reports not found",
			petName: "pets/missing",
			getter: stubPetGetter{
				err: &pet.NotFoundError{Message: "pets/missing not found"},
			},
			wantCode:     "NOT_FOUND",
			wantGetCalls: []string{"pets/missing"},
		},
		{
			name:    "MUST return Internal when Get fails for a non-not-found reason",
			petName: "pets/fido",
			getter: stubPetGetter{
				err: errors.New("connection refused"),
			},
			wantCode:     "INTERNAL",
			wantGetCalls: []string{"pets/fido"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			names := tc.examples
			if len(names) == 0 {
				names = []string{tc.petName}
			}
			for _, name := range names {
				t.Run(name, func(t *testing.T) {
					t.Parallel()
					g := tc.getter

					got, err := functional.BuildPetView(context.Background(), name, &g)

					assert.Equal(t, tc.wantGetCalls, g.calls, "Get call args (injected leaf)")

					if tc.wantCode != "" {
						require.Error(t, err)
						var known pet.KnownError
						require.True(t, errors.As(err, &known), "error MUST be KnownError, got %T: %v", err, err)
						assert.Equal(t, tc.wantCode, known.Code())
						assert.Equal(t, functional.PetView{}, got)
						return
					}
					require.NoError(t, err)
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}
