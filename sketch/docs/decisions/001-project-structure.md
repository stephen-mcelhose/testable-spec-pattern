---
type: decision
title: "ADR-001: Project Structure"
description: Establishes the top-level folder layout and required files for the repo.
tags: [adr, structure, skills, examples, docs]
timestamp: 2026-10-23T23:10:06Z
---

# ADR-001: Project Structure

## Status

Accepted

## Y-Statement

In the context of building a Go behavioral-spec workshop with agent skills and runnable examples, facing the need to separate skills, documentation, and executable code, we decided on the folder layout below to achieve clear separation of concerns and discoverability, accepting that examples are standalone modules rather than a shared one.

## Decision

### Folders

```
.agents/skills/{skill-name}/    ← one folder per skill; flat (no subdirectory layer)
docs/                           ← wiki; managed by docs-wiki skill
examples/{example-name}/        ← standalone Go modules (own go.mod / go.sum)
AGENTS.md                       ← agent steering document
README.md                       ← human-facing overview
```

### Required docs/ files

| File                   | Purpose                                      |
| ---------------------- | -------------------------------------------- |
| `docs/conventions.md`  | Wiki rules; read by docs-wiki skill first    |
| `docs/index.md`        | Catalog of all pages; updated on every write |
| `docs/log.md`          | Append-only operations log                   |

## Consequences

- Examples must be self-contained Go modules — no cross-example imports.
- New skills go directly under `.agents/skills/{name}/` — no category subdirectory.
