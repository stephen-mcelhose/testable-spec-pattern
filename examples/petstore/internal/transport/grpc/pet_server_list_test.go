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

func TestPetServer_ListPets(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string

		req   *petstorev1.ListPetsRequest
		store stubStore

		want         *petstorev1.ListPetsResponse
		wantCode     string
		wantListArgs []listCall
	}

	testCases := []testCase{
		{
			name: "MUST return InvalidArgument when page_size fails shape validation",
			req: &petstorev1.ListPetsRequest{
				PageSize: -1,
			},
			wantCode:     "INVALID_ARGUMENT",
			wantListArgs: nil,
		},
		{
			name: "MUST use the default page_size when page_size is unspecified",
			req:  &petstorev1.ListPetsRequest{},
			store: stubStore{
				listPets: []pet.Pet{},
			},
			want: &petstorev1.ListPetsResponse{
				Pets: []*petstorev1.Pet{},
			},
			wantListArgs: []listCall{{pageSize: 50}},
		},
		{
			name: "MUST coerce page_size above the maximum to the maximum",
			req:  &petstorev1.ListPetsRequest{PageSize: 101},
			store: stubStore{
				listPets: []pet.Pet{},
			},
			want: &petstorev1.ListPetsResponse{
				Pets: []*petstorev1.Pet{},
			},
			wantListArgs: []listCall{{pageSize: 100}},
		},
		{
			name: "MUST return the stored page and next_page_token for a valid ListPetsRequest",
			req:  &petstorev1.ListPetsRequest{PageSize: 10, PageToken: "t0"},
			store: stubStore{
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
			name: "MUST return an empty page when there are no pets",
			req:  &petstorev1.ListPetsRequest{PageSize: 10},
			store: stubStore{
				listPets: nil,
			},
			want: &petstorev1.ListPetsResponse{
				Pets: []*petstorev1.Pet{},
			},
			wantListArgs: []listCall{{pageSize: 10}},
		},
		{
			name: "MUST return Internal when list fails for a non-domain reason",
			req:  &petstorev1.ListPetsRequest{PageSize: 10},
			store: stubStore{
				listErr: errors.New("connection refused"),
			},
			wantCode:     "INTERNAL",
			wantListArgs: []listCall{{pageSize: 10}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			store := tc.store
			srv := petgrpc.PetServer{Pets: &service.PetService{Store: &store}}

			got, err := srv.ListPets(context.Background(), tc.req)

			assert.Equal(t, tc.wantListArgs, store.listCalls, "Store.List args")
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
