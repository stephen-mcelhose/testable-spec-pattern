---
type: example-plan
title: Canonical example — Store over injected generic DB (KnownError contract)
description: Plan for the storage-leaf example teaching how to spec a domain Store whose public contract is KnownErrors, delegating to a generic injected DB interface stubbed in tests.
status: done
skill: go-behavioral-spec
depends_on: lessons-learned; reuses types from 01/03 for convenience (not a layered demo)
tags: [go-behavioral-spec, examples, store, known-error, injected-dependency, rfc-2119]
timestamp: 2026-10-23T22:45:43Z
---

# 04 — Store over injected DB (storage leaf)

## Goal

Document the discipline for **spec’ing a domain Store**: tests ≈ contract, and that contract is **KnownErrors + returned resources**. The Store **delegates** to an injected **generic DB** interface (KV/row-ish — not Pet-shaped). Specs mock the DB; optional tiny in-memory DB impl is fine for local play, not the teaching centerpiece.

Examples in this series are **mostly independent**. Reuse `pet.KnownError` (and optionally `Pet`) for convenience. An agent should open `04-…/` alone and learn: Store owns KnownErrors toward callers; DB is the next injectable leaf.

Read [lessons-learned.md](./lessons-learned.md).

## Why this exists (workflow / discipline)

| Unit under test | Spec focus |
|-----------------|------------|
| **Store** | Persist/return Pet; **emit KnownErrors** (`AlreadyExists`, `NotFound`, `Aborted`, `Internal`) |
| **DB** (injected) | Generic get/put/delete/list — stubbed in Store tests (presence/absence of calls + canned bytes/errors) |

Callers of Store should not parse driver strings. Store maps DB outcomes → KnownErrors. Same pattern as 02/03 (inject leaf, assert boundary) with the **outcome contract** being typed domain errors.

Do **not** frame 04 as “the missing half of 03.” Independence > wiring into `PetService`.

## Approach

```go
// Usage-site narrow DB (generic — keys/values, not Pet).
type DB interface {
    Get(ctx context.Context, key string) ([]byte, error)
    Put(ctx context.Context, key string, value []byte) error
    Delete(ctx context.Context, key string) error
    // List or scan as needed for ListPets — keep tiny
}

// Domain store: KnownError contract lives here.
type Store struct {
    DB DB
}

func (s *Store) Create(ctx context.Context, p Pet) (Pet, error)
// Get, Update, Delete, List …
```

- Encode/decode Pet (JSON is fine) inside Store.
- Key scheme boring: e.g. `pets/{id}` or `pet:` + name.
- **Create:** Get-then-Put or Put-if-absent; existing key → `AlreadyExists`; assign etag on write.
- **Get:** missing → `NotFound`; decode → Pet.
- **Update/Delete:** load; missing → `NotFound`; etag provided and mismatch → `Aborted`; else write/delete.
- **DB errors** that aren’t “not found” / conflict → `Internal`.
- Prefer a small `ErrNotFound` (or sentinel) from the DB stub/impl that Store maps — do not leak it as the Store contract.

No production `MemoryDB` — build against the interface; tests supply a stub.

### KnownError contract (Store → caller)

| Method | MUST return (KnownError) | Notes |
|--------|--------------------------|--------|
| **Create** | `AlreadyExists` if present; else Pet with etag | Never NotFound |
| **Get** | `NotFound` if absent; else Pet | |
| **Update** | `NotFound`; `Aborted` if etag mismatch; else updated Pet | Empty etag = skip freshness |
| **Delete** | `NotFound`; `Aborted` if etag mismatch; else nil | Empty etag = skip freshness |
| **List** | page + next token; empty OK | Keep pagination tiny |

Do **not** return raw `errors.New("…")` for domain cases. Use `pet.*Error` from 01.

### Shape validation

**Out of scope** for Store specs (service concern). Assume valid names/etags.

## Deliverable

```
examples/
  04-in-memory-store/             # or 04-store-db/ — package memstore / petstoredb
    README.md                     # KnownError contract; DB injection; agent lookup
    db.go                         # DB interface (usage site)
    store.go                      # Store{DB} + CRUD
    store_create_test.go          # stub DB; KnownError MUSTs + call recording
    store_get_test.go
    store_update_test.go
    store_delete_test.go
    store_list_test.go
    stubs_test.go                 # recording stubDB
```

Path lookup: `store_<method>_test.go` / `TestStore_<Method>`.

Update `examples/README.md` index.

## Suite layout

1. **MUST** return the stored resource / succeed for the happy path (assert DB Put/Get args as leaf boundary)
2. **MUST** return `<KnownError>` when preconditions fail (table above) — assert DB interactions as needed (e.g. no Put on Aborted)
3. Fixtures via `examples` slices; behaviour-first MUST titles
4. Create is the teaching centerpiece

## Acceptance criteria

- [x] Standalone folder on shared `examples/go.mod`
- [x] `Store` injects usage-site `DB`; tests use a recording stub
- [x] Domain failures are **only** KnownError types (`pet` from 01)
- [x] Create/Get/Update/Delete/List Store suites; Create teaches the pattern
- [x] No real Spanner/SQL driver; no MemoryDB — stub `DB` in tests only
- [x] README: tests≈spec; KnownError = Store↔caller contract; DB = injected leaf
- [x] `examples/README.md` + lessons-learned updated

## Verification

```bash
cd .agents/skills/go-behavioral-spec/examples
go test ./04-in-memory-store/... -count=1
```

## Out of scope

Real cloud DB, migrations, transactions, soft delete, cascade, `allow_missing`, counterfeiter generate, wiring to 03 `PetService`, re-teaching AIP service shape or name format.

## Agent notes

Copy discipline: **KnownError MUSTs on Store**, stub the generic DB, assert leaf calls. Do not turn the suite into a test of JSON encoding minutiae unless a MUST claims it.
