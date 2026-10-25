---
type: concept
title: Contracts
description: Formal definitions of what a system must provide — covering both wire contracts (proto, OpenAPI) and code contracts (Go interfaces at layer boundaries).
tags: [contracts, proto, openapi, interfaces, clean-architecture, artifacts]
timestamp: 2026-10-25T21:24:53Z
---

# Contracts

A contract is a formal definition of *what* something must provide — without specifying *how*. Contracts are enforced at boundaries: between processes (wire contracts) or between layers within a process (code contracts).

## Two Kinds of Contract

### Wire Contracts

Contracts that cross process boundaries — machine-readable, enforced at runtime:

- **Proto files** — gRPC service and resource definitions (AIP-compliant)
- **OpenAPI specs** — REST API definitions
- **ROD resource definitions** — Resource Oriented Design resource shapes

### Code Contracts

Contracts that cross layer boundaries within a process — expressed in Go, enforced at compile time:

- **Go interfaces** — define what a dependency must do without specifying its implementation
- **Layered interfaces** (clean architecture style) — each architectural layer defines the interface its outer layer must satisfy; nothing leaks inward
- **Strategy function signatures** — functions as contracts for pluggable behaviour

A [Domain Model](./domain-modeling.md) is a set of code contracts: it names the domain's interfaces and strategy functions, leaving implementations as stubs. The model *is* the contract expressed in Go.

## What Contracts Are Not

Contracts define shape and obligation — not behaviour. How the system behaves under those contracts is captured in [Specifications](../specifications/index.md).

## Inputs

- [Requirements](../requirements-gathering/requirements.md)
- Or a [Domain Model](./domain-modeling.md) — the model surfaces the code contracts that inform wire contract design

## Outputs

- Contracts consumed by:
  - [Specifications](../specifications/index.md)
  - [Design](../design/index.md)
  - [Implementation](../implementation/index.md)
  - [Runnable Specifications](../runnable-specifications/index.md)
