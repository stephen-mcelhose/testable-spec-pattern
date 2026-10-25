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

func TestPetServer_UpdatePet(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string

		req   *petstorev1.UpdatePetRequest
		store stubStore

		want           *petstorev1.Pet
		wantCode       string
		wantUpdateArgs []updateCall
	}

	validReq := &petstorev1.UpdatePetRequest{
		Pet: &petstorev1.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
	}

	testCases := []testCase{
		{
			name:           "MUST return InvalidArgument when UpdatePetRequest fails shape validation",
			req:            &petstorev1.UpdatePetRequest{Pet: nil},
			wantCode:       "INVALID_ARGUMENT",
			wantUpdateArgs: nil,
		},
		{
			name: "MUST return InvalidArgument when UpdatePetRequest fails business rules",
			req: &petstorev1.UpdatePetRequest{
				Pet: &petstorev1.Pet{Name: "fido", DisplayName: "Fido"},
			},
			wantCode:       "INVALID_ARGUMENT",
			wantUpdateArgs: nil,
		},
		{
			name: "MUST update the pet and return the stored Pet",
			req:  validReq,
			store: stubStore{
				updatePet: pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v2"},
			},
			want: &petstorev1.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v2"},
			wantUpdateArgs: []updateCall{
				{pet: pet.Pet{Name: "pets/fido", DisplayName: "Fido"}, etag: "v1"},
			},
		},
		{
			name: "MUST update the pet when etag is omitted",
			req: &petstorev1.UpdatePetRequest{
				Pet: &petstorev1.Pet{Name: "pets/fido", DisplayName: "Fido"},
			},
			store: stubStore{
				updatePet: pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v2"},
			},
			want: &petstorev1.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v2"},
			wantUpdateArgs: []updateCall{
				{pet: pet.Pet{Name: "pets/fido", DisplayName: "Fido"}, etag: ""},
			},
		},
		{
			name: "MUST return NotFound when the pet does not exist",
			req:  validReq,
			store: stubStore{
				updateErr: &pet.NotFoundError{Message: "pets/fido not found"},
			},
			wantCode: "NOT_FOUND",
			wantUpdateArgs: []updateCall{
				{pet: pet.Pet{Name: "pets/fido", DisplayName: "Fido"}, etag: "v1"},
			},
		},
		{
			name: "MUST return Aborted when a provided etag does not match",
			req:  validReq,
			store: stubStore{
				updateErr: &pet.AbortedError{Message: "etag mismatch"},
			},
			wantCode: "ABORTED",
			wantUpdateArgs: []updateCall{
				{pet: pet.Pet{Name: "pets/fido", DisplayName: "Fido"}, etag: "v1"},
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
				{pet: pet.Pet{Name: "pets/fido", DisplayName: "Fido"}, etag: "v1"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			store := tc.store
			srv := petgrpc.PetServer{Pets: &service.PetService{Store: &store}}

			got, err := srv.UpdatePet(context.Background(), tc.req)

			assert.Equal(t, tc.wantUpdateArgs, store.updateCalls, "Store.Update args")
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
