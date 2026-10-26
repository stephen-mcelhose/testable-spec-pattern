---
type: decision
title: "ADR-006: Skill eval artifacts live inside the skill directory"
description: Where rubrics, scorecards, and other skill evaluation artifacts are stored.
tags: [adr, evals, skills, docs]
timestamp: 2026-10-24
---

# ADR-006: Skill eval artifacts live inside the skill directory

## Status

Accepted

## Y-Statement

In the context of building an eval loop for agent skills (rubric → run → score → improve), facing the question of where to store rubrics and test case definitions, we decided to follow the agentskills convention and place them under `{skill}/evals/` to achieve portability (the skill and its evals travel together), accepting that run outputs (workspace) live separately and are not committed alongside the skill.

## Decision

- **Eval definitions** (evals.json, assertions) go in `{skill}/evals/`
- **Research artifacts** (coverage audits, gap analyses, raw normative extracts) go in `docs/research/{skill-name}/` — they are repo knowledge, not skill-portable
- **Run outputs** (grading, benchmarks, feedback) go in an adjacent workspace directory, not committed
- Skills remain self-contained: only `evals.json` and directly skill-relevant fixtures travel with the skill

## Alternatives considered

| Option                             | Rejected because                                                              |
| ---------------------------------- | ----------------------------------------------------------------------------- |
| `docs/evals/`                      | Breaks skill portability — evals should travel with the skill they test       |
| Separate top-level `evals/` folder | Adds a new top-level concept and breaks skill portability                     |
| All research inside skill          | Research artifacts are repo knowledge, not skill-portable                     |

## Consequences

- `docs/index.md` does NOT need an `evals/` section — evals are a skill concern, not a wiki concern
- Research artifacts for a skill live in `docs/research/{skill-name}/` (established by moving coverage-audit.md, gap-audit.md, normative-raw.md, trigger-eval.json there)
- If the agentskills convention evolves, revisit
