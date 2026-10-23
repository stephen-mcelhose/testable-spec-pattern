---
type: decision
title: "ADR-003: Kanban Workflow via GitHub Issues"
description: Establishes a lightweight Kanban protocol — board state, Definition of Ready, session protocol, and Definition of Done — that works equally for human contributors and AI agents.
tags: [adr, workflow, kanban, github, issues]
timestamp: 2026-10-23T23:54:28Z
---

# ADR-003: Kanban Workflow via GitHub Issues

## Status

Accepted

## Y-Statement

In the context of a repo with a growing backlog of skill-building work that spans multiple sessions and contributors (human or agent), facing the risk of context loss, duplicated effort, and issues that aren't actionable when picked up cold, we decided on a Kanban workflow using GitHub Issues as the board with an explicit Definition of Ready and session protocol, to achieve continuity and clarity regardless of who (or what) picks up the work, accepting the overhead of maintaining issue quality.

## Decision

The full protocol is in `AGENTS.md § Workflow`. The decisions behind it are recorded here.

### Why Kanban, not Scrum sprints

No fixed cadence. Work happens in sessions of variable length across human and agent contributors. Kanban's pull-based, continuous-flow model fits better than time-boxed sprints.

### Why WIP limit = 1

An agent context-switching between two in-progress issues is no different from a human doing the same: neither produces good work. One issue in progress at a time prevents shallow half-done work and keeps sessions focused.

### Why a Definition of Ready (DoR)

An issue that isn't ready to work wastes a session. The DoR is not bureaucracy — it is the minimum information needed to start without asking clarifying questions. The threshold: **could someone cold-start from this issue and produce the right thing?** If not, the issue isn't ready.

The human/agent parity principle is the root constraint: if a step in the workflow requires tooling only an agent has, or intuition only a human has, it is wrong. Every step must be executable by either.

### Why DoR is checked before assigning, not after

Assigning signals "I am working this." If the issue isn't ready, assigning it blocks it without producing value. The check comes first; if the issue fails the check, the right action is to comment on the gap and move on.

### Priority model

Human contributors set priority. In the absence of explicit priority, any contributor (human or agent) chooses the most unblocked issue — the one with the fewest unresolved upstream dependencies. The question "is this actually ready to work?" subsumes priority: a high-priority unready issue waits; a low-priority ready issue can be worked.

## Consequences

- Issues that fail the Definition of Ready are not started; a comment explaining the gap is left instead.
- "In Progress" is signalled by assignment, making the board readable without a separate tool.
- Incomplete sessions MUST leave a comment and a committed branch — no invisible in-progress state.
