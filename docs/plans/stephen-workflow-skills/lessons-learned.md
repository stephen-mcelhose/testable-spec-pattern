---
type: concept
title: "Lessons Learned — go-behavioral-spec examples (01–03)"
description: Distilled norms from shipping examples 01–03 covering form, normative-vs-fixture discipline, suite layout, and storage leaf contracts — applied to 04 and beyond.
tags: [go-behavioral-spec, lessons-learned, examples, discipline]
timestamp: 2026-10-23T22:45:43Z
---

# Lessons learned — go-behavioral-spec examples (from 01–03)

Side note for agents and humans implementing 04+. Distilled from shipping `01-pure-function`, `02-injected-dependency`, and `03-petstore-api`.

**Series intent:** each numbered example teaches a slice of the discipline for a workflow where **tests ≈ spec**. Examples are **mostly independent**; they may import types from earlier folders for convenience (KnownErrors, `Pet`), not to form a required layered product demo.

## Form

- **One `go.mod` at `examples/`**; packages in example folders. Directory names may not match `package` clauses (`pet` under `01-pure-function`). No per-example modules and no `replace` chains.
- **Agent lookup by path:** when two wirings share contracts (functional vs method injection), put each under its own subdir (`functional/`, `method/`) so an agent opens one folder. Prefer **self-contained suites per package** over a shared helper across packages (duplicate MUST tables are fine; agents should not chase imports to read the contract).
- **RPC / leaf lookup by file name:** flat package + `service_<method>_test.go` / `store_<method>_test.go` and `TestPetService_<Method>` is enough — no per-RPC subdirs unless suites get huge.
- Define **consumer-owned interfaces at the usage site** (`Store` next to `PetService`, not a separate `store.go` package of interfaces).
- Real `.go` / `_test.go`. Pattern READMEs may point at the package. Do not ship paste-only `.go.md` as the artifact.
- Keep example code **clean enough to copy** (boring, idiomatic). Janky char-loops undermine the example even if tests pass.
- Commit **examples separately** from discipline/skill doc churn when iterating the sample.
- Odd import paths are OK: `import pet "example.com/go-behavioral-spec/examples/01-pure-function"`.

## Normative vs fixture

- **MUST / MUST NOT / SHOULD / MAY** name **contracts only**.
- Fixtures that illustrate one rule go in an `examples` slice (or nested `t.Run(in)` with a **plain** label — the input string or a short descriptive tag).
- **Do not** write `MUST accept a single-character pet id` — that dresses an example as a requirement. The contract is `MUST accept when {pet} matches the id format` with `pets/a` as a fixture.

## Suite layout (what users should copy)

- **Own `TestFoo_<Invariant>`** (same `_test.go`) when an invariant is important and will grow (resource ID format, money scale, enums) — not because every axis needs a file.
- Kitchen-sink single `TestFoo` for unrelated invariants (e.g. shape + id format) teaches the wrong default; success and error rows for **one** concern can share a suite/table when the row shape stays uniform (`want` / `wantCode` / stub setup).
- Split suites when a **different invariant** deserves its own name/order (01: shape vs PetIDFormat). Split *tables* when assertion shapes truly diverge. Do **not** split merely because some rows succeed and some fail, and do **not** mint a MUST per regex branch or bad string.

## One rule, many fixtures

```go
{
    name: "MUST return InvalidArgument when {pet} does not match the id format",
    examples: []string{"pets/Fido", "pets/-fido", "pets/fido-"},
    wantInvalidArg: true,
},
```

One MUST; loop fixtures; nested subtest name = input (non-normative).

## Injected dependency (from 02)

- Spec the **unit under test**, including the leaf boundary: assert **presence and absence** of calls (e.g. Get with canonical name; Get not called when validation fails). Outcome-only rows miss the injection lesson.
- Prefer a short MUST that names both leaf args and the returned projection over a bare “MUST return PetView” plus a separate call-count MUST.
- Hand-roll a local stub that records call args. Comment that counterfeiter is preferred when the interface is package-owned and you want `make generate-fakes` — do not generate fakes in this example tree.
- Reuse typed errors from `01` (`KnownError`, `InvalidArgument`, `NotFound`, `Internal`, …); map leaf “other” failures to Internal.
- Method form may wrap the functional core; MUST **strings** stay identical across wirings.

