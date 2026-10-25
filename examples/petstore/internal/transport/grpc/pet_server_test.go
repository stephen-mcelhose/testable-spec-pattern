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

// stubPets records PetAPI calls so transport specs assert mapping, not business rules.
type stubPets struct {
	createPet pet.Pet
	createErr error
	getPet    pet.Pet
	getErr    error
	updatePet pet.Pet
	updateErr error
	deleteErr error
	listPets  []pet.Pet
	listNext  string
	listErr   error

	createCalls []createCall
	getCalls    []string
	updateCalls []updateCall
	deleteCalls []deleteCall
	listCalls   []listCall
}

type createCall struct {
	petID       string
	displayName string
}

type updateCall struct {
	name        string
	displayName string
	etag        string
}

type deleteCall struct {
	name string
	etag string
}

type listCall struct {
	pageSize  int32
	pageToken string
}

func (s *stubPets) CreatePet(_ context.Context, petID, displayName string) (pet.Pet, error) {
	s.createCalls = append(s.createCalls, createCall{petID: petID, displayName: displayName})
	return s.createPet, s.createErr
}

func (s *stubPets) GetPet(_ context.Context, name string) (pet.Pet, error) {
	s.getCalls = append(s.getCalls, name)
	return s.getPet, s.getErr
}

func (s *stubPets) UpdatePet(_ context.Context, name, displayName, etag string) (pet.Pet, error) {
	s.updateCalls = append(s.updateCalls, updateCall{name: name, displayName: displayName, etag: etag})
	return s.updatePet, s.updateErr
}

func (s *stubPets) DeletePet(_ context.Context, name, etag string) error {
	s.deleteCalls = append(s.deleteCalls, deleteCall{name: name, etag: etag})
	return s.deleteErr
}

func (s *stubPets) ListPets(_ context.Context, pageSize int32, pageToken string) ([]pet.Pet, string, error) {
	s.listCalls = append(s.listCalls, listCall{pageSize: pageSize, pageToken: pageToken})
	return s.listPets, s.listNext, s.listErr
}

// TestPetServer_CreatePet covers transport concerns only: shape + field mapping to PetAPI.
// Business rules and Store behaviour live in service_test.
func TestPetServer_CreatePet(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string

		req  *petstorev1.CreatePetRequest
		pets stubPets

		want           *petstorev1.Pet
		wantCode       string
		wantCreateArgs []createCall
	}

	testCases := []testCase{
		{
			name:           "MUST return InvalidArgument when CreatePetRequest fails shape validation",
			req:            &petstorev1.CreatePetRequest{PetId: "fido", Pet: nil},
			wantCode:       "INVALID_ARGUMENT",
			wantCreateArgs: nil,
		},
		{
			name: "MUST map pet_id and display_name to PetAPI and map the domain Pet to proto",
			req: &petstorev1.CreatePetRequest{
				PetId: "fido",
				Pet:   &petstorev1.Pet{Name: "pets/other", DisplayName: "Fido"},
			},
			pets: stubPets{
				createPet: pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
			},
			want: &petstorev1.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
			wantCreateArgs: []createCall{
				{petID: "fido", displayName: "Fido"},
			},
		},
		{
			name: "MUST forward PetAPI errors without calling further mapping",
			req: &petstorev1.CreatePetRequest{
				PetId: "fido",
				Pet:   &petstorev1.Pet{DisplayName: "Fido"},
			},
			pets: stubPets{
				createErr: &pet.AlreadyExistsError{Message: "pets/fido already exists"},
			},
			wantCode: "ALREADY_EXISTS",
			wantCreateArgs: []createCall{
				{petID: "fido", displayName: "Fido"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			pets := tc.pets
			srv := petgrpc.PetServer{Pets: &pets}

			got, err := srv.CreatePet(context.Background(), tc.req)

			assert.Equal(t, tc.wantCreateArgs, pets.createCalls, "PetAPI.CreatePet args")
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
