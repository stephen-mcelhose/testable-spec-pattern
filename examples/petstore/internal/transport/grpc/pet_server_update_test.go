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

func TestPetServer_UpdatePet(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string

		req  *petstorev1.UpdatePetRequest
		pets stubPets

		want           *petstorev1.Pet
		wantCode       string
		wantUpdateArgs []updateCall
	}

	testCases := []testCase{
		{
			name:           "MUST return InvalidArgument when UpdatePetRequest fails shape validation",
			req:            &petstorev1.UpdatePetRequest{Pet: nil},
			wantCode:       "INVALID_ARGUMENT",
			wantUpdateArgs: nil,
		},
		{
			name: "MUST map pet fields to PetAPI and map the domain Pet to proto",
			req: &petstorev1.UpdatePetRequest{
				Pet: &petstorev1.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
			},
			pets: stubPets{
				updatePet: pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v2"},
			},
			want: &petstorev1.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v2"},
			wantUpdateArgs: []updateCall{
				{name: "pets/fido", displayName: "Fido", etag: "v1"},
			},
		},
		{
			name: "MUST forward PetAPI errors",
			req: &petstorev1.UpdatePetRequest{
				Pet: &petstorev1.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
			},
			pets: stubPets{
				updateErr: &pet.AbortedError{Message: "etag mismatch"},
			},
			wantCode: "ABORTED",
			wantUpdateArgs: []updateCall{
				{name: "pets/fido", displayName: "Fido", etag: "v1"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			pets := tc.pets
			srv := petgrpc.PetServer{Pets: &pets}

			got, err := srv.UpdatePet(context.Background(), tc.req)

			assert.Equal(t, tc.wantUpdateArgs, pets.updateCalls, "PetAPI.UpdatePet args")
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
