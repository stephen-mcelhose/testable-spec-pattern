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

func TestPetServer_GetPet(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string

		req  *petstorev1.GetPetRequest
		pets stubPets

		want        *petstorev1.Pet
		wantCode    string
		wantGetArgs []string
	}

	testCases := []testCase{
		{
			name:        "MUST return InvalidArgument when GetPetRequest fails shape validation",
			req:         &petstorev1.GetPetRequest{Name: ""},
			wantCode:    "INVALID_ARGUMENT",
			wantGetArgs: nil,
		},
		{
			name: "MUST map name to PetAPI and map the domain Pet to proto",
			req:  &petstorev1.GetPetRequest{Name: "pets/fido"},
			pets: stubPets{
				getPet: pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
			},
			want:        &petstorev1.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
			wantGetArgs: []string{"pets/fido"},
		},
		{
			name: "MUST forward PetAPI errors",
			req:  &petstorev1.GetPetRequest{Name: "pets/missing"},
			pets: stubPets{
				getErr: &pet.NotFoundError{Message: "pets/missing not found"},
			},
			wantCode:    "NOT_FOUND",
			wantGetArgs: []string{"pets/missing"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			pets := tc.pets
			srv := petgrpc.PetServer{Pets: &pets}

			got, err := srv.GetPet(context.Background(), tc.req)

			assert.Equal(t, tc.wantGetArgs, pets.getCalls, "PetAPI.GetPet args")
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
