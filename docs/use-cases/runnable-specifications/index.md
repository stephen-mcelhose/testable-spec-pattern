---
type: use-case
title: Runnable Specifications
description: Use cases for creating runnable specifications — Go tests named as RFC 2119 requirements that verify an implementation satisfies its contracts.
tags: [runnable-spec, testing, go-test, rfc-2119, behavioral-spec]
timestamp: 2026-10-25T21:24:53Z
---

# Runnable Specifications

A runnable specification is the executable form of a [Specification](../specifications/index.md) — Go tests named as RFC 2119 requirements. A passing run means the [Implementation](../implementation/index.md) satisfies its [Contracts](../contracts-and-domain-modeling/contracts.md).

## Use Cases

### Create a Runnable Specification from a Specification

As a software engineer, I want to [Create a Runnable Specification](../runnable-specifications/index.md) from a [Specification](../specifications/index.md) in order to verify that [Implementation](../implementation/index.md) satisfies its behavioral contracts.

### Create a Runnable Specification from an Implementation

As a software engineer, I want to [Create a Runnable Specification](../runnable-specifications/index.md) from an [Implementation](../implementation/index.md) in order to retroactively document and lock the behavior of existing code.

> [!NOTE]
> The `go-behavioral-spec` skill is the primary tool for producing runnable specifications in this workflow. A passing runnable specification is the definition of a verified project.
