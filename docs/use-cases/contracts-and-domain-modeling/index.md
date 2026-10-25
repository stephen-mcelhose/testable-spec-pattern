---
type: use-case
title: Contracts & Domain Modeling
description: Use cases for producing machine-readable contracts and domain models — the formal artifacts that bridge requirements to implementation.
tags: [contracts, domain-modeling, proto, openapi, model-as-code]
timestamp: 2026-10-25T21:24:53Z
---

# Contracts & Domain Modeling

This phase produces two kinds of formal artifact from [Requirements](../requirements-gathering/requirements.md) or a raw Idea:

- **[Contracts](./contracts.md)** — machine-readable wire formats (proto, OpenAPI) defining resource shapes and service methods
- **[Domain Models](./domain-modeling.md)** — partial code artifacts expressing the domain in Go types, interfaces, and strategy functions

## Use Cases

- [Create Contracts](./contracts.md) from Requirements
- [Model the Domain](./domain-modeling.md) from an Idea — producing a partial code artifact
- [Model the Domain](./domain-modeling.md) from Requirements — refining the model against gathered requirements

> [!NOTE]
> A domain model is a *partial* code artifact — it expresses structure and intent but is not a full implementation. It is closer in nature to a contract than to a specification or implementation.
