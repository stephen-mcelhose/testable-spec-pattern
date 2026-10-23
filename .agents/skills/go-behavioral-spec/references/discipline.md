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

**When flat `t.Run`s are correct**

Use flat tests when each subtest exercises a genuinely different code path — not different inputs to the same path, but different branches of behaviour. The distinction: a table shows the full case space of one path; flat tests show the independent branches of a function's structure. If you'd need a conditional in the table assertion, the cases are different paths and belong as flat tests.

Anti-pattern: a table where half the cases need `if tc.wantFoo != nil { ... }`. That's flat tests wearing a table costume. Pull them out.

---

## The public interface principle

Test the public interface, not the implementation.

External test packages (`package foo_test`) enforce this mechanically — you literally cannot access unexported fields. This is a feature. If you find yourself needing to reach into unexported state to assert something, ask: is that state actually observable through the public interface? If not, the test is asserting an implementation detail that can change without breaking the contract.

Internal test packages (`package foo`, white-box) are sometimes legitimate — for pure functions that are unexported but too complex to skip. When you do this, add a comment explaining why the public interface is insufficient.

---

## One rule per test

If a test name contains "and", it tests two requirements. Split it.

This is harder than it sounds because it forces you to clarify what you're actually specifying. "MUST set Observable and MUST advance CommitableOffset" looks like one test but it's two contracts that can fail independently.

The benefit: when the test fails, you know exactly *which* contract was violated. When a combined test fails, you have to inspect assertion output to figure out which part broke.

---

## Presence AND absence

For every optional field or conditional side effect, write both directions:

```go
t.Run("non-empty AuditURI MUST annotate spec", ...)
t.Run("empty AuditURI MUST NOT add annotation", ...)
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

**Descriptive names without keywords.** `"skips delete events"` — is this a MUST or a MAY? Is skipping required or incidental? The reader cannot tell without reading the code.

**Testing implementation instead of contract.** `"calls GetLabMethod with the correct observable name"` — this asserts *how* the function works, not *what it produces*. If the implementation changes (e.g., batches the call), the test breaks even if the contract is intact.

**Silent gaps.** A function has 10 conditional paths; the test covers 4. Use `t.Skip("OPEN: ...")` for any known gap so the absence is explicit and tracked.

**Tables for orthogonal cases.** Two cases that exercise completely different code paths stuffed into the same table to "reduce boilerplate." The table obscures the fact that these are separate specifications.

**MUST on things you haven't verified.** Writing `MUST` to sound authoritative when you're not sure the behaviour is actually required. If you're not sure, write `SHOULD` and document why, or clarify the requirement first.

---

## Upgrade path for existing descriptive tests

Do not upgrade tests in place. Write a new spec test suite for the function — a fresh `Test<Foo>` (or a new `_test.go` file if the old one is dense) that covers the interface properly with RFC 2119 names and a chosen ordering. Once the spec suite exists and passes, evaluate whether the old descriptive tests add anything. If they duplicate coverage without adding clarity, retire them. If they catch something the spec missed, fold that case into the spec first, then retire.

This approach avoids the half-upgraded file problem — a mix of MUST names and descriptive names in the same test function, with no coherent ordering, is harder to read than either style alone.

---

## Open questions

These require team agreement, not just individual skill:

1. **Stub placement.** Stubs inline in `_test.go` vs a dedicated `<pkg>_stubs_test.go` (as in the pipeline packages). What's the rule for when a stub warrants its own file?

2. **One test file per source file, or per logical grouping?** When a source file has multiple exported things, do the tests all live in `<file>_test.go`, or do they split by concern?
