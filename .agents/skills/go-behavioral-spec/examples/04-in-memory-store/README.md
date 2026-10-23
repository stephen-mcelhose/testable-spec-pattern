# 04 — Store over injected DB

Domain `Store` with a **KnownError** caller-facing contract. Delegates bytes to an injected generic `DB`. Copy the **discipline**, not the pet domain.

Standalone example (tests ≈ spec). Reuses `pet.KnownError` and `petstore.Pet` / `petstore.Store` method set for convenience (`var _ petstore.Store = (*Store)(nil)`). Store assigns Etag on Create/Update from a hash of the resource; Update/Delete take a precondition etag argument.

## Agent lookup

| Need | Open |
|------|------|
| Create (centerpiece) | `store_create_test.go` → `TestStore_Create` |
| Get / Update / Delete / List | `store_*_test.go` |
| `DB` interface | `db.go` |
| `Store` | `store.go` |
| Test stub | `stubs_test.go` |

## Store ↔ caller contract (KnownErrors)

| Method | KnownErrors |
|--------|-------------|
| Create | AlreadyExists, Internal |
| Get | NotFound, Internal |
| Update | NotFound, Aborted (etag), Internal |
| Delete | NotFound, Aborted (etag), Internal |
| List | Internal (empty page is OK) |

`DB` may return `ErrNotFound`; Store maps that — callers of Store never see `ErrNotFound` as the contract.

## Verify

```bash
cd .agents/skills/go-behavioral-spec/examples
go test ./04-in-memory-store/... -count=1
```
