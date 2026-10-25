---
type: concept
title: Requirements
description: The structured output of requirements gathering — the set of actor goals, domain entities, and behaviours that define what needs to be built.
tags: [requirements, artifacts]
timestamp: 2026-10-25T21:24:53Z
---

# Requirements

Requirements are the structured output of the [requirements gathering phase](./index.md). They are produced from a raw Idea and consumed by every downstream phase.

## What Requirements Are

Requirements capture *what* needs to be built and *for whom* — not how. They are:

- **Use cases** — actor–goal interactions with the system
- **Domain sketches** — named entities and their relations
- **Gherkin scenarios** — Given/When/Then behaviours grounded in domain names

These are views of requirements, not a fixed composition. The strategy determines which views are produced.

## Inputs

- An Idea — a product concept, feature request, domain description, or user story

## Outputs

- A set of Requirements consumed by:
  - [Contracts & Domain Modeling](../contracts-and-domain-modeling/index.md)
  - [Specifications](../specifications/index.md)
  - [Design](../design/index.md)
  - [Implementation](../implementation/index.md)
