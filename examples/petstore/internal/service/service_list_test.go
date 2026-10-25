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

func TestPetService_ListPets(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string

		pageSize  int32
		pageToken string
		store     stubStore

		wantPets     []pet.Pet
		wantNext     string
		wantCode     string
		wantListArgs []listCall
	}

	testCases := []testCase{
		{
			name: "MUST use the default page_size when page_size is unspecified",
			store: stubStore{
				listPets: []pet.Pet{},
			},
			wantPets:     []pet.Pet{},
			wantListArgs: []listCall{{pageSize: 50}},
		},
		{
			name:     "MUST coerce page_size above the maximum to the maximum",
			pageSize: 101,
			store: stubStore{
				listPets: []pet.Pet{},
			},
			wantPets:     []pet.Pet{},
			wantListArgs: []listCall{{pageSize: 100}},
		},
		{
			name:      "MUST return the stored page and next_page_token",
			pageSize:  10,
			pageToken: "t0",
			store: stubStore{
				listPets: []pet.Pet{{Name: "pets/fido", DisplayName: "Fido"}},
				listNext: "t1",
			},
			wantPets:     []pet.Pet{{Name: "pets/fido", DisplayName: "Fido"}},
			wantNext:     "t1",
			wantListArgs: []listCall{{pageSize: 10, pageToken: "t0"}},
		},
		{
			name:     "MUST return an empty page when there are no pets",
			pageSize: 10,
			store: stubStore{
				listPets: nil,
			},
			wantPets:     []pet.Pet{},
			wantListArgs: []listCall{{pageSize: 10}},
		},
		{
			name:     "MUST return Internal when list fails for a non-domain reason",
			pageSize: 10,
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
			srv := service.PetService{Store: &store}

			got, next, err := srv.ListPets(context.Background(), tc.pageSize, tc.pageToken)

			assert.Equal(t, tc.wantListArgs, store.listCalls, "Store.List args")
			if tc.wantCode != "" {
				require.Error(t, err)
				var known pet.KnownError
				require.True(t, errors.As(err, &known), "error MUST be KnownError, got %T: %v", err, err)
				assert.Equal(t, tc.wantCode, known.Code())
				assert.Nil(t, got)
				assert.Empty(t, next)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.wantPets, got)
			assert.Equal(t, tc.wantNext, next)
		})
	}
}
