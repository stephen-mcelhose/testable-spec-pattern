---
type: concept
title: "Test Discipline: Behavioral Specification Style"
description: The discipline for sustaining RFC 2119-style behavioral spec tests — the why and the harder questions that the skill alone cannot answer.
tags: [go, behavioral-spec, rfc-2119, testing, discipline]
timestamp: 2026-10-23T22:45:43Z
---

# Test Discipline: Behavioral Specification Style


This document captures the disciplines required to sustain RFC 2119-style behavioral spec tests. The `go-behavioral-spec` skill handles the *how* — this document covers the *why* and the harder questions that skill can't answer for you.

---

## The core covenant

**Test names are specifications, not descriptions.**

A description says what the code does. A specification says what the code *must* do. That distinction sounds philosophical but has a concrete consequence: when you change the implementation, a description-named test still passes silently even if the behaviour is now wrong. A specification-named test fails visibly — or forces you to consciously update the requirement.

```
// description — can silently become untrue:
"skips items that already have CommitOnlyReasonError set"

// specification — explicitly contracts the requirement:
"MUST skip items that are commit-only (Observable and ObservableErr remain nil)"
```

When a test fails, the requirement-statement name tells you *what contract was violated*, not just *what code broke*. That is the entire point.

---

## RFC 2119 keyword discipline

The keywords are not decorative. Use them precisely or the spec degrades.

| Keyword    | Contract                                                              | Common misuse to avoid                                |
| ---------- | --------------------------------------------------------------------- | ----------------------------------------------------- |
| `MUST`     | Required. This test verifies it. There is no exception.               | Using MUST when the behaviour is actually conditional |
| `MUST NOT` | Required to be absent. This test verifies the absence.                | Writing `MUST NOT` without asserting `Nil`/`Zero`     |
| `SHOULD`   | Expected but not yet fully verifiable — a conscious deferral.         | Using SHOULD to soften a MUST you're uncertain about  |
| `MAY`      | Optional. May or may not be present depending on inputs.              | Using MAY for things that are actually MUST            |

`SHOULD` is **not** a softer MUST. It specifically means: this requirement is real and correct, but this test cannot fully verify it yet (missing fixture, external side effect, known gap). Mark the reason in the skip body.

If you catch yourself writing `SHOULD` because you're uncertain whether the requirement is real — stop. Clarify the requirement first, then write the test.

---

## Ordering is a decision, not a default

Every test function must choose one of two orderings before the first test is written:

**Spec-natural:** Happy path first, then precondition skips, then field-by-field behaviour, then error paths, then gaps. Reads like a spec — the reader sees the normal case before edge cases. Best when spec-ing a new interface from scratch.

**Code-flow:** Mirror the function's branches top-to-bottom. The reader who has the source open can follow along. Best when adding coverage to an existing function, or when the function has a strong sequential structure.

Document the choice:

```go
// TestFetchObservablesHandler covers NewFetchObservablesHandler.
// Order: precondition skips → happy-path fetch → error capture → mixed batch.
func TestFetchObservablesHandler(t *testing.T) {
```

Why this matters: an unordered test suite looks complete but is actually arbitrary. When the next engineer adds a case, they don't know where it belongs. Over time the suite becomes unsearchable. An explicit ordering gives every new case a natural landing spot.

---

## Tables as a design signal

We prefer table-driven tests. This is not just a style preference — it is a feedback mechanism on code design.

A function is naturally table-drivable when: the same code path runs for all cases, inputs are data not procedures, outputs are deterministic given inputs, and any observable side effect is controlled through a single injected dependency with a uniform stub. When you can express every case as a row with the same assertion shape, the table is the right tool.

When a function *resists* table-driving — setup varies wildly per case, assertions need `if tc.wantFoo != nil` conditionals, each case needs a different stub configuration — that resistance is a design signal: the function is mixing concerns, or side effects are entangled with logic that should be pure.

**The principle: push testable side effects to leaves, and expose those leaves through narrow injected interfaces.**

