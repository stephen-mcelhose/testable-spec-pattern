---
type: use-case
title: Contracts & Domain Modeling
description: Use cases for producing machine-readable contracts and domain models — the formal artifacts that bridge requirements to implementation.
tags: [contracts, domain-modeling, proto, openapi, model-as-code]
timestamp: 2026-10-25T21:24:53Z
---

# Contracts & Domain Modeling

This phase produces two kinds of formal artifact from [Requirements](../requirements-gathering/requirements.md) or a raw [Idea](../ideas/index.md):

- **[Contracts](./contracts.md)** — machine-readable wire formats (proto, OpenAPI) defining resource shapes and service methods
- **[Domain Models](./domain-modeling.md)** — partial code artifacts expressing the domain in Go types, interfaces, and strategy functions

## Use Cases

### Build Contracts from Requirements

As a software engineer, I want to [Build Contracts](./contracts.md) from [Requirements](../requirements-gathering/requirements.md) in order to inform [Specifications](../specifications/index.md), [Planning](../planning/index.md), [Implementation](../implementation/index.md), and [Testing](../testing/index.md) of software.

### Create Proto Contracts from Requirements

As a software engineer, I want to [Create Proto Contracts](./proto-contracts.md) from [Requirements](../requirements-gathering/requirements.md) in order to define the wire surface of the service and inform code generation during [Implementation](../implementation/index.md).

### Create OpenAPI Contracts from Requirements

> [!WARNING]
> WILL NOT IMPLEMENT.

### Model the Domain in Go from an Idea

As a software engineer, I want to [Model the Domain in Go](./model-in-go.md) from an [Idea](../ideas/index.md) in order to express structure and intent as Go types before [Requirements](../requirements-gathering/requirements.md) are fully formed.

### Model the Domain in Go from Requirements

As a software engineer, I want to [Model the Domain in Go](./model-in-go.md) from [Requirements](../requirements-gathering/requirements.md) in order to refine the domain model and inform [Contracts](./contracts.md), [Specifications](../specifications/index.md), and [Implementation](../implementation/index.md).

> [!NOTE]
> A domain model is a *partial* code artifact — it expresses structure and intent but is not a full implementation. It is closer in nature to a contract than to a specification or implementation.
