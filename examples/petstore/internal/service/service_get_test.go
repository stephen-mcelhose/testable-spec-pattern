package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore/internal/pet"
	"github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore/internal/service"
)

// TestPetService_GetPet covers GetPet.
// Order: shape validation → return stored Pet → NotFound → Internal.
func TestPetService_GetPet(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string

		petName  string
		examples []string
		store    stubStore

		want        pet.Pet
		wantCode    string
		wantGetArgs []string
	}

	testCases := []testCase{
		{
			name: "MUST return InvalidArgument when the pet name fails shape validation",
			examples: []string{
				"fido",
				"pets/",
				"pets/Fido",
			},
			wantCode:    "INVALID_ARGUMENT",
			wantGetArgs: nil,
		},
		{
			name:    "MUST return the stored Pet for a valid pet name",
			petName: "  pets/fido  ",
			store: stubStore{
				getPet: pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
			},
			want:        pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
			wantGetArgs: []string{"pets/fido"},
		},
		{
			name:    "MUST return NotFound when the pet does not exist",
			petName: "pets/missing",
			store: stubStore{
				getErr: &pet.NotFoundError{Message: "pets/missing not found"},
			},
			wantCode:    "NOT_FOUND",
			wantGetArgs: []string{"pets/missing"},
		},
		{
			name:    "MUST return Internal when get fails for a non-not-found reason",
			petName: "pets/fido",
			store: stubStore{
				getErr: errors.New("connection refused"),
			},
			wantCode:    "INTERNAL",
			wantGetArgs: []string{"pets/fido"},
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
					store := tc.store
					srv := service.PetService{Store: &store}

					got, err := srv.GetPet(context.Background(), name)

					assert.Equal(t, tc.wantGetArgs, store.getCalls, "Store.Get args")
					if tc.wantCode != "" {
						require.Error(t, err)
						var known pet.KnownError
						require.True(t, errors.As(err, &known), "error MUST be KnownError, got %T: %v", err, err)
						assert.Equal(t, tc.wantCode, known.Code())
						assert.Equal(t, pet.Pet{}, got)
						return
					}
					require.NoError(t, err)
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}
