---
type: concept
title: Runnable Specification
description: The executable form of a specification — Go tests named as RFC 2119 requirements whose passing run constitutes proof that an implementation satisfies its behavioral contracts.
tags: [runnable-spec, testing, go-test, rfc-2119, behavioral-spec]
timestamp: 2026-10-26T00:00:00Z
---

# Runnable Specification

A runnable specification is the executable form of a [Specification](../specifications/specification.md) — Go tests named as RFC 2119 requirements. A passing run means the [Implementation](../implementation/index.md) satisfies its [Contracts](../contracts-and-domain-modeling/contracts.md). **A passing runnable specification is the definition of done.**

## Inputs

- A [Specification](../specifications/specification.md) — the behavioral contracts to verify
- Or an [Implementation](../implementation/index.md) — when writing retroactively to lock existing behavior

## Outputs

- A passing test run — the verification that implementation satisfies specification

## Method

See [Go Runnable Specification](./go/runnable-spec.md) for the specific format and discipline used in this workflow.

## Key properties

A runnable specification:

- Names each test case as a behavioral requirement, not a description (`MUST skip items that are commit-only`, not `skips commit-only items`)
- Covers all meaningful specable behaviors including both directions of every conditional behavior
- Records known gaps as named skips (`t.Skip("OPEN: ...")`) rather than silently omitting them
- Uses table-driven tests when multiple cases share the same assertion shape
- Uses a declared ordering (happy path → errors → gaps, or code-flow ordering)
