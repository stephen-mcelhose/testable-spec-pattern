---
type: decision
title: "ADR-002: GitHub Issues as Work Tracker"
description: Use GitHub issues to track skill-building work because the scope has grown beyond what fits in a single agent session.
tags: [adr, workflow, github, issues]
timestamp: 2026-10-23T23:44:23Z
---

# ADR-002: GitHub Issues as Work Tracker

## Status

Accepted

## Y-Statement

In the context of a repo where scope has grown to encompass multiple skills, examples, and conventions that cannot be completed in a single session, facing the risk of losing intent and context between sessions, we decided to file GitHub issues for each discrete skill or piece of work before starting it, to achieve continuity and a clear audit trail, accepting the overhead of issue management.

## Decision

All planned work MUST be captured as a GitHub issue before execution begins.

An issue MUST include:
- **What** — the skill or artefact being built, in one sentence
- **Why** — the use-case it serves in the workflow
- **Acceptance criteria** — specific, testable statements of done
- **Out of scope** — explicit exclusions to keep issues focused

Use labels to categorise: `skill`, `example`, `docs`, `adr`.

Agent sessions pick up work by reading the open issues. The first act of any session is to check open issues and resume the most relevant one.

## Consequences

- No work starts without an issue. "I'll just do it quickly" is not acceptable.
- Issues accumulate the project backlog, making it readable after a gap of days or weeks.
- Completed issues are closed with a reference to the commit or PR.
