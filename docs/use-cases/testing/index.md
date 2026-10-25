---
type: use-case
title: Testing
description: Use cases for the testing phase — verifying that an implementation satisfies its contracts and specifications through runnable specifications.
tags: [testing, runnable-spec, behavioral-spec, verification]
timestamp: 2026-10-25T23:20:03Z
---

# Testing

The testing phase verifies that an [Implementation](../implementation/index.md) satisfies its [Contracts](../contracts-and-domain-modeling/contracts.md) and [Specifications](../specifications/index.md). The primary artifact is a [Runnable Specification](../runnable-specifications/index.md) — Go tests named as RFC 2119 requirements.

## Use Cases

### Create a Runnable Specification from a Specification

As a software engineer, I want to [Create a Runnable Specification](../runnable-specifications/index.md) from a [Specification](../specifications/index.md) in order to verify that [Implementation](../implementation/index.md) satisfies its behavioral contracts.

### Create a Runnable Specification from an Implementation

As a software engineer, I want to [Create a Runnable Specification](../runnable-specifications/index.md) from an [Implementation](../implementation/index.md) in order to retroactively document and lock the behavior of existing code.

> [!NOTE]
> Testing in this workflow means behavioral verification via runnable specifications, not unit tests in isolation. A passing runnable specification is the definition of a verified project.