## Resource method skeleton (from 03)

Every thin service RPC that validates then performs a store-backed operation tends toward the same contracts (spec-natural order). Name the **behaviour**; prove the leaf with stub asserts.

1. **MUST** return InvalidArgument when the request fails **shape validation** (resource name / required fields). Assert the leaf was not invoked.
2. **MUST** perform the operation and return the resulting resource for a valid request (create / get / update / delete / list). Assert the stub received validated / canonical args.
3. **MUST** return the appropriate KnownError when the operation cannot complete (domain codes + Internal for other).
4. **MAY** post-process on success — only if the method actually does; do not invent rows.

Do **not** title contracts `MUST call Store.<Op>…` — that is an implementation detail. Deep format corpora belong in pure helpers (01), not re-listed under every RPC.

When aligning to AIP (132–135, 154, 158), put a **claimed vs out of scope** table in the example README and only MUST what you claim (e.g. List coerces over-max `page_size`; Delete etag optional → `ABORTED` on mismatch).

### Who owns shape validation?

- The **RPC’s public contract** still includes step 1 (`CreatePet` MUST return InvalidArgument when the request fails shape). Callers care about `CreatePet`, not a Validator type.
- **Reusable / deep rules** (resource name format) → pure helper with its own suite (01). RPC suites only sample that the helper is wired (a few fixtures), not the full corpus.
- **Request-local checks** (required `display_name`, `pet_id`, page_size bounds) may live inline on the method. Update/Delete etag is optional (AIP-134/135). Extract `validateCreatePetRequest` when it grows or is shared.
- **Do not** introduce an injected `Validator` interface for teaching unless validation is truly swappable. That splits the skeleton across two specs without buying a leaf boundary.

## Storage leaf contract (for 04)

For a domain Store, **tests ≈ spec** means callers rely on returned resources and **KnownErrors**, not DB/driver mechanics.

- Spec **Store** methods with MUSTs that name KnownErrors (`AlreadyExists`, `NotFound`, `Aborted`, …) as the **caller-facing contract**.
- **Inject a generic DB** (usage-site interface: get/put/delete/…). Stub it in Store tests; assert call presence/absence like 02. No need for a real/map DB impl in the example — the interface + stub is enough.
- Store maps DB outcomes → KnownErrors; domain cases must not leak as raw driver strings. Opaque DB failure → `Internal` when claimed.
- Keep AIP/request shape out of the store suite. Assume valid inputs.
- Reuse KnownError types from 01 for ease; do not require wiring into 03’s `PetService`.

## Errors

- Assert typed errors with `errors.As` (KnownError / InvalidArgument), not string match alone.
- Prefer **outcome + leaf-boundary** contracts over bare call-count MUSTs (`MUST call Get exactly once`). Stub `t.Fatal` on unexpected call is optional non-spec detail, not a MUST name.
- One shared KnownError vocabulary in `01` (`package pet`); extend there when Create/Update need new codes (`AlreadyExists`, `Aborted`, …).

## Verify

```bash
cd .agents/skills/go-behavioral-spec/examples
go test ./... -count=1
```

## Pointers

- Pure function: `examples/01-pure-function/` (`package pet`)
- Injected dep: `examples/02-injected-dependency/functional/` and `…/method/`
- Resource CRUD: `examples/03-petstore-api/` — `service_<method>_test.go` / `TestPetService_<Method>`
- Store + DB: `examples/04-in-memory-store/` — KnownError Store contract; stub injected `DB`
- Discipline: `../../skills/go-behavioral-spec/references/discipline.md` (do not invent new discipline bullets mid-example unless Stephen asks)
