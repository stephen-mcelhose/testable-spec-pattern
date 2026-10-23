package memstore_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pet "example.com/go-behavioral-spec/examples/01-pure-function"
	"example.com/go-behavioral-spec/examples/03-petstore-api"
	"example.com/go-behavioral-spec/examples/04-in-memory-store"
)

// TestStore_Create covers Store.Create — KnownError contract toward callers.
// Order: create + return Pet with etag → AlreadyExists → Internal (DB get failure).
func TestStore_Create(t *testing.T) {
	t.Parallel()

	t.Run("MUST create the pet and return the stored Pet including an etag", func(t *testing.T) {
		t.Parallel()
		db := &stubDB{getErr: memstore.ErrNotFound}
		s := memstore.Store{DB: db}

		got, err := s.Create(context.Background(), petstore.Pet{
			Name: "pets/fido", DisplayName: "Fido",
		})
		require.NoError(t, err)
		assert.Equal(t, "pets/fido", got.Name)
		assert.Equal(t, "Fido", got.DisplayName)
		assert.NotEmpty(t, got.Etag) // freshness token for later Update/Delete; how it is computed is not the contract

		assert.Equal(t, []string{"pet:pets/fido"}, db.getCalls)
		require.Len(t, db.putCalls, 1)
		assert.Equal(t, "pet:pets/fido", db.putCalls[0].key)
		var stored petstore.Pet
		require.NoError(t, json.Unmarshal(db.putCalls[0].value, &stored))
		assert.Equal(t, got, stored)
	})

	t.Run("MUST return AlreadyExists when the pet name already exists", func(t *testing.T) {
		t.Parallel()
		existing, _ := json.Marshal(petstore.Pet{Name: "pets/fido", DisplayName: "Old", Etag: "1"})
		db := &stubDB{getValue: existing}
		s := memstore.Store{DB: db}

		got, err := s.Create(context.Background(), petstore.Pet{
			Name: "pets/fido", DisplayName: "Fido",
		})
		require.Error(t, err)
		var known pet.KnownError
		require.True(t, errors.As(err, &known))
		assert.Equal(t, "ALREADY_EXISTS", known.Code())
		assert.Equal(t, petstore.Pet{}, got)
		assert.Empty(t, db.putCalls)
	})

	t.Run("MUST return Internal when the DB get fails for a non-not-found reason", func(t *testing.T) {
		t.Parallel()
		db := &stubDB{getErr: errors.New("connection refused")}
		s := memstore.Store{DB: db}

		got, err := s.Create(context.Background(), petstore.Pet{
			Name: "pets/fido", DisplayName: "Fido",
		})
		require.Error(t, err)
		var known pet.KnownError
		require.True(t, errors.As(err, &known))
		assert.Equal(t, "INTERNAL", known.Code())
		assert.Equal(t, petstore.Pet{}, got)
		assert.Empty(t, db.putCalls)
	})
}
