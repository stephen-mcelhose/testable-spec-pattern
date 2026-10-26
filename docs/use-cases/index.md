---
type: use-case
title: Use Cases
description: The set of use cases that define a software engineer's workflow — from an idea through a verified, runnable implementation.
tags: [use-cases, workflow, software-engineering]
timestamp: 2026-10-26T00:00:00Z
---

# Use Cases

As a [Software Engineer](../personas/software-engineer.md), I want a [Process](./index.md) to follow to [Complete a Software Project](./index.md) from an [Idea](./ideas/idea.md) in order to produce a verified, runnable implementation.

The phases below are a **summative** decomposition of that use case — each phase is a required step. Entry can happen at any phase if prior artifacts already exist.

---

## Ideas

As a software engineer, I want to [Capture an Idea](./ideas/index.md) from a problem or opportunity in order to begin [Gathering Requirements](./requirements-gathering/index.md) and initiate the workflow.

## Requirements Gathering

As a software engineer, I want to [Gather Requirements](./requirements-gathering/index.md) from an [Idea](./ideas/idea.md) in order to inform [Contracts & Domain Modeling](./contracts-and-domain-modeling/index.md), [Specifications](./specifications/index.md), [Planning](./planning/index.md), [Implementation](./implementation/index.md), and [Testing](./testing/index.md) of software.

## Contracts & Domain Modeling

As a software engineer, I want to [Build Contracts](./contracts-and-domain-modeling/index.md) from [Requirements](./requirements-gathering/requirements.md) in order to inform [Specifications](./specifications/index.md), [Planning](./planning/index.md), [Implementation](./implementation/index.md), and [Testing](./testing/index.md) of software.

As a software engineer, I want to [Model the Domain](./contracts-and-domain-modeling/index.md) from an [Idea](./ideas/idea.md) in order to express structure and intent before Requirements are fully formed.

As a software engineer, I want to [Model the Domain](./contracts-and-domain-modeling/index.md) from [Requirements](./requirements-gathering/requirements.md) in order to refine the model and inform [Contracts](./contracts-and-domain-modeling/contracts.md), [Specifications](./specifications/index.md), and [Implementation](./implementation/index.md).

## Specifications

As a software engineer, I want to [Write a Specification](./specifications/index.md) from [Requirements](./requirements-gathering/requirements.md) in order to inform [Planning](./planning/index.md), [Implementation](./implementation/index.md), and [Testing](./testing/index.md) of software.

As a software engineer, I want to [Write a Specification](./specifications/index.md) from [Requirements](./requirements-gathering/requirements.md) and [Contracts](./contracts-and-domain-modeling/contracts.md) in order to inform [Planning](./planning/index.md), [Implementation](./implementation/index.md), and [Testing](./testing/index.md) of software.

## Planning

As a software engineer, I want to [Plan](./planning/index.md) a new feature from [Requirements](./requirements-gathering/requirements.md) in order to inform [Implementation](./implementation/index.md) of software.

As a software engineer, I want to [Plan](./planning/index.md) a new feature from [Requirements](./requirements-gathering/requirements.md) and [Contracts](./contracts-and-domain-modeling/contracts.md) in order to inform [Implementation](./implementation/index.md) of software.

## Implementation

As a software engineer, I want to [Implement](./implementation/index.md) a new feature from [Contracts](./contracts-and-domain-modeling/contracts.md) and [Specifications](./specifications/index.md) in order to produce a working project verified through [Testing](./testing/index.md).

As a software engineer, I want to [Implement](./implementation/index.md) a new feature from [Requirements](./requirements-gathering/requirements.md), [Contracts](./contracts-and-domain-modeling/contracts.md), and [Specifications](./specifications/index.md) in order to produce a working project verified through [Testing](./testing/index.md).

## Testing

As a software engineer, I want to [Create a Runnable Specification](./testing/runnable-specification.md) from a [Specification](./specifications/specification.md) in order to verify that [Implementation](./implementation/index.md) satisfies its behavioral contracts.

As a software engineer, I want to [Create a Runnable Specification](./testing/runnable-specification.md) from an [Implementation](./implementation/index.md) in order to retroactively document and lock the behavior of existing code.
