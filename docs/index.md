---
type: concept
title: Docs Index
description: Structured catalog of all docs/ knowledge base pages — updated on every write operation.
tags: [index, docs]
timestamp: 2026-10-23T22:52:54Z
---

# Docs Index

| Page                                                                                                 | Type           | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | -------------- | ---------------------------------------------------------------------------------------------------- |
| [Docs Conventions](./conventions.md)                                                                                 | concept        | Canonical conventions for the docs/ knowledge base — folder taxonomy, OKF frontmatter, naming.      |
| [ADR-001: Project Structure](./decisions/001-project-structure.md)                                                  | decision       | Establishes the top-level layout — skills in .agents/skills/, docs as a wiki, examples/ standalone modules. |
| [ADR-002: GitHub Issues as Work Tracker](./decisions/002-github-issues-as-work-tracker.md)                         | decision       | Use GitHub issues to track skill-building work across sessions.                                              |
| [ADR-003: Kanban Workflow via GitHub Issues](./decisions/003-kanban-workflow.md)                                    | decision       | Board state, Definition of Ready, session protocol, and WIP limit — applies to humans and agents equally.    |
| [ADR-004: Use internal/ not pkg/ for Go packages](./decisions/004-internal-over-pkg.md)                            | decision       | Why go-project-scaffold prescribes internal/ and rejects pkg/; links to the community dispute (issue #117). |
| [Test Discipline: Behavioral Specification Style](./discipline.md)                                   | concept        | The discipline for sustaining RFC 2119-style behavioral spec tests.                                  |
| [Project: Testable Spec Pattern](./ideation/task.md)                                                 | design-doc     | Scope and goals for the Go behavioral-spec workflow project.                                         |
| [Domain Names as the Bridge](./ideation/domain-names-as-bridge.md)                                  | proposal       | Hypothesis that domain names are a necessary precondition for requirement-traceable testable specs.  |
| [Contract-First as the Foundation](./ideation/contract-first.md)                                    | proposal       | Hypothesis that contract-first (OpenAPI, proto, ROD) naturally produces testable spec targets.       |
| [Composable Skills Architecture](./research/composable-skills.md)                                   | spike          | Research into Matt Pocock's composable skills patterns and how to apply them here.                   |
| [Plans: go-behavioral-spec canonical examples](./plans/stephen-workflow-skills/README.md)            | playbook       | Execution order and shared decisions for the go-behavioral-spec example series.                      |
| [Lessons Learned — go-behavioral-spec examples (01–03)](./plans/stephen-workflow-skills/lessons-learned.md) | concept | Distilled norms from examples 01–03 for agents and humans implementing future examples.              |
| [01 — Pure function behavioral spec](./plans/stephen-workflow-skills/01-pure-function.md)            | example-plan   | Plan for the NormalizePetName example — the series baseline.                                        |
| [02a — Injected dependency (functional)](./plans/stephen-workflow-skills/02-injected-dependency-functional.md) | example-plan | Plan for the functional-injection example (BuildPetView).                               |
| [02b — Injected dependency (method / struct)](./plans/stephen-workflow-skills/02-injected-dependency-method.md) | example-plan | Plan for the struct-method injection example (PetViewer).                              |
| [03 — PetStore resource-oriented API](./plans/stephen-workflow-skills/03-petstore-api.md)            | example-plan   | Plan for the AIP-aligned PetStore CRUD example.                                                     |
| [04 — Store over injected DB](./plans/stephen-workflow-skills/04-in-memory-store.md)                 | example-plan   | Plan for the storage-leaf KnownError contract example.                                              |
