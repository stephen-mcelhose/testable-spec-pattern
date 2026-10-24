# 01 — Pure function

**Teaches by example:** RFC 2119 table-driven spec for a pure helper — resource name normalize/validate.

Package clause is `pet`; import path is this directory (odd but fine for the examples module).

## Layout

| Path | Role |
|------|------|
| [`../go.mod`](../go.mod) | Single module for all examples |
| `normalize_pet_name.go` | `NormalizePetName` + KnownError types (`package pet`) |
| `normalize_pet_name_test.go` | Shape suite + `TestNormalizePetName_PetIDFormat` |

## Contracts (summary)

- MUST accept and return `pets/{pet}`; MUST reject when the name is not that shape
- MUST accept when `{pet}` matches the id format; MUST return InvalidArgument when it does not
- Errors are `*InvalidArgumentError` (`errors.As`)

MUST names are contracts only. Fixtures live in `examples` / nested `t.Run(in)` with plain labels — do not mint a MUST per sample string.

## Verify

```bash
cd .agents/skills/go-behavioral-spec/examples
go test ./01-pure-function/... -count=1
```

