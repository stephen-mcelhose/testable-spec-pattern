---
type: decision
title: "ADR-007: Issue title prefix scheme"
description: Which prefix scheme to use for GitHub issue titles to make the backlog scannable and discoverable via gh issue list.
tags: [adr, issues, workflow, write-issue, kanban]
timestamp: 2026-10-25
---

# ADR-007: Issue title prefix scheme

## Status

Accepted — Verb-first imperative (Option 2)

## Context

The `write-issue` skill encodes a title formula:

```
<type>: <domain-artifact> — <specific-change>
```

The `<type>` prefix is the first token a reader sees when scanning `gh issue list`. It signals the nature of the work before anything else is read. The choice of prefix scheme has consequences for:

- **Scannability** — how quickly a human or agent can categorise a list entry
- **Parity** — whether issue prefixes align with commit message prefixes (reducing cognitive switching)
- **Expressiveness** — how well the prefix captures the distinction that actually matters (delivery vs exploration vs maintenance vs correction)
- **Learnability** — how much a new contributor needs to know before filing correctly

The current scheme is Conventional Commits. This ADR evaluates alternatives before committing.

---

## Options

### Option 1: Conventional Commits *(current; community standard)*

`feat:` `fix:` `chore:` `docs:` `refactor:` `spike:` `test:`

