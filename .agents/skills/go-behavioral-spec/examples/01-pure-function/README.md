# 01 — Pure function

**Teaches by example:** RFC 2119 table-driven spec for a pure helper — resource name normalize/validate.

**Code lives in** [`../pet/`](../pet/) (shared examples module). This folder is the index entry / orientation for the pure-function pattern.

## Layout

| Path | Role |
|------|------|
| [`../go.mod`](../go.mod) | Single module for all examples |
| [`../pet/normalize_pet_name.go`](../pet/normalize_pet_name.go) | `NormalizePetName` + KnownError types |
| [`../pet/normalize_pet_name_test.go`](../pet/normalize_pet_name_test.go) | Shape suite + `TestNormalizePetName_PetIDFormat` |

## Contracts (summary)

- MUST accept and return `pets/{pet}`; MUST reject when the name is not that shape
- MUST accept when `{pet}` matches the id format; MUST return InvalidArgument when it does not
- Errors are `*InvalidArgumentError` (`errors.As`)

MUST names are contracts only. Fixtures live in `examples` / nested `t.Run(in)` with plain labels — do not mint a MUST per sample string.

## Verify

```bash
cd .agents/skills/go-behavioral-spec/examples
go test ./pet/... -count=1
```
