package memstore

import (
	"context"
	"errors"
)

// ErrNotFound is returned by DB when a key is absent. Store maps this to pet.NotFound
// (or AlreadyExists checks). It is not part of Store's caller-facing KnownError contract.
var ErrNotFound = errors.New("db: not found")

// DB is a generic key/value leaf injected into Store (usage-site interface).
// Production might wrap Spanner/SQL; tests inject a recording stub.
type DB interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Put(ctx context.Context, key string, value []byte) error
	Delete(ctx context.Context, key string) error
	// List returns up to limit entries with the given key prefix, after pageToken (empty = start).
	// nextToken is empty when there is no further page.
	List(ctx context.Context, prefix string, limit int, pageToken string) (keys []string, values [][]byte, nextToken string, err error)
}
