---
type: research
title: Issue title naming conventions
description: Survey of prefix and naming schemes for issue/task titles across OSS, enterprise, and community standards. Research to support ADR-007.
tags: [research, issues, naming, write-issue, kanban]
timestamp: 2026-10-25
---

# Issue title naming conventions

Research supporting [ADR-007](../decisions/007-issue-title-prefix-scheme.md).

---

## Finding 1 — No documented standard exists for issue title naming

No verifiable, citable specification prescribes a naming convention specifically for issue or task titles (GitHub Issues, Jira, Linear, Shortcut, Azure DevOps). The conventions that do exist are documented for **commit messages**, not issue titles. Teams apply them to issues by analogy.

---

## Finding 2 — Modern PM tools treat type as a separate field, not a title prefix

Every current PM tool surveyed puts work item type in a **structured metadata field**, not the title:

| Tool          | Type model                                     | Titles                         |
| ------------- | ---------------------------------------------- | ------------------------------ |
| Jira          | Issue type field (Epic, Story, Task, Bug)      | Freeform — no title convention |
| GitHub Issues | Issue Types field (2024, separate from labels) | Freeform — no title convention |
| Linear        | Type as a property on each issue               | Freeform — no title convention |
| Shortcut      | Story type field (Feature, Bug, Chore)         | Freeform — no title convention |
| Azure DevOps  | Work item type (Epic, Feature, Story, Bug)     | Freeform — no title convention |

None of these tools prescribe type-in-title. The bracket notation (`[BUG]`, `[FEATURE]`) is a **community workaround**, not a tool convention — teams embedded type in titles because type metadata wasn't visible in email notifications, third-party integrations, or CLI output.

**This is directly relevant to our context.** `gh issue list` doesn't surface label or type metadata in its default output — only title, number, and state. Encoding type in the title is meaningful *because* we're CLI-first. But there's no PM tool giving us a standard to borrow; we're solving a problem those tools don't have.

---

## Shortcut's branch naming as a data point

Shortcut's Git integration auto-generates branch names in the format `feature/sc-123-story-title`, `bug/sc-123-story-title`, `chore/sc-123-story-title`. The type appears as a **path segment**, not a title prefix. This is the closest any tool gets to type-in-title — and it's in branch names, not issue titles.

---

## Commit conventions surveyed

These are documented for commits, not issues. Included for completeness since they're the source of Options 1, 2, and 5 in ADR-007.

### Conventional Commits

- **Spec:** https://www.conventionalcommits.org/en/v1.0.0/
- **Scope:** Commit messages. Designed to drive automated semantic versioning and changelog generation.
- **Format:** `<type>[optional scope]: <description>` — types: `feat`, `fix`, `chore`, `ci`, `docs`, `refactor`, `perf`, `test`, `build`
- **Application to issues:** By analogy only.

### Linux kernel subsystem prefix

- **Spec:** https://www.kernel.org/doc/html/latest/process/submitting-patches.html
- **Scope:** Patch email subject lines (which become commit subjects via `git am`).
- **Format:** `[PATCH] subsystem: summary phrase` — subsystem identifies which kernel area is being patched (e.g. `ext2`, `x86`, `net`). Imperative mood required in the description.
- **Application to issues:** None. The subsystem prefix is for patch submissions, not task tracking.

### Verb-first imperative (Beams)

- **Spec:** https://cbea.ms/git-commit/
- **Scope:** Git commit subject lines only. The article explicitly does not extend this to issue titles — it notes that issues follow different grammatical conventions (e.g. "User profile picture does not load on Safari").
- **Rule:** Subject uses imperative mood. Test: "If applied, this commit will [subject line]."
- **Application to issues:** By analogy only. Not a citable standard for issue titles.

---

## Summary for ADR-007

- Options 1, 2, and 5 are commit conventions applied to issue titles by analogy — none are issue title standards.
- Option 3 (bracket notation) has no spec but is native to issue titles as a community workaround for metadata-visibility gaps.
- Option 4 (deliver vs explore) is entirely custom.
- No further PM tool exploration is warranted — all modern tools converge on type-as-separate-field. Bracket notation is the only "issue-native" encoding pattern and it comes from that same gap, not from tool design.
- The strongest argument for any prefix scheme in our context is not "it's a standard" but "it solves the `gh issue list` visibility problem" — that's the real requirement to evaluate against.
