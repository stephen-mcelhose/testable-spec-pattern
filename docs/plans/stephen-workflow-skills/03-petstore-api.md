---
type: example-plan
title: Canonical example — PetStore resource-oriented API
description: Plan for the AIP-aligned PetStore CRUD example teaching the four-step resource spec skeleton (shape → operate → failures → post-processing) with a stubbed Store leaf.
status: done
skill: go-behavioral-spec
depends_on: 01-pure-function + 02-injected-dependency (patterns); lessons-learned
tags: [go-behavioral-spec, examples, petstore, aip, resource-oriented, rfc-2119]
timestamp: 2026-10-23T22:45:43Z
---

# 03 — PetStore API (resource-oriented CRUD)

## Goal

Approachable, **runnable** mini PetStore showing behavioral specs for a resource-oriented service: narrow `Store` leaf, deep Create suite, thinner siblings for other methods. Same single `examples/go.mod` as 01–02 (no nested module).

Read [lessons-learned.md](./lessons-learned.md) before writing. Pattern refs: `examples/01-pure-function/`, `examples/02-injected-dependency/method/` (struct + injected leaf).

## Approach

AIP-flavoured surface (no proto/codegen). Prefer **method form** (struct field injection) — matches how real Go services look; agents who want function-arg wiring already have 02/`functional/`.

| Method | Depth |
|--------|--------|
| CreatePet | Full RFC 2119 suite (validation order, name/`pets/{id}`, store write, leaf boundary) |
| GetPet | Thinner suite (found / not found / invalid name skips Store) |
| UpdatePet | Light or sketched |
| DeletePet | Light or sketched |
| ListPets | Light or sketched |

```go
type Store interface {
    Create(ctx context.Context, pet Pet) (Pet, error)
    Get(ctx context.Context, name string) (Pet, error)
    Update(ctx context.Context, pet Pet, etag string) (Pet, error) // etag may live on Pet instead
    Delete(ctx context.Context, name string, etag string) error
    List(ctx context.Context, pageSize int32, pageToken string) ([]Pet, string, error)
}

type PetService struct {
    Store Store
}
```

### KnownError model (think through before coding)

**Server-only (before Store):** `InvalidArgument` — bad resource name / required fields / bad page size. Assert Store not called.

**From Store (or opaque → map):** domain codes below. Unknown/non-KnownError from Store → `Internal`.

