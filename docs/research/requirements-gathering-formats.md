---
type: spike
title: "Requirements Gathering Formats — Options for Issue #3"
description: Survey of industry formats for requirements gathering, evaluated against the goal of surfacing domain concepts that feed model-as-code, behavioral specs, and proto contracts.
tags: [requirements, use-cases, domain-concepts, skill-design]
timestamp: 2026-10-25T03:14:00Z
---

# Requirements Gathering Formats

Research for [issue #3](https://github.com/stephen-mcelhose/testable-spec-pattern/issues/3) — selecting a format and approach for the `domain-use-cases` skill.

## The goal

The skill MUST help gather requirements for a project. It MAY directionally point to runnable specs.

The primary value is **concept identification**: surfacing the named domain concepts (actors, objects, lifecycle outcomes, events) that downstream artefacts — `model-as-code`, `go-behavioral-spec`, `resource-oriented-design` — will consume. The requirements format is the vehicle; the concept docs are the deliverable.

---

## Option 1: Use-Cases as Concept Extraction (our approach)

See [issue comment](https://github.com/stephen-mcelhose/testable-spec-pattern/issues/3#issuecomment-6392794343).

**Core idea:** Format-agnostic. Whatever style the author writes in, the skill's job is to:
1. Identify named concepts (actors, objects, lifecycle outcomes, events)
2. Produce a **concept document** for each one

**Concept doc shape:**
```markdown
## Concept: Order
**Kind:** Lifecycle Object
**Definition:** A customer's commitment to purchase specific items from a restaurant.
**Also known as:** (synonyms to avoid — canonicalize one name)
**Lifecycle states:** Placed → Accepted → Picked Up → Delivered
**Produced by:** Checkout
**Consumed by:** AssignCourier, DeliverOrder
**Relations:** belongs to Customer, sourced from Restaurant, fulfilled by Courier
```

**Concept docs are format-agnostic** — they can be extracted from any use-case style. They also double as a domain glossary, an LLM wiki, and a proto resource inventory.

**Downstream feeds:**
- `model-as-code`: each concept → candidate type
- `go-behavioral-spec`: lifecycle states → test scenario names
- `resource-oriented-design`: concepts → proto resources and methods
- `domain-gherkin-scenarios` (#4): concept vocabulary anchors scenario language

**Dogfooding path:** Write the first use-cases about this repo's own workflow. Concepts produced: Skill, Eval, Use-Case, Concept, Contract, Spec, Implementation.

**Open questions:**
- What is the minimum viable concept doc?
- Should concept docs live as files in the repo?
- How does the skill handle concept conflicts across use-cases?

---

## Option 2: Cockburn Fully-Dressed Use Cases

**Origin:** Alistair Cockburn, *Writing Effective Use Cases* (2000). Widely adopted in enterprise software.

**Format:**
```
Use Case: <name>
Actor: <primary actor>
Goal: <in user's words>
Preconditions: <what must be true>
Main Success Scenario:
  1. Actor does X
  2. System responds with Y
  ...
Extensions:
  2a. If Y fails: ...
Technology & Data Variations: ...
```

**Concept surfacing:** Actors named explicitly. Domain objects appear in step verbs and nouns. Alternate flows surface error states and edge concepts. Not structured to extract concepts — you read them out of the prose.

**Strongest for:** Complex business processes with many alternate flows and stakeholder concerns. Traditional BA/PM workflow.

**Weakness for this skill:** Very verbose. The extension numbering (2a, 2a.1, 2a.1.i) is hostile to agents and hard to parse. No explicit concept extraction step. Goal is documentation, not vocabulary construction.

---

## Option 3: User Stories (Agile/XP)

**Origin:** Kent Beck / Ron Jeffries, Extreme Programming. Ron Jeffries' "3 Cs" (Card, Conversation, Confirmation). Widely used in Agile/Scrum.

**Format:**
```
As a [role]
I want [goal/feature]
So that [benefit]

Acceptance criteria:
- Given... When... Then...
```

**Concept surfacing:** Roles (actors) named in the "As a" clause. Goals tend toward features rather than domain objects ("I want to see a list of orders" leaks UI). Benefits in "So that" surface outcomes but often in user-value language, not domain language.

**Strongest for:** Product backlogs, sprint planning, stakeholder communication. Universally understood.

**Weakness for this skill:** The "I want [feature]" clause is implementation-leaning — it names UI elements and system actions, not domain concepts. Acceptance criteria in Given/When/Then (Gherkin) is a separate skill (#4). Personas ("As a premium user") can be vague and may not map to a named domain type.

---

## Option 4: Job Stories (Jobs-to-be-Done)

**Origin:** Paul Adams (Intercom), derived from Clayton Christensen's Jobs-to-be-Done theory. Popularised as an alternative to User Stories for complex or context-sensitive systems.

**Format:**
```
When [situation/trigger]
I want to [motivation/capability]
So I can [outcome/benefit]
```

**Concept surfacing:**
- **When** clause → domain events and triggers (strong)
- **I want to** clause → capabilities/goals (avoids feature bias)
- **So I can** clause → outcomes and post-conditions (maps to `DeliveryConfirmation`-style types)

Job Stories naturally surface **domain events** (the trigger) and **success states** (the outcome) without committing to UI or implementation. The "When" clause often names the exact event that EventStorming would put on an orange sticky.

**Strongest for:** Systems with complex contextual triggers (IoT, real-time, event-driven). Surfaces event vocabulary cleanly.

**Weakness for this skill:** Loses explicit actor naming — "I" is the implicit actor, not a named type. Doesn't produce a structured concept vocabulary on its own. Works best in combination with an explicit concept-extraction step.

---

## Option 5: EventStorming (Alberto Brandolini)

**Origin:** Alberto Brandolini, 2012. Rooted in Domain-Driven Design. Workshop format for rapid domain discovery.

**Format:** Collaborative post-it exercise on a long wall:
- 🟠 **Domain Events** (past tense): `OrderPlaced`, `CourierAssigned`, `OrderDelivered`
- 🔵 **Commands** (imperative): `PlaceOrder`, `AssignCourier`, `DeliverOrder`
- 🟡 **Aggregates**: `Order`, `Courier`, `Customer`
- 🔴 **Actors**: `Customer`, `Restaurant`, `Courier`
- 🟣 **External Systems**: third-party services involved
- 🔴 **Hotspots** (problems): conflicts, ambiguities, open questions

**Flow:** Events → Commands → Aggregates → Bounded Contexts

**Concept surfacing:** The strongest of all options. Events, commands, aggregates, and actors are all explicitly named in a structured sequence. The output maps directly to:
- Events → domain event types / Kafka topics
- Aggregates → model-as-code types
- Commands → proto methods / function names
- Actors → actor types

**Strongest for:** DDD bounded context discovery, microservice decomposition, event-sourced systems. Produces the richest vocabulary.

**Weakness for this skill:** Workshop format — designed for a group of humans in a room. Hard to execute solo or agent-first. Requires domain experts and developers together. Output lives on a wall (or Miro board), not as structured text. Significant facilitation overhead.

---

## Option 6: EARS (Easy Approach to Requirements Syntax)

**Origin:** Alistair Mavin, Rolls-Royce PLC. Structured natural language templates for testable system requirements.

**5 patterns:**
```
Ubiquitous:    The <system> shall <response>.
Event-driven:  When <trigger>, the <system> shall <response>.
State-driven:  While <state>, the <system> shall <response>.
Exception:     If <condition>, then the <system> shall <response>.
Optional:      Where <feature>, the <system> shall <response>.
```

**Concept surfacing:** System name (one named system). Triggers (domain events). States. Responses (capabilities). Conditions (exception paths). Structured enough to be machine-parsed.

**Strongest for:** Safety-critical, embedded, regulated systems (avionics, automotive, medical devices). Produces highly testable, unambiguous requirements. Low learning curve.

**Weakness for this skill:** "The `<system>` shall" framing treats the whole system as one actor — loses actor diversity. Focused on system behaviour, not domain object vocabulary. `shall` instead of RFC 2119 `MUST` is a cosmetic mismatch. Designed for system requirements, not domain discovery.

---

## Comparison

| Dimension                        | Option 1 (Concept Extraction) | Option 2 (Cockburn) | Option 3 (User Stories) | Option 4 (Job Stories) | Option 5 (EventStorming) | Option 6 (EARS) |
|----------------------------------|:-----------------------------:|:-------------------:|:-----------------------:|:---------------------:|:------------------------:|:---------------:|
| Surfaces actor types             | ✅                             | ✅                  | ✅ (weak)               | ❌ (implicit "I")     | ✅                        | ❌               |
| Surfaces domain object types     | ✅ (explicit extraction)       | ⚠️ (in prose)      | ❌                      | ⚠️ (in outcome)       | ✅ (aggregates)           | ⚠️               |
| Surfaces domain events           | ✅                             | ⚠️                 | ❌                      | ✅ (When clause)       | ✅ (primary output)       | ✅ (When)        |
| Surfaces lifecycle/outcomes      | ✅                             | ✅ (extensions)     | ⚠️                      | ✅ (So I can)          | ✅                        | ⚠️               |
| Produces structured concept docs | ✅ (by design)                 | ❌                  | ❌                      | ❌                     | ⚠️ (sticky notes)        | ❌               |
| Agent-executable (solo)          | ✅                             | ✅                  | ✅                      | ✅                     | ❌ (workshop)             | ✅               |
| Feeds model-as-code directly     | ✅                             | ⚠️                 | ❌                      | ⚠️                    | ✅                        | ❌               |
| Points toward runnable specs     | ✅ (via concept docs)          | ⚠️ (via extensions) | ✅ (AC → Gherkin)       | ⚠️                    | ⚠️                        | ✅ (testable)    |
| Non-engineer readable            | ✅                             | ⚠️ (verbose)       | ✅                      | ✅                     | ✅ (visual)               | ⚠️               |
| Format-agnostic                  | ✅ (by design)                 | ❌                  | ❌                      | ❌                     | ❌                        | ❌               |

---

## Recommendation

No single format is ideal on its own. Two plausible paths:

**Path A — Option 1 + Job Stories as input format**
Job Stories (Option 4) surface events and outcomes cleanly; pair with Option 1's concept extraction step to produce structured concept docs. The "When" clause names domain events; the "So I can" clause names outcome types. Explicit concept extraction adds what Job Stories lack (named actors, named objects).

**Path B — Option 1 + lightweight EventStorming vocabulary**
Borrow EventStorming's colour-coded artifact vocabulary (events in past tense, commands in imperative, aggregates as nouns) without the workshop format. Agent produces an ordered list of: events → commands → aggregates → actors. Concept docs follow from aggregates and actors.

Both paths are format-agnostic at the concept extraction layer and produce the same downstream feeds.

---

*Research status: initial survey. Pros/cons comparison pending agent review.*
