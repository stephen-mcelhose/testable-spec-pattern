---
type: concept
title: Plan
description: An ordered implementation approach — the artifact that bridges specifications and contracts to working code.
tags: [plan, planning, artifacts]
timestamp: 2026-10-26T00:00:00Z
---

# Plan

> [!WARNING]
> Incomplete. No planning skill exists yet. The format of a plan in this workflow is not yet settled. This doc captures the concept and known alternatives; the preferred form will be determined when a planning skill is built.

A plan is an ordered implementation approach — it bridges [Specifications](../specifications/specification.md) and [Contracts](../contracts-and-domain-modeling/contracts.md) to working code. It answers: what will be built, in what order, and with what approach?

## Inputs

- [Requirements](../requirements-gathering/requirements.md)
- [Contracts](../contracts-and-domain-modeling/contracts.md)
- [Specification](../specifications/specification.md)

## Outputs

A plan informs:

- [Implementation](../implementation/index.md)

## Known forms

| Form                                                                                  | Notes                                               |
| ------------------------------------------------------------------------------------- | --------------------------------------------------- |
| Implementation plan (this workflow)                                                   | Ordered tasks with rationale; the preferred form    |
| [Scrum Sprint Plan](https://en.wikipedia.org/wiki/Scrum_(software_development))      | Iteration-scoped; story points and sprint goals     |
| [Technology Roadmap](https://en.wikipedia.org/wiki/Technology_roadmap)               | Long-horizon; milestone-based                       |
| [Gantt Chart](https://en.wikipedia.org/wiki/Gantt_chart)                             | Time-boxed; dependency-aware                        |
| [Kanban Board](https://en.wikipedia.org/wiki/Kanban_board)                           | Flow-based; pull system                             |
