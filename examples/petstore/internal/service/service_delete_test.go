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
