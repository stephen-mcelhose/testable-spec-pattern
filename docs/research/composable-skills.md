---
type: spike
title: Composable Skills Architecture
description: A research-backed guide on building composable AI agent skills using Matt Pocock's "skills" repository patterns, focusing on orchestrator/discipline separation and markdown-as-state.
resource: https://github.com/mattpocock/skills
tags: [skills, composable, agent, workflow]
timestamp: 2026-10-23T22:45:43Z
---

# Research: Composable Skills Architecture

## Goal
This spike investigates how to structure and organize composable AI agent skills, leaning on Matt Pocock's approach. The goal is to provide a blueprint for growing the `testable-spec-pattern` repo from a single skill into a multi-skill architecture comprising orchestrators, discipline skills, and worked examples.

## Approach
1.  **Codebase Analysis**: Examined the current repo structure (`docs/discipline.md`, `.agents/skills/`) to understand the existing "go-behavioral-spec" discipline.
2.  **Web Research**: Analyzed Matt Pocock's `skills` repository (README, SCOPE, AGENTS).
3.  **Pattern Extraction**: Identified key architectural patterns: orchestrator vs. discipline separation, hierarchical composition rules, and the "markdown-as-state-substrate" pattern.
4.  **Synthesis**: Applied these patterns to the specific Go testable-spec-pattern context.

## Findings

### Repository Structure & Folder Layout
Matt Pocock organizes skills into domain-specific buckets to manage complexity:
-   **`skills/engineering/`**: Reusable technical disciplines (e.g., `tdd`, `code-review`).
-   **`skills/productivity/`**: Workflow-oriented skills (e.g., `grill-me`, `to-tickets`).
-   **Skill Folders**: Each skill lives in its own directory containing a `SKILL.md` file, which defines the agent's behavior for that specific skill.

### Orchestrator vs. Discipline Skills
A critical distinction in this architecture is the invocation model:
-   **Orchestrator Skills (User-invoked)**: Triggered directly by the user (e.g., `/grill-me`, `/implement`). These skills manage the high-level workflow, interface with the user, and decide which discipline skills to call.
-   **Discipline Skills (Model-invoked)**: Can be reached automatically by the agent when appropriate (e.g., `/tdd`, `/go-behavioral-spec`). They hold the "hard rules" and technical patterns.

### Composition Rules
-   **Hierarchical Invocation**: Orchestrators (User-invoked) can invoke Discipline skills (Model-invoked).
-   **Flat Orchestration**: A user-invoked skill should **never** invoke another user-invoked skill. This prevents infinite loops and complex state transitions that confuse the agent.
-   **Narrow Focus**: Skills are kept "narrow" by focusing on a single, observable outcome or discipline. If a skill does too much (contains "and" in its purpose), it should be split.

### Markdown-as-State-Substrate
Rather than using external databases, state is managed through localized Markdown files in the workspace:
-   **`GLOSSARY.md`**: Stores a "ubiquitous language" to ensure the agent and human use consistent terms, reducing token waste and confusion.
-   **Task/Ticket Files**: The `/to-tickets` pattern writes plan state directly to the repo, allowing the agent to "see" its progress in the file system.
-   **`handoff.md`**: Condenses conversation history into a structured artifact so a new agent session can resume without losing context.

### Key Governance Files
-   **`SCOPE.md`**: A human-facing governance doc. It enforces a "high bar" for new skills: they must address a documented agent failure, not just a hypothetical improvement.
-   **`AGENTS.md` / `CLAUDE.md`**: Agent-facing steering documents. They act as the "entry point" or "system prompt extension" for the repo, pointing the agent to relevant skills and the `GLOSSARY.md`.

## Recommendation

To grow the `testable-spec-pattern` repo, I recommend the following steps:

1.  **Restructure `.agents/skills/`**:
    -   Move the existing `go-behavioral-spec` to `.agents/skills/discipline/go-behavioral-spec/SKILL.md`.
    -   Create `.agents/skills/orchestrator/spec-loop/SKILL.md` to handle the end-to-end workflow (e.g., "Grill the user on requirements → Write spec tests → Implement → Verify").

2.  **Establish the State Substrate**:
    -   Create a `docs/GLOSSARY.md` (or similar) to capture the RFC 2119 keywords and domain terms (MUST, SHOULD, etc.) defined in `discipline.md`.
    -   Use a `docs/research/tickets/` or `docs/tasks/` folder for the agent to track its own implementation progress.

3.  **Implement Steering**:
    -   Add an `AGENTS.md` at the root that instructs any agent entering the repo to read `docs/discipline.md` and the `GLOSSARY.md` first.

4.  **Composition Implementation**:
    -   The `spec-loop` orchestrator should be instructed to call the `go-behavioral-spec` discipline skill specifically during the "Write spec tests" phase.

## References
- [Matt Pocock's Skills Repo](https://github.com/mattpocock/skills)
- [Matt Pocock's README.md](https://raw.githubusercontent.com/mattpocock/skills/main/README.md)
- [Matt Pocock's SCOPE.md](https://github.com/mattpocock/skills/blob/main/SCOPE.md)
