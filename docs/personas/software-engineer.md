---
type: concept
title: Software Engineer
description: The primary actor in the use-cases-progressive-discovery workflow — responsible for the full lifecycle from idea to verified implementation.
tags: [persona, software-engineer, actor]
timestamp: 2026-10-26T00:00:00Z
---

# Software Engineer

## Who they are

A software engineer takes an idea from concept to verified implementation. They are comfortable working at every layer of the workflow — gathering requirements, modeling the domain, writing contracts and specifications, planning, implementing, and verifying through runnable specifications. In this workflow, a software engineer is the primary actor at every phase.

## Goals

- Turn a raw idea into a working, verified implementation
- Produce artifacts at each phase that make downstream phases easier
- Ensure that behavioral contracts are explicit, named, and testable

## Produces

| Artifact                                                                                       | Phase                     |
| ---------------------------------------------------------------------------------------------- | ------------------------- |
| [Requirements](../use-cases/requirements-gathering/requirements.md)                           | Requirements Gathering    |
| [Use Cases](../use-cases/requirements-gathering/use-cases.md)                                 | Requirements Gathering    |
| [Domain Types Sketch](../use-cases/requirements-gathering/domain-types.md)                    | Requirements Gathering    |
| [Gherkin Scenarios](../use-cases/requirements-gathering/gherkin-scenarios.md)                 | Requirements Gathering    |
| [Contracts](../use-cases/contracts-and-domain-modeling/contracts.md)                          | Contracts & Domain Modeling |
| [Domain Model](../use-cases/contracts-and-domain-modeling/domain-modeling.md)                 | Contracts & Domain Modeling |
| [Specification](../use-cases/specifications/index.md)                                         | Specifications            |
| [Plan](../use-cases/planning/index.md)                                                        | Planning                  |
| [Implementation](../use-cases/implementation/index.md)                                        | Implementation            |
| [Runnable Specification](../use-cases/runnable-specifications/index.md)                       | Testing                   |

## Consumes

| Input                                                                                          | From                      |
| ---------------------------------------------------------------------------------------------- | ------------------------- |
| [Idea](../use-cases/ideas/index.md)                                                           | External / prior session  |
| [Requirements](../use-cases/requirements-gathering/requirements.md)                           | Requirements Gathering    |
| [Contracts](../use-cases/contracts-and-domain-modeling/contracts.md)                          | Contracts & Domain Modeling |
| [Specification](../use-cases/specifications/index.md)                                         | Specifications            |
| [Plan](../use-cases/planning/index.md)                                                        | Planning                  |
| [Implementation](../use-cases/implementation/index.md)                                        | Implementation            |

## Use cases

All current use cases in the workflow are owned by the software engineer. See [Use Cases](../use-cases/index.md).
