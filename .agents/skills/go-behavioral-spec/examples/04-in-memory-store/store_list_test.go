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

// TestStore_List covers Store.List.
// Order: return page → empty page → Internal.
func TestStore_List(t *testing.T) {
	t.Parallel()

	t.Run("MUST return the stored page and next_page_token", func(t *testing.T) {
		t.Parallel()
		p := petstore.Pet{Name: "pets/fido", DisplayName: "Fido", Etag: "1"}
		raw, _ := json.Marshal(p)
		db := &stubDB{
			listKeys: []string{"pet:pets/fido"},
			listVals: [][]byte{raw},
			listNext: "tok",
		}
		s := memstore.Store{DB: db}

		got, next, err := s.List(context.Background(), 10, "t0")
		require.NoError(t, err)
		assert.Equal(t, []petstore.Pet{p}, got)
		assert.Equal(t, "tok", next)
		assert.Equal(t, []listCall{{prefix: "pet:", limit: 10, pageToken: "t0"}}, db.listCalls)
	})

	t.Run("MUST return an empty page when there are no pets", func(t *testing.T) {
		t.Parallel()
		db := &stubDB{}
		s := memstore.Store{DB: db}

		got, next, err := s.List(context.Background(), 10, "")
		require.NoError(t, err)
		assert.Empty(t, got)
		assert.Empty(t, next)
	})

	t.Run("MUST return Internal when the DB list fails", func(t *testing.T) {
		t.Parallel()
		db := &stubDB{listErr: errors.New("connection refused")}
		s := memstore.Store{DB: db}

		got, next, err := s.List(context.Background(), 10, "")
		require.Error(t, err)
		var known pet.KnownError
		require.True(t, errors.As(err, &known))
		assert.Equal(t, "INTERNAL", known.Code())
		assert.Nil(t, got)
		assert.Empty(t, next)
	})
}
