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

func TestPetServer_DeletePet(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string

		req  *petstorev1.DeletePetRequest
		pets stubPets

		wantCode       string
		wantDeleteArgs []deleteCall
	}

	testCases := []testCase{
		{
			name: "MUST return InvalidArgument when DeletePetRequest fails shape validation",
			req: &petstorev1.DeletePetRequest{
				Name: "",
				Etag: "v1",
			},
			wantCode:       "INVALID_ARGUMENT",
			wantDeleteArgs: nil,
		},
		{
			name: "MUST map name and etag to PetAPI",
			req: &petstorev1.DeletePetRequest{
				Name: "pets/fido",
				Etag: "v1",
			},
			wantDeleteArgs: []deleteCall{
				{name: "pets/fido", etag: "v1"},
			},
		},
		{
			name: "MUST forward PetAPI errors",
			req: &petstorev1.DeletePetRequest{
				Name: "pets/fido",
				Etag: "stale",
			},
			pets: stubPets{
				deleteErr: &pet.AbortedError{Message: "etag mismatch"},
			},
			wantCode: "ABORTED",
			wantDeleteArgs: []deleteCall{
				{name: "pets/fido", etag: "stale"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			pets := tc.pets
			srv := petgrpc.PetServer{Pets: &pets}

			_, err := srv.DeletePet(context.Background(), tc.req)

			assert.Equal(t, tc.wantDeleteArgs, pets.deleteCalls, "PetAPI.DeletePet args")
			if tc.wantCode != "" {
				require.Error(t, err)
				var known pet.KnownError
				require.True(t, errors.As(err, &known), "error MUST be KnownError, got %T: %v", err, err)
				assert.Equal(t, tc.wantCode, known.Code())
				return
			}
			require.NoError(t, err)
		})
	}
}
