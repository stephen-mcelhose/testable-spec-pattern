---
type: concept
title: Specification
description: A behavioral description of what a system must do — expressed in a form precise enough to inform implementation and verify correctness.
tags: [specification, rfc-2119, behavioral-spec, artifacts]
timestamp: 2026-10-26T00:00:00Z
---

# Specification

A specification is a behavioral description of what a system must do — expressed in a form precise enough to inform implementation and to serve as a basis for verification. A specification is distinct from a [Contract](../contracts-and-domain-modeling/contracts.md) (which defines shape) and from a [Runnable Specification](../testing/runnable-specification.md) (which verifies behaviour at runtime).

## Inputs

- [Requirements](../requirements-gathering/requirements.md)
- [Contracts](../contracts-and-domain-modeling/contracts.md) — optional; ground the specification in named domain types and resource shapes

## Outputs

A specification informs:

- [Planning](../planning/index.md)
- [Implementation](../implementation/index.md)
- [Testing](../testing/index.md) — the runnable specification is the executable form of this document

## Known forms

This workflow uses RFC 2119 prose specifications. Other forms are listed for context.

| Form                                                                                          | Notes                                               |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------- |
| [RFC 2119 prose spec](https://datatracker.ietf.org/doc/html/rfc2119)                        | **Used in this workflow.** MUST/SHOULD/MAY language |
| [IEEE 29148 / SRS](https://en.wikipedia.org/wiki/Software_requirements_specification)        | Formal; common in regulated industries              |
| [OpenAPI Specification](https://en.wikipedia.org/wiki/OpenAPI_Specification)                 | Machine-readable; HTTP API contracts                |
| [Gherkin / BDD](https://en.wikipedia.org/wiki/Behavior-driven_development)                   | Executable; Given/When/Then scenarios               |
| [C4 Model](https://c4model.com/)                                                             | Architecture-level; context, container, component   |
| [AsyncAPI](https://www.asyncapi.com/)                                                        | Machine-readable; event-driven API contracts        |

See [rfc2119.md](./rfc2119.md) for the format used in this workflow.
