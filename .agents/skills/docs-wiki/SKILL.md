---
name: docs-wiki
description: Manage the docs/ knowledge base for this repo. The LLM is the programmer; docs/ is the codebase. Invoke this skill whenever the user wants to add a new idea, article, URL, or conversation snippet to the docs ("add this to the docs", "document this", "ingest this", "put this in the knowledge base", "update the docs with"), ask what the docs already say ("what do the docs say about X", "query the docs", "what do we know about X"), or audit the knowledge base for gaps and staleness ("lint the docs", "audit the docs", "are the docs consistent"). Also invoke when the user shares a URL or file alongside any mention of documentation or knowledge base.
license: MIT
allowed-tools: Read Glob Grep WebFetch WebSearch Bash(git add:*) Bash(git commit:*) Bash(git status:*) Bash(date:*) Write Edit
metadata:
  version: "1.0.0"
  based-on: llm-wiki (https://gist.github.com/karpathy/442a6bf555914893e9891c11519de94f)
---

# docs-wiki — Repo Docs Manager

Manages the `docs/` knowledge base following the Karpathy LLM-wiki pattern: the LLM is the programmer; `docs/` is the codebase.

**Before every operation:**
1. Check that `docs/conventions.md`, `docs/index.md`, and `docs/log.md` exist. If any are missing, stop and tell the user — do not proceed without them.
2. Read `docs/conventions.md`. It defines the folder taxonomy, OKF types, naming rules, and cross-reference conventions. The conventions file is authoritative — if it contradicts anything in this skill, the conventions file wins.

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

   Body scaffold by type — use the matching template below. Focus on synthesis, not transcription. Cross-reference related docs with relative markdown links.

   **`concept`** — explanation of a principle or domain idea:
   ```markdown
   ## Overview
   <2–3 paragraphs synthesizing the concept>

   ## Key Points
   - …

   ## Sources
   - `<url-or-path>`
   ```

   **`proposal`** — hypothesis or early-stage idea:
   ```markdown
   ## The Observation
   ## The Hypothesis
   ## Open Questions
   ## What This Is NOT Yet
   ```

   **`spike`** — research findings:
   ```markdown
   ## Goal
   ## Approach
   ## Findings
   ## Recommendation
   ## References
   ```

   **`design-doc`** — detailed design:
   ```markdown
   ## Context
   ## Goals
   ## Non-Goals
   ## Design
   ## Alternatives
   ## Open Questions
   ```

   **`decision`** — ADR-style:
   ```markdown
   ## Status
   ## Y-Statement
   ## Context
   ## Decision
   ## Consequences
   ```

5. **Propagate to related pages.** Use `docs/index.md` to shortlist candidates (scan titles and descriptions), then read those pages to confirm relevance. For each confirmed related page:
   - Integrate what the new source adds, confirms, or contradicts.
   - Edit in-place (new section, updated claim, added cross-reference).
   - Upgrade incomplete OKF frontmatter while editing (preserve body exactly).

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

1. Read `docs/index.md` to map the docs. Shortlist 3–6 most relevant pages by title and description, then read them. Follow relative links to related pages when they seem relevant.
2. Write a cited answer — every claim traces back to a specific page:

   > <Answer.> ([Page Title](./relative/path.md))

3. **Write-back rule:** if synthesis produced a new insight not captured anywhere, write it as a new `concept` doc to `docs/research/synthesis-<topic>-<YYYY-MM-DD>.md` with OKF frontmatter. Don't write back routine answers — only genuinely new syntheses.

4. Append to `docs/log.md`:

   ```
   ## [YYYY-MM-DD] query | <question truncated to 80 chars>
   ```

---

## Operation: lint

Audit `docs/` and fix what can be fixed automatically.

### Steps

1. `Glob("**/*.md", "docs/")` to inventory all pages. Read `docs/index.md`.

2. Check each content page for (skip `docs/log.md` — it has no OKF frontmatter by design):
   - **Orphans** — no inbound relative link from any other page
   - **Contradictions** — claims that conflict across pages
   - **Stale claims** — tied to a source superseded by a newer page
   - **Missing cross-references** — mentions a concept that has its own page but doesn't link to it
   - **Index gaps** — not listed in `docs/index.md`
   - **Missing OKF frontmatter** — no `---` block or missing required fields (`type`, `title`, `description`, `timestamp`)

3. Fix automatically:
   - Add missing relative links inline
   - Add orphans to `docs/index.md`
   - Mark contradictions: `> ⚠️ Contradiction with [Page](./relative/path.md) — needs resolution.`
   - Upgrade incomplete OKF frontmatter (preserve body exactly)

4. Report what requires human judgment — contradictions needing resolution, gaps addressable only by new sources.

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
| **log.md**         | Append-only — no OKF frontmatter; never rewrite or delete entries                         |
| **Synthesis**      | Integrate and cross-reference; never just copy a source verbatim                          |
| **Timestamps**     | Frontmatter: `date -u +%Y-%m-%dT%H:%M:%SZ`. Log entries: `YYYY-MM-DD`                    |

---

## Rules

- **Never modify source files** pointed at by the user. Read; never write.
- **Conventions file wins.** If `docs/conventions.md` contradicts this skill, follow the conventions file and note the discrepancy.
- **Propagate aggressively.** Ingest is a ripple, not a point write — touch every related page.
- **Cite everything.** Every claim traceable to a source must say which source.
- **index.md is always current.** Every write finishes by updating the index.
- **log.md is append-only.** New entries go at the bottom. Never add OKF frontmatter to it.
- **Git is opt-in.** After writing, offer to commit the changes — don't commit silently.
