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

// TestStore_Update covers Store.Update.
// Order: update + new store etag → NotFound → Aborted (precondition etag) → omit precondition.
func TestStore_Update(t *testing.T) {
	t.Parallel()

	existing := petstore.Pet{Name: "pets/fido", DisplayName: "Old", Etag: "1"}
	existingRaw, _ := json.Marshal(existing)

	t.Run("MUST update the pet and return the stored Pet with a new etag", func(t *testing.T) {
		t.Parallel()
		db := &stubDB{getValue: existingRaw}
		s := memstore.Store{DB: db}

		got, err := s.Update(context.Background(), petstore.Pet{
			Name: "pets/fido", DisplayName: "Fido",
		}, "1")
		require.NoError(t, err)
		assert.Equal(t, "pets/fido", got.Name)
		assert.Equal(t, "Fido", got.DisplayName)
		assert.NotEmpty(t, got.Etag)
		assert.NotEqual(t, "1", got.Etag)
		require.Len(t, db.putCalls, 1)
	})

	t.Run("MUST update the pet when precondition etag is omitted", func(t *testing.T) {
		t.Parallel()
		db := &stubDB{getValue: existingRaw}
		s := memstore.Store{DB: db}

		got, err := s.Update(context.Background(), petstore.Pet{
			Name: "pets/fido", DisplayName: "Fido",
		}, "")
		require.NoError(t, err)
		assert.Equal(t, "Fido", got.DisplayName)
		assert.NotEqual(t, "1", got.Etag)
		require.Len(t, db.putCalls, 1)
	})

	t.Run("MUST return NotFound when the pet does not exist", func(t *testing.T) {
		t.Parallel()
		db := &stubDB{getErr: memstore.ErrNotFound}
		s := memstore.Store{DB: db}

		got, err := s.Update(context.Background(), petstore.Pet{
			Name: "pets/missing", DisplayName: "X",
		}, "1")
		require.Error(t, err)
		var known pet.KnownError
		require.True(t, errors.As(err, &known))
		assert.Equal(t, "NOT_FOUND", known.Code())
		assert.Equal(t, petstore.Pet{}, got)
		assert.Empty(t, db.putCalls)
	})

	t.Run("MUST return Aborted when a provided etag does not match", func(t *testing.T) {
		t.Parallel()
		db := &stubDB{getValue: existingRaw}
		s := memstore.Store{DB: db}

		got, err := s.Update(context.Background(), petstore.Pet{
			Name: "pets/fido", DisplayName: "Fido",
		}, "stale")
		require.Error(t, err)
		var known pet.KnownError
		require.True(t, errors.As(err, &known))
		assert.Equal(t, "ABORTED", known.Code())
		assert.Equal(t, petstore.Pet{}, got)
		assert.Empty(t, db.putCalls)
	})
}