A testable side effect is any observable behaviour other than logging: writing to a database, calling an API, emitting an event, mutating shared state. Logging is explicitly excluded — we do not test log output.

Pushing side effects to a leaf is only half the work. If the leaf is a concrete type, the core still cannot be unit-tested without real infrastructure. The interface is what makes the boundary injectable — the core depends on an abstraction, the test injects a stub, production injects the real implementation.

**Narrow interfaces matter.** Defining the interface at the boundary forces you to articulate exactly what the core needs from the leaf: not "a Spanner client" but `Write(records) error`. That narrowness is what makes stubs trivial, keeps the core's dependencies honest, and prevents the core from reaching into infrastructure concerns it shouldn't know about.

Code structured for specification-as-tests tends toward:

- **Core functions that are nearly pure:** they take a data structure, apply logic, return a new data structure. Any leaf dependency is injected as a narrow interface and controlled by a uniform stub in tests. These are naturally table-drivable.
- **Leaf functions where side effects live:** storage writes, API calls, event emission. They implement the narrow interface; they have few enough code paths that tables rarely add value. Tested with integration tests or stub-captured argument verification.

This is a design discipline, not just a test discipline. If you find yourself reaching for conditional assertions inside a table row, ask: has a side effect leaked out of a leaf? If so, the fix is in the code — extract the pure transformation, leave the I/O call in the leaf behind its interface. The table preference surfaces this continuously, without requiring a separate design review.

**When one table is the wrong shape**

A table shows the full case space of one assertion shape. When cases need divergent shapes — `if tc.wantFoo != nil`, different stub setups, different success vs error checks — more than one specification is sharing a costume.

Prefer splitting by assertion shape into separate table-driven suites (e.g. success vs error). Use flat `t.Run`s when the branches are not uniform enough for a table. Prefer tables when a group of cases still shares one shape — see preference above.

One optional nil-check does not automatically require a split; several or most cases needing different assertion shapes does.

---

## Spec the unit whose behaviour you are dictating

The covenant is about **contracts**, not about export visibility. Spec the function (or logical unit) whose rules you are writing down. Do not restating a leaf's rules on its caller.

| Spec target | Test package | Why |
| ----------- | ------------ | --- |
| Exported handler / API entry point | `package foo_test` preferred | Callers depend on this surface; the external package is a cheap fence against private fields |
| Unexported helper with real behavioural rules | `package foo` + short why-comment | The helper *is* the unit of behaviour; forcing every case through the caller bloats the caller suite or under-specs the leaf |
| Either | Assert outcomes on that unit's contract | Call counts and private wiring are still implementation details unless the contract is about that call |

`package foo_test` is a useful default for exported entry points — you cannot reach unexported state, which pushes assertions toward observable results. It is not a universal law. White-box specs for unexported helpers are first-class when those helpers own the rules (e.g. audit-time selection, parsing fallbacks). Prefer an internal suite over bleeding leaf cases into the caller's table.

Anti-pattern: only testing the fat orchestrator, then either (a) omitting leaf contracts or (b) exporting helpers solely so `foo_test` can see them. Spec the leaf where it lives.

White-box access does not turn incidental internals into RFC 2119 requirements. If it isn't part of the unit's contract, don't assert it as MUST/SHOULD/MAY. A separate, non-spec test *may* document an implementation detail in plain language when that's useful (coverage, regression, wiring example). Alternatively, if it IS a part of the unit's contract, assert it as MUST/SHOULD/MAY.

---

## One rule per test

If a test name contains "and", it tests two requirements. Split it.

This is harder than it sounds because it forces you to clarify what you're actually specifying. "MUST set Observable and MUST advance CommitableOffset" looks like one test but it's two contracts that can fail independently.

The benefit: when the test fails, you know exactly *which* contract was violated. When a combined test fails, you have to inspect assertion output to figure out which part broke.

---

## Presence AND absence

For every optional field or conditional side effect, write both directions:

