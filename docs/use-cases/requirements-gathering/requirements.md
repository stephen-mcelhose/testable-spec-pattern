---
type: concept
title: Requirements
description: What a requirement is — any artifact that sufficiently communicates what needs to be built to inform the downstream phases of software development.
tags: [requirements, artifacts, decomposition]
timestamp: 2026-10-26T00:00:00Z
---

# Requirements

A requirement is any artifact that sufficiently communicates what needs to be built to inform the downstream phases of software development. Requirements are not defined by their format — they are defined by what they enable.

This means that [Use Cases](./use-cases.md), [Domain Type Sketches](./domain-types.md), [Gherkin Scenarios](./gherkin-scenarios.md), user stories, or a whiteboard photo are all requirements — as long as they communicate enough for the next phase to proceed. The form is a choice; the sufficiency is the constraint.

## Inputs

- An [Idea](../ideas/index.md) — a product concept, feature request, domain description, or problem statement

## Outputs

Requirements are consumed by:

- [Contracts & Domain Modeling](../contracts-and-domain-modeling/index.md)
- [Specifications](../specifications/index.md)
- [Planning](../planning/index.md)
- [Implementation](../implementation/index.md)
- [Testing](../testing/index.md)

## Decomposition

Requirements gathering can be decomposed in three ways. The decomposition type is always explicit — it tells the reader whether they must do all sub-cases or any one of them.

### Summative (A + B + C = Requirements)

Sub-artifacts are dependent steps that together constitute requirements. All are required. Omitting one leaves a gap.

Use when sub-artifacts cover non-overlapping concerns and each is necessary for the downstream phase to proceed.

### Composable (A = Requirements, B = Requirements; any combination = Requirements)

Each sub-artifact independently satisfies the requirements goal. Any single one is sufficient. Combining them produces richer requirements, but none is mandatory.

> [Use Cases](./use-cases.md) = Requirements ✓  
> [Domain Types](./domain-types.md) = Requirements ✓  
> [Gherkin Scenarios](./gherkin-scenarios.md) = Requirements ✓  
> Any combination = Requirements ✓

This is the correct decomposition for requirements gathering in this workflow: use cases, domain type sketches, and Gherkin scenarios are each independently valid requirements.

### Methods (A or B achieves Requirements)

Sub-artifacts are mutually exclusive approaches — pick one. Combining them would be redundant or contradictory.

Use when sub-artifacts are genuinely alternative techniques where choosing one precludes the others.

## Known forms

| Form                                         | When to reach for it                                              |
| -------------------------------------------- | ----------------------------------------------------------------- |
| [Use Cases](./use-cases.md)                  | Actors, goals, and interactions are the primary concern           |
| [Domain Types](./domain-types.md)            | Named entities and relationships are the primary concern          |
| [Gherkin Scenarios](./gherkin-scenarios.md)  | Concrete Given/When/Then behaviours are the primary concern       |
| User Stories                                 | Lightweight; good for single-actor, single-goal slices            |
| [SRS](https://en.wikipedia.org/wiki/Software_requirements_specification) | Formal; good for regulated or large-scale contexts |
