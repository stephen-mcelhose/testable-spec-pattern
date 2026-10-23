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

// TestStore_Delete covers Store.Delete.
// Order: delete → NotFound → Aborted → delete when etag omitted.
func TestStore_Delete(t *testing.T) {
	t.Parallel()

	existing := petstore.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "1"}
	existingRaw, _ := json.Marshal(existing)

	t.Run("MUST delete the pet when it exists", func(t *testing.T) {
		t.Parallel()
		db := &stubDB{getValue: existingRaw}
		s := memstore.Store{DB: db}

		err := s.Delete(context.Background(), "pets/fido", "1")
		require.NoError(t, err)
		assert.Equal(t, []string{"pet:pets/fido"}, db.deleteCalls)
	})

	t.Run("MUST delete the pet when etag is omitted", func(t *testing.T) {
		t.Parallel()
		db := &stubDB{getValue: existingRaw}
		s := memstore.Store{DB: db}

		err := s.Delete(context.Background(), "pets/fido", "")
		require.NoError(t, err)
		assert.Equal(t, []string{"pet:pets/fido"}, db.deleteCalls)
	})

	t.Run("MUST return NotFound when the pet does not exist", func(t *testing.T) {
		t.Parallel()
		db := &stubDB{getErr: memstore.ErrNotFound}
		s := memstore.Store{DB: db}

		err := s.Delete(context.Background(), "pets/missing", "1")
		require.Error(t, err)
		var known pet.KnownError
		require.True(t, errors.As(err, &known))
		assert.Equal(t, "NOT_FOUND", known.Code())
		assert.Empty(t, db.deleteCalls)
	})

	t.Run("MUST return Aborted when a provided etag does not match", func(t *testing.T) {
		t.Parallel()
		db := &stubDB{getValue: existingRaw}
		s := memstore.Store{DB: db}

		err := s.Delete(context.Background(), "pets/fido", "stale")
		require.Error(t, err)
		var known pet.KnownError
		require.True(t, errors.As(err, &known))
		assert.Equal(t, "ABORTED", known.Code())
		assert.Empty(t, db.deleteCalls)
	})
}
