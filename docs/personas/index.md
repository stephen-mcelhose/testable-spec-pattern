---
type: concept
title: Personas
description: The actors in the use-cases-progressive-discovery workflow — who they are, what they own, and how to write a persona document.
tags: [personas, actors, use-cases-progressive-discovery]
timestamp: 2026-10-26T00:00:00Z
---

# Personas

A persona is a named actor in the [use-cases-progressive-discovery](../use-cases-progressive-discovery.md) workflow. Each use case begins with an actor:

> As a [**software engineer**](./software-engineer.md), I want to...

Personas are linked, not prose. An actor that appears in a use case MUST have a corresponding persona document.

## Persona document format

A persona document covers:

| Section         | Content                                                                 |
| --------------- | ----------------------------------------------------------------------- |
| **Who they are** | A one-paragraph description of the actor's role and context            |
| **Goals**        | What the actor is trying to achieve in the workflow                    |
| **Produces**     | The artifacts or outputs the actor creates                             |
| **Consumes**     | The inputs the actor depends on from upstream                          |
| **Use cases**    | Links to every use case in which this actor appears                    |

## Personas

| Persona                                          | Role                                                             |
| ------------------------------------------------ | ---------------------------------------------------------------- |
| [Software Engineer](./software-engineer.md)      | Designs, builds, and verifies software through the full workflow |
