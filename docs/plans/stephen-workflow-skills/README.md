---
type: playbook
title: "Plans: go-behavioral-spec canonical examples"
description: Execution order, shared decisions, and handoff notes for the go-behavioral-spec canonical example series (01–04).
tags: [go-behavioral-spec, examples, planning, workflow]
timestamp: 2026-10-23T22:45:43Z
---

# Plans: go-behavioral-spec canonical examples

**Parent:** `.agents/skills/go-behavioral-spec/`  
**Branch context:** `stephen-workflow-skills` (discipline + skill gathering)  
**Status:** 01–04 done. Decide per plan: work here vs hand to another agent.

## Order

| # | Plan | Artifact form | Depends on |
|---|------|---------------|------------|
| 1 | [01-pure-function.md](./01-pure-function.md) | Nested runnable — **done** | — |
| 2 | [02-injected-dependency-functional.md](./02-injected-dependency-functional.md) | `examples/02-injected-dependency/functional/` — **done** | 01 + [lessons-learned](./lessons-learned.md) |
| 3 | [02-injected-dependency-method.md](./02-injected-dependency-method.md) | `…/method/` (wraps functional) — **done** | 02-functional (same contracts) |
| 4 | [03-petstore-api.md](./03-petstore-api.md) | `examples/03-petstore-api/` — **done** | 01–02 + [lessons-learned](./lessons-learned.md) |
| 5 | [04-in-memory-store.md](./04-in-memory-store.md) | `examples/04-in-memory-store/` — **done** | lessons-learned; types from 01 for ease |

## Side notes

| Doc | Role |
|-----|------|
| [lessons-learned.md](./lessons-learned.md) | Normative vs fixture, suite split, AIP claimed/out-of-scope, KnownError leaf contract — apply when implementing 04+ |

## Shared decisions (locked)

- Examples live under `.agents/skills/go-behavioral-spec/examples/`.
- One `go.mod` at `examples/`; packages in example dirs (`01-pure-function` as `package pet`). Odd import paths are OK.
- Injected-dep wiring is **path-routed**: `02-injected-dependency/functional/` vs `…/method/` (not one flat package).
- PetStore for 03 (approachable CRUD) on the **same** `examples/go.mod`; method-form server + Store leaf; reuse `pet` from 01.
- Examples are mostly independent (tests ≈ spec discipline); reuse types across folders for convenience.
- 04 specs a Store over an **injected generic DB**: KnownErrors are the Store↔caller contract; DB is stubbed in tests.
- Do **not** overhaul `SKILL.md` into a workflow yet — gathering.
- Follow `references/discipline.md` and [lessons-learned.md](./lessons-learned.md).

## Handoff

Each plan is self-contained enough for a fresh agent. Give them: this README + **lessons-learned.md** + the one plan file + `discipline.md` + `SKILL.md` + path to `examples/01-pure-function/` as the pattern reference. One plan in flight at a time unless explicitly parallelizing 02a/02b after 01.