Drawn from the [Conventional Commits](https://www.conventionalcommits.org/) specification (v1.0.0), extended with `spike:` for exploration work.

**What it expresses:** whether the work delivers user-visible capability (`feat`), corrects a defect (`fix`), or is internal/maintenance (`chore`, `refactor`, `docs`, `test`).

**Parity:** direct — the same vocabulary is used in commit messages, so an issue title and its eventual commit message share a prefix.

**Tradeoffs:**
- `feat` / `fix` are semantically loaded toward deployed software; they're awkward for skill authoring or ADR writing
- `chore` is a catch-all that frequently obscures more than it reveals
- The scheme was designed for release automation, not issue discovery — the distinction between `chore` and `refactor` matters for changelogs but less for a kanban board

---

### Option 2: Verb-first imperative *(borrowed from commit convention)*

No explicit type token. The title begins with a strong imperative verb. The verb *is* the type signal.

```
Define billing Invoice proto — AIP-compliant resource with standard CRUD methods
Wire BillingService.CreateInvoice RPC — implementation to generated proto types
Implement orders-created Kafka consumer — with dead-letter queue on processing failure
Fix shipping Shipment proto — correct State enum nesting per AIP-216
Spike: orders-created Kafka consumer — evaluate dead-letter vs retry-with-backoff
```

Imperative mood is a well-established convention for **commit messages** — Git's own documentation and Chris Beams' widely-cited post ([cbea.ms/git-commit](https://cbea.ms/git-commit/)) both prescribe it. Its application to *issue titles* is by analogy, not by documented standard. No citable specification for verb-first issue titles was found.

**What it expresses:** the action being taken. Readers scan the verb to triage.

**Addressing variance — a curated verb list per work class:**

Without discipline, verb choice is inconsistent (Add vs Implement vs Wire — meaningfully different things). A per-work-class vocabulary resolves this:

| Work class         | Canonical verbs                                        |
| ------------------ | ------------------------------------------------------ |
| Proto contract     | `Define` (new resource), `Annotate` (field behaviors), `Restructure` (file layout), `Fix` (linter violations) |
| gRPC method        | `Wire` (impl to generated types), `Spec` (behavioral tests), `Add` (new method), `Fix` (error codes, pagination) |
| Kafka consumer     | `Implement` (new consumer), `Add` (handler logic), `Fix` (error path/retry) |
| PubSub consumer    | `Implement` (new consumer), `Add` (handler logic), `Fix` (ack/nack strategy) |
| Go service / pkg   | `Implement` (new service), `Wire` (dependencies), `Extract` (interface/helper), `Refactor` |
| Skill              | `Add` (new guidance), `Update` (existing guidance), `Refactor` |
| ADR / docs         | `Document` (decision/ADR), `Add` (reference doc), `Update` |
| Investigation      | `Spike:` (breaks the pattern intentionally — keeps explorations visually distinct) |

**Tradeoffs:**
- No reliable `gh issue list --search` filter by type (no common token across all issues of a type)
- `Spike:` is a pragmatic exception that introduces a de-facto prefix for the exploration case
- High expressiveness when the verb list is followed; high variance when it isn't — requires the reference to be consulted
- No commit parity, though verb-first commit messages (imperative mood) are themselves a well-established convention

---

### Option 3: Bracket notation *(enterprise / Jira-adjacent)*

`[feature]` `[bug]` `[task]` `[spike]` `[doc]` `[refactor]`

```
[feature] BillingService.CreateInvoice RPC — wire implementation to generated proto types
[spike] orders-created Kafka consumer — evaluate dead-letter vs retry-with-backoff
[doc] ADR-008 — adopt buf as proto enforcement gate
```

Common in Jira-heavy environments and teams that treat issue type as a formal field separate from the title text.

**What it expresses:** same taxonomy as Conventional Commits but with a different visual grammar — label-then-content rather than `prefix:content`.

**Tradeoffs:**
- Bracket scanning is slower than colon-prefix scanning at a glance
- No commit parity
- `[feature]` is more legible than `feat:` to non-technical stakeholders
- Works well if the prefix vocabulary intentionally diverges from commit vocabulary

---

### Option 4: Delivery vs exploration *(minimal)*

Two prefixes only. Everything else is expressed in the artifact and the action verb.

`deliver:` `explore:`

```
deliver: BillingService.CreateInvoice RPC — wire implementation to generated proto types
explore: orders-created Kafka consumer — evaluate dead-letter vs retry-with-backoff strategy
```

The single most useful distinction on a kanban board is whether an issue will produce a committed artefact or a decision/finding. Everything else (feat vs fix vs chore) is secondary and can be inferred from the artifact name and verb.

**What it expresses:** the fundamental work mode — building or learning.

**Tradeoffs:**
- Loses the feat/fix/chore distinction — cannot tell from the prefix whether work is new capability or correction
- Very low ambiguity — "is this deliver or explore?" is almost always unambiguous
- No commit parity, no community standard
- May need extension (e.g. `maintain:`) as the backlog grows

---

### Option 5: Area-first *(Linux kernel style)*

The prefix is the **domain area**, not the work type. Area is already encoded in the artifact name, so this collapses two signals into one token.

`proto:` `grpc:` `kafka:` `pubsub:` `skill:` `docs:` `infra:`

```
proto: billing Invoice — define AIP-compliant resource with standard CRUD methods
grpc: BillingService.CreateInvoice — wire implementation to generated proto types
skill: write-issue — add title formula and label vocabulary
docs: ADR-008 — adopt buf as proto enforcement gate
```

Used in the Linux kernel (`net:`, `mm:`, `fs:`). The area is the first filter; work type is inferred from the action verb.

**What it expresses:** which domain the work touches. Optimises for domain filtering over type filtering.

**Tradeoffs:**
- Overlaps with the label vocabulary — `proto:` prefix and `proto` label are redundant
- No commit parity unless area-scoped commits are adopted across the board
- Loses the delivery-vs-exploration signal without a secondary convention
- `docs:` and `infra:` work well; `grpc:` and `kafka:` are more domain-specific than most teams' commit scopes

---

## Decision

**Option 2: Verb-first imperative.**

Issue titles begin with a canonical imperative verb from the per-work-class vocabulary. No type prefix token. `Spike:` is the only permitted prefix, used to keep explorations visually distinct from delivery work.

The title formula becomes:

```
<Verb> <domain-artifact> — <specific-change>
```

Examples:
```
Define billing Invoice proto — AIP-compliant resource with standard CRUD methods
Wire BillingService.CreateInvoice RPC — implementation to generated proto types
Spike: orders-created Kafka consumer — evaluate dead-letter vs retry-with-backoff
```

The verb list per work class (see Option 2 above) is the canonical reference. When in doubt, consult it before choosing a verb.

## Success criteria

The decision is working if the following access patterns produce usable results using [`gh issue list`](https://cli.github.com/manual/gh_issue_list):

| Access pattern                        | Command                                              | Expected result                            |
| ------------------------------------- | ---------------------------------------------------- | ------------------------------------------ |
| All open issues                       | `gh issue list`                                      | Verb-led titles scan cleanly by eye        |
| Find work on a specific artifact      | `gh issue list --search "CreateInvoice"`             | Returns issues touching that artifact      |
| Find all explorations                 | `gh issue list --search "Spike:"`                    | Returns only investigation issues          |
| Find work in a domain                 | `gh issue list --label grpc`                         | Returns domain-filtered set                |
| Find work of a specific action type   | `gh issue list --search "Wire"`                      | Returns issues with that verb              |

Labels carry the domain-filter load that a type prefix would otherwise carry. Search carries the verb-filter load. Both must work in practice.

## Revisit condition

Revisit this decision if:

- `gh issue list --search "<verb>"` returns noisy results — issues from unrelated contexts where the verb appears incidentally in the body, making verb-based filtering unreliable
- A work class emerges for which no canonical verb cleanly distinguishes it from adjacent work
- The per-work-class verb list requires more than one lookup per filing — indicating the vocabulary is too large to internalise

Reference: [`gh issue list` manual](https://cli.github.com/manual/gh_issue_list)

## Consequences

- The `write-issue` skill title formula and verb table encode this decision. If the decision changes, update the skill.
- Labels become load-bearing for domain filtering — the label vocabulary in the skill must be kept current as new work classes are added.
- Commit message style is unaffected — commits continue to use Conventional Commits for release automation. Issue titles and commit messages now use different conventions intentionally.
