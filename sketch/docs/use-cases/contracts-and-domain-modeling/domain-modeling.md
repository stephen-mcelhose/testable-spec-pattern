---
type: concept
title: Domain Modeling
description: Expressing domain understanding as a partial Go code artifact — types, interfaces, and strategy functions that make the structure of a system explicit before implementation.
tags: [domain-modeling, model-as-code, artifacts, go]
timestamp: 2026-10-25T21:24:53Z
---

# Domain Modeling

Domain modeling is the activity of expressing domain understanding in Go — using types, interfaces, and strategy functions as the modeling primitives. The output is a **partial code artifact**: it expresses structure and intent but is not a runnable implementation.

## Why "Partial"

A domain model is contract-level, not implementation-level. It:

- Names the things in the domain and their relationships
- Makes entry points and data flows explicit
- Leaves leaf functions as `ErrUnimplemented` stubs

This is the same discipline as the `internal/modeling` package in this repo.

## Inputs

- An Idea — the model can be built directly from a raw concept
- [Requirements](../requirements-gathering/requirements.md) — the model can be refined once requirements are structured

## Outputs

- A `CodeArtifact` — a Go package that acts as a partial implementation, consumed by:
  - [Contracts](./contracts.md) — the model informs contract shapes
  - [Implementation](../implementation/index.md) — the model scaffolds the implementation
