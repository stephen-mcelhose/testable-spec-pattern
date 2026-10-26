---
type: use-case
title: Requirements Gathering
description: Use cases for the requirements gathering phase — converting an idea into structured requirements through use cases, domain sketches, and Gherkin scenarios.
tags: [requirements, use-cases, gherkin, domain-sketch]
timestamp: 2026-10-25T21:24:53Z
---

# Requirements Gathering

The requirements gathering phase takes a raw [Idea](../ideas/index.md) and produces [Requirements](./requirements.md) — structured enough to inform contracts, domain modeling, and specifications.

## Use Cases

### Gather Requirements from an Idea

As a software engineer, I want to [Gather Requirements](./index.md) from an [Idea](../ideas/index.md) in order to inform [Contracts & Domain Modeling](../contracts-and-domain-modeling/index.md), [Specifications](../specifications/index.md), [Planning](../planning/index.md), [Implementation](../implementation/index.md), and [Testing](../testing/index.md) of software.

### Write Use Cases from an Idea

> [!NOTE]
> Composable — satisfies [Gather Requirements](#gather-requirements-from-an-idea) independently. Combine with [Domain Types](./domain-types.md) or [Gherkin Scenarios](./gherkin-scenarios.md) for richer requirements.

As a software engineer, I want to [Write Use Cases](./use-cases.md) from an [Idea](../ideas/index.md) in order to surface the actors, goals, and interactions that inform the downstream phases of software development.

### Sketch Domain Types from an Idea and Use Cases

> [!NOTE]
> Composable — satisfies [Gather Requirements](#gather-requirements-from-an-idea) independently. Combine with [Use Cases](./use-cases.md) or [Gherkin Scenarios](./gherkin-scenarios.md) for richer requirements.

As a software engineer, I want to [Sketch Domain Types](./domain-types.md) from an [Idea](../ideas/index.md) and [Use Cases](./use-cases.md) in order to name the entities and relationships that ground [Contracts & Domain Modeling](../contracts-and-domain-modeling/index.md) and [Specifications](../specifications/index.md).

### Write Gherkin Scenarios from Use Cases and a Domain Sketch

> [!NOTE]
> Composable — satisfies [Gather Requirements](#gather-requirements-from-an-idea) independently. Combine with [Use Cases](./use-cases.md) or [Domain Types](./domain-types.md) for richer requirements.

As a software engineer, I want to [Write Gherkin Scenarios](./gherkin-scenarios.md) from [Use Cases](./use-cases.md) and a [Domain Sketch](./domain-types.md) in order to produce concrete Given/When/Then scenarios that inform [Specifications](../specifications/index.md) and [Testing](../testing/index.md).

> [!NOTE]
> The workflow does not have to start here. If requirements already exist from a prior session or external source, proceed directly to [Contracts & Domain Modeling](../contracts-and-domain-modeling/index.md) or [Specifications](../specifications/index.md).
