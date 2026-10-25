---
name: domain-use-cases
description: Write domain-focused use-cases that surface named domain concepts — actors, domain objects, events, and actions — in the ubiquitous language of the domain. Use when gathering requirements for a feature or system, before writing behavioral specs or designing contracts. Trigger on phrases like "write a use-case for X", "gather requirements for X", "what are the domain concepts for X", "let's spec out what X does".
license: MIT
allowed-tools: Read Glob Grep Write Edit
metadata:
  version: "0.1.0"
---

# Domain Use-Cases

Write use-cases that surface named domain concepts. The use-case text is the vehicle; the concept documents are the deliverable.

See `references/glossary.md` for term definitions and examples of what belongs — and what does not.

---

## Use-case format

Stripped-down Cockburn style. Three fields plus a numbered main flow:

```
Use Case: [Name — verb phrase in domain language]
Actor:    [named actor type]
Goal:     [what is true when this succeeds — outcome, not a step]

Main flow:
  1. [step in domain language]
  2. [step in domain language]
  3. ...
```

**Good:**
```
Use Case: Adopt a Pet
Actor:    Customer
Goal:     Pet is reserved and an order is placed

Main flow:
  1. Customer selects an available pet
  2. Customer places an order for the pet
  3. Pet state changes to pending
  4. Order is confirmed
```

**Bad:**
```
Use Case: Create Order           ← CRUD verb, not domain language
Actor:    User                   ← too generic, not a domain type
Goal:     POST /v1/orders → 201  ← implementation detail

Main flow:
  1. Call CreateOrder endpoint   ← technical step
  2. Return order ID             ← wire detail
```

---

## Concept extraction

After writing use-cases, read each step and field. Every noun that is a named domain thing is a candidate concept. Extract them and produce one concept document per concept.

Four concept kinds:

| Kind          | What it is                                              | Naming convention          |
|---------------|---------------------------------------------------------|----------------------------|
| Actor         | A named person or system that initiates or receives events | Noun (`Customer`)        |
| Domain object | A thing that holds its own state and has a lifecycle    | Noun (`Order`)             |
| Domain event  | Something that happened in the domain                   | Past tense (`OrderPlaced`) |
| Action        | Something an actor does, causing an event               | Imperative (`PlaceOrder`)  |

See `references/glossary.md` for definitions and exclusion examples.

---

## Concept document shape

One file per concept with OKF frontmatter:

```markdown
---
type: concept
title: "[Concept Name]"
description: [one sentence definition in domain language — no implementation terms]
tags: [<kind-tag>, <additional-ddd-tags>]
timestamp: <ISO-8601 UTC>
---

## [Concept Name]

**Kind:** Actor | Domain Object | Domain Event | Action

**Definition:** one sentence in domain language

**Lifecycle states:** (domain objects only) comma-separated named states in domain language
```

### Kind tags

| Kind          | Primary tag    | Additional tags                                  |
|---------------|----------------|--------------------------------------------------|
| Actor         | `actor`        | `external-system` if the actor is a third party  |
| Domain object | `entity`       | `aggregate-root` if it owns other domain objects |
| Domain object | `value-object` | when it has no identity and is immutable         |
| Domain event  | `domain-event` |                                                  |
| Action        | `action`       |                                                  |

When uncertain between `entity` and `value-object`, use `entity` — refine later.

---

## Rules

- Names MUST come from the domain, not from implementation (`Order` not `OrderEntity`, `Customer` not `UserService`)
- Actor names MUST be specific domain roles (`Customer` not `User`, `StoreStaff` not `Admin`)
- Domain object names MUST be nouns; domain event names MUST be past tense; action names MUST be imperative
- Lifecycle state names MUST be in domain language (`placed`, `pending`, `delivered` — not `CREATED`, `ACTIVE`, `DONE`)
- Every use-case MUST produce at least one concept document
- If a definition cannot be written without mentioning a class, method, or endpoint — the name is wrong

---

## Worked example

See `examples/petstore.md`.
