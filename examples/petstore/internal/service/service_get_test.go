package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	petstorev1 "github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore/internal/gen/petstore/v1"
	"github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore/internal/pet"
	"github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore/internal/service"
)

// TestPetService_GetPet covers GetPet.
// Order: proto validation → business rules → return stored Pet → NotFound → Internal.
func TestPetService_GetPet(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string

		req      *petstorev1.GetPetRequest
		examples []*petstorev1.GetPetRequest
		store    stubStore

		want        *petstorev1.Pet
		wantCode    string
		wantGetArgs []string
	}

	testCases := []testCase{
		{
			name: "MUST return InvalidArgument when the pet name fails shape validation",
			examples: []*petstorev1.GetPetRequest{
				{Name: "fido"},
				{Name: "pets/"},
				{Name: "pets/Fido"},
			},
			wantCode:    "INVALID_ARGUMENT",
			wantGetArgs: nil,
		},
		{
			name: "MUST return the stored Pet for a valid pet name",
			req:  &petstorev1.GetPetRequest{Name: "  pets/fido  "},
			store: stubStore{
				getPet: &petstorev1.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
			},
			want:        &petstorev1.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
			wantGetArgs: []string{"pets/fido"},
		},
		{
			name: "MUST return NotFound when the pet does not exist",
			req:  &petstorev1.GetPetRequest{Name: "pets/missing"},
			store: stubStore{
				getErr: &pet.NotFoundError{Message: "pets/missing not found"},
			},
			wantCode:    "NOT_FOUND",
			wantGetArgs: []string{"pets/missing"},
		},
		{
			name: "MUST return Internal when get fails for a non-not-found reason",
			req:  &petstorev1.GetPetRequest{Name: "pets/fido"},
			store: stubStore{
				getErr: errors.New("connection refused"),
			},
			wantCode:    "INTERNAL",
			wantGetArgs: []string{"pets/fido"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			reqs := tc.examples
			if len(reqs) == 0 {
				reqs = []*petstorev1.GetPetRequest{tc.req}
			}
			for _, req := range reqs {
				label := "nil"
				if req != nil {
					label = req.Name
				}
				t.Run(label, func(t *testing.T) {
					t.Parallel()
					store := tc.store
					srv := service.PetService{Store: &store}

					got, err := srv.GetPet(context.Background(), req)

					assert.Equal(t, tc.wantGetArgs, store.getCalls, "Store.Get args")
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
		})
	}
}
