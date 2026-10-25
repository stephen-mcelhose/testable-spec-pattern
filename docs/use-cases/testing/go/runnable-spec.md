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

**In scope:** test naming, coverage completeness, ordering, example style, and gap handling.

**Out of scope:** Go tooling configuration, CI setup, and test infrastructure.

## 2. Terminology

The key words "MUST", "MUST NOT", "REQUIRED", "SHOULD", "SHOULD NOT", "RECOMMENDED", "MAY", and "OPTIONAL" in this document are to be interpreted as described in [RFC 2119](https://datatracker.ietf.org/doc/html/rfc2119).

**Specable behavior** — an observable, externally verifiable behavior of the unit under test; something that could fail a black-box test.  
**Conditional behavior** — a behavior that differs depending on the presence, absence, or value of an input.  
**Known gap** — a specable behavior that is identified but not yet verifiable.

## 3. Behavioral Requirements

### 3.1 Coverage

1. A runnable specification MUST specify all meaningful specable behaviors of the unit under test.
2. For every conditional behavior, a runnable specification MUST specify both directions: the behavior when the condition is met and the behavior when it is not.

### 3.2 Naming

3. Every `t.Run` name MUST include a capitalized RFC 2119 keyword (`MUST`, `MUST NOT`, `SHOULD`, `MAY`).
4. Each `t.Run` MUST encode exactly one rule. If a name requires the word "and", it MUST be split into separate cases.

### 3.3 Ordering

5. A runnable specification MUST use a prescribed ordering — the ordering MUST be stated in a comment on the enclosing test function.
6. A runnable specification MAY use specification ordering (happy path → skips → errors → gaps).
7. A runnable specification MAY use code-flow ordering (mirroring the implementation top-to-bottom).

### 3.4 Examples

8. A runnable specification MAY include illustrative examples.
9. Example cases MUST NOT use RFC 2119 keywords in their names.

### 3.5 Known Gaps

10. A specable behavior that cannot yet be verified SHOULD be recorded as a named skip: `t.Skip("OPEN: <reason>")`.
11. A known gap MUST NOT be silently omitted — omission is indistinguishable from ignorance.

## 4. Data Models / Interfaces

A conforming test function has this shape:

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
