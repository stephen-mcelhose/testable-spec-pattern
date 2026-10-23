---
type: example-plan
title: Canonical example — pure function behavioral spec
status: done
skill: go-behavioral-spec
---

# 01 — Pure function

## Goal

Canonical markdown example of an RFC 2119 table-driven spec for a **pure** exported function (no injected deps, no I/O): normalize a Pet resource name in resource-oriented (AIP-ish) form.

## Approach

Single function:

```go
func NormalizePetName(name string) (string, error)
```

Resource-oriented contracts (locked):

| Requirement | Detail |
|-------------|--------|
| MUST accept prefixed name | Input MUST be `pets/{pet}` (collection `pets`, one path segment after) |
| MUST return prefixed name | Output MUST be canonical `pets/{pet}` (trimmed; ID normalized if we define case rules) |
| MUST validate `{pet}` format | See **Pet ID format** below; assert accept/reject in the table |
| MUST return InvalidArgument on format error | Error MUST be a **KnownError** that is / wraps **InvalidArgument** (`errors.As` / `errors.Is` — pick one and stick to it in the example) |

Drop `ParsePetStatus` from this plan — one function only.

Prefer 5–8 table rows. Spec-natural order: valid → boundary valid → format errors. `package pet_test` (or `example_test`).

### Pet ID format (locked for the example)

`{pet}` is a single path segment (no `/`):

- Characters: URL-safe unreserved subset — `[a-z0-9-]` only (lowercase)
- Length: 1–63
- MUST NOT start or end with `-`
- Example valid: `pets/fido`, `pets/pet-42`
- Example invalid: `pets/Fido` (uppercase), `pets/-fido`, `pets/fi/do` (extra segment), `pet/fido` (wrong collection), `fido` (unprefixed), `pets/` (empty id)

Canonicalization: trim surrounding whitespace on the full name; reject rather than silently rewrite illegal IDs (keep the pure function boring).

### Error model (sketch)

Small teaching types — not full gRPC status:

```go
type KnownError interface {
    error
    Code() string // e.g. "INVALID_ARGUMENT"
}

type InvalidArgumentError struct {
    Message string
}

func (e *InvalidArgumentError) Error() string { return e.Message }
func (e *InvalidArgumentError) Code() string  { return "INVALID_ARGUMENT" }
```

NormalizePetName MUST return an error that `errors.As` into `*InvalidArgumentError` (or satisfies `KnownError` with that code) on any format failure. Spec asserts with `errors.As`, not string match alone.

## Deliverable

```
.agents/skills/go-behavioral-spec/examples/
  README.md
  01-pure-function/
    README.md
    go.mod                         # nested module
    normalize_pet_name.go
    normalize_pet_name_test.go
```

## Types / signatures (sketch)

```go
func NormalizePetName(name string) (string, error)
```

No interfaces beyond the tiny KnownError/InvalidArgument teaching types. No stubs.

## Suggested cases

| name (RFC 2119) | fixtures |
|-----------------|----------|
| MUST accept a valid `pets/{pet}` … | single |
| MUST trim surrounding whitespace … | single |
| MUST return InvalidArgument when the name is not `pets/{pet}` | `examples`: unprefixed, wrong collection, empty id, extra segment |
| MUST return InvalidArgument when `{pet}` does not match the id format | `examples`: case, leading/trailing hyphen |

One MUST per **rule**; vary fixtures with `examples` (see discipline “One rule per test”).

## Acceptance criteria

- [x] Contracts above are encoded as MUST cases (prefixed in/out, ID format, InvalidArgument KnownError)
- [x] Test names use RFC 2119; one rule per case
- [x] Table-driven; shared assertion shape where possible
- [x] Ordering documented on `TestNormalizePetName`
- [x] README: pure / resource-name helper / copy structure not domain
- [x] Nested `go.mod`; no p360 imports; `go test ./...` green
- [x] Examples index lists 01

## Verification

```bash
cd .agents/skills/go-behavioral-spec/examples/01-pure-function
go test ./... -count=1
```

## Out of scope

Injected deps, counterfeiter, full PetsServer, SKILL.md workflow rewrite, accepting unprefixed IDs (explicitly rejected).

## Agent notes

Read `discipline.md` + `SKILL.md` before writing. Prefer table. Assert `errors.As` to InvalidArgument — do not only `assert.Error`. Do not invent OPEN skips for unreachable paths.
