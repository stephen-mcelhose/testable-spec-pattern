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

func TestPetService_UpdatePet(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string

		petName     string
		displayName string
		etag        string
		store       stubStore

		want           pet.Pet
		wantCode       string
		wantUpdateArgs []updateCall
	}

	testCases := []testCase{
		{
			name:           "MUST return InvalidArgument when UpdatePet fails business rules",
			petName:        "fido",
			displayName:    "Fido",
			etag:           "v1",
			wantCode:       "INVALID_ARGUMENT",
			wantUpdateArgs: nil,
		},
		{
			name:        "MUST update the pet and return the stored Pet",
			petName:     "pets/fido",
			displayName: "Fido",
			etag:        "v1",
			store: stubStore{
				updatePet: pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v2"},
			},
			want: pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v2"},
			wantUpdateArgs: []updateCall{
				{pet: pet.Pet{Name: "pets/fido", DisplayName: "Fido"}, etag: "v1"},
			},
		},
		{
			name:        "MUST update the pet when etag is omitted",
			petName:     "pets/fido",
			displayName: "Fido",
			store: stubStore{
				updatePet: pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v2"},
			},
			want: pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v2"},
			wantUpdateArgs: []updateCall{
				{pet: pet.Pet{Name: "pets/fido", DisplayName: "Fido"}, etag: ""},
			},
		},
		{
			name:        "MUST return NotFound when the pet does not exist",
			petName:     "pets/fido",
			displayName: "Fido",
			etag:        "v1",
			store: stubStore{
				updateErr: &pet.NotFoundError{Message: "pets/fido not found"},
			},
			wantCode: "NOT_FOUND",
			wantUpdateArgs: []updateCall{
				{pet: pet.Pet{Name: "pets/fido", DisplayName: "Fido"}, etag: "v1"},
			},
		},
		{
			name:        "MUST return Aborted when a provided etag does not match",
			petName:     "pets/fido",
			displayName: "Fido",
			etag:        "v1",
			store: stubStore{
				updateErr: &pet.AbortedError{Message: "etag mismatch"},
			},
			wantCode: "ABORTED",
			wantUpdateArgs: []updateCall{
				{pet: pet.Pet{Name: "pets/fido", DisplayName: "Fido"}, etag: "v1"},
			},
		},
		{
			name:        "MUST return Internal when update fails for a non-domain reason",
			petName:     "pets/fido",
			displayName: "Fido",
			etag:        "v1",
			store: stubStore{
				updateErr: errors.New("connection refused"),
			},
			wantCode: "INTERNAL",
			wantUpdateArgs: []updateCall{
				{pet: pet.Pet{Name: "pets/fido", DisplayName: "Fido"}, etag: "v1"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			store := tc.store
			srv := service.PetService{Store: &store}

			got, err := srv.UpdatePet(context.Background(), tc.petName, tc.displayName, tc.etag)

			assert.Equal(t, tc.wantUpdateArgs, store.updateCalls, "Store.Update args")
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
}
