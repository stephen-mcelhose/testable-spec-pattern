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

// TestStore_Get covers Store.Get.
// Order: return stored Pet → NotFound → Internal.
func TestStore_Get(t *testing.T) {
	t.Parallel()

	t.Run("MUST return the stored Pet", func(t *testing.T) {
		t.Parallel()
		want := petstore.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "1"}
		raw, _ := json.Marshal(want)
		db := &stubDB{getValue: raw}
		s := memstore.Store{DB: db}

		got, err := s.Get(context.Background(), "pets/fido")
		require.NoError(t, err)
		assert.Equal(t, want, got)
		assert.Equal(t, []string{"pet:pets/fido"}, db.getCalls)
	})

	t.Run("MUST return NotFound when the pet does not exist", func(t *testing.T) {
		t.Parallel()
		db := &stubDB{getErr: memstore.ErrNotFound}
		s := memstore.Store{DB: db}

		got, err := s.Get(context.Background(), "pets/missing")
		require.Error(t, err)
		var known pet.KnownError
		require.True(t, errors.As(err, &known))
		assert.Equal(t, "NOT_FOUND", known.Code())
		assert.Equal(t, petstore.Pet{}, got)
	})

	t.Run("MUST return Internal when the DB get fails for a non-not-found reason", func(t *testing.T) {
		t.Parallel()
		db := &stubDB{getErr: errors.New("connection refused")}
		s := memstore.Store{DB: db}

		got, err := s.Get(context.Background(), "pets/fido")
		require.Error(t, err)
		var known pet.KnownError
		require.True(t, errors.As(err, &known))
		assert.Equal(t, "INTERNAL", known.Code())
		assert.Equal(t, petstore.Pet{}, got)
	})
}
