---
type: example-plan
title: Canonical example — injected dependency (method / struct)
status: done
skill: go-behavioral-spec
depends_on: 02-injected-dependency-functional
---

# 02b — Injected dependency (method on struct)

## Goal

**Same behavioural contracts as 02a**, wired as a struct field + method. Readers should see that only construction differs; MUST names and suite split stay aligned with functional/.

Read [lessons-learned.md](./lessons-learned.md) and complete 02a first.

## Approach

```go
type PetViewer struct {
    Getter PetGetter // same interface as 02a
}

func (v *PetViewer) BuildPetView(ctx context.Context, name string) (PetView, error)
```

Shared types live in `functional/` (`Pet`, `PetView`, `PetGetter`). `method` imports `functional` and wraps `BuildPetView` — no forked contracts.

Setup difference only: `PetViewer{Getter: stub}` vs `BuildPetView(ctx, name, stub)`.

## Deliverable

```
examples/02-injected-dependency/
  README.md                    # MUST parity + path index
  functional/                  # 02a
  method/                      # package method — open this for struct wiring
    pet_viewer.go
    pet_viewer_test.go         # same MUST strings; self-contained suite
```

Same `examples/go.mod`. Path-based agent lookup: open `method/`, not the parent alone.

## Suite layout

Mirror 02a in one suite (`TestPetViewer_BuildPetView`) with the same four MUST strings. Self-contained table (do not share a helper across packages — agents should read one folder).

- Copy MUST **strings** from parent README parity list — do not rename for style.
- Same `examples` fixture approach; plain nested labels; no new normative cases.
- Skip nil-receiver / “struct zero value” unless it is a real product contract (prefer omit for parity).

## Acceptance criteria

- [x] Same interface + same MUST names as 02a
- [x] Method receiver; dependency on struct field; wraps `functional.BuildPetView`
- [x] lessons-learned: no fixture MUSTs; leaf-boundary focused
- [x] `go test ./...` green from module root
- [x] Index / parent README path-routed to `functional/` and `method/`

## Verification

```bash
cd .agents/skills/go-behavioral-spec/examples
go test ./02-injected-dependency/... -count=1
```

Side-by-side skim: MUST strings in functional vs method tests match.

## Out of scope

New domain behaviour. Counterfeiter. PetStore (03). Divergent error mapping from 02a.

## Agent notes

If 02a is missing or MUST list is unclear, stop — do not invent divergent contracts.
