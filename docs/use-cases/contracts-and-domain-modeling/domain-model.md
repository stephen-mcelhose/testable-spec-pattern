---
type: concept
title: Domain Model
description: A named representation of the entities, relationships, and behaviours in a problem domain — expressed in a form that makes structure explicit before implementation.
tags: [domain-model, model-as-code, artifacts]
timestamp: 2026-10-26T00:00:00Z
---

# Domain Model

A domain model is a named representation of the entities, relationships, and behaviours in a problem domain. It makes structure and intent explicit before implementation begins. The output is a **partial artifact**: it expresses structure but is not a runnable implementation.

## Known forms

| Form                                                                                          | When to reach for it                                               |
| --------------------------------------------------------------------------------------------- | ------------------------------------------------------------------ |
| [Model in Go](./model-in-go.md)                                                               | Go types and interfaces as the modeling language (**preferred**)   |
| [ERD](https://en.wikipedia.org/wiki/Entity%E2%80%93relationship_model)                       | Data-centric; entities, attributes, relationships                  |
| [UML Class Diagram](https://en.wikipedia.org/wiki/Class_diagram)                             | Language-neutral OO structure                                      |
| [Context Map (DDD)](https://en.wikipedia.org/wiki/Domain-driven_design)                      | Bounded context relationships across a large system                |
| [Event Storming output](https://en.wikipedia.org/wiki/Event_storming)                        | Collaboratively discovered aggregates and domain events            |

## Why "Partial"

A domain model is contract-level, not implementation-level. It:

- Names the things in the domain and their relationships
- Makes entry points and data flows explicit
- Leaves leaf functions as `ErrUnimplemented` stubs

## Inputs

- An [Idea](../ideas/idea.md) — the model can be built directly from a raw concept
- [Requirements](../requirements-gathering/requirements.md) — the model can be refined once requirements are structured

## Outputs

A domain model is consumed by:

- [Contracts](./contracts.md) — the model informs contract shapes
- [Implementation](../implementation/index.md) — the model scaffolds the implementation
