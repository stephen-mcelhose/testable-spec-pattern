# Issue-Writing Patterns — Reference Summary

A survey of established formats for writing software issues. Each entry covers the structure, origin, strengths, and limitations.

---

## 1. Task / Guardrails / Verification

**Structure**

| Section        | Question it answers                                      |
| -------------- | -------------------------------------------------------- |
| **Task**       | What should the agent/engineer produce?                  |
| **Guardrails** | What must the agent NOT do? What is explicitly excluded? |
| **Verification** | How does anyone confirm the task is done?              |

**Origin and context**

An emerging pattern in agent-oriented software development (2024–2025). Not attributed to a single author. Appears in practitioner writing on agentic workflows (e.g., axonbuild.com, mergeloom.ai) under related names ("finishable tickets", "agent-ready issues"). The core insight is that agents expand into undefined space: explicit guardrails prevent overreach, and an explicit verification condition lets the agent self-check before presenting work.

**Strengths**
- Forces the author to think adversarially: what could go wrong if not constrained?
- Verification section doubles as the agent's self-review checklist.
- Minimal by design — three sections is a low friction target.

**Limitations**
- Does not capture *why* the work matters. An agent can produce exactly what was tasked without understanding the intent, leading to technically correct but misaligned output.
- No use-case section — relies on the task description alone, which is often not enough for an agent to infer intent when edge cases arise.
- Most decisioning is left to the agent. Without a *Why* or use-case, the agent has no basis to resolve ambiguity — it either halts to ask or proceeds on assumption. Both outcomes increase friction.

---

## 2. Need / Use-Case / Acceptance Criteria

**Structure**

| Section                  | Question it answers                                   |
| ------------------------ | ----------------------------------------------------- |
| **Need**                 | What pain, gap, or risk motivates this work?          |
| **Use-case**             | In what concrete scenario is the need encountered?    |
| **Acceptance criteria**  | What specific, verifiable outcomes constitute done?   |

**Origin and context**

