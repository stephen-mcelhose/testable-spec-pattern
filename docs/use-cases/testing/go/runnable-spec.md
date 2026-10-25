---
type: specification
title: Go Runnable Specification
description: Normative RFC 2119 specification for a Go runnable specification — what a behavioral test file must satisfy to qualify as a runnable specification in this workflow.
tags: [runnable-spec, go, rfc-2119, behavioral-spec, testing]
timestamp: 2026-10-25T23:20:03Z
---

# Go Runnable Specification

## 1. Introduction & Scope

This specification defines what a Go test file must satisfy to be a **runnable specification** in this workflow — a test file whose names read as behavioral contracts and whose passing run constitutes proof that an implementation satisfies its [Specification](../../specifications/index.md).

**In scope:** test naming, coverage completeness, ordering, example style, table-driven tests, and gap handling.

**Out of scope:** Go tooling configuration, CI setup, and test infrastructure.

## 2. Terminology

The key words "MUST", "MUST NOT", "REQUIRED", "SHOULD", "SHOULD NOT", "RECOMMENDED", "MAY", and "OPTIONAL" in this document are to be interpreted as described in [RFC 2119](https://datatracker.ietf.org/doc/html/rfc2119).

**Specable behavior** — an observable, externally verifiable behavior of the unit under test; something that could fail a black-box test.  
**Conditional behavior** — a behavior that differs depending on the presence, absence, or value of an input.  
**Known gap** — a specable behavior that is identified but not yet verifiable.  
**Assertion shape** — the set of fields asserted and the structure of the assertion block for a given case.

## 3. Behavioral Requirements

### 3.1 Coverage

1. A runnable specification MUST specify all meaningful specable behaviors of the unit under test.
2. For every conditional behavior, a runnable specification MUST specify both directions: the behavior when the condition is met and the behavior when it is not.

### 3.2 Naming

3. Each `t.Run` MUST encode exactly one rule. If a name requires the word "and", it MUST be split into separate cases.

### 3.3 Ordering

4. A runnable specification MUST use a prescribed ordering — the ordering MUST be stated in a comment on the enclosing test function.
5. A runnable specification MAY use specification ordering (happy path → skips → errors → gaps).
6. A runnable specification MAY default to code-flow ordering (mirroring the implementation top-to-bottom).

### 3.4 Examples

7. A runnable specification MAY include illustrative examples.
8. Example cases MUST NOT use RFC 2119 keywords in their names.

### 3.5 Known Gaps

9. A specable behavior that cannot yet be verified SHOULD be recorded as a named skip: `t.Skip("OPEN: <reason>")`.
10. A known gap MUST NOT be silently omitted — omission is indistinguishable from ignorance.

### 3.6 Table-Driven Tests

11. A runnable specification SHOULD use table-driven tests when multiple cases share the same assertion shape.
12. When cases require divergent assertion shapes, they MUST be split into separate table suites or flat `t.Run` blocks — they MUST NOT be forced into a single table.
13. A single optional `nil`-check does not require a split; cases where most rows need different assertion structures do.

## 4. Data Models / Interfaces

A conforming test function:

```go
func TestFoo(t *testing.T) {
    t.Parallel()
    // Order: happy path → errors → gaps

    cases := []struct {
        name    string
        // inputs ...
        // expected ...
    }{
        {name: "MUST ..."},
        {name: "MUST NOT ..."},
        {name: "SHOULD ... (skipped if not verifiable)"},
    }

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            t.Parallel()
            // assert ...
        })
    }
}
```

A known gap:

```go
t.Run("OPEN: Foo MUST ... when ...", func(t *testing.T) {
    t.Skip("OPEN: not yet implemented — remove when Foo handles X")
})
```

## 5. Discipline

### Test names are specifications, not descriptions

A description says what the code does. A specification says what the code *must* do. When you change an implementation, a description-named test can pass silently even if the behavior is now wrong. A specification-named test fails visibly — or forces a conscious update to the requirement.

```
// description — can silently become untrue:
"skips items that already have CommitOnlyReasonError set"

// specification — explicitly contracts the requirement:
"MUST skip items that are commit-only (Observable and ObservableErr remain nil)"
```

When a test fails, a requirement-statement name tells you *what contract was violated*, not just *what code broke*.

### RFC 2119 keyword precision

The keywords are not decorative. Each one has a specific meaning:

| Keyword    | Meaning                                                                                                             |
| ---------- | ------------------------------------------------------------------------------------------------------------------- |
| `MUST`     | Required. This test verifies it. No exception.                                                                      |
| `MUST NOT` | Required to be absent. This test verifies the absence.                                                              |
| `SHOULD`   | Real requirement, not yet fully verifiable — a conscious deferral. Always include a skip body explaining why.       |
| `MAY`      | Optional. May or may not be present depending on inputs.                                                            |

`SHOULD` is not a softer `MUST`. If you write `SHOULD` because you are uncertain whether the requirement is real — stop. Clarify the requirement first.

### Table-driven tests as a design signal

Preferring table-driven tests is not just a style choice — it is a feedback mechanism on code design. A function is naturally table-drivable when inputs are data, outputs are deterministic, and any observable side effect is controlled through a single injected dependency with a uniform stub.

When a function *resists* table-driving — setup varies wildly, assertions need `if tc.wantFoo != nil` conditionals, each case needs a different stub — that resistance is a design signal: the function is mixing concerns, or side effects are entangled with logic that should be pure.

**The principle: push testable side effects to leaves and expose those leaves through narrow injected interfaces.** Code structured for runnable specifications tends toward core functions that are nearly pure (naturally table-drivable) and leaf functions where side effects live (tested with integration tests or stub-captured argument verification).

### Maintaining specs as code evolves

A spec test has two failure modes:

1. The test fails — the contract was violated. Fix the implementation.
2. The test passes but the name is wrong — the contract changed without updating the spec. **Worse than a failure because it is invisible.**

When you change a function's behavior, check the test names, not just the test results. A passing test with a wrong name actively misleads the next engineer.

### Anti-patterns

**Testing implementation instead of contract.** Asserting *how* the function works (`"calls GetLabMethod with the correct observable name"`) rather than *what it produces* freezes an incidental call shape as a contract. If the implementation changes, the test breaks even if the contract is intact.

**`MUST` on things you have not verified.** Writing `MUST` to sound authoritative when the behavior is not confirmed. Use `SHOULD` with a skip body, or clarify the requirement first.

**Descriptive names in a behavioral suite.** In an RFC 2119 suite, `"skips delete events"` is ambiguous — is this a `MUST` or incidental? Outside the behavioral spec (coverage, regression, wiring examples), descriptive names without keywords are correct — do not dress incidentals in `MUST`/`SHOULD`/`MAY`.
