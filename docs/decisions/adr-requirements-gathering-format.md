---
type: decision
title: "ADR — Requirements Gathering Format for domain-use-cases skill"
description: Which requirements format (or hybrid) should the domain-use-cases skill be built around to best surface domain concepts for downstream artefacts?
tags: [adr, requirements, use-cases, skill-design, issue-3]
timestamp: 2026-10-25T03:14:00Z
---

# ADR: Requirements Gathering Format for `domain-use-cases` skill

**Status:** ACCEPTED  
**Issue:** [#3 Add write domain-focused use-cases skill](https://github.com/stephen-mcelhose/testable-spec-pattern/issues/3)  
**Research:** [requirements-gathering-formats.md](../research/requirements-gathering-formats.md)

---

## Context

We are building a `domain-use-cases` skill (issue #3) to help gather requirements for a project. The skill has two hard constraints:

- **MUST** help gather requirements
- **MAY** directionally point to runnable specs

The underlying insight: the primary value of the skill is not the use-case text — it is **surfacing named domain concepts** (actors, objects, lifecycle outcomes, events) that feed downstream artefacts:

- `model-as-code` → types
- `go-behavioral-spec` → test subjects and scenario names
- `resource-oriented-design` → proto resources and methods
- `domain-gherkin-scenarios` (#4) → scenario vocabulary

The format question is: what requirements style best surfaces those concepts with the least overhead, in a way an agent can execute solo?

---

## Options

### Option 1: Concept-extraction-first (our approach)
Format-agnostic. Whatever style the author writes in, the skill explicitly extracts named concepts and produces a **concept document** for each. Concept docs are the primary deliverable; the use-case text is the input vehicle.

*See [issue comment](https://github.com/stephen-mcelhose/testable-spec-pattern/issues/3#issuecomment-6392794343) for full detail.*

### Option 2: Cockburn Fully-Dressed Use Cases
Traditional enterprise format. Actor, Goal, Preconditions, Main Success Scenario (numbered steps), Extensions. Rich but verbose. Concept vocabulary read out of prose — not extracted explicitly.

### Option 3: User Stories (Agile/XP)
"As a [role] I want [goal] so that [benefit]" + acceptance criteria. Universally understood. Implementation-leaning ("I want to see a list") — tends to name features rather than domain objects. Acceptance criteria can be Gherkin (covered by skill #4).

### Option 4: Job Stories (Jobs-to-be-Done)
"When [situation] I want to [motivation] so I can [outcome]." Surfaces domain events (When) and success states (outcome) cleanly. Avoids persona bias. Loses explicit actor naming. Works well paired with a concept extraction step.

### Option 5: EventStorming vocabulary (adapted)
Borrows EventStorming's artifact taxonomy (domain events in past tense, commands in imperative, aggregates as nouns, actors) without the workshop format. Agent produces an ordered list: events → commands → aggregates → actors. Richest concept vocabulary. Requires DDD familiarity.

### Option 6: EARS (Easy Approach to Requirements Syntax)
Structured templates: `When <trigger>, the <system> shall <response>`. Highly testable. Designed for safety-critical system requirements. "The system" framing loses actor diversity; not suited to domain concept discovery.

---

## Decision

### Hard-constraint check

Both constraints apply to every option before scoring:

| Option                              | MUST help gather requirements | MAY point to runnable specs | Passes? |
|-------------------------------------|:-----------------------------:|:---------------------------:|:-------:|
| 1 — Concept extraction              | ✅ (any input vehicle)         | ✅ lifecycle states → names  | ✅       |
| 2 — Cockburn Fully-Dressed          | ✅                             | ⚠️ extensions → edge cases  | ✅       |
| 3 — User Stories                    | ✅                             | ✅ AC → Gherkin              | ✅       |
| 4 — Job Stories                     | ✅                             | ⚠️ outcome → state names    | ✅       |
| 5 — EventStorming (adapted, text)   | ✅                             | ✅ events/aggregates → types | ✅       |
| 6 — EARS                            | ✅                             | ✅ `shall` → testable reqs  | ✅       |
| Path A — Option 1 + Job Stories     | ✅                             | ✅                           | ✅       |
| Path B — Option 1 + EventStorming   | ✅                             | ✅                           | ✅       |

All options clear the constraints. Differentiation comes from the five evaluation dimensions.

---

### Evaluation: five dimensions

Scoring key: **●●●** strong &nbsp;·&nbsp; **●●○** moderate &nbsp;·&nbsp; **●○○** weak &nbsp;·&nbsp; **○○○** absent

#### Dimension 1 — Concept surface area
*How many of the four concept kinds (actors, objects, events, outcomes) does the format naturally name?*

| Option                            | Actors       | Objects      | Events       | Outcomes     | Score  |
|-----------------------------------|:------------:|:------------:|:------------:|:------------:|:------:|
| 1 — Concept extraction            | ●●● explicit | ●●● explicit | ●●● explicit | ●●● explicit | **4/4** |
| 2 — Cockburn                      | ●●● named    | ●●○ in prose | ●●○ implicit | ●●● main flow| **3/4** |
| 3 — User Stories                  | ●●○ persona  | ○○○          | ○○○          | ●●○ benefit  | **1.5/4** |
| 4 — Job Stories                   | ○○○ "I"      | ●○○ outcome  | ●●● When     | ●●● So I can | **2/4** |
| 5 — EventStorming (adapted)       | ●●● named    | ●●● aggregates | ●●● past-tense | ●●● states | **4/4** |
| 6 — EARS                          | ○○○ "system" | ○○○          | ●●● When     | ●●○ response | **1.5/4** |
| **Path A** — 1 + Job Stories      | ●●○ prompted | ●●● extracted | ●●● When    | ●●● So I can | **3.5/4** |
| **Path B** — 1 + EventStorming    | ●●● named    | ●●● aggregates | ●●● past-tense | ●●● states | **4/4** |

Options 1, 5, and Path B are the only ones that naturally name all four concept kinds. Option 1 does it through extraction (needs input); Option 5 and Path B do it structurally. Path A loses actor explicitness — the "I" must be elicited by the extraction step.

---

#### Dimension 2 — Agent-executability
*Can a single agent produce useful output from this format solo, without a workshop?*

| Option                          | Verdict | Rationale |
|---------------------------------|:-------:|-----------|
| 1 — Concept extraction          | ●●●     | Reads any input, extracts via defined schema; no prerequisites |
| 2 — Cockburn                    | ●●○     | Agent can produce the format but the extension-numbering tree (2a, 2a.1, 2a.1.i) is structurally complex; verbosity creates noise without a human to trim it |
| 3 — User Stories                | ●●●     | Trivially producible; but low-quality output for domain concept work |
| 4 — Job Stories                 | ●●●     | Three-clause structure; clean to produce and review |
| 5 — EventStorming (adapted)     | ●●●     | Text-based adaptation removes the workshop dependency; agent produces ordered list: events → commands → aggregates → actors |
| 6 — EARS                        | ●●●     | Template-fill; easy to produce, easy to validate mechanically |
| **Path A** — 1 + Job Stories    | ●●●     | Two clean steps: write Job Stories, then run extraction |
| **Path B** — 1 + EventStorming  | ●●●     | Two clean steps: produce EventStorming artifact list, then produce concept docs |

EARS and User Stories score high here but their outputs are weak on downstream value. The meaningful differentiator in this dimension is Cockburn, which is the only format that is structurally hostile to agent execution.

---

#### Dimension 3 — Domain language discipline
*Does the format resist implementation-leaning vocabulary?*

| Option                          | Verdict | Rationale |
|---------------------------------|:-------:|-----------|
| 1 — Concept extraction          | ●●●     | Extraction step explicitly surfaces and challenges non-domain names; concept doc schema forces named types |
| 2 — Cockburn                    | ●●○     | Format doesn't constrain vocabulary in steps — "System sends HTTP 200" and "UI renders dropdown" can appear without violating the format |
| 3 — User Stories                | ●○○     | "I want to see a list", "I want to filter by status" — the "I want" clause is an implementation magnet; feature language is the norm, not the exception |
| 4 — Job Stories                 | ●●●     | "When [situation]" stays situational and trigger-level; avoids feature and UI language by design; no slot for implementation intent |
| 5 — EventStorming (adapted)     | ●●●     | DDD naming conventions act as structural constraints: past-tense events (`OrderPlaced` not `callPlaceOrderEndpoint`), noun aggregates, imperative commands — violations are obvious |
| 6 — EARS                        | ●●○     | "The system shall" constrains to behaviour but loses actor and domain object vocabulary; templates don't resist feature language in the response clause |
| **Path A** — 1 + Job Stories    | ●●●     | Job Stories surface situational triggers in domain language; extraction step reviews for drift — two filters |
| **Path B** — 1 + EventStorming  | ●●●     | EventStorming vocabulary is the most structurally resistant of any format; extraction adds a second review pass |

User Stories (Option 3) is the only option that actively invites implementation vocabulary. Path A and Path B are the strongest — both have two resistance layers (format + extraction).

---

#### Dimension 4 — Path to runnable specs
*How directly does the output feed `go-behavioral-spec` test names or Gherkin scenarios?*

| Option                          | Verdict | Rationale |
|---------------------------------|:-------:|-----------|
| 1 — Concept extraction          | ●●●     | Lifecycle states → scenario names; concept types → test subjects; one indirection (needs concept docs first) |
| 2 — Cockburn                    | ●●○     | Extensions → edge cases, but the extension tree is not structured for Gherkin or spec names; parsing overhead |
| 3 — User Stories                | ●●○     | AC → Gherkin is direct, but Gherkin is skill #4; the use-case text itself is weak on domain concept names that would become test subject identifiers |
| 4 — Job Stories                 | ●●○     | Outcome → success state names; trigger → Gherkin `When` clause; requires extraction step to complete the mapping |
| 5 — EventStorming (adapted)     | ●●●     | Domain events directly name test scenario triggers (`TestOrderPlaced`); aggregates → test subjects; commands → SUT method names; most direct structural mapping |
| 6 — EARS                        | ●●○     | Requirements are testable by design; `shall` → `MUST`; but "the system" is not a typed domain subject — test names would be generic |
| **Path A** — 1 + Job Stories    | ●●●     | Outcome → success states → `go-behavioral-spec` scenario names; When → Gherkin `When`; concept docs bridge to test subjects |
| **Path B** — 1 + EventStorming  | ●●●     | Strongest: events directly name test scenario triggers; aggregates name test subjects; concept docs confirm the mapping; fewest translation steps |

Path B produces the most literal mapping to test names. `CourierAssigned` (event) → `TestCourier_Assigned` (spec test) is a rename, not a design step.

---

#### Dimension 5 — Non-engineer legibility
*Can a product owner or domain expert read and validate the output?*

| Option                          | Verdict | Rationale |
|---------------------------------|:-------:|-----------|
| 1 — Concept extraction (docs)   | ●●●     | Concept docs are readable markdown; the domain glossary shape is familiar to any business stakeholder |
| 2 — Cockburn                    | ●○○     | The extension-numbering system (2a, 2a.1) is hostile to non-engineers; verbosity loses attention; perceived as "developer documentation" |
| 3 — User Stories                | ●●●     | Universally understood; anyone who has worked in Agile can read and validate |
| 4 — Job Stories                 | ●●●     | Simple, outcome-focused, business language; easier to read than Cockburn, comparable to User Stories |
| 5 — EventStorming (intermediate)| ●●○     | Event lists (`OrderPlaced`, `CourierAssigned`) are readable; "aggregate" is unfamiliar to non-engineers; concept docs from extraction step are fully legible |
| 6 — EARS                        | ●●○     | `shall` language reads as legal/formal; "the system" loses business context; templates are readable but sterile |
| **Path A** — 1 + Job Stories    | ●●●     | Job Stories are accessible; concept docs are accessible; both artifacts legible without DDD background |
| **Path B** — 1 + EventStorming  | ●●○     | Intermediate EventStorming list requires terminology explanation ("aggregate"); concept docs are fully legible; overall requires one DDD glossary note in the skill |

Path A has a clear advantage here if the audience lacks DDD familiarity. Path B can close the gap with a one-paragraph terminology note bundled in the skill.

---

### Summary scorecard

| Option                          | Concept area | Agent-exec | Domain discipline | Spec path | Legibility | **Total** |
|---------------------------------|:------------:|:----------:|:-----------------:|:---------:|:----------:|:---------:|
| 1 — Concept extraction          | ●●●●         | ●●●        | ●●●               | ●●●       | ●●●        | **strong** |
| 2 — Cockburn                    | ●●●○         | ●●○        | ●●○               | ●●○       | ●○○        | **weak**   |
| 3 — User Stories                | ●●○○         | ●●●        | ●○○               | ●●○       | ●●●        | **weak**   |
| 4 — Job Stories                 | ●●○○         | ●●●        | ●●●               | ●●○       | ●●●        | **moderate** |
| 5 — EventStorming (adapted)     | ●●●●         | ●●●        | ●●●               | ●●●       | ●●○        | **strong** |
| 6 — EARS                        | ●●○○         | ●●●        | ●●○               | ●●○       | ●●○        | **weak**   |
| **Path A** — 1 + Job Stories    | ●●●○         | ●●●        | ●●●               | ●●●       | ●●●        | **strong** |
| **Path B** — 1 + EventStorming  | ●●●●         | ●●●        | ●●●               | ●●●       | ●●○        | **strongest** |

---

### Ranked recommendation

**Rank 1 — Path B: Option 1 + EventStorming vocabulary (adapted, text-based)**

Path B is the only option that scores strong across all five dimensions simultaneously. The EventStorming vocabulary acts as a structural discipline that enforces domain language at the point of authoring — before extraction runs. The four concept kinds are not inferred; they are the named output slots. The mapping from EventStorming artifacts to downstream feeds (`go-behavioral-spec`, `model-as-code`, `resource-oriented-design`) is the most direct of any option: domain event names become test scenario names with a rename. The one legibility gap ("aggregate" terminology) is addressable with a glossary note inside the skill — it is not architectural.

**Rank 2 — Path A: Option 1 + Job Stories**

Path A scores nearly as well and has a legibility advantage. If the intended author population has no DDD background, Path A is the safer choice — Job Stories are universally understood and the "When/I want/So I can" structure stays domain-focused without requiring vocabulary onboarding. The cost is actor coverage: the implicit "I" must be explicitly elicited during extraction, adding an extra prompt step. Path A is the fallback if Path B's terminology proves a barrier in practice.

**Rank 3 — Option 4: Job Stories alone**

Useful as a lightweight standalone for small-scope use-cases where concept docs are not needed. Does not scale to the full downstream feed requirement.

**Options 2, 6 — Eliminated**

Cockburn is too verbose and structurally hostile to agent execution. EARS erases actor diversity, which is fatal for the domain concept extraction goal.

**Option 3 — Eliminated**

User Stories are the only format that actively invites implementation vocabulary. The "I want" clause is an implementation magnet. Concept extraction from User Stories is fighting the format.

---

---

## Decision

**We adopt stripped-down Cockburn use-cases as the input format, not Path B.**

The research ranked Path B (EventStorming vocabulary + concept extraction) highest. We deviated from that recommendation for the following reasons:

**Use-case orientation fits the goal better than EventStorming.** Cockburn use-cases are system-focused — they describe what happens in the domain when an actor pursues a goal, not what a user desires. That orientation naturally resists implementation vocabulary and naturally surfaces domain events and objects as named nouns and verbs in the main flow steps. EventStorming's four artifact slots (event, actor, aggregate, command) are the same four concept kinds we extract — but they require a separate workshop artifact as input. The Cockburn main flow produces an equivalent input without the extra step.

**Stripped-down Cockburn removes the verbosity problem.** The research ranked Option 2 (Cockburn) poorly on agent-executability due to extension-numbering complexity. We remove that: no extensions, no alternate flows, no preconditions in v0.1. The format is: Use Case name, Actor, Goal, and a numbered main flow. That is structurally as clean as Job Stories.

**Actor explicitness is preserved.** The "I" problem that eliminated Job Stories and User Stories is solved — Actor is a named first-class field in the format header, not embedded in a sentence.

**DDD terminology is deferred.** Path B's legibility gap (the word "aggregate") disappears entirely. Concept kinds are named in plain English: Actor, Domain Object, Domain Event, Action. The glossary carries the DDD precision without the onboarding overhead.

**The format chosen:**

```
Use Case: [verb phrase in domain language]
Actor:    [named actor type]
Goal:     [what is true when this succeeds]

Main flow:
  1. [step in domain language]
  2. [step in domain language]
  ...
```

---

## Consequences

### What this means for the skill

The `domain-use-cases` skill operates in two phases:

**Phase 1 — Use-case authoring.** The author (human or agent) writes use-cases in stripped-down Cockburn format. Each use-case names an Actor, states a Goal, and lists a main flow of domain-language steps. No extensions, no alternate flows.

**Phase 2 — Concept extraction.** Each named noun and verb in the use-cases is a candidate concept. The skill produces one concept document per concept with OKF frontmatter, a kind tag (`entity`, `value-object`, `domain-event`, `actor`, `action`), a one-sentence domain-language definition, and lifecycle states for domain objects.

### Trade-offs accepted

| Trade-off | Accepted because |
|-----------|-----------------|
| Use-case main flow is less structurally disciplined than EventStorming vocabulary | The goal / main flow combination is sufficient to distinguish domain events (what happened) from actions (what the actor did) — violations are visible and correctable in the extraction step |
| No alternate flows or extensions in v0.1 | Keeps the format agent-executable; extensions can be added if evals show gaps |
| Domain object / entity distinction deferred to tagging | The glossary carries the distinction; authors use `entity` when uncertain and refine later |

### Downstream feed mapping

| Concept kind  | `model-as-code`     | `go-behavioral-spec`           | `resource-oriented-design` | `domain-gherkin-scenarios` (#4) |
|---------------|---------------------|-------------------------------|----------------------------|----------------------------------|
| Domain event  | event type          | scenario name (`TestXxx_Yyy`) | n/a (cross-cutting)        | `When` / `Then` clause           |
| Actor         | actor type          | test actor setup              | caller role                | `Given` actor context            |
| Domain object | struct type         | subject under test            | proto resource             | domain noun in scenario          |
| Action        | method signature    | SUT function name             | proto method               | `When` clause verb               |

### Revisit criteria

Reopen this decision if any of the following occur:

- **ROD translation issues.** If translating concept documents to Resource Oriented Design (proto resources and methods) causes systematic friction — e.g., actors don't map to caller roles, or domain objects don't map cleanly to proto resources — the format may not produce the right granularity of concepts.
- **Implementation vocabulary leaks.** If evals consistently show main flow steps containing HTTP verbs, endpoint paths, or CRUD language despite the format's constraints, the Cockburn main flow is not providing sufficient discipline and EventStorming vocabulary should be reconsidered.
- **Concept extraction ambiguity.** If it is regularly unclear which nouns in the main flow are domain objects versus wire shapes, the extraction step needs a more structured input (EventStorming's explicit artifact slots).
- **Actor granularity mismatch.** If the same actor name covers meaningfully different roles in practice (as the petstore `User` resource covers both Customer and StoreStaff), the format needs an actor decomposition step before use-case authoring.
- **Alternate flows are consistently missing.** If evals show that skipping extensions causes concept gaps — concepts that only appear in failure paths — add extensions as an optional Phase 1 step.
