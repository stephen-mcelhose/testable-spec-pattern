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

// TestPetService_CreatePet covers CreatePet against AIP-133 behaviours we claim:
// user-specified pet_id; ignore Pet.Name on the body; return the stored Pet;
// ALREADY_EXISTS on duplicate id. Out of scope: parent, LRO, PERMISSION_DENIED,
// system-generated ids when pet_id omitted.
// Order: shape validation → create (ignore body name) → AlreadyExists → Internal.
func TestPetService_CreatePet(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string

		req      pet.CreatePetRequest
		examples []pet.CreatePetRequest
		store    stubStore

		want           pet.Pet
		wantCode       string
		wantCreateArgs []pet.Pet
	}

	testCases := []testCase{
		{
			name: "MUST return InvalidArgument when CreatePetRequest fails shape validation",
			examples: []pet.CreatePetRequest{
				{PetID: "", Pet: pet.Pet{DisplayName: "Fido"}},
				{PetID: "Fido", Pet: pet.Pet{DisplayName: "Fido"}},
				{PetID: "fido", Pet: pet.Pet{DisplayName: ""}},
				{PetID: "fido", Pet: pet.Pet{DisplayName: "   "}},
			},
			wantCode:       "INVALID_ARGUMENT",
			wantCreateArgs: nil,
		},
		{
			name: "MUST create the pet from pet_id and return the stored Pet",
			req: pet.CreatePetRequest{
				PetID: "  fido  ",
				Pet:   pet.Pet{DisplayName: "  Fido  "},
			},
			store: stubStore{
				createPet: pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
			},
			want: pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
			wantCreateArgs: []pet.Pet{
				{Name: "pets/fido", DisplayName: "Fido"},
			},
		},
		{
			// AIP-133: name on the resource body must be ignored; pet_id wins.
			name: "MUST ignore Pet.Name on the create body and use pet_id for the resource name",
			req: pet.CreatePetRequest{
				PetID: "fido",
				Pet:   pet.Pet{Name: "pets/other", DisplayName: "Fido"},
			},
			store: stubStore{
				createPet: pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
			},
			want: pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
			wantCreateArgs: []pet.Pet{
				{Name: "pets/fido", DisplayName: "Fido"},
			},
		},
		{
			// AIP-133: duplicate user-specified id → ALREADY_EXISTS.
			name: "MUST return AlreadyExists when the pet already exists",
			req: pet.CreatePetRequest{
				PetID: "fido",
				Pet:   pet.Pet{DisplayName: "Fido"},
			},
			store: stubStore{
				createErr: &pet.AlreadyExistsError{Message: "pets/fido already exists"},
			},
			wantCode: "ALREADY_EXISTS",
			wantCreateArgs: []pet.Pet{
				{Name: "pets/fido", DisplayName: "Fido"},
			},
		},
		{
			name: "MUST return Internal when create fails for a non-domain reason",
			req: pet.CreatePetRequest{
				PetID: "fido",
				Pet:   pet.Pet{DisplayName: "Fido"},
			},
			store: stubStore{
				createErr: errors.New("connection refused"),
			},
			wantCode: "INTERNAL",
			wantCreateArgs: []pet.Pet{
				{Name: "pets/fido", DisplayName: "Fido"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			reqs := tc.examples
			if len(reqs) == 0 {
				reqs = []pet.CreatePetRequest{tc.req}
			}
			for _, req := range reqs {
				t.Run(req.PetID+"/"+req.Pet.DisplayName, func(t *testing.T) {
					t.Parallel()
					store := tc.store
					srv := service.PetService{Store: &store}

					got, err := srv.CreatePet(context.Background(), req)

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
