---
type: example-plan
title: Canonical example — injected dependency (functional)
description: Plan for the functional-injection example (BuildPetView) teaching narrow interface injection, leaf boundary assertions, and stub discipline with a function-arg wiring.
status: done
skill: go-behavioral-spec
depends_on: 01-pure-function
tags: [go-behavioral-spec, examples, injected-dependency, functional, rfc-2119]
timestamp: 2026-10-23T22:45:43Z
---

# 02a — Injected dependency (functional)

## Goal

Nested **runnable** module: table-drivable core with a **narrow interface** passed as a function parameter (leaf may I/O in prod; stubbed in tests). Same MUST contracts as 02b (method form) — only wiring differs.

Read [lessons-learned.md](./lessons-learned.md) before writing.

## Approach

```go
type PetGetter interface {
    Get(ctx context.Context, name string) (Pet, error)
}

func BuildPetView(ctx context.Context, name string, getter PetGetter) (PetView, error)
```

Flow (keep boring):

1. `pet.NormalizePetName(name)` from `01-pure-function` (same module — `import pet "…/examples/01-pure-function"`).
2. On invalid name → return InvalidArgument **without** calling getter.
3. `getter.Get` → map not-found / other errors to KnownError types; on success map `Pet` → `PetView`.

Hand-roll a **local** stub in the test file. Comment that counterfeiter is preferred when call-count/arg capture or a package-owned interface warrants `make generate-fakes` — do not generate fakes in this example tree.

## Deliverable

```
examples/                         # go.mod at examples/ root
  02-injected-dependency/
    README.md                     # MUST parity + path index
    functional/                   # package functional — open this for function-arg wiring
      build_pet_view.go
      build_pet_view_test.go
    method/                       # 02b — do not put method tests here
```

Same module as 01. Path-based lookup: agents open `functional/` or `method/`. No nested `go.mod`.

## Suite layout (copy 01’s pattern)

**Default:** one `TestBuildPetView` covering success, invalid name, and getter errors if the row shape stays uniform (`want` / `wantCode` / stub setup).

**Optional second suite** `TestBuildPetView_<Concern>` only if a concern needs its own ordering/fixtures (same bar as 01’s PetIDFormat — important invariant, not “errors go elsewhere”).

- MUST names = contracts. Fixtures for “several invalid names” use `examples` + plain nested `t.Run(in)` — **no** `MUST fail on empty id` style names (01 lesson).
- Invalid-name corpus can be a small `examples` slice under **one** MUST (“MUST return InvalidArgument when the pet name is invalid”) rather than one MUST per bad string.
- Do **not** put `MUST call Get exactly once` in the spec suite. Optional: stub `Fatal` on unexpected call as non-spec detail.

List final MUST strings in `functional/README.md` for 02b parity.

## Suggested contracts (stable names for 02b)

| MUST (normative) | Notes |
|------------------|--------|
| MUST call Get with the canonical pet name and return a PetView from the Get result | leaf args + view from stub data |
| MUST return InvalidArgument without calling Get when the pet name is invalid | boundary; `examples` of bad names |
| MUST return NotFound when Get reports not found | leaf error surfaces |
| MUST return Internal when Get fails for a non-not-found reason | leaf error mapped |

Prefer **one** `TestBuildPetView` when success and error rows share a uniform shape (e.g. `wantErr` / `wantCode` + same asserts) — as in 01’s shape suite. Split into another `TestBuildPetView_<Concern>` only for a **different invariant** (e.g. getter-error mapping that needs its own fixtures/order), not merely because some rows succeed and some fail. Split *tables* inside a suite if assertion shapes truly diverge (`if tc.wantX != nil` soup).

## Acceptance criteria

- [x] Code in `02-injected-dependency/functional/` under single `examples/go.mod`; `go test ./...` green
- [x] Narrow `PetGetter`; local hand stub; no counterfeiter generate in tree
- [x] Suite split / MUST-vs-`examples` follows [lessons-learned.md](./lessons-learned.md)
- [x] MUST names listed on parent README for 02b mirror
- [x] Uses `pet.NormalizePetName` (no duplicate name rules)
- [x] `examples/README.md` path index → `functional/` / `method/`
- [x] README: functional injection; stub = leaf stand-in

## Verification

```bash
cd .agents/skills/go-behavioral-spec/examples
go test ./02-injected-dependency/... -count=1
```

## Out of scope

Struct/method form (02b). PetStore CRUD (03). Discipline.md edits unless asked.

## Agent notes

Prefer outcome assertions. If tempted to add a MUST per stub call count or per invalid-name sample, re-read lessons-learned.
