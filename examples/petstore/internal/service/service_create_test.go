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

// TestPetService_CreatePet covers application concerns: business rules → Store.
// Proto shape lives in transport/grpc tests.
func TestPetService_CreatePet(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string

		petID       string
		displayName string
		examples    []struct{ petID, displayName string }
		store       stubStore

		want           pet.Pet
		wantCode       string
		wantCreateArgs []pet.Pet
	}

	testCases := []testCase{
		{
			name: "MUST return InvalidArgument when CreatePet fails business rules",
			examples: []struct{ petID, displayName string }{
				{petID: "", displayName: "Fido"},
				{petID: "Fido", displayName: "Fido"},
				{petID: "fido", displayName: ""},
				{petID: "fido", displayName: "   "},
			},
			wantCode:       "INVALID_ARGUMENT",
			wantCreateArgs: nil,
		},
		{
			name:        "MUST create the pet from pet_id and return the stored Pet",
			petID:       "  fido  ",
			displayName: "  Fido  ",
			store: stubStore{
				createPet: pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
			},
			want: pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
			wantCreateArgs: []pet.Pet{
				{Name: "pets/fido", DisplayName: "Fido"},
			},
		},
		{
			name:        "MUST return AlreadyExists when the pet already exists",
			petID:       "fido",
			displayName: "Fido",
			store: stubStore{
				createErr: &pet.AlreadyExistsError{Message: "pets/fido already exists"},
			},
			wantCode: "ALREADY_EXISTS",
			wantCreateArgs: []pet.Pet{
				{Name: "pets/fido", DisplayName: "Fido"},
			},
		},
		{
			name:        "MUST return Internal when create fails for a non-domain reason",
			petID:       "fido",
			displayName: "Fido",
			store: stubStore{
				createErr: errors.New("connection refused"),
			},
			wantCode: "INTERNAL",
			wantCreateArgs: []pet.Pet{
				{Name: "pets/fido", DisplayName: "Fido"},
			},
		},
	}

	type pair struct{ petID, displayName string }

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			inputs := make([]pair, 0, len(tc.examples)+1)
			for _, ex := range tc.examples {
				inputs = append(inputs, pair{ex.petID, ex.displayName})
			}
			if len(inputs) == 0 {
				inputs = []pair{{tc.petID, tc.displayName}}
			}
			for _, in := range inputs {
				t.Run(in.petID+"/"+in.displayName, func(t *testing.T) {
					t.Parallel()
					store := tc.store
					srv := service.PetService{Store: &store}

					got, err := srv.CreatePet(context.Background(), in.petID, in.displayName)

					assert.Equal(t, tc.wantCreateArgs, store.createCalls, "Store.Create args")
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
