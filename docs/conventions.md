---
type: concept
title: Docs Conventions
description: Canonical conventions for the docs/ knowledge base — folder taxonomy, OKF frontmatter rules, naming, and cross-reference patterns.
tags: [conventions, docs, okf, wiki]
timestamp: 2026-10-23T22:52:54Z
---

# Docs Conventions

This file is the authoritative reference for how the `docs/` knowledge base is structured. The `docs-wiki` skill reads it at the start of every operation.

---

## Docs root

All knowledge base files live under `docs/`. Paths below are relative to `docs/`.

---

## Folder taxonomy

| Folder              | OKF types that go here                     | When to use                                                    |
| ------------------- | ------------------------------------------ | -------------------------------------------------------------- |
| `docs/` (root)      | `concept`, `how-to`                        | Cross-cutting discipline, conventions, and foundational guides |
| `docs/ideation/`    | `proposal`, `design-doc`                   | Hypotheses, ideas, and early-stage designs not yet in discipline |
| `docs/research/`    | `spike`                                    | Time-boxed investigations with findings and recommendations    |
| `docs/plans/`       | `playbook`, `example-plan`, `design-doc`   | Execution plans, worked-example plans, and implementation guides |
| `docs/decisions/`   | `decision`                                 | Architectural and design decisions (ADR-style)                 |
| `docs/how-to/`      | `how-to`, `runbook`                        | Step-by-step guides and operational runbooks                   |

**Rule:** when in doubt, use the folder that matches the OKF type. If no folder fits, place in `docs/` root and note it in the log.

---

## OKF frontmatter

Every file in `docs/` MUST begin with an OKF frontmatter block:

```yaml
---
type: <see table below>
title: "<concise display title>"
description: <one sentence — what this concept is and why it matters>
resource: <URL or path to canonical source — omit if none>
tags: [<relevant>, <keywords>]
timestamp: <ISO-8601 UTC — run `date -u +%Y-%m-%dT%H:%M:%SZ`>
---
```

**Required fields:** `type`, `title`, `description`, `timestamp`  
**Optional fields:** `resource` (set when a canonical source exists), `tags` (set when 2+ meaningful keywords apply)

---

## OKF types used in this repo

| Type           | Meaning                                                                 |
| -------------- | ----------------------------------------------------------------------- |
| `concept`      | An explanation of a domain term, principle, or discipline rule          |
| `proposal`     | An idea or hypothesis floated for exploration — not yet discipline      |
| `design-doc`   | A detailed technical or structural design                               |
| `decision`     | An ADR-style architectural or design decision with rationale            |
| `spike`        | Time-boxed research findings and recommendations                        |
| `playbook`     | A higher-level guide covering a workflow or category of tasks           |
| `how-to`       | Step-by-step guide for a specific task (audience: practitioners)        |
| `runbook`      | Operational procedure with verification and rollback steps              |
| `example-plan` | Plan for a worked example — captures goal, approach, and acceptance criteria |

Custom types are valid — OKF is minimally opinionated. Use `example-plan` for example plans.

---

## Naming conventions

- **Filenames:** `kebab-case.md` derived from the concept title
- **No generic names:** prefer `contract-first.md` over `idea-2.md`
- **Index and log:** always `index.md` and `log.md` at `docs/` root — never rename

---

## Cross-references

Use standard relative markdown links. Paths are relative to the linking file:

```markdown
[Domain Names as the Bridge](./ideation/domain-names-as-bridge.md)
[Composable Skills research](./research/composable-skills.md)
```

Never use `[[wikilinks]]` or absolute paths.

---

## index.md

`docs/index.md` is the structured catalog of all docs. Updated on every write operation.

Format:
```markdown
| [Title](relative/path.md) | type | Description |
```

---

## log.md

`docs/log.md` is append-only. Never rewrite or delete existing entries.

Format:
```
## [YYYY-MM-DD] <operation> | <detail>
```

Operations: `ingest`, `query`, `lint`, `add`, `edit`, `init`
