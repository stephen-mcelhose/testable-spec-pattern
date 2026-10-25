---
name: domain-use-case-interview
description: >
  Conduct a domain discovery interview to surface actors, domain objects,
  flows, and events before writing use-cases or designing contracts.
  Use when someone says "help me figure out the domain for X", "I'm not sure
  what the concepts are in X", "let's explore what X involves", or "I want to
  understand the domain before we build". Produces a structured domain summary
  that feeds directly into the domain-use-cases skill.
license: personal
version: 0.1.0
tags:
  - domain-discovery
  - requirements
  - interviewing
user-invocable: true
allowed-tools:
  - Write
---

# Domain Use-Case Interview

A discovery conversation that surfaces domain vocabulary — actors, objects,
events, flows — through natural storytelling. The output feeds into
`domain-use-cases` to produce Cockburn-format use-cases and concept documents.

---

## What this skill does not do

- Write use-cases. Pass the discovery summary to `domain-use-cases` for that.
- Extract or tag concept documents. That too is `domain-use-cases`.
- Ask schema-driven questions ("what states does X go through?"). Structure
  emerges from what the user says; categorisation happens after the story,
  not during it.

---

## The interview protocol

### Phase 1 — Open with a story prompt

Ask one question only. Do not introduce yourself, explain what you are doing,
or ask multiple questions at once.

> "Tell me what happens when [the primary thing this system does]. Walk me
> through it as if I'm watching over someone's shoulder."

If the user has not named the domain's primary scenario, ask first:
> "In one sentence — what does this system do for someone?"
Then follow immediately with the story prompt.

After the user answers, ask one story-extending question. Do not name or
classify anything you heard. Do not say "so [X] would be an actor" or "it
sounds like [Y] has its own lifecycle." Keep a running note of the nouns,
verbs, and state transitions mentioned — you will classify these in Phase 3,
once you have enough context to do so accurately.

---

### Phase 2 — Follow the story

Ask one follow-up question at a time, staying inside the user's narrative.
Good follow-up patterns:

- **Extend the arc:** "And then what? What happens after that?"
- **Unpack a noun:** "You mentioned [X] — what is that exactly?" — this is
  not categorisation. You are using the user's own word back at them and asking
  them to elaborate. Categorisation would be saying "so [X] sounds like it has
  its own lifecycle" or "is [X] a separate thing?" — those assign a category.
  "What is [X]?" just asks the user to keep talking.
- **Surface a second actor:** "You've been describing this from [role A]'s
  perspective. What does [role B] experience at that same moment?"
- **Find the edges:** "What happens if that goes wrong? What does [role] do
  then?"

Stay in the user's vocabulary. If they say "booking," say "booking." If they
say "reservation," say "reservation." Do not switch terms without confirmation.

Continue until you have one complete story arc — from trigger to resolution —
for the primary scenario.

---

### Phase 3 — Draft and probe loop

After hearing the primary story, draft a terse one-liner per use-case extracted
so far. Show it and ask if it is complete.

**Format:**

```
Here's what I have so far:

- [Actor]: [goal, one line]
- [Actor]: [goal, one line]
- [Actor]: [goal, one line]

Is that everything? A few things I'm not sure about:
- [Targeted probe about a specific actor or flow gap]
- [Targeted probe about a specific actor or flow gap]
```

**Targeted probes use subdivision reasoning — not open-ended questions:**

- *Actor subdivision:* "Is that the only thing a [Actor] does, or are there
  other interactions they have with the system?"
- *Flow subdivision:* "You've described what happens when everything goes well.
  What happens when [thing] fails or is rejected?"
- *Hidden object probe:* "You mentioned [noun] several times — does the system
  track that separately, or is it just part of [other thing]?"
- *Actor gap probe:* "Who actually [does the thing just described]? Is that
  [existing actor] or someone else?"

Ask at most two probes per round. Wait for answers before asking more.

Update the draft after each round. Repeat Phase 3 until the user confirms the
draft is complete.

---

### Phase 4 — Write the discovery summary

Once the user confirms the draft is complete, write a structured discovery
summary. This is the output of this skill — it feeds `domain-use-cases` if
the user wants use-cases next, or stands alone as a domain reference.

```markdown
# {Domain} — Discovery Summary

## Actors

| Actor | Goal |
|-------|------|
| {Actor} | {one-line goal in user's words} |

## Scenarios (confirmed)

- {Actor}: {goal}
- {Actor}: {goal}

## Domain objects

| Name | Lifecycle hints |
|------|----------------|
| {Name} | {states mentioned verbatim, e.g. "requested → matched → completed"} |

## Vocabulary

| Term used | Notes |
|-----------|-------|
| {term} | {any synonyms or ambiguities noted} |

## Open questions

{Anything unresolved — gaps the user couldn't answer, terms that need
clarifying, flows that were mentioned but not walked through.}
```

---

## Interview discipline

- **One question at a time.** Never bundle. Each question must be answerable in
  1–3 sentences.
- **Never invent.** Every name in the draft must come from the user's words or
  be explicitly confirmed. If you need a name and the user hasn't given one,
  ask: "what would you call that?"
- **Categorise in Phase 4, not during the conversation.** During Phases 1 and
  2, accumulate what the user says without classifying it. In Phase 3, surface
  candidates for confirmation. In Phase 4, assign categories when writing the
  summary. Classifying mid-conversation ("so that would be a domain object")
  leads the user and produces invented categories rather than discovered ones.
- **No DDD vocabulary in questions.** Do not say "aggregate," "entity," "value
  object," "bounded context," or "command" to the user.
- **No technology vocabulary.** Do not ask about endpoints, tables, services,
  or schemas. Redirect: "Forget the implementation for a moment — what does the
  business need to happen here?"
- **Stop when confirmed.** Once the user says the draft is complete, produce the
  summary and stop. Do not keep asking.

---

## Output

A single discovery summary document: actors, confirmed scenarios, domain objects
with lifecycle hints, vocabulary, and open questions.

Pass it to `domain-use-cases` to write use-cases and extract concept documents.
Or use it as-is — the summary stands alone as a domain reference.
