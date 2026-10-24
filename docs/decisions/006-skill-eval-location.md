---
type: decision
title: "ADR-006: Skill eval artifacts live inside the skill directory"
description: Where rubrics, scorecards, and other skill evaluation artifacts are stored.
tags: [adr, evals, skills, docs]
timestamp: 2025
---

# ADR-006: Skill eval artifacts live inside the skill directory

## Status

Accepted

## Y-Statement

In the context of building an eval loop for agent skills (rubric → run → score → improve), facing the question of where to store rubrics and test case definitions, we decided to follow the [agentskills standard](https://agentskills.io/skill-creation/evaluating-skills) and place them under `{skill}/evals/` to achieve portability (the skill and its evals travel together) and consistency with the established standard, accepting that run outputs (workspace) live separately and are not committed alongside the skill.

## Decision

- **Test case definitions and rubrics** go in `{skill}/evals/` — e.g. `.agents/skills/resource-oriented-design/evals/rubric.md`
- **Run outputs** (grading, benchmarks, feedback) go in an adjacent workspace directory: `.agents/skills/resource-oriented-design-workspace/iteration-N/` — not committed to the repo by default
- Skills remain self-contained and copy-paste portable: the `evals/` folder travels with the skill

## Alternatives considered

| Option                             | Rejected because                                                                 |
| ---------------------------------- | -------------------------------------------------------------------------------- |
| `docs/evals/`                      | Initially considered; overruled by the agentskills standard which is explicit    |
| Separate top-level `evals/` folder | Adds a new top-level concept and breaks skill portability                        |
| External eval repo                 | Premature; this pattern is still being established                               |

## Consequences

- `docs/index.md` does NOT need an `evals/` section — evals are a skill concern, not a wiki concern
- If the agentskills standard evolves, revisit
