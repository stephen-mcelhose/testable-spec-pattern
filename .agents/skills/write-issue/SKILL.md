---
name: write-issue
description: Write a clear, focused, and actionable GitHub issue body. Use when given a topic, feature idea, bug, or skill to build and asked to produce a ready-to-file issue. Trigger on "write an issue for X", "draft an issue about X", "create a ticket for X", "I need an issue for X".
license: MIT
allowed-tools: Read Write Bash(gh issue list:*) Bash(gh issue view:*)
---

# write-issue

Produce a GitHub issue body that a human or agent can pick up cold and deliver without back-and-forth. The output must satisfy the Definition of Ready (see below) before it is presented.

See `references/issue-patterns.md` for the research behind the format choice.

---

## Preferred format

The format is a synthesis of **Need / Use-Case / Acceptance Criteria** and **Task / Guardrails / Verification**, stripped to the minimum that satisfies the Definition of Ready and extended with an explicit out-of-scope section.

```markdown
## What
One sentence. The concrete artefact or change being delivered. Start with a noun, not a verb. No "investigate", "look into", or "consider" — every issue delivers something.

## Why
The pain, gap, or risk that makes this necessary. One paragraph. A new contributor should be able to read this without asking follow-up questions.

## Use-case
A concrete scenario: who invokes this, what they do, and what they get. Shows the deliverable in context, not in the abstract.

## Acceptance criteria
- [ ] Specific, individually verifiable statement
- [ ] Another checkable statement
- [ ] …

## Out of scope
- Explicit exclusion — what a reasonable implementer might assume is included but isn't
- Another exclusion
```

**Optional sections** (add only when genuinely needed):

```markdown
## Dependencies
External requirements that must be in place before work can start: upstream issues, required API access, data, or other external blockers.
```

---

## Why this format

**Why not User Stories?** "As a [role], I want…" requires a persona. In this repo, humans and agents are interchangeable workers — there is no meaningful role distinction. Personas also fragment naturally unified motivations into artificial per-role splits.

**Why not Job Stories?** "When / I want to / So I can" captures situation and motivation well, but the sentence format doesn't compose with checkable acceptance criteria. It is an input to issue-writing, not the issue format itself.

**Why not RFC-style?** RFC templates are thorough but heavyweight. A Kanban ticket with nine sections ("Prior Art", "Future Possibilities", "Reference-level explanation"…) will be partially completed most of the time. Partial completion is worse than a shorter format done completely.

**Why not Task / Guardrails / Verification alone?** This three-part agentic format is tight and self-reviewing, but it omits *why* — an agent that doesn't understand the motivation will satisfy the task literally while missing the intent when edge cases arise.

**Why the synthesis?** The four questions every good issue must answer are: *What, Why, Scenario, Done*. The preferred format maps one section to each question, adds explicit out-of-scope (the most consistently missing guardrail), and keeps total sections at five or fewer to maximise completion rate.

---

## Steps

### 1. Understand the topic

Before writing anything, identify:
- The **artefact**: what is being built, changed, or fixed? If the topic is vague, ask one clarifying question.
- The **motivation**: why does this matter? If not stated, infer from context and state your inference explicitly in the Why section so a reviewer can correct it.
- Whether any **existing issues** cover this topic (`gh issue list` to check for duplicates).

### 2. Draft the What

Write one sentence. Test it: can you hand this sentence to a new contributor and have them understand what they are supposed to produce?

- ✓ "A `write-issue` skill that guides an agent to write focused, actionable GitHub issues."
- ✗ "Improve how we write issues." (no artefact)
- ✗ "Investigate whether we need a better issue format." (delivers nothing)

### 3. Draft the Why

Write one paragraph explaining the pain or gap. Include evidence where possible (observed failures, inconsistency, a concrete past example). Do not describe the solution — save that for the use-case.

### 4. Draft the Use-case

Write a short narrative (2–5 sentences) of the scenario in which the deliverable is used. Name the actor (agent, human, user), the trigger, the action, and the result.

Example: *"Before any work begins, an agent invokes `write-issue` with a topic and receives a ready-to-file issue body. The agent reviews the body for completeness and files it; a human reviewer can confirm the issue satisfies the Definition of Ready without asking clarifying questions."*

### 5. Draft the Acceptance criteria

