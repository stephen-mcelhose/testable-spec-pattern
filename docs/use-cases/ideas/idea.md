---
type: concept
title: Idea
description: The amorphic entry point to the workflow — any unstructured starting point sufficient to begin producing requirements.
tags: [idea, entry-point, workflow]
timestamp: 2026-10-26T00:00:00Z
---

# Idea

An Idea is the amorphic entry point to the workflow — any unstructured starting point sufficient to begin producing requirements. It carries no prescribed form. It may be a sentence, a conversation, a slide deck, or a half-finished document. Its only requirement is that it represents a problem or opportunity worth pursuing.

An Idea is distinct from Requirements — it is pre-structured. Once an Idea produces any requirements artifact (use cases, domain types, Gherkin scenarios, etc.) it has served its purpose and become Requirements.

## Inputs

- None — an Idea is the origin. It may be received from a stakeholder, discovered through observation, or self-generated.

## Outputs

An Idea feeds:

- [Requirements Gathering](../requirements-gathering/index.md) — the primary next phase
- [Contracts & Domain Modeling](../contracts-and-domain-modeling/index.md) — if the Idea arrives as a Design Doc or schema sketch

## Known forms

Some document types straddle the Idea/Requirements boundary — they could reasonably be treated as either depending on how much structure they contain. Use whichever phase they most naturally feed.

| Form                                                                                          | Likely feeds                 |
| --------------------------------------------------------------------------------------------- | ---------------------------- |
| Verbal description / conversation                                                             | Idea → Requirements          |
| [Problem Statement](https://en.wikipedia.org/wiki/Problem_statement)                         | Idea → Requirements          |
| [Opportunity Brief](https://en.wikipedia.org/wiki/Opportunity_management)                    | Idea → Requirements          |
| [RFC (Request for Comments)](https://en.wikipedia.org/wiki/Request_for_Comments)             | Idea → Specifications        |
| [PRD](https://en.wikipedia.org/wiki/Product_requirements_document)                           | Idea **or** Requirements     |
| [SRS](https://en.wikipedia.org/wiki/Software_requirements_specification)                     | Idea **or** Requirements     |
| [Design Doc](https://en.wikipedia.org/wiki/Design_document)                                  | Idea → Contracts             |

PRD and SRS are ambiguous — if they arrive with enough structure they feed Requirements directly; if they are a starting sketch they are an Idea.
