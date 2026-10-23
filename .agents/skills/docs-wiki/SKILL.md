---
name: docs-wiki
description: Manage the docs/ knowledge base for this repo. Based on the llm-wiki pattern — the LLM is the programmer, docs/ is the codebase. Use when ingesting new sources, querying the docs, or linting for stale content. Trigger on "docs-wiki ingest", "add this to the docs", "update the docs with", "query the docs about", "lint the docs", or "what do the docs say about".
license: MIT
allowed-tools: Read Glob Grep Bash(git add:*) Bash(git commit:*) Bash(git status:*) Bash(date:*) Write Edit
metadata:
  version: "1.0.0"
  based-on: llm-wiki (https://gist.github.com/karpathy/442a6bf555914893e9891c11519de94f)
---

# docs-wiki — Repo Docs Manager

Manages the `docs/` knowledge base following the Karpathy LLM-wiki pattern: the LLM is the programmer; `docs/` is the codebase.

**Before every operation:** read `docs/conventions.md`. It defines the folder taxonomy, OKF types, naming rules, and cross-reference conventions. The conventions file is authoritative — if it contradicts anything in this skill, the conventions file wins.

**Docs root:** always `docs/` relative to the repo root. Never prompt the user to locate it.

---

## Operation: ingest

Add a new source to the docs and propagate its ideas to related pages.

### Steps

1. **Read the source.**
   - URL → `WebFetch`
   - File path → `Read` (never write to source files)
   - Conversation content → treat as-is; no external fetch needed

2. **Surface 3–5 key ideas in chat.** Give the user a chance to redirect focus before any writes happen.

3. **Determine placement.** Derive a `kebab-case` filename from the source title. Choose the folder using the taxonomy in `docs/conventions.md`. When in doubt, match the OKF `type` to the folder.

4. **Write the doc page** with OKF frontmatter:

   ```yaml
   ---
   type: <type from conventions>
   title: "<Title>"
   description: <one sentence — what and why it matters>
   resource: <url-or-path if a single canonical source exists; omit otherwise>
   tags: [<relevant>, <keywords>]
   timestamp: <run `date -u +%Y-%m-%dT%H:%M:%SZ`>
   ---
   ```

   Body scaffold by type — use the matching template from the OKF skill. For `spike`:

   ```markdown
   ## Goal
   ## Approach
   ## Findings
   ## Recommendation
   ## References
   ```

   Focus on synthesis, not transcription. Cross-reference related docs with relative markdown links.

5. **Propagate to related pages.** Read `docs/index.md` to identify related pages. For each:
   - Read the page.
   - Integrate what the new source adds, confirms, or contradicts.
   - Edit in-place (new section, updated claim, added cross-reference).
   - Upgrade incomplete OKF frontmatter while editing.

6. **Update `docs/index.md`.** Add a row for the new page:

   ```markdown
   | [Title](./relative/path.md) | type | One-line description |
   ```

7. **Append to `docs/log.md`:**

   ```
   ## [YYYY-MM-DD] ingest | <title>
   ```

---

## Operation: query

Synthesize an answer from existing docs.

### Steps

1. Read `docs/index.md` to map the docs. Identify 3–6 most relevant pages.
2. Read those pages. Follow relative links to related pages when relevant.
3. Write a cited answer — every claim traces back to a specific page:

   > <Answer.> ([Page Title](./relative/path.md))

4. **Write-back rule:** if synthesis produced a new insight not captured anywhere, write it as a new doc (`type: concept`, slug `synthesis-<topic>-<date>.md`) with OKF frontmatter. Don't write back routine answers.

5. Append to `docs/log.md`:

   ```
   ## [YYYY-MM-DD] query | <question truncated to 80 chars>
   ```

---

## Operation: lint

Audit `docs/` and fix what can be fixed automatically.

### Steps

1. `Glob("**/*.md", "docs/")` to inventory all pages. Read `docs/index.md`.

2. Check each page for:
   - **Orphans** — no inbound relative link from any other page
   - **Contradictions** — claims that conflict across pages
   - **Stale claims** — tied to a source superseded by a newer page
   - **Missing cross-references** — mentions a concept that has its own page but doesn't link to it
   - **Index gaps** — not listed in `docs/index.md`
   - **Missing OKF frontmatter** — no `---` block or missing required fields

3. Fix automatically:
   - Add missing relative links inline
   - Add orphans to `docs/index.md`
   - Mark contradictions: `> ⚠️ Contradiction with [Page](./relative/path.md) — needs resolution.`
   - Upgrade incomplete OKF frontmatter (preserve body exactly)

4. Report what requires human judgment — contradictions needing resolution, gaps needing new sources.

5. Append to `docs/log.md`:

   ```
   ## [YYYY-MM-DD] lint | <N pages checked, M issues found, K fixed>
   ```

---

## Conventions (summary — `docs/conventions.md` is authoritative)

| Convention         | Rule                                                                                       |
| ------------------ | ------------------------------------------------------------------------------------------ |
| **Docs root**      | Always `docs/` — never prompt the user for the location                                   |
| **Conventions**    | Read `docs/conventions.md` first; it is the authoritative source for all rules            |
| **Folder**         | Match OKF `type` to folder per `docs/conventions.md` taxonomy                             |
| **Filenames**      | `kebab-case.md` from the concept title                                                     |
| **Frontmatter**    | OKF required (`type`, `title`, `description`, `timestamp`); `resource` and `tags` optional |
| **Cross-refs**     | Standard relative markdown links; paths relative to linking file; no `[[wikilinks]]`       |
| **index.md**       | Updated on every write operation                                                           |
| **log.md**         | Append-only — never rewrite or delete entries                                              |
| **Synthesis**      | Integrate and cross-reference; never just copy a source verbatim                           |
| **Timestamps**     | Frontmatter: `date -u +%Y-%m-%dT%H:%M:%SZ`. Log entries: `YYYY-MM-DD`                    |

---

## Rules

- **Never modify source files** pointed at by the user. Read; never write.
- **Conventions file wins.** If `docs/conventions.md` contradicts this skill, follow the conventions file and note the discrepancy.
- **Propagate aggressively.** Ingest is a ripple, not a point write — touch every related page.
- **Cite everything.** Every claim traceable to a source must say which source.
- **index.md is always current.** Every write finishes by updating the index.
- **log.md is append-only.** New entries go at the bottom.