Each criterion must be:
- **Specific** — names the exact behaviour, file, output, or condition
- **Verifiable** — a reviewer can check it without interpretation
- **Independent** — one thing per bullet; if a bullet needs "and", split it

Write the AC as a checklist. Aim for 3–7 items. More than seven suggests the issue bundles multiple concerns.

### 6. Draft the Out of scope

List at least two explicit exclusions. These are things a reasonable implementer might assume are included based on the What/Why, but are not part of this issue. If you cannot name any exclusions, the scope is probably still underspecified.

### 7. Apply the Definition of Ready checklist

Before presenting the issue, verify every item:

- [ ] Single deliverable — no bundled concerns
- [ ] The "why" is stated (the pain or gap being addressed)
- [ ] At least one concrete use-case
- [ ] Acceptance criteria are specific and individually checkable
- [ ] Out of scope is explicit
- [ ] Dependencies on other issues are listed (or explicitly none)
- [ ] No open questions that would block an implementer from starting

If any item fails, revise before presenting.

---

## Anti-patterns

### 1. Implementation prescription

**Symptom:** The issue describes *how* to build something rather than *what* to deliver.

> ✗ "Add a `FormatIssueBody` function in `pkg/issue/format.go` that calls the template engine with the preferred format…"

**Problem:** Prescribing implementation locks the implementer in before they have read the code. It conflates the issue (what/why/done) with the design (how). Design decisions belong in the implementation, not the issue.

**Fix:** Describe the outcome. "A skill that produces a workable issue body when given a topic" — not the internal function signature.

---

### 2. No use-case

**Symptom:** The issue states a need and acceptance criteria but skips the concrete scenario.

> ✗ **What:** A rate-limiter for the ingestion pipeline.  
> ✗ **Why:** The pipeline occasionally falls over under load.  
> ✗ **AC:** The pipeline handles bursts without error.

**Problem:** Without a use-case, the implementer has no model of *when* the rate-limiter fires, *what* triggers it, or *what the caller sees*. Two implementers reading the same issue will build different things.

**Fix:** Add a scenario: *"When the ingestion service receives more than 500 events per second, it queues excess events rather than dropping them or returning errors to the caller. Under normal load the queue is empty."*

---

### 3. Missing out of scope

**Symptom:** The issue has no explicit exclusions. Everything adjacent to the deliverable is implicitly in scope.

**Problem:** Implementers expand into undefined space. An issue for "a rate-limiter" without exclusions will attract: retry logic, backpressure signalling, metrics dashboards, configuration UI, and a tuning guide. Each is adjacent; none was intended.

**Fix:** List explicit exclusions immediately after the acceptance criteria: "Out of scope: retry logic (separate issue), metrics export, operator-facing configuration."

---

### 4. Bundled concerns

**Symptom:** The issue covers two or more independent deliverables.

> ✗ "Add a rate-limiter and update the runbook to document the new thresholds."

**Problem:** Two concerns → two reviewers → two code-review threads → one issue that is never fully done because one part blocks the other. If the runbook update requires the rate-limiter to be live first, that dependency is hidden inside a single issue.

**Fix:** One issue per deliverable. Use the "Dependencies" section to chain them explicitly: the runbook issue lists the rate-limiter issue as a dependency.

---

### 5. Template bloat leading to partial completion

**Symptom:** The issue template has many sections; most are empty or contain placeholder text.

> ✗ An issue with "Testing Scenarios", "Links", "Prior Art", "Dependencies", "Migration Notes", and "Future Possibilities" — where only "Description" and "AC" are filled in.

**Problem:** Blank sections signal to the reader (and agent) that the issue was not fully thought through. They also create ambiguity: is the section blank because it was skipped, or because it genuinely does not apply?

**Fix:** Keep required sections to five or fewer. Omit optional sections entirely when they do not apply — do not leave them blank.

---

## Rules

- **One issue, one deliverable.** If the What sentence needs "and", split the issue.
- **State the why.** An issue without a Why cannot be prioritised or deprioritised — it just sits.
- **Criteria are verifiable.** If you cannot check a criterion without human judgement, rewrite it.
- **Out of scope is mandatory.** At least two explicit exclusions. If you can't name any, the scope is underspecified.
- **Definition of Ready is the gate.** Do not present the issue until all seven items are checked.
