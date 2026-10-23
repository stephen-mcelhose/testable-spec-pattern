# AGENTS.md

Steering document for any agent working in this repo. Read this first.

## What this repo is

A personal workshop for building a **Go behavioral-spec workflow** — tests that read like RFC 2119 contracts, grounded in domain names that bridge spec, implementation, and test.

## North star

**The goal is a workflow expressed as composable, agentskills-standard skills.** Every skill in `.agents/skills/` should be a valid `agentskills.io` skill — proper frontmatter, `allowed-tools`, narrow focus, and a clear trigger description. Skills that don't meet the standard should be treated as drafts.

Skills should be composable in the style of [Matt Pocock's skills repo](https://github.com/mattpocock/skills): narrow discipline skills (model-invoked) composed by orchestrator skills (user-invoked). An orchestrator can call a discipline skill; disciplines never call each other. If a skill's purpose contains "and", split it. See [`docs/research/composable-skills.md`](docs/research/composable-skills.md) for the full research.

## Where to look

| Resource | What you'll find |
| -------- | ---------------- |
| [`docs/index.md`](docs/index.md) | Full catalog of the knowledge base |
| [`docs/discipline.md`](docs/discipline.md) | The core behavioral-spec discipline — read before writing any tests |
| [`docs/conventions.md`](docs/conventions.md) | Wiki conventions for `docs/` — OKF frontmatter, folder taxonomy |
| [`docs/research/agentskills-standard.md`](docs/research/agentskills-standard.md) | What the agentskills standard requires (read before creating or editing skills) |

## Skills in this repo

### Exists

| Skill | Purpose |
| ----- | ------- |
| [`go-behavioral-spec`](.agents/skills/go-behavioral-spec/SKILL.md) | Write Go tests as RFC 2119 behavioral specs |
| [`docs-wiki`](.agents/skills/docs-wiki/SKILL.md) | Manage the `docs/` knowledge base |

### Planned

| Skill | Purpose |
| ----- | ------- |
| `contract-first-with-domain-names` | Guide contract-first design (OpenAPI, proto, ROD) grounded in domain names — the bridge between spec, implementation, and test |
| `planning` | Consistent planning format with known sections, subdivided by class of problem |

Skills are added when a documented need or failure demands them — not speculatively.

## Skill conventions

- Skills are **model-invoked** unless explicitly marked otherwise
- Skill frontmatter MUST include `name`, `description`, `license`, `allowed-tools`
- `allowed-tools` should be as narrow as possible — don't grant broad `Bash` access when specific patterns will do
- **Skills must be self-contained and packageable.** A skill may compose over other skills in the workflow, but MUST NOT depend on docs, examples, or any other file in this repo. If a skill needs reference material (examples, discipline docs), bundle it inside the skill directory. The goal: any skill could be copy-pasted to a different repo and work as-is.
