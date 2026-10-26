# AGENTS.md

Steering document for any agent working in this repo. Read this first.

> [!WARNING]
> **TODO — Issue [#48](https://github.com/stephen-mcelhose/testable-spec-pattern/issues/48)**
>
> The project structure, documentation, and skills have not yet been trimmed to what is actually useful.
> There is a `sketch/` directory (gitignored) exploring what the project should look like organized by
> lifecycle stage. Open questions include: where does shared vocabulary (concepts) live, what documentation
> is genuinely user-facing vs. build artifact, and which skills earn a place in the final layout.
>
> **Do not treat the current `docs/` and `.agents/skills/` layout as settled.** Revisit issue #48 before
> adding new documentation or skills.

## What this repo is

A personal workshop for building a **Go behavioral-spec workflow** — tests that read like RFC 2119 contracts, grounded in domain names that bridge spec, implementation, and test.

## North star

**The goal is a workflow expressed as composable, agentskills-standard skills.** Every skill in `.agents/skills/` should be a valid `agentskills.io` skill — proper frontmatter, `allowed-tools`, narrow focus, and a clear trigger description. Skills that don't meet the standard should be treated as drafts.

Skills should be composable in the style of [Matt Pocock's skills repo](https://github.com/mattpocock/skills): narrow discipline skills (model-invoked) composed by orchestrator skills (user-invoked). An orchestrator can call a discipline skill; disciplines never call each other. If a skill's purpose contains "and", split it. See [`docs/research/composable-skills.md`](docs/research/composable-skills.md) for the full research.

## Workflow

This repo uses GitHub Issues as a lightweight Kanban board. The workflow applies equally to human contributors and AI agents — every step must be executable by either.

### Board state

| Issue state       | Kanban column   |
| ----------------- | --------------- |
| Open, unassigned  | Backlog         |
| Open, assigned    | In Progress     |
| Closed            | Done            |

**WIP limit: 1.** Only one issue In Progress at a time. Finish or hand off before pulling the next.

### Start of session

1. **Resume before starting.** Check for any issue already assigned to you (`gh issue list --assignee @me`). If one exists, check out its branch (`git checkout {n}-{slug}`) and continue.
2. **Pick from backlog.** If nothing is in progress, look at open unassigned issues (`gh issue list`). Human contributors set priority. In the absence of explicit priority, choose the most unblocked issue — the one with the fewest unresolved upstream dependencies.
3. **Ask: is this ready to work?** Apply the Definition of Ready below. If the issue is not ready, comment identifying the gap and stop — do not start work on an unready issue.
4. **Signal In Progress** by assigning yourself: `gh issue edit {n} --add-assignee @me`.
5. **Create a branch:** `git checkout -b {n}-{short-slug}` (e.g. `git checkout -b 5-go-scaffold`). All commits go to this branch — never directly to `main`.

### Definition of Ready

An issue is ready to work when all of the following are true:

- [ ] Single deliverable — no bundled concerns
- [ ] The "why" is stated (the pain or gap being addressed)
- [ ] At least one concrete use-case
- [ ] Acceptance criteria are specific and individually checkable
- [ ] Out of scope is explicit
- [ ] Dependencies on other issues are listed; each is either closed or explicitly waived
- [ ] No open questions in the issue that would block progress

If any item is unchecked: comment on the issue with what is missing (`gh issue comment {n} --body "..."`). Do not start work.

### In flight

- All commits go to the issue branch — never directly to `main`.
- Reference the issue in every commit (`refs #N`).
- If you discover scope not in the acceptance criteria, comment on the issue before expanding. Do not silently extend the work.

### End of session

- **Done:** open a PR (`gh pr create --fill`) with `closes #N` in the body. Merge after review.
- **Incomplete:** push the branch (`git push -u origin {n}-{slug}`); comment on the issue with current state and what remains (`gh issue comment {n} --body "..."`). Never leave untracked in-progress state.

---

## Where to look

| Resource | What you'll find |
| -------- | ---------------- |
| [`docs/index.md`](docs/index.md) | Full catalog of the knowledge base |
| [`docs/discipline.md`](docs/discipline.md) | The core behavioral-spec discipline — read before writing any tests |
| [`docs/conventions.md`](docs/conventions.md) | Wiki conventions for `docs/` — OKF frontmatter, folder taxonomy |
| [`docs/research/agentskills-standard.md`](docs/research/agentskills-standard.md) | What the agentskills standard requires (read before creating or editing skills) |

## Skills in this repo

### Exists

| Skill | Purpose |
| ----- | ------- |
| [`go-behavioral-spec`](.agents/skills/go-behavioral-spec/SKILL.md) | Write Go tests as RFC 2119 behavioral specs |
| [`rfc-2119-spec-writer`](.agents/skills/rfc-2119-spec-writer/SKILL.md) | Draft prose technical specs with RFC 2119 keywords (behaviors only) |
| [`docs-wiki`](.agents/skills/docs-wiki/SKILL.md) | Manage the `docs/` knowledge base |

### Planned

| Skill | Purpose |
| ----- | ------- |
| `contract-first-with-domain-names` | Guide contract-first design (OpenAPI, proto, ROD) grounded in domain names — the bridge between spec, implementation, and test |
| `planning` | Consistent planning format with known sections, subdivided by class of problem |

Skills are added when a documented need or failure demands them — not speculatively.

## Skill conventions

- Skills are **model-invoked** unless explicitly marked otherwise
- Skill frontmatter MUST include `name`, `description`, `license`, `allowed-tools`
- `allowed-tools` should be as narrow as possible — don't grant broad `Bash` access when specific patterns will do
- **Skills must be self-contained and packageable.** A skill may compose over other skills in the workflow, but MUST NOT depend on docs, examples, or any other file in this repo. If a skill needs reference material (examples, discipline docs), bundle it inside the skill directory. The goal: any skill could be copy-pasted to a different repo and work as-is.
