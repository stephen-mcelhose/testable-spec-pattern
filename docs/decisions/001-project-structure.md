---
type: decision
title: "ADR-001: Project Structure"
description: Establishes the top-level layout of the repo — skills in .agents/skills/, docs as a wiki, and examples as standalone Go modules under examples/.
tags: [adr, structure, skills, examples, docs]
timestamp: 2026-10-23T23:10:06Z
---

# ADR-001: Project Structure

## Status

Accepted

## Y-Statement

In the context of building a Go behavioral-spec workshop with agent skills and runnable examples, facing the need to separate skills, documentation, and executable code without a prescribed standard, we decided to use `.agents/skills/{skill-name}/` for workflow skills, `docs/` as an OKF wiki, and `examples/{example-name}/` as standalone Go modules, to achieve clear separation of concerns and discoverability, accepting that examples live outside the skills directory (breaking from the initial arrangement where they lived under `go-behavioral-spec/examples/`).

## Context

The repo has three distinct kinds of artefact:

- **Agent skills** — markdown-based instructions that shape how an agent behaves (`SKILL.md` + optional bundled resources)
- **Documentation** — a growing knowledge base of discipline, research, ideation, and decisions
- **Runnable examples** — Go code that demonstrates the behavioral-spec pattern in practice

Early in the project these were mixed together: examples lived inside the skill directory, and docs lived inside `.agents/plans/`. This created friction — the examples are code, not skill metadata, and the docs belong to the project, not to a specific skill.

## Decision

### Skills

```
.agents/skills/
  {skill-name}/
    SKILL.md
```

One folder per skill, named directly (no `discipline/` or `orchestrator/` subdirectory layer — that breaks csgdaa-code's skill loader). The orchestrator/discipline distinction is expressed through the skill's `description` and `user-invocable` frontmatter field.

### Docs

```
docs/
  conventions.md        ← wiki rules; read by docs-wiki skill
  discipline.md         ← behavioral-spec discipline
  index.md              ← catalog of all pages
  log.md                ← append-only operations log
  decisions/
    001-project-structure.md
  ideation/
    contract-first.md
    domain-names-as-bridge.md
    task.md
  plans/
    stephen-workflow-skills/
      README.md
      lessons-learned.md
      01-pure-function.md
      02-injected-dependency-functional.md
      02-injected-dependency-method.md
      03-petstore-api.md
      04-in-memory-store.md
  research/
    agentskills-standard.md
    composable-skills.md
```

### Examples

```
examples/
  {example-name}/
    go.mod
    go.sum
    …
```

Each example is a standalone Go module. This keeps CI simple (each `go test ./...` is scoped to one example), avoids cross-example dependency pollution, and makes each example self-contained for readers.

First example: `examples/petstore/`.

## Consequences

- Examples must be moved from `.agents/skills/go-behavioral-spec/examples/` to `examples/` — the CI workflow and any relative links in the skill will need updating.
- The `go-behavioral-spec` skill's `examples/` references (in `SKILL.md` and `plan.md`) must be updated to point to `examples/{name}/`.
- New examples get their own `go.mod` — no shared module.
- The `docs/` wiki is the canonical home for all non-code artefact; nothing lives in `.agents/plans/` anymore.
