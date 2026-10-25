---
type: decision
title: "ADR — Requirements Gathering Format for domain-use-cases skill"
description: Which requirements format (or hybrid) should the domain-use-cases skill be built around to best surface domain concepts for downstream artefacts?
tags: [adr, requirements, use-cases, skill-design, issue-3]
timestamp: 2026-10-25T03:14:00Z
---

# ADR: Requirements Gathering Format for `domain-use-cases` skill

**Status:** PROPOSED — options under review  
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

*Pending pros/cons review.*

---

## Consequences

*To be completed after decision.*
