---
type: concept
title: Use Cases — Progressive Discovery
description: A use case format that expresses workflow steps as linked, chainable actor-goal-input-output statements. Each step's output becomes the next step's input.
tags: [use-cases, progressive-discovery, methodology, actors, personas]
timestamp: 2026-10-26T00:00:00Z
---

# Use Cases — Progressive Discovery

## What it is

Progressive discovery is a use case format where each step in a workflow is expressed as a linked, navigable statement:

> As a [**Actor**](./personas/software-engineer.md), I want to [**Verb Noun**](link-to-concept) from [**Input**](link-to-input) in order to inform [**Output**](link-to-output).

Every slot is a named, linked concept — not prose. This makes the workflow navigable: you can follow any link to understand what it means, where it comes from, or what it produces.

The "progressive" part is the chain: each use case's output becomes the next use case's input. Start from an idea, gather requirements, build contracts, write specifications, plan, implement, verify. Each step is explicit about what it needs and what it produces.

## Format

```
As a [Actor](personas/actor.md), I want to [Verb Noun](concept.md) from [Input](input.md) in order to inform [Output](output.md).
```

| Slot       | What it names                                           | Must be linked? |
| ---------- | ------------------------------------------------------- | --------------- |
| **Actor**  | Who performs the step — a named [persona](./personas/index.md) | Yes |
| **Verb Noun** | The action and its concept document                  | Yes             |
| **Input**  | What the actor starts from                              | Yes             |
| **Output** | What downstream phase this step informs                 | Yes             |

### Multiple inputs

When a step takes more than one input, list them with "and":

> As a [software engineer](./personas/software-engineer.md), I want to [Write a Specification](./use-cases/specifications/index.md) from [Requirements](./use-cases/requirements-gathering/requirements.md) and [Contracts](./use-cases/contracts-and-domain-modeling/contracts.md) in order to inform [Planning](./use-cases/planning/index.md), [Implementation](./use-cases/implementation/index.md), and [Testing](./use-cases/testing/index.md).

### Multiple outputs

When a step informs more than one downstream phase, list them comma-separated:

> ...in order to inform [Contracts & Domain Modeling](link), [Specifications](link), [Planning](link), [Implementation](link), and [Testing](link).

### Optional inputs

When the starting point varies (the workflow can begin from different places), express each variant as a separate use case — one per starting point. Each variant gets its own `###` heading.

## Decomposition

A high-level use case can be decomposed into sub-use-cases. The decomposition type must be explicit — readers need to know whether they must do all sub-cases or any one of them.

See [Requirements — Decomposition](./use-cases/requirements-gathering/requirements.md#decomposition) for the three patterns: **Summative**, **Composable**, and **Methods**.

Annotate the decomposition type at the top of each sub-use-case section with a `> [!NOTE]` so readers know which pattern applies at the point of use:

```markdown
### Write Use Cases from an Idea

> [!NOTE]
> Composable — satisfies [Gather Requirements](./index.md) independently. Combine with [Domain Types](./domain-types.md) or [Gherkin Scenarios](./gherkin-scenarios.md) for richer requirements.

As a [software engineer](../../personas/software-engineer.md), I want to...
```

## Document structure

A progressive discovery use case document uses:

- `##` headings for **phases** — the major stages of the workflow (Requirements Gathering, Contracts & Domain Modeling, etc.)
- `###` headings for **individual use cases** — one heading per use case statement, named after the action ("Gather Requirements from an Idea")
- A single use case statement per heading — if a heading contains "and", split it

```markdown
## Requirements Gathering

### Gather Requirements from an Idea

As a [software engineer](../personas/software-engineer.md), I want to [Gather Requirements](./index.md) from an [Idea](../ideas/index.md) in order to inform [Contracts & Domain Modeling](../contracts-and-domain-modeling/index.md), [Specifications](../specifications/index.md), [Planning](../planning/index.md), [Implementation](../implementation/index.md), and [Testing](../testing/index.md).
```

## How it differs from domain-use-cases

| Dimension         | Progressive Discovery                                    | Domain Use Cases                                        |
| ----------------- | -------------------------------------------------------- | ------------------------------------------------------- |
| **Focus**         | Workflow steps — what a practitioner does and in what order | Domain interactions — what actors do within a system boundary |
| **Chain**         | Explicit: each output is a next step's input             | Not inherently chained                                  |
| **Actor**         | A practitioner persona (engineer, lead, etc.)            | A domain actor (user, service, external system)         |
| **Output**        | An artifact that feeds the next phase                    | A system response or state change                       |
| **Granularity**   | One step per phase of the development workflow           | One scenario per observable system behavior             |
| **Navigation**    | Every slot is a linked concept                           | Narrative prose, sometimes linked                       |

Use progressive discovery to document **how engineers work**. Use domain use cases to document **how a system behaves**.

## Personas

Every actor in a progressive discovery use case is a [persona](./personas/index.md). A persona document defines who the actor is, what they produce, what they consume, and which use cases they own.

Current personas:

- [Software Engineer](./personas/software-engineer.md)
