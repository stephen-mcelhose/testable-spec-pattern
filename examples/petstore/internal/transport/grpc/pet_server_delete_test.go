package grpc_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	petstorev1 "github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore/internal/gen/petstore/v1"
	"github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore/internal/pet"
	"github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore/internal/service"
	petgrpc "github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore/internal/transport/grpc"
)

func TestPetServer_DeletePet(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string

		req   *petstorev1.DeletePetRequest
		store stubStore

		wantCode       string
		wantDeleteArgs []deleteCall
	}

	testCases := []testCase{
		{
			name: "MUST return InvalidArgument when the pet name fails business rules",
			req: &petstorev1.DeletePetRequest{
				Name: "fido",
				Etag: "v1",
			},
			wantCode:       "INVALID_ARGUMENT",
			wantDeleteArgs: nil,
		},
		{
			name: "MUST delete the pet when it exists",
			req: &petstorev1.DeletePetRequest{
				Name: "  pets/fido  ",
				Etag: "v1",
			},
			wantDeleteArgs: []deleteCall{
				{name: "pets/fido", etag: "v1"},
			},
		},
		{
			name: "MUST delete the pet when etag is omitted",
			req: &petstorev1.DeletePetRequest{
				Name: "pets/fido",
			},
			wantDeleteArgs: []deleteCall{
				{name: "pets/fido", etag: ""},
			},
		},
		{
			name: "MUST return NotFound when the pet does not exist",
			req: &petstorev1.DeletePetRequest{
				Name: "pets/missing",
				Etag: "v1",
			},
			store: stubStore{
				deleteErr: &pet.NotFoundError{Message: "pets/missing not found"},
			},
			wantCode: "NOT_FOUND",
			wantDeleteArgs: []deleteCall{
				{name: "pets/missing", etag: "v1"},
			},
		},
		{
			name: "MUST return Aborted when a provided etag does not match",
			req: &petstorev1.DeletePetRequest{
				Name: "pets/fido",
				Etag: "stale",
			},
			store: stubStore{
				deleteErr: &pet.AbortedError{Message: "etag mismatch"},
			},
			wantCode: "ABORTED",
			wantDeleteArgs: []deleteCall{
				{name: "pets/fido", etag: "stale"},
			},
		},
		{
			name: "MUST return Internal when delete fails for a non-domain reason",
			req: &petstorev1.DeletePetRequest{
				Name: "pets/fido",
				Etag: "v1",
			},
			store: stubStore{
				deleteErr: errors.New("connection refused"),
			},
			wantCode: "INTERNAL",
			wantDeleteArgs: []deleteCall{
				{name: "pets/fido", etag: "v1"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			store := tc.store
			srv := petgrpc.PetServer{Pets: &service.PetService{Store: &store}}

			_, err := srv.DeletePet(context.Background(), tc.req)

			assert.Equal(t, tc.wantDeleteArgs, store.deleteCalls, "Store.Delete args")
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
