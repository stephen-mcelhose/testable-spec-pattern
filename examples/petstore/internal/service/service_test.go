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

// stubStore records Store calls so service specs assert the leaf boundary.
type stubStore struct {
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

	createCalls []pet.Pet
	getCalls    []string
	updateCalls []updateCall
	deleteCalls []deleteCall
	listCalls   []listCall
}

type updateCall struct {
	pet  pet.Pet
	etag string
}

type deleteCall struct {
	name string
	etag string
}

type listCall struct {
	pageSize  int32
	pageToken string
}

func (s *stubStore) Create(_ context.Context, p pet.Pet) (pet.Pet, error) {
	s.createCalls = append(s.createCalls, p)
	return s.createPet, s.createErr
}

func (s *stubStore) Get(_ context.Context, name string) (pet.Pet, error) {
	s.getCalls = append(s.getCalls, name)
	return s.getPet, s.getErr
}

func (s *stubStore) Update(_ context.Context, p pet.Pet, etag string) (pet.Pet, error) {
	s.updateCalls = append(s.updateCalls, updateCall{pet: p, etag: etag})
	return s.updatePet, s.updateErr
}

func (s *stubStore) Delete(_ context.Context, name, etag string) error {
	s.deleteCalls = append(s.deleteCalls, deleteCall{name: name, etag: etag})
	return s.deleteErr
}

func (s *stubStore) List(_ context.Context, pageSize int32, pageToken string) ([]pet.Pet, string, error) {
	s.listCalls = append(s.listCalls, listCall{pageSize: pageSize, pageToken: pageToken})
	return s.listPets, s.listNext, s.listErr
}

