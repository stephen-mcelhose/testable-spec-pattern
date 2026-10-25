# 03 — PetStore API (resource-oriented CRUD)

Thin `PetService` + injected `Store` (mock in tests only). Copy the **skeleton and file names**, not the domain.

AIP refs: [132 List](https://google.aip.dev/132) · [133 Create](https://google.aip.dev/133) · [134 Update](https://google.aip.dev/134) · [135 Delete](https://google.aip.dev/135) · [154 etag](https://google.aip.dev/154) · [158 pagination](https://google.aip.dev/158)

## Agent lookup

One folder. Open by **test function name** inside `service_test.go` (one file per source type — not one file per method):

| Need | Open |
|------|------|
| Create (full skeleton) | `service_test.go` → `TestPetService_CreatePet` |
| Get / Update / Delete / List | `service_test.go` → matching `TestPetService_*` |
| Types / Store / impl | `pet.go`, `service.go` (`Store` at usage site) |

No per-RPC subdirs or `service_<method>_test.go` splits — MUST names + `TestPetService_<Method>` are enough for lookup.

## Contract skeleton (every method)

1. **MUST** return InvalidArgument when the request fails shape validation (assert Store not called)
2. **MUST** perform the operation and return the result for a valid request (assert stub args)
3. **MUST** return the method’s domain KnownErrors / Internal
4. **MAY** post-process on success (omit if none)

| Method | Domain KnownErrors from Store |
|--------|-------------------------------|
| Create | AlreadyExists, Internal |
| Get | NotFound, Internal |
| Update | NotFound, Aborted (etag), Internal |
| Delete | NotFound, Aborted (etag when provided), Internal |
| List | Internal (empty page is OK; negative page_size → InvalidArgument; over-max coerced) |

### Create ([AIP-133](https://google.aip.dev/133)) — claimed vs out of scope

| Behaviour | In this example |
|-----------|-----------------|
| User-specified `pet_id` builds `pets/{pet_id}` | YES |
| Ignore `Pet.Name` on the create body | YES |
| Return the stored (fully populated) Pet | YES |
| Duplicate id → `ALREADY_EXISTS` | YES |
| `parent` collection field | NO (top-level pets) |
| Optional system-generated id when `pet_id` omitted | NO (`pet_id` required) |
| LRO create | NO |
| `PERMISSION_DENIED` instead of AlreadyExists when duplicate invisible | NO |

### Update ([AIP-134](https://google.aip.dev/134)) — claimed vs out of scope

| Behaviour | In this example |
|-----------|-----------------|
| `Pet.Name` identifies the resource; return stored Pet | YES |
| `NOT_FOUND` when resource does not exist | YES |
| Optional etag on resource; mismatch → `ABORTED` | YES |
| `update_mask` / implied mask / `*` full replace | NO (always applies DisplayName) |
| `allow_missing` upsert | NO |
| LRO update | NO |
| `PERMISSION_DENIED` | NO |

### Delete ([AIP-135](https://google.aip.dev/135)) — claimed vs out of scope

| Behaviour | In this example |
|-----------|-----------------|
| Succeed iff resource existed and was deleted | YES — happy-path MUST |
| `NOT_FOUND` when resource does not exist | YES |
| Optional etag; mismatch → `ABORTED` | YES (etag omitted still deletes) |
| Cascading delete / `force` → `FAILED_PRECONDITION` if children | NO |
| `allow_missing` no-op delete | NO |
| Soft delete (AIP-164) | NO |
| `PERMISSION_DENIED` before existence check | NO |

### List ([AIP-132](https://google.aip.dev/132) / [AIP-158](https://google.aip.dev/158)) — claimed vs out of scope

| Behaviour | In this example |
|-----------|-----------------|
| Top-level collection (no `parent`) | YES |
| `page_size` + `page_token` on request; `next_page_token` on response | YES |
| Default `page_size` when 0/unspecified (50) | YES |
| Coerce `page_size` above max (100) to max | YES |
| `INVALID_ARGUMENT` for negative `page_size` | YES |
| Empty collection is success (not `NOT_FOUND`) | YES |
| `filter` / `order_by` / `show_deleted` / `total_size` | NO |
| `PERMISSION_DENIED` | NO |

Resource names use `pet.NormalizePetName` from `01-pure-function`. KnownError types live there too.

## What not to copy

- Nested `go.mod` or a real Store implementation
- One mega-table for every RPC
- Re-specifying the full pet-id format corpus (that is 01)

## Verify

```bash
cd .agents/skills/go-behavioral-spec/examples
go test ./03-petstore-api/... -count=1
```