Rooted in requirements engineering practice. The use-case layer bridges the gap between abstract need and testable criteria — a pattern advocated in [IEEE 830-1998 (Software Requirements Specifications)](https://ieeexplore.ieee.org/document/720574) and in practitioner literature on [acceptance-test-driven development (ATDD)](https://gojko.net/books/specification-by-example/) (see Gojko Adzic, *Specification by Example*, Manning 2011). The three-part structure is often used in product management without a formal name.

**Strengths**
- Use-case gives the implementer a concrete mental model, reducing interpretive drift.
- Acceptance criteria are directly convertible to tests or verification steps.
- The "need" section ensures the issue survives scope negotiation: if the need disappears, so does the issue.

**Limitations**
- Does not explicitly exclude scope — authors must remember to add an "out of scope" note or it gets omitted.
- "Need" and "use-case" are sometimes conflated, collapsing the distinction between motivation and scenario.

---

## 3. Job Stories

**Structure**

> **When** [situation], **I want to** [motivation], **so I can** [outcome].

**Origin and context**

Invented at Intercom (circa mid-2013) by the product and research team led by Paul Adams, while adapting the Jobs-to-be-Done (JTBD) framework to software. The format was officially named "Job Stories" by Alan Klement. References: [Paul Adams, *"The dribbblisation of design"*, Intercom blog (2013)](https://www.intercom.com/blog/the-dribbblisation-of-design/); [Alan Klement, *"Replacing the User Story with the Job Story"* (2013)](https://medium.com/the-job-to-be-done/replacing-the-user-story-with-the-job-story-af7cdee10c27).

**Strengths**
- Situation-first framing captures context that "As a [role]" masks. Two users with different demographics but identical situations produce identical job stories.
- Forces articulation of motivation, not just feature request.
- Research-driven: each story should be grounded in observed user behaviour.

**Limitations**
- The sentence format is expressive but not structured enough to be a complete issue. Usually requires additional context, acceptance criteria, and scope notes.
- "When/I want/so I can" can feel grammatically awkward for internal engineering tasks where there is no user persona.
- Acceptance criteria are absent — must be added as a separate section.
- "Motivation" and "outcome" blur in practice. "I want to [motivation]" captures the immediate desire; "so I can [outcome]" captures the functional progress sought. Writers often collapse the two — "I want to submit the form so I can submit the form" — restating the same idea twice and removing the layered intent the format is designed to capture.

---

## 4. User Stories

**Structure**

> **As a** [role], **I want** [capability], **so that** [benefit].

Often supplemented with: Acceptance criteria (sometimes formatted as "Given/When/Then" from BDD).

**Origin and context**

Originated at Connextra (2001) — attributed to Rachel Davies and Ron Jeffries. Popularised by Mike Cohn in *"User Stories Applied"* (2004, Addison-Wesley). The "As a / I want / So that" template is now ubiquitous in Agile teams. "Given/When/Then" comes from Cucumber/BDD (Dan North, 2006).

**Strengths**
- Well-understood by most engineering teams.
- Role-first framing prompts the author to think about *who* is affected.
- Pairs naturally with Given/When/Then AC for testability.

**Limitations**
- Persona fragmentation: "As a [role]" tempts authors to write separate stories for different user types even when the motivation is identical.
- Engineering-driven teams often write stories that describe implementation ("As a system, I want to call the API…") rather than user intent.
- No explicit out-of-scope section; scope creep is common.
- For agent-authored or human-authored internal engineering tasks, there is no meaningful "role" — the agent and human are interchangeable workers.

---

## 5. RFC-Style

**Structure** (Rust RFC template as the canonical example)

| Section                      | Purpose                                                                 |
| ---------------------------- | ----------------------------------------------------------------------- |
| Summary                      | One paragraph; what is being proposed                                   |
| Motivation                   | The problem and why it matters; includes use-cases                     |
| Guide-level explanation      | How to teach the feature; examples; impact on existing users            |
| Reference-level explanation  | Technical detail; interactions with existing design; edge cases         |
| Drawbacks                    | Arguments against accepting this RFC                                    |
| Rationale and alternatives   | Why this design over others; consequences of not doing it               |
| Prior art                    | Related work in other languages/systems                                 |
| Unresolved questions         | Known gaps to resolve during discussion or implementation               |
| Future possibilities         | Out-of-scope but related ideas this proposal enables                   |

**Origin and context**

Rust RFC (Request for Comments) process, established by the Rust core team circa 2014. Template at: https://github.com/rust-lang/rfcs/blob/master/0000-template.md. Adopted (with variations) by Ember.js, React (for major features), TC39 (ECMAScript proposals), and others.

**Strengths**
- Extremely thorough. Forces the author to reason about drawbacks, alternatives, and prior art.
- The "unresolved questions" section prevents false confidence.
- Scales to large, contentious language or API changes where community consensus is needed.

**Limitations**
- Heavyweight for a Kanban ticket. Most sections ("Prior art", "Future possibilities", "Reference-level explanation") add noise for a small, contained engineering task.
- The extensive structure discourages completion — authors fill in what they know and leave the rest blank, defeating the purpose.
- Not designed for agent consumption: the narrative prose style requires inference; agents prefer declarative, checkable criteria.

---

## 6. BDD / Gherkin Scenarios

**Structure**

Acceptance criteria written as executable scenarios using the Given / When / Then vocabulary:

```
Given [an initial context or precondition]
When  [an action or event occurs]
Then  [an observable outcome is expected]
And   [additional outcome or step — chained onto any clause]
```

Scenarios are grouped under a **Feature** (the deliverable) and optionally tagged with labels for filtering. Multiple scenarios cover the happy path, edge cases, and failure cases.

**Origin and context**

Behaviour-Driven Development (BDD) was introduced by Dan North in 2006 ("Introducing BDD", dannorth.net) as a reframing of Test-Driven Development in plain language. The Gherkin vocabulary (Given/When/Then) was formalised in the Cucumber test framework (Aslak Hellesøy, 2008). Today Gherkin is the lingua franca of BDD tools across most languages: Cucumber (Java/Ruby/JS), Behat (PHP), Behave (Python), Godog (Go).

**Strengths**
- Scenarios are simultaneously documentation, specification, and (when wired to a framework) executable tests. The issue body *is* the test spec.
- Plain language bridges product, engineering, and QA. Non-engineers can read and challenge scenarios before implementation starts.
- Given/When/Then enforces a specific structure: precondition → trigger → outcome. This prevents vague criteria like "the system should handle errors gracefully".
- Edge cases and error paths are first-class citizens, not afterthoughts.

**Limitations**
- Gherkin describes *behaviour at a boundary* (what the system does when given an input). It is awkward for issues about internal tooling, skill authoring, or infrastructure where there is no clear user-facing boundary to test.
- Scenarios describe observable outputs, not the artefact being delivered. An issue with only Gherkin scenarios often omits the *Why* — the motivation disappears into the scenario preconditions.
- For agent-authored issues in this repo, the "feature" level of abstraction (what users experience) is often coarser than the issue level of abstraction (what is delivered). An agent picking up a Gherkin-only issue may not know what to *build*, only what to *verify*.
- Writing good Gherkin requires discipline — poorly written scenarios ("Given the system is running", "Then it works") provide the structure with none of the benefit.

**Relationship to other patterns**

Gherkin is best understood as an *acceptance criteria format* that can be embedded in any issue structure, not as a standalone issue format. It pairs well with Need/Use-Case/AC: the use-case narrows to a scenario, and the AC section becomes a set of Gherkin scenarios rather than plain checkboxes. It is also the natural complement to User Stories ("As a… I want… So that…" followed by "Given/When/Then" scenarios).

---

## 7. GitHub's Recommendations

**Summary**

GitHub provides two template mechanisms:
- **Markdown templates** (`.github/ISSUE_TEMPLATE/*.md`): free-form guidance; contributors can ignore structure.
- **Issue forms** (`.github/ISSUE_TEMPLATE/*.yml`): structured web form fields; required fields are enforced.

GitHub's template builder prompts for: **Title**, **Description**, and **Template Content**. Required frontmatter: `name:` and `about:` (Markdown) or `name:` and `description:` (YAML form).

Reference: https://docs.github.com/en/communities/using-templates-to-encourage-useful-issues-and-pull-requests/about-issue-and-pull-request-templates

**Strengths**
- Issue forms with required fields enforce completion.
- Flexible enough to embed any of the above patterns.

**Limitations**
- GitHub recommends *mechanics*, not content. The choice of sections — and what makes them good — is left entirely to the team.
- Markdown templates are easily ignored; contributors submit blank or partially-filled issues.

---

## 8. Description / Dependencies / Testing Scenarios / Acceptance Criteria / Links

**Structure**

| Section               | Question it answers                                              |
| --------------------- | ---------------------------------------------------------------- |
| **Description**       | What value is delivered? Why does the need exist?                |
| **Dependencies**      | External requirements (API access, test data, upstream systems)  |
| **Testing scenarios** | Scenarios for thinking through business logic and catching bugs  |
| **Acceptance criteria** | Validates the value defined in the description; includes "Tests Added", "Documentation Updated" |
| **Links**             | Supporting information: wikis, discussions, service tickets      |

**Origin and context**

A pragmatic team-defined template, commonly found in enterprise and product-team environments. The version documented here is from a real-world team template shared in context.

**Strengths**
- Dependencies section is explicit — prevents "I didn't know I needed X" surprises mid-sprint.
- Testing scenarios as a distinct section prompts business-logic thinking before implementation.
- Links section keeps the issue self-contained as a knowledge artifact.

**Limitations**
- Five sections is a high bar; partial completion is the observed failure mode. The more sections a template has, the more likely some will be skipped.
- "Testing scenarios" and "acceptance criteria" can overlap, creating confusion about which drives implementation and which gates completion.
- No explicit "out of scope" section — scope creep risk.
- "Documentation Updated" checkbox in AC conflates delivery with housekeeping.

---

## 9. RFC 2119 Behavioral Language

**Structure**

Acceptance criteria written as normative clauses using the modal vocabulary defined in [IETF RFC 2119](https://www.rfc-editor.org/rfc/rfc2119) (updated for case-sensitivity by [RFC 8174](https://www.rfc-editor.org/rfc/rfc8174)):

| Keyword      | Meaning                                                                 |
| ------------ | ----------------------------------------------------------------------- |
| **MUST**     | An absolute requirement. Deviation is a conformance failure.            |
| **MUST NOT** | An absolute prohibition.                                                |
| **SHOULD**   | Recommended; valid reasons to depart exist but must be understood.      |
| **MAY**      | Optional; implementations may include or omit.                          |

Applied to acceptance criteria, this replaces plain checkboxes with normative statements:

```
- The skill MUST produce an issue body that satisfies the Definition of Ready before presenting it.
- The What section MUST begin with a noun phrase, not a verb.
- The issue SHOULD include an Out of scope section with at least two explicit exclusions.
- The issue MAY include a Dependencies section when upstream blockers exist.
```

**Origin and context**

RFC 2119 was authored by Scott Bradner and published by the IETF in March 1997. Its intent is to give precise, interoperable meaning to the modal verbs used in protocol specifications. RFC 8174 (2017) clarified that the normative keywords apply only in **UPPERCASE** — lowercase "must" and "should" retain their ordinary English meaning. The vocabulary has since been adopted beyond networking: in W3C web standards, OpenAPI specifications, and behavioral testing disciplines (including the `go-behavioral-spec` skill in this repo).

**Strengths**
- Removes interpretive wiggle room. "The response MUST include an `ETag` header" is unambiguous; "the response should include an ETag header" is not.
- Directly closes the gap between issue and behavioral spec. If acceptance criteria are already written as `MUST`/`SHOULD`/`MAY` clauses, they translate word-for-word into test assertions — no reinterpretation step required.
- `SHOULD` and `MAY` make trade-offs explicit. A `SHOULD` criterion signals "this matters but has known exceptions"; a `MAY` signals "optional but valid". Plain checkboxes cannot express this.
- Well-suited to the preferred format in this repo: the What/Why/Use-case structure provides context; RFC 2119 sharpens the AC section into the same register the behavioral spec uses.

**Limitations**
- Unfamiliar to many contributors. Teams not versed in standards writing may misuse the keywords — writing `SHOULD` when they mean `MUST`, or using uppercase for emphasis rather than normative intent.
- Overhead for obvious criteria. "The file MUST be saved as `.md`" adds no precision over a plain checkbox. The keyword form pays off when the criterion involves conditional behaviour, ordering, or known exceptions.
- Ill-suited to process or tooling issues where the deliverable is not a system behaviour. "The ADR MUST be reviewed by two engineers" is a process rule, not a behavioural assertion — RFC 2119 language works but feels forced.
- `SHOULD` is frequently misread as weaker than intended. RFC 2119 defines it as "strong recommendation with known exceptions"; contributors often treat it as "nice to have". Usage requires team literacy.

**Relationship to the preferred format**

RFC 2119 language is a refinement of the **Acceptance criteria** section, not a replacement for the full format. It is compatible with What/Why/Use-case/AC/Out-of-scope as-is: the AC section becomes a set of normative `MUST`/`SHOULD`/`MAY` clauses rather than plain checkboxes. This is the natural next step for issues that will feed directly into a behavioral spec or agent workflow — starting in the same register closes the translation gap entirely.

---

## Synthesis

Across all patterns, four questions recur in every effective issue format:

1. **What** — the concrete artefact or change being delivered
2. **Why** — the motivation; what pain or gap this addresses
3. **Scenario** — a concrete situation showing the deliverable in use
4. **Done** — specific, verifiable conditions that confirm completion

Patterns diverge on:
- How much **design detail** to include (RFC: extensive; Task/Guardrails/Verification: none)
- Whether **scope exclusions** are explicit (usually not, but this is a consistent failure mode)
- Whether a **role/persona** is required (User Stories: yes; Job Stories: no; everything else: no)
- How many **sections** to require (more sections → lower completion rate)
- Whether **acceptance criteria** are plain checkboxes or structured Gherkin scenarios (BDD/Gherkin adds rigour at the cost of verbosity; plain checkboxes are faster to write and sufficient for most engineering tasks)
