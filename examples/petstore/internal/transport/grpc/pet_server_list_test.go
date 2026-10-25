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

func TestPetServer_ListPets(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string

		req  *petstorev1.ListPetsRequest
		pets stubPets

		want         *petstorev1.ListPetsResponse
		wantCode     string
		wantListArgs []listCall
	}

	testCases := []testCase{
		{
			name: "MUST return InvalidArgument when ListPetsRequest fails shape validation",
			req: &petstorev1.ListPetsRequest{
				PageSize: -1,
			},
			wantCode:     "INVALID_ARGUMENT",
			wantListArgs: nil,
		},
		{
			name: "MUST map page_size and page_token to PetAPI and map domain pets to proto",
			req:  &petstorev1.ListPetsRequest{PageSize: 10, PageToken: "t0"},
			pets: stubPets{
				listPets: []pet.Pet{{Name: "pets/fido", DisplayName: "Fido"}},
				listNext: "t1",
			},
			want: &petstorev1.ListPetsResponse{
				Pets:          []*petstorev1.Pet{{Name: "pets/fido", DisplayName: "Fido"}},
				NextPageToken: "t1",
			},
			wantListArgs: []listCall{{pageSize: 10, pageToken: "t0"}},
		},
		{
			name: "MUST forward PetAPI errors",
			req:  &petstorev1.ListPetsRequest{PageSize: 10},
			pets: stubPets{
				listErr: &pet.InternalError{Message: "store: connection refused"},
			},
			wantCode:     "INTERNAL",
			wantListArgs: []listCall{{pageSize: 10}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			pets := tc.pets
			srv := petgrpc.PetServer{Pets: &pets}

			got, err := srv.ListPets(context.Background(), tc.req)

			assert.Equal(t, tc.wantListArgs, pets.listCalls, "PetAPI.ListPets args")
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
