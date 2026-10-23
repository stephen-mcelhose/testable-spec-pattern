package memstore

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	pet "example.com/go-behavioral-spec/examples/01-pure-function"
	"example.com/go-behavioral-spec/examples/03-petstore-api"
)

// Store matches petstore.Store: KnownError contract toward callers, bytes via injected DB.
// Store assigns Etag on Create/Update from a hash of the resource. Update/Delete take a
// precondition etag argument.
type Store struct {
	DB DB
}

// Compile-time check: same method set as 03's usage-site Store.
var _ petstore.Store = (*Store)(nil)

// Create stores a new pet and assigns an Etag. Returns AlreadyExists if present.
func (s *Store) Create(ctx context.Context, p petstore.Pet) (petstore.Pet, error) {
	key := petKey(p.Name) // [IMPLEMENTATION_DETAIL - not behavior]
	_, err := s.DB.Get(ctx, key)
	if err == nil {
		return petstore.Pet{}, &pet.AlreadyExistsError{Message: fmt.Sprintf("%s already exists", p.Name)}
	}
	if !errors.Is(err, ErrNotFound) {
		return petstore.Pet{}, internalDB("get", err)
	}

	p.Etag = etagFor(p)
	if err := s.putPet(ctx, key, p); err != nil {
		return petstore.Pet{}, err
	}
	return p, nil
}

// Get returns the pet or NotFound.
func (s *Store) Get(ctx context.Context, name string) (petstore.Pet, error) {
	return s.load(ctx, petKey(name))
}

// Update persists fields when the precondition etag matches (empty skips the check),
// and assigns a new Etag.
func (s *Store) Update(ctx context.Context, p petstore.Pet, etag string) (petstore.Pet, error) {
	key := petKey(p.Name)
	cur, err := s.load(ctx, key)
	if err != nil {
		return petstore.Pet{}, err
	}
	if etag != "" && etag != cur.Etag {
		return petstore.Pet{}, &pet.AbortedError{Message: "etag mismatch"}
	}
	cur.DisplayName = p.DisplayName
	cur.Etag = etagFor(cur)
	if err := s.putPet(ctx, key, cur); err != nil {
		return petstore.Pet{}, err
	}
	return cur, nil
}

// Delete removes the pet when the precondition etag matches (empty skips the check).
func (s *Store) Delete(ctx context.Context, name, etag string) error {
	key := petKey(name)
	cur, err := s.load(ctx, key)
	if err != nil {
		return err
	}
	if etag != "" && etag != cur.Etag {
		return &pet.AbortedError{Message: "etag mismatch"}
	}
	if err := s.DB.Delete(ctx, key); err != nil {
		if errors.Is(err, ErrNotFound) {
			return &pet.NotFoundError{Message: fmt.Sprintf("%s not found", name)}
		}
		return internalDB("delete", err)
	}
	return nil
}

// List returns a page of pets. Empty collection is success.
func (s *Store) List(ctx context.Context, pageSize int32, pageToken string) ([]petstore.Pet, string, error) {
	limit := int(pageSize)
	if limit <= 0 {
		limit = 50
	}
	_, values, next, err := s.DB.List(ctx, "pet:", limit, pageToken)
	if err != nil {
		return nil, "", internalDB("list", err)
	}
	pets := make([]petstore.Pet, 0, len(values))
	for _, v := range values {
		var p petstore.Pet
		if err := json.Unmarshal(v, &p); err != nil {
			return nil, "", &pet.InternalError{Message: fmt.Sprintf("decode pet: %v", err)}
		}
		pets = append(pets, p)
	}
	return pets, next, nil
}

func (s *Store) load(ctx context.Context, key string) (petstore.Pet, error) {
	raw, err := s.DB.Get(ctx, key)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			name := strings.TrimPrefix(key, "pet:")
			return petstore.Pet{}, &pet.NotFoundError{Message: fmt.Sprintf("%s not found", name)}
		}
		return petstore.Pet{}, internalDB("get", err)
	}
	var p petstore.Pet
	if err := json.Unmarshal(raw, &p); err != nil {
		return petstore.Pet{}, &pet.InternalError{Message: fmt.Sprintf("decode pet: %v", err)}
	}
	return p, nil
}

func (s *Store) putPet(ctx context.Context, key string, p petstore.Pet) error {
	raw, err := json.Marshal(p)
	if err != nil {
		return &pet.InternalError{Message: fmt.Sprintf("encode pet: %v", err)}
	}
	if err := s.DB.Put(ctx, key, raw); err != nil {
		return internalDB("put", err)
	}
	return nil
}

// etagFor hashes the resource body (excluding Etag) so freshness tracks content.
func etagFor(p petstore.Pet) string {
	sum := sha256.Sum256([]byte(p.Name + "\x00" + p.DisplayName))
	return hex.EncodeToString(sum[:8])
}

func petKey(name string) string {
	return "pet:" + name
}

func internalDB(op string, err error) error {
	return &pet.InternalError{Message: fmt.Sprintf("db %s: %v", op, err)}
}
