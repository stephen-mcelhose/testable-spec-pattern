# AGENTS.md — docs/use-cases/

This directory documents a software engineer's workflow as progressive-discovery use cases.
Read this before adding, moving, or renaming anything here.

## The organizing principle: work, not taxonomy

Every file here serves a **named use case**. The primary unit is:

> As a [actor], I want to [Verb Noun] from [Input] in order to inform [Output].

Files exist because a use case references them — not because a concept deserves documentation.
Before adding a file, identify the use case it serves and confirm it appears in the relevant
`index.md`. If you cannot point to a use case, do not add the file.

## Document types

Every `.md` file in this directory (except this one) MUST have one of these types,
expressed as the OKF `type` frontmatter field.

| Type         | `type` value | What it is                                                                    | Required sections                                                |
| ------------ | ------------ | ----------------------------------------------------------------------------- | ---------------------------------------------------------------- |
| **use-case** | `use-case`   | Progressive-discovery chain of actor → goal → input → output statements       | `##` for phases, `###` for individual use cases                  |
| **concept**  | `concept`    | What a named artifact IS — definition, inputs, outputs, known forms            | `## Inputs`, `## Outputs`, `## Known forms` (where applicable)   |
| **method**   | `how-to`     | How to produce an artifact using a specific technique or format               | Prose + examples; MUST link to the parent concept doc            |

No other types belong here. Research lives in `docs/research/`. Decisions in
`docs/decisions/`. Plans in `docs/plans/`.

OKF frontmatter minimum fields: `type`, `title`, `description`, `timestamp`.
See [conventions.md](../../conventions.md) for full field requirements.

## Structure

Each phase folder contains:

- **`index.md`** (`use-case`) — progressive-discovery use cases for that phase
- **Concept docs** (`concept`) — one per named artifact the phase produces.
  Named as singular nouns: `plan.md`, `specification.md`, `domain-model.md`.
- **Method docs** (`how-to`) — one per specific technique for producing a concept.
  Named after the technique: `model-in-go.md`, `rfc2119.md`, `proto-contracts.md`.
  MUST open with a link to the parent concept doc.

**Known forms tables:** every concept doc that has multiple format options MUST include
a Known forms table. Every external link in that table MUST point to Wikipedia or the
format's canonical source. Do not add rows with unverified links.

## What `tree .` should tell you

Running `tree docs/use-cases` should show what this workflow **can do** — not what was
explored while building it. If a file would confuse a new reader about the project's
capabilities, it does not belong here.

## Rules

- **One folder per lifecycle phase.** No sub-phases unless a use case introduces one.
- **Concept docs are singular nouns.** `plan.md` not `planning.md`. `domain-model.md` not `domain-modeling.md`.
- **Method docs are named after the technique.** `model-in-go.md`, `rfc2119.md`, `proto-contracts.md`.
- **No speculative files.** Use case first, concept doc second.
- **No duplicate coverage.** One source of truth per concept, linked everywhere else.
- **Reorganize around use cases.** The structure follows the work, not familiarity.
