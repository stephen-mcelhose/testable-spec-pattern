# Plan: Disciplined Approach to Clean Domain Code (with Agent Support)

**Branch:** `stephen-workflow-skills`  
**Status:** In progress — stopped mid-flow  

---

## What we are doing

Documenting a coherent, opinionated discipline for writing Go code that is:

- **Legible** — names rooted in domain language, not implementation detail
- **Testable** — pure functional core with side effects isolated at leaves
- **Specified** — tests that read as behavioral contracts, not implementation descriptions

The output is a set of agent-executable skills and supporting reference documents. The discipline drives the skills, not the other way around — we write the philosophy down first, then encode it.

---

## What we have done

### 1. `go-behavioral-spec` skill (committed on `stephen-workflow-skills`)

Location: `.agents/skills/go-behavioral-spec/`

- **`SKILL.md`** — executable skill for writing RFC 2119 behavioral spec tests
  - Prefer table-driven tests
  - Flat `t.Run`s only when branches are genuinely orthogonal
  - `t.Parallel()` everywhere, testify only, named skips for gaps
  - Links to discipline doc for the "why"

- **`references/discipline.md`** — the design philosophy behind the skill
  - Test names are specifications, not descriptions
  - RFC 2119 keyword discipline (MUST / MUST NOT / SHOULD / MAY)
  - Table preference as a design signal — resistance to tabling = design problem
  - **The principle:** push testable side effects to leaves, expose them through narrow injected interfaces
  - Pure core + leaf interfaces = naturally table-drivable code
  - Upgrade path: write new spec suite, then retire old tests if redundant
  - Open questions: stub placement, one-file-per-source vs per-concern

**Note:** This skill may want to move to personal skills (`~/.cursor/skills/` or equivalent) rather than staying in the repo — it's a cross-project discipline, not p360-specific. Decision deferred.

### 2. Research doc

Location: `research/agentic-task-patterns.md` (this project output)

Key findings:
- Industry shift: **task + guardrails + definition of done** over heavy upfront planning
- Harper Reed: spec → prompt_plan.md → todo.md → self-correcting execution loop
- Anthropic: Explore → Plan → Implement → Commit; skip planning for obvious fixes; DoD must be programmatically verifiable
- Martin Fowler series: TDD inside agent loop as iterative DoD; harness = context + architectural constraints + garbage collection
- **Our skills already follow this pattern** — `go-behavioral-spec` is a task + guardrails + DoD, not a plan

---

## Where we want to go

The full discipline covers four pillars. We have one skill covering the fourth. The first three are undocumented and have no corresponding skills yet.

### Pillar 1: Domain-rooted naming

**Principle:** Function, type, and variable names should come from the domain model (the ubiquitous language of the problem), not from the implementation mechanism.

- `InitializeEventContext` not `ParseMsg`
- `BatchEventsHandler` not `ProcessBatch`
- Names derived from the proto/domain lexicon in `AGENTS.md`

**Why it matters for agents:** Domain names make the behavioral spec write itself. A precisely named function has an obvious MUST statement. Vague names produce vague tests.

**To do:**
- Write `references/naming.md` inside `go-behavioral-spec/` (or as a standalone) covering the discipline
- Consider a `go-domain-naming` skill: read code + domain model, identify implementation-leaking names, suggest domain-rooted alternatives
- This skill requires the agent to know the domain lexicon — tractable in this codebase (AGENTS.md + protos)

### Pillar 2: Nearly-pure functional core

**Principle:** Logic that decides *what* to do should be separated from logic that *does* it. Core functions take data, return data. Decisions live in the core; I/O lives at the leaves.

**Why it matters:** Pure core is table-drivable. Mixed core resists tables, which signals the design problem.

**To do:**
- Document this in a `references/design.md` (or extend `discipline.md` in the skill)
- Consider a `go-testable-design` skill: review a package, identify side effects entangled with logic, suggest extraction points
- This is a design review skill, not a code-writing skill — output is suggestions, not code

### Pillar 3: Narrow injected interfaces at leaf boundaries

**Principle:** Leaf functions (I/O, storage, API calls) are exposed through narrow interfaces that express exactly what the core needs — not "a Spanner client" but `Write(records) error`. The interface is what makes the leaf injectable and the core testable without infrastructure.

**Why it matters:** Narrow interfaces = trivial stubs. Wide interfaces = complex test setup that resists table-driving.

**To do:**
- Document as part of `design.md`
- Potentially extend the `go-behavioral-spec` skill to flag when a test requires non-trivial stub setup (design signal)

### Pillar 4: Tests as behavioral specification ← **DONE (draft)**

Covered by `go-behavioral-spec`. Known gaps in the current draft:
- Open question: stub placement convention (inline vs `<pkg>_stubs_test.go`)
- Open question: one-file-per-source vs per-concern
- The skill currently lives in `.agents/skills/` — should it move to personal skills?

---

## Connecting thread

The four pillars form a closed loop:

```
Domain names          →  Precise function identity
                          makes MUST statements obvious

Pure functional core  →  Table-drivable tests
                          make case space visible

Narrow interfaces     →  Trivial stubs
                          remove friction from table setup

Behavioral spec tests →  Continuous pressure back on design
                          (resistance to tables = design signal)
```

Each pillar reinforces the others. The testing discipline is the feedback loop that keeps the design discipline honest — continuously, without requiring a separate design review.

---

## Immediate next steps (when resuming)

1. **Decide skill home** — does `go-behavioral-spec` stay in the repo or move to personal skills? If personal, how does the discipline doc travel with it?
2. **Close the open questions** in `discipline.md` — stub placement and file-per-source conventions
3. **Write `references/naming.md`** — the domain naming discipline, with examples from this codebase
4. **Write `references/design.md`** — pure core + narrow interfaces, as a unified design discipline
5. **Evaluate `go-domain-naming` skill** — is it agent-executable, or does it require too much domain context to be reliable?
6. **Evaluate `go-testable-design` skill** — design review, not code writing; what does a useful output look like?
7. **Revisit the research findings** — specifically: are our `### Verify` sections in breakdown.md programmatically executable? Should we adopt stop hooks?

---

## Files produced this session

| File | Location |
| ---- | -------- |
| `go-behavioral-spec/SKILL.md` | `.agents/skills/go-behavioral-spec/` (committed) |
| `go-behavioral-spec/references/discipline.md` | `.agents/skills/go-behavioral-spec/references/` (committed) |
| `agentic-task-patterns.md` | `research/` (this project output) |
| This plan | `plans/stephen-workflow-discipline.md` (this project output) |