```go
{ name: "non-empty AuditURI MUST annotate spec", ... },
{ name: "empty AuditURI MUST NOT add annotation", ... },
```

The absence test is easy to skip — "obviously if the field is empty, nothing happens." But that's exactly when it matters. If someone adds code that accidentally writes an empty annotation for missing URIs, only the absence test catches it. Without it, the spec is asymmetric: you've specified that the field *can* be set, but not that it *only* gets set when appropriate.

---

## Maintaining specs as code evolves

A spec test has two failure modes:

1. The test fails — the contract was violated. Fix the implementation.
2. The test passes but the name is wrong — the contract changed without updating the spec. **Worse than a failure because it's invisible.**

When you change a function's behaviour, check the test names, not just the test results. If a `MUST` statement is now untrue (even for a good reason), update the test name to reflect the new contract. A passing test with a wrong name actively misleads the next engineer.

---

## Anti-patterns

**Descriptive names in a spec suite.** In an RFC 2119 suite, `"skips delete events"` is ambiguous — is this a MUST or incidental? Prefer a keyword so the contract is explicit. Outside the spec (coverage, regression, wiring examples), descriptive names without keywords are preferred — do not dress incidentals in MUST/SHOULD/MAY.

**Testing implementation instead of contract (in a spec suite).** `"calls GetLabMethod with the correct observable name"` asserts *how* the function works, not *what it produces*. In a behavioral spec that freezes the contract to an incidental call shape — if the implementation changes (e.g., batches the call), the test breaks even if the contract is intact. A separate non-spec test may document call wiring in plain language when that is useful; do not elevate it to MUST/SHOULD/MAY.

**Silent gaps.** Prefer covering the path. If a contract is known but not yet verifiable, you MAY mark it with `t.Skip("OPEN: …")` (or SHOULD + skip body). Do not add OPEN skips for every uncovered line or for paths you intentionally won't chase. Many paths are a prompt to ask whether this unit owns too much or a leaf should own some rules — path count alone is not a smell.

**Tables for orthogonal cases.** Two cases that exercise completely different code paths (or need different assertion shapes) stuffed into the same table to "reduce boilerplate." Split into separate table suites, or use flat tests when a table no longer fits — see above.

**MUST on things you haven't verified.** Writing `MUST` to sound authoritative when you're not sure the behaviour is actually required. If you're not sure, write `SHOULD` and document why, or clarify the requirement first.

---

## Upgrade path for existing descriptive tests

Do not upgrade tests in place. Write a new spec test suite for the function — a fresh `Test<Foo>` (or a new `_test.go` file if the old one is dense) that covers the interface properly with RFC 2119 names and a chosen ordering. Once the spec suite exists and passes, evaluate whether the old descriptive tests add anything. If they duplicate coverage without adding clarity, retire them. If they catch something the spec missed, fold that case into the spec first, then retire.

This approach avoids the half-upgraded file problem — a mix of MUST names and descriptive names in the same test function, with no coherent ordering, is harder to read than either style alone.

---

## File layout

**Stubs and fakes.** Prefer counterfeiter fakes (`*fakes/`, `make generate-fakes`) for package interfaces that already have them or warrant generation (call counts, argument capture, shared across packages). Hand-roll a local stub in the `_test.go` that uses it when a tiny test-only stand-in is enough. Extract a shared hand stub to `<pkg>_stubs_test.go` (or similar) only when that same stub is used across multiple test files — do not hand-roll a parallel fake for an interface that already has counterfeiter coverage.

**Test files.** Default: one `*_test.go` per source file (`foo.go` → `foo_test.go`; white-box helpers MAY use `foo_internal_test.go`). For a multi-method type (e.g. `PetService`), keep all `TestPetService_<Method>` suites in `service_test.go` — lookup by function name, not `service_<method>_test.go`. You MAY collapse several source files into one test file when that grouping is also one logical spec (same contracts, same ordering, same fixtures). Do not merge unrelated specs just to reduce file count. Do not split one file per method unless the suite is genuinely huge.
