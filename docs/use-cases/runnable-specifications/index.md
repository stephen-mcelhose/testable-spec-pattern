---
type: use-case
title: Runnable Specifications
description: Use cases for creating runnable specifications — Go tests named as RFC 2119 requirements that verify an implementation satisfies its contracts.
tags: [runnable-spec, testing, go-test, rfc-2119, behavioral-spec]
timestamp: 2026-10-25T21:24:53Z
---

# Runnable Specifications

A runnable specification is the executable form of a [Specification](../specifications/index.md) — Go tests named as RFC 2119 requirements. A passing run means the implementation satisfies its contracts.

## Use Cases

- Create Runnable Specifications from [Specifications](../specifications/index.md)
- Create Runnable Specifications from [Implementations](../implementation/index.md) — deriving the spec retroactively from working code

## What Makes a Specification "Runnable"

- Test functions named as RFC 2119 requirements (e.g. `Test_It_MUST_...`)
- Grounded in domain names established during requirements and modeling
- A passing run is the definition of a verified project

> [!NOTE]
> The `go-behavioral-spec` skill is the primary tool for producing runnable specifications in this workflow.
