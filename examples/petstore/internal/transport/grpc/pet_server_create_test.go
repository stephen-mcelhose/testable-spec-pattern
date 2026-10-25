package grpc_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	petstorev1 "github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore/internal/gen/petstore/v1"
	"github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore/internal/pet"
	petgrpc "github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore/internal/transport/grpc"
)

// TestPetServer_CreatePet covers transport concerns only: shape + field mapping to PetAPI.
// Business rules and Store behaviour live in service_test.
func TestPetServer_CreatePet(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string

		req  *petstorev1.CreatePetRequest
		pets stubPets

		want           *petstorev1.Pet
		wantCode       string
		wantCreateArgs []createCall
	}

	testCases := []testCase{
		{
			name:           "MUST return InvalidArgument when CreatePetRequest fails shape validation",
			req:            &petstorev1.CreatePetRequest{PetId: "fido", Pet: nil},
			wantCode:       "INVALID_ARGUMENT",
			wantCreateArgs: nil,
		},
		{
			name: "MUST map pet_id and display_name to PetAPI and map the domain Pet to proto",
			req: &petstorev1.CreatePetRequest{
				PetId: "fido",
				Pet:   &petstorev1.Pet{Name: "pets/other", DisplayName: "Fido"},
			},
			pets: stubPets{
				createPet: pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
			},
			want: &petstorev1.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
			wantCreateArgs: []createCall{
				{petID: "fido", displayName: "Fido"},
			},
		},
		{
			name: "MUST forward PetAPI errors without calling further mapping",
			req: &petstorev1.CreatePetRequest{
				PetId: "fido",
				Pet:   &petstorev1.Pet{DisplayName: "Fido"},
			},
			pets: stubPets{
				createErr: &pet.AlreadyExistsError{Message: "pets/fido already exists"},
			},
			wantCode: "ALREADY_EXISTS",
			wantCreateArgs: []createCall{
				{petID: "fido", displayName: "Fido"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			pets := tc.pets
			srv := petgrpc.PetServer{Pets: &pets}

			got, err := srv.CreatePet(context.Background(), tc.req)

			assert.Equal(t, tc.wantCreateArgs, pets.createCalls, "PetAPI.CreatePet args")
			if tc.wantCode != "" {
				require.Error(t, err)
				var known pet.KnownError
				require.True(t, errors.As(err, &known), "error MUST be KnownError, got %T: %v", err, err)
				assert.Equal(t, tc.wantCode, known.Code())
				assert.Nil(t, got)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
