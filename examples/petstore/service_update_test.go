package petstore_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	petstore "github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore"
)

// TestPetService_UpdatePet covers UpdatePet against AIP-134 behaviours we claim:
// Pet.Name identifies the resource; return the stored Pet; NOT_FOUND if absent;
// optional Pet.Etag → ABORTED on mismatch (AIP-154). Out of scope: update_mask,
// allow_missing upsert, LRO, PERMISSION_DENIED, PUT full-replace semantics.
// Order: shape validation → update (etag optional) → NotFound → Aborted → Internal.
func TestPetService_UpdatePet(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string

		req   petstore.UpdatePetRequest
		store stubStore

		want           petstore.Pet
		wantCode       string
		wantUpdateArgs []updateCall
	}

	validReq := petstore.UpdatePetRequest{
		Pet: petstore.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
	}

	testCases := []testCase{
		{
			name: "MUST return InvalidArgument when UpdatePetRequest fails shape validation",
			req: petstore.UpdatePetRequest{
				Pet: petstore.Pet{Name: "fido", DisplayName: "Fido"},
			},
			wantCode:       "INVALID_ARGUMENT",
			wantUpdateArgs: nil,
		},
		{
			name: "MUST update the pet and return the stored Pet",
			req:  validReq,
			store: stubStore{
				updatePet: petstore.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v2"},
			},
			want: petstore.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v2"},
			wantUpdateArgs: []updateCall{
				{pet: petstore.Pet{Name: "pets/fido", DisplayName: "Fido"}, etag: "v1"},
			},
		},
		{
			name: "MUST update the pet when etag is omitted",
			req: petstore.UpdatePetRequest{
				Pet: petstore.Pet{Name: "pets/fido", DisplayName: "Fido"},
			},
			store: stubStore{
				updatePet: petstore.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v2"},
			},
			want: petstore.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v2"},
			wantUpdateArgs: []updateCall{
				{pet: petstore.Pet{Name: "pets/fido", DisplayName: "Fido"}, etag: ""},
			},
		},
		{
			// AIP-134: missing resource → NOT_FOUND (unless allow_missing, which we omit).
			name: "MUST return NotFound when the pet does not exist",
			req:  validReq,
			store: stubStore{
				updateErr: &petstore.NotFoundError{Message: "pets/fido not found"},
			},
			wantCode: "NOT_FOUND",
			wantUpdateArgs: []updateCall{
				{pet: petstore.Pet{Name: "pets/fido", DisplayName: "Fido"}, etag: "v1"},
			},
		},
		{
			// AIP-134 / AIP-154: provided etag that does not match → ABORTED.
			name: "MUST return Aborted when a provided etag does not match",
			req:  validReq,
			store: stubStore{
				updateErr: &petstore.AbortedError{Message: "etag mismatch"},
			},
			wantCode: "ABORTED",
			wantUpdateArgs: []updateCall{
				{pet: petstore.Pet{Name: "pets/fido", DisplayName: "Fido"}, etag: "v1"},
			},
		},
		{
			name: "MUST return Internal when update fails for a non-domain reason",
			req:  validReq,
			store: stubStore{
				updateErr: errors.New("connection refused"),
			},
			wantCode: "INTERNAL",
			wantUpdateArgs: []updateCall{
				{pet: petstore.Pet{Name: "pets/fido", DisplayName: "Fido"}, etag: "v1"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			store := tc.store
			srv := petstore.PetService{Store: &store}

			got, err := srv.UpdatePet(context.Background(), tc.req)

			assert.Equal(t, tc.wantUpdateArgs, store.updateCalls, "Store.Update args")
			if tc.wantCode != "" {
				require.Error(t, err)
				var known petstore.KnownError
				require.True(t, errors.As(err, &known), "error MUST be KnownError, got %T: %v", err, err)
				assert.Equal(t, tc.wantCode, known.Code())
				assert.Equal(t, petstore.Pet{}, got)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