| Method | Domain KnownErrors from Store | Notes |
|--------|-------------------------------|--------|
| **Create** | `AlreadyExists`, `Internal` | [AIP-133](https://google.aip.dev/133): `pet_id` + body; ignore `Pet.Name`; duplicate → AlreadyExists. |
| **Get** | `NotFound`, `Internal` | Missing resource → NotFound. No AlreadyExists. No etag. |
| **Update** | `NotFound`, `Aborted`, `Internal` | [AIP-134](https://google.aip.dev/134): name on resource; optional etag → Aborted ([AIP-154](https://google.aip.dev/154)). |
| **Delete** | `NotFound`, `Aborted`, `Internal` | [AIP-135](https://google.aip.dev/135): NotFound if absent; optional etag → Aborted. Cascade `FailedPrecondition` out of scope. |
| **List** | `Internal` | [AIP-132](https://google.aip.dev/132)/[158](https://google.aip.dev/158): default + coerce page_size; negative → InvalidArgument; empty page OK. |

**Reuse / extend types:** keep `InvalidArgument`, `NotFound`, `Internal` from 01. Add `AlreadyExists`, `Aborted`, and `FailedPrecondition` (cascade only / unused in suites).

**AIP claimed vs out of scope:** see `examples/03-petstore-api/README.md` tables for Create/Update/Delete.

**Do not invent for this example:** PermissionDenied, Unauthenticated, ResourceExhausted, etc.

Every RPC follows the same contract skeleton (spec-natural order). Word MUST names as **observable behaviour**, not leaf method calls (`Store.Create` is how the test proves it, not the contract title).

| Step | Keyword | Contract shape (fill in per method) |
|------|---------|-------------------------------------|
| 1 | **MUST** | return InvalidArgument when the request fails **shape validation** (resource name format, required fields — reuse `pet.NormalizePetName` for names; do not re-spec the full id corpus from 01). Assert the leaf was **not** invoked. |
| 2 | **MUST** | perform the resource operation and return the resulting Pet (or list) for a valid request — e.g. create / get / update / delete. Assert the stub received validated / canonical args. |
| 3 | **MUST** | return the method’s domain KnownErrors (table above) when the operation cannot complete; `Internal` for other Store failures. |
| 4 | **MAY** | apply post-processing on success (defaults, derived fields). Omit rows if the method has none — do not invent busywork. |

**Store is mocked in tests.** Ship the `Store` interface in production code; put a hand stub in `*_test.go` that returns canned Pet/error and records calls. No Store implementation outside tests (no in-memory map, no persistence package). Comment where counterfeiter would apply.

## Deliverable

```
examples/                         # existing go.mod — do NOT add 03/go.mod
  03-petstore-api/                # package petstore (or pets)
    README.md                     # what to copy (Create shape); what not to
    pet.go                        # Pet, CreatePetRequest, … (domain types)
    service.go                    # Store interface (usage site) + PetService + CRUD
    service_create_test.go        # deep Create spec; stub Store lives here (or stubs_test.go)
    service_get_test.go           # thinner Get spec
    # optional: update/delete/list tests with OPEN/MAY
```

No nested `petstore/` package unless it earns its keep — keep the example openable as one folder (same path-lookup lesson as 02).

Update `examples/README.md` index row from “planned” to the real path.

## Suite layout

**Create** teaches the skeleton in full; **Get** (and light siblings) copy the same four steps with fewer fixtures.

- One primary suite per method (`TestPetService_CreatePet`, `TestPetService_GetPet`, …) with an order comment that lists the skeleton steps.
- Prefer uniform rows (`want` / `wantCode` / `wantStoreCalls`) when possible; split *tables* only if assertion shapes diverge.
- Shape-validation fixtures → `examples` slice + plain nested labels; no MUST-per-bad-string. Deep id-format corpus stays in 01.
- Leaf boundary stays in **assertions** (no Store call on bad shape; stub args on happy path) — not in MUST titles (02 lesson, behaviour-first naming).

### Create (centerpiece — AIP-133)

| Contract | Example MUST string |
|----------|---------------------|
| Shape | `MUST return InvalidArgument when CreatePetRequest fails shape validation` |
| Operate | `MUST create the pet from pet_id and return the stored Pet` |
| AIP-133 | `MUST ignore Pet.Name on the create body and use pet_id for the resource name` |
| Failures | `MUST return AlreadyExists when the pet already exists`; Internal for other |

### Get (thinner — same skeleton)

| Contract | Example MUST string |
|----------|---------------------|
| Shape | `MUST return InvalidArgument when the pet name fails shape validation` |
| Operate | `MUST return the stored Pet for a valid pet name` |
| Failures | `MUST return NotFound when the pet does not exist`; Internal for other |

### Update / Delete / List (light — AIP-134 / AIP-135)

| Method | Failure MUST themes |
|--------|---------------------|
| Update | NotFound; Aborted when provided etag mismatches; etag optional; Internal |
| Delete | NotFound when absent; Aborted when provided etag mismatches; etag optional; Internal |
| List | InvalidArgument for negative page_size; default + coerce-above-max; empty page OK; Internal |

Implement enough to compile; spec at least one row per domain code you claim, or leave OPEN — do not invent a mega-table per RPC.

Prefer covering real paths over OPEN spam. Many validation paths → pure helper (01), not another MUST row here.

## Acceptance criteria

- [x] Code under `examples/03-petstore-api/` on the **shared** `examples/go.mod`
- [x] Create suite teaches the full skeleton (shape → operate → failures), with leaf asserts under the hood
- [x] Get + light Update/Delete/List suites use the same skeleton
- [x] Store is injected; interface in prod code; **mock/stub only in tests**; stub records calls
- [x] Reuses `pet.NormalizePetName`; KnownError set includes `AlreadyExists` + `Aborted` (+ `FailedPrecondition` type for cascade docs)
- [x] Error matrix respected (Create≠NotFound; Get≠AlreadyExists; Update/Delete etag → Aborted per AIP-135/154)
- [x] README: skeleton + agent lookup via `service_<method>_test.go`
- [x] No p360 imports; testify OK
- [x] File layout: one test file per method + `stubs_test.go`
- [x] `examples/README.md` index updated

**Lookup verdict:** flat `03-petstore-api/` + `service_<method>_test.go` / `TestPetService_<Method>` is enough — no per-RPC subdirs unless suites grow large.

## Verification

```bash
cd .agents/skills/go-behavioral-spec/examples
go test ./03-petstore-api/... -count=1
# or
go test ./... -count=1
```

## Out of scope

Any Store implementation outside tests (including in-memory maps). HTTP/gRPC servers, OpenAPI, counterfeiter generate, Helm, nested `go.mod`, SKILL workflow rewrite, re-teaching pure-function / injection wiring (point at 01 / 02).

## Agent notes

Keep Create approachable — the four skeleton steps, not a kitchen sink. Side-by-side with 02/`method/`: same stub-recording pattern; 03 adds resource shape + richer leaf errors. If tempted to nest another go.mod or duplicate NormalizePetName, stop and re-read lessons-learned.