// TestPetService_CreatePet covers application concerns: business rules → Store.
// Proto shape lives in transport/grpc tests.
func TestPetService_CreatePet(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string

		petID       string
		displayName string
		examples    []struct{ petID, displayName string }
		store       stubStore

		want           pet.Pet
		wantCode       string
		wantCreateArgs []pet.Pet
	}

	testCases := []testCase{
		{
			name: "MUST return InvalidArgument when CreatePet fails business rules",
			examples: []struct{ petID, displayName string }{
				{petID: "", displayName: "Fido"},
				{petID: "Fido", displayName: "Fido"},
				{petID: "fido", displayName: ""},
				{petID: "fido", displayName: "   "},
			},
			wantCode:       "INVALID_ARGUMENT",
			wantCreateArgs: nil,
		},
		{
			name:        "MUST create the pet from pet_id and return the stored Pet",
			petID:       "  fido  ",
			displayName: "  Fido  ",
			store: stubStore{
				createPet: pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
			},
			want: pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
			wantCreateArgs: []pet.Pet{
				{Name: "pets/fido", DisplayName: "Fido"},
			},
		},
		{
			name:        "MUST return AlreadyExists when the pet already exists",
			petID:       "fido",
			displayName: "Fido",
			store: stubStore{
				createErr: &pet.AlreadyExistsError{Message: "pets/fido already exists"},
			},
			wantCode: "ALREADY_EXISTS",
			wantCreateArgs: []pet.Pet{
				{Name: "pets/fido", DisplayName: "Fido"},
			},
		},
		{
			name:        "MUST return Internal when create fails for a non-domain reason",
			petID:       "fido",
			displayName: "Fido",
			store: stubStore{
				createErr: errors.New("connection refused"),
			},
			wantCode: "INTERNAL",
			wantCreateArgs: []pet.Pet{
				{Name: "pets/fido", DisplayName: "Fido"},
			},
		},
	}

	type pair struct{ petID, displayName string }

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			inputs := make([]pair, 0, len(tc.examples)+1)
			for _, ex := range tc.examples {
				inputs = append(inputs, pair{ex.petID, ex.displayName})
			}
			if len(inputs) == 0 {
				inputs = []pair{{tc.petID, tc.displayName}}
			}
			for _, in := range inputs {
				t.Run(in.petID+"/"+in.displayName, func(t *testing.T) {
					t.Parallel()
					store := tc.store
					srv := service.PetService{Store: &store}

					got, err := srv.CreatePet(context.Background(), in.petID, in.displayName)

					assert.Equal(t, tc.wantCreateArgs, store.createCalls, "Store.Create args")
					if tc.wantCode != "" {
						require.Error(t, err)
						var known pet.KnownError
						require.True(t, errors.As(err, &known), "error MUST be KnownError, got %T: %v", err, err)
						assert.Equal(t, tc.wantCode, known.Code())
						assert.Equal(t, pet.Pet{}, got)
						return
					}
					require.NoError(t, err)
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}

func TestPetService_GetPet(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string

		petName  string
		examples []string
		store    stubStore

		want        pet.Pet
		wantCode    string
		wantGetArgs []string
	}

	testCases := []testCase{
		{
			name: "MUST return InvalidArgument when the pet name fails business rules",
			examples: []string{
				"fido",
				"pets/",
				"pets/Fido",
			},
			wantCode:    "INVALID_ARGUMENT",
			wantGetArgs: nil,
		},
		{
			name:    "MUST return the stored Pet for a valid pet name",
			petName: "  pets/fido  ",
			store: stubStore{
				getPet: pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
			},
			want:        pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v1"},
			wantGetArgs: []string{"pets/fido"},
		},
		{
			name:    "MUST return NotFound when the pet does not exist",
			petName: "pets/missing",
			store: stubStore{
				getErr: &pet.NotFoundError{Message: "pets/missing not found"},
			},
			wantCode:    "NOT_FOUND",
			wantGetArgs: []string{"pets/missing"},
		},
		{
			name:    "MUST return Internal when get fails for a non-not-found reason",
			petName: "pets/fido",
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
			names := tc.examples
			if len(names) == 0 {
				names = []string{tc.petName}
			}
			for _, name := range names {
				t.Run(name, func(t *testing.T) {
					t.Parallel()
					store := tc.store
					srv := service.PetService{Store: &store}

					got, err := srv.GetPet(context.Background(), name)

					assert.Equal(t, tc.wantGetArgs, store.getCalls, "Store.Get args")
					if tc.wantCode != "" {
						require.Error(t, err)
						var known pet.KnownError
						require.True(t, errors.As(err, &known), "error MUST be KnownError, got %T: %v", err, err)
						assert.Equal(t, tc.wantCode, known.Code())
						assert.Equal(t, pet.Pet{}, got)
						return
					}
					require.NoError(t, err)
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}

func TestPetService_UpdatePet(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string

		petName     string
		displayName string
		etag        string
		store       stubStore

		want           pet.Pet
		wantCode       string
		wantUpdateArgs []updateCall
	}

	testCases := []testCase{
		{
			name:           "MUST return InvalidArgument when UpdatePet fails business rules",
			petName:        "fido",
			displayName:    "Fido",
			etag:           "v1",
			wantCode:       "INVALID_ARGUMENT",
			wantUpdateArgs: nil,
		},
		{
			name:        "MUST update the pet and return the stored Pet",
			petName:     "pets/fido",
			displayName: "Fido",
			etag:        "v1",
			store: stubStore{
				updatePet: pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v2"},
			},
			want: pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v2"},
			wantUpdateArgs: []updateCall{
				{pet: pet.Pet{Name: "pets/fido", DisplayName: "Fido"}, etag: "v1"},
			},
		},
		{
			name:        "MUST update the pet when etag is omitted",
			petName:     "pets/fido",
			displayName: "Fido",
			store: stubStore{
				updatePet: pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v2"},
			},
			want: pet.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "v2"},
			wantUpdateArgs: []updateCall{
				{pet: pet.Pet{Name: "pets/fido", DisplayName: "Fido"}, etag: ""},
			},
		},
		{
			name:        "MUST return NotFound when the pet does not exist",
			petName:     "pets/fido",
			displayName: "Fido",
			etag:        "v1",
			store: stubStore{
				updateErr: &pet.NotFoundError{Message: "pets/fido not found"},
			},
			wantCode: "NOT_FOUND",
			wantUpdateArgs: []updateCall{
				{pet: pet.Pet{Name: "pets/fido", DisplayName: "Fido"}, etag: "v1"},
			},
		},
		{
			name:        "MUST return Aborted when a provided etag does not match",
			petName:     "pets/fido",
			displayName: "Fido",
			etag:        "v1",
			store: stubStore{
				updateErr: &pet.AbortedError{Message: "etag mismatch"},
			},
			wantCode: "ABORTED",
			wantUpdateArgs: []updateCall{
				{pet: pet.Pet{Name: "pets/fido", DisplayName: "Fido"}, etag: "v1"},
			},
		},
		{
			name:        "MUST return Internal when update fails for a non-domain reason",
			petName:     "pets/fido",
			displayName: "Fido",
			etag:        "v1",
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
			srv := service.PetService{Store: &store}

			got, err := srv.UpdatePet(context.Background(), tc.petName, tc.displayName, tc.etag)

			assert.Equal(t, tc.wantUpdateArgs, store.updateCalls, "Store.Update args")
			if tc.wantCode != "" {
				require.Error(t, err)
				var known pet.KnownError
				require.True(t, errors.As(err, &known), "error MUST be KnownError, got %T: %v", err, err)
				assert.Equal(t, tc.wantCode, known.Code())
				assert.Equal(t, pet.Pet{}, got)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPetService_DeletePet(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string

		petName string
		etag    string
		store   stubStore

		wantCode       string
		wantDeleteArgs []deleteCall
	}

	testCases := []testCase{
		{
			name:           "MUST return InvalidArgument when the pet name fails business rules",
			petName:        "fido",
			etag:           "v1",
			wantCode:       "INVALID_ARGUMENT",
			wantDeleteArgs: nil,
		},
		{
			name:    "MUST delete the pet when it exists",
			petName: "  pets/fido  ",
			etag:    "v1",
			wantDeleteArgs: []deleteCall{
				{name: "pets/fido", etag: "v1"},
			},
		},
		{
			name:    "MUST delete the pet when etag is omitted",
			petName: "pets/fido",
			etag:    "",
			wantDeleteArgs: []deleteCall{
				{name: "pets/fido", etag: ""},
			},
		},
		{
			name:    "MUST return NotFound when the pet does not exist",
			petName: "pets/missing",
			etag:    "v1",
			store: stubStore{
				deleteErr: &pet.NotFoundError{Message: "pets/missing not found"},
			},
			wantCode: "NOT_FOUND",
			wantDeleteArgs: []deleteCall{
				{name: "pets/missing", etag: "v1"},
			},
		},
		{
			name:    "MUST return Aborted when a provided etag does not match",
			petName: "pets/fido",
			etag:    "stale",
			store: stubStore{
				deleteErr: &pet.AbortedError{Message: "etag mismatch"},
			},
			wantCode: "ABORTED",
			wantDeleteArgs: []deleteCall{
				{name: "pets/fido", etag: "stale"},
			},
		},
		{
			name:    "MUST return Internal when delete fails for a non-domain reason",
			petName: "pets/fido",
			etag:    "v1",
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
			srv := service.PetService{Store: &store}

			err := srv.DeletePet(context.Background(), tc.petName, tc.etag)

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
