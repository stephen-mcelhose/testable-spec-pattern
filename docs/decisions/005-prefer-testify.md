---
type: decision
title: "ADR-005: Prefer testify for assertions"
description: Why we default to testify/assert + testify/require over the standard library, and when to revisit.
tags: [adr, go, testing, testify]
timestamp: 2026-10-25T00:00:00Z
---

# ADR-005: Prefer testify for assertions

## Status

Accepted — with a named revisit trigger (see below).

## Context

When writing Go tests we need an assertion style. The two real options are:

**Standard library (`testing.T`)** — hand-roll comparisons with `if` statements and `t.Errorf` / `t.Fatalf`. It ships with Go, adds no dependency, and is what the Go standard library itself uses.

**testify** (`github.com/stretchr/testify`) — provides `assert` and `require` sub-packages with readable failure messages, deep equality, and a `require` variant that calls `t.FailNow()` immediately.

The standard library argument has real weight: zero dependencies, no import graph, and Go's own codebase as a proof point. But the ergonomic gap matters in practice:

| Concern                        | Standard library                          | testify                                      |
| ------------------------------ | ----------------------------------------- | -------------------------------------------- |
| Failure message quality        | Whatever you write in `t.Errorf`          | Diffs values automatically                   |
| Stop-on-failure                | `t.Fatal` — easy to forget                | `require.*` — explicit and named             |
| Deep equality                  | `reflect.DeepEqual` + manual message      | `assert.Equal` — prints diff                 |
| Table-driven test readability  | Verbose per-assertion boilerplate         | One line per assertion                        |
| Dependency                     | None                                      | One well-maintained module                   |

The cost of testify is a single, stable, widely-adopted dependency. The benefit is tests that read as documentation without spending tokens on assertion boilerplate.

## Decision

Default to **testify** (`assert` + `require`) for all assertion code in this repo and in projects scaffolded by the `go-project-scaffold` skill.

Conventions:

- Use `require.*` when a failure should stop the test immediately (setup, preconditions, single-step tests).
- Use `assert.*` when the test should continue to surface all failures in one run (table-driven rows, multi-step validations).
- Do not mix hand-rolled `if/t.Errorf` assertions with testify in the same file — pick one style per file.

## Rationale

The `go-behavioral-spec` skill targets tests that read like RFC 2119 contracts. Assertion boilerplate is noise in that reading. testify's failure diffs also produce better signal when a spec test fails — you see *what* diverged, not just *that* it did.

## Revisit trigger

**Revisit this decision if the Go standard library gains built-in assertion support** — either via a first-party `testing/assert` package or a significant improvement to `t.Errorf` ergonomics — to the point where testify's advantage over stdlib is marginal. At that point, evaluate whether the zero-dependency benefit outweighs the remaining ergonomic gap.

Until that trigger fires, the default is testify.

## Consequences

- The `go-project-scaffold` skill adds `github.com/stretchr/testify` as a default test dependency.
- New tests in this repo use `assert` or `require` — not hand-rolled `if/t.Errorf` blocks.
- If the standard library reaches parity, a follow-up ADR supersedes this one.
