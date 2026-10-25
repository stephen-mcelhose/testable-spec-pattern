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

// TestPetServer_CreatePet covers CreatePet against AIP-133 behaviours we claim.
// Order: protovalidate shape → business rules → create → AlreadyExists → Internal.
func TestPetServer_CreatePet(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string

		req      *petstorev1.CreatePetRequest
		examples []*petstorev1.CreatePetRequest
		store    stubStore

		want           *petstorev1.Pet
		wantCode       string
		wantCreateArgs []pet.Pet
	}

	testCases := []testCase{
		{
			name: "MUST return InvalidArgument when CreatePetRequest fails shape validation",
			examples: []*petstorev1.CreatePetRequest{
				{PetId: "fido", Pet: nil},
			},
			wantCode:       "INVALID_ARGUMENT",
			wantCreateArgs: nil,
		},
		{
			name: "MUST return InvalidArgument when CreatePetRequest fails business rules",
			examples: []*petstorev1.CreatePetRequest{
				{PetId: "", Pet: &petstorev1.Pet{DisplayName: "Fido"}},
				{PetId: "Fido", Pet: &petstorev1.Pet{DisplayName: "Fido"}},
				{PetId: "fido", Pet: &petstorev1.Pet{DisplayName: ""}},
				{PetId: "fido", Pet: &petstorev1.Pet{DisplayName: "   "}},
			},
			wantCode:       "INVALID_ARGUMENT",
			wantCreateArgs: nil,
		},
		{
			name: "MUST create the pet from pet_id and return the stored Pet",
			req: &petstorev1.CreatePetRequest{
				PetId: "  fido  ",
				Pet:   &petstorev1.Pet{DisplayName: "  Fido  "},
			},
			store: stubStore{
				createPet: pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
			},
			want: &petstorev1.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
			wantCreateArgs: []pet.Pet{
				{Name: "pets/fido", DisplayName: "Fido"},
			},
		},
		{
			name: "MUST ignore Pet.Name on the create body and use pet_id for the resource name",
			req: &petstorev1.CreatePetRequest{
				PetId: "fido",
				Pet:   &petstorev1.Pet{Name: "pets/other", DisplayName: "Fido"},
			},
			store: stubStore{
				createPet: pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
			},
			want: &petstorev1.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
			wantCreateArgs: []pet.Pet{
				{Name: "pets/fido", DisplayName: "Fido"},
			},
		},
		{
			name: "MUST return AlreadyExists when the pet already exists",
			req: &petstorev1.CreatePetRequest{
				PetId: "fido",
				Pet:   &petstorev1.Pet{DisplayName: "Fido"},
			},
			store: stubStore{
				createErr: &pet.AlreadyExistsError{Message: "pets/fido already exists"},
			},
			wantCode: "ALREADY_EXISTS",
			wantCreateArgs: []pet.Pet{
				{Name: "pets/fido", DisplayName: "Fido"},
			},
		},
		{
			name: "MUST return Internal when create fails for a non-domain reason",
			req: &petstorev1.CreatePetRequest{
				PetId: "fido",
				Pet:   &petstorev1.Pet{DisplayName: "Fido"},
			},
			store: stubStore{
				createErr: errors.New("connection refused"),
			},
			wantCode: "INTERNAL",
			wantCreateArgs: []pet.Pet{
				{Name: "pets/fido", DisplayName: "Fido"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			reqs := tc.examples
			if len(reqs) == 0 {
				reqs = []*petstorev1.CreatePetRequest{tc.req}
			}
			for _, req := range reqs {
				label := "nil"
				if req != nil {
					display := ""
					if req.Pet != nil {
						display = req.Pet.DisplayName
					}
					label = req.PetId + "/" + display
				}
				t.Run(label, func(t *testing.T) {
					t.Parallel()
					store := tc.store
					srv := petgrpc.PetServer{Pets: &service.PetService{Store: &store}}

					got, err := srv.CreatePet(context.Background(), req)

					assert.Equal(t, tc.wantCreateArgs, store.createCalls, "Store.Create args")
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
