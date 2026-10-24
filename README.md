# testable-spec-pattern

A personal workshop for building an **end-to-end AI-assisted software development workflow** expressed as composable agentic skills.

## The pre-baked assumption

**A runnable spec is more valuable than a written one.**

A spec that lives in a document drifts. A spec that lives in a test suite fails loudly when the contract is violated. This repo assumes that without argument — if it weren't true, there'd be no point building the workflow. The workflow is about *how* to make runnable specs easy to write and maintain, not *whether* to.

## The core thesis

Domain-named contracts — proto, OpenAPI, ROD — are the bridge between requirements and runnable specs. When you start with a well-named contract, the spec writes itself. The workflow makes that chain explicit and executable at every step.

```
Requirements → Contract → Implementation → Runnable Spec → Formal Spec
     ↑                                                           |
     └─────────────────── living documentation ─────────────────┘
```

## Workflow phases

The workflow is organized into six skill groups. Each group is a composable collection of discipline skills, invocable individually or chained.

| Phase                        | What it produces                                               | Skills                                          |
| ---------------------------- | -------------------------------------------------------------- | ----------------------------------------------- |
| **Requirements Gathering**   | Use cases, Gherkin scenarios, domain type sketches             | *(planned)*                                     |
| **Contract Design**          | Proto / OpenAPI / ROD definitions, AIP-compliant services      | `resource-oriented-design`                      |
| **Implementation**           | Idiomatic Go project structure, proto-wired services           | `go-project-scaffold`                           |
| **Testing & Evaluation**     | AIP evals, score-cards, behavioral spec test suites            | `go-behavioral-spec`                            |
| **Specification Writing**    | RFC 2119 behavioral specs, Gherkin, issue-to-RFC escalation    | *(planned)*                                     |
| **Project Management**       | Structured GitHub issues, machine-readable translations, RFCs  | `write-issue`                                   |

## Skills

| Skill                                                                    | Phase                | Purpose                                                  |
| ------------------------------------------------------------------------ | -------------------- | -------------------------------------------------------- |
| [`go-behavioral-spec`](.agents/skills/go-behavioral-spec/SKILL.md)       | Testing              | Write Go tests as RFC 2119 behavioral specs              |
| [`resource-oriented-design`](.agents/skills/resource-oriented-design/SKILL.md) | Contract Design | Guide API design using Resource Oriented Design (AIP-compliant proto) |
| [`go-project-scaffold`](.agents/skills/go-project-scaffold/SKILL.md)     | Implementation       | Scaffold a new Go project with idiomatic structure       |
| [`write-issue`](.agents/skills/write-issue/SKILL.md)                     | Project Management   | Write clear, focused, actionable GitHub issues           |
| [`docs-wiki`](.agents/skills/docs-wiki/SKILL.md)                         | *(meta)*             | Manage the `docs/` knowledge base                        |

## Docs

See [`docs/index.md`](docs/index.md) for the full knowledge base — discipline, research, ideation, and plans.

Key reading:
- [`docs/discipline.md`](docs/discipline.md) — the behavioral-spec discipline (read before writing any tests)
- [`docs/ideation/contract-first.md`](docs/ideation/contract-first.md) — why contract-first naturally produces testable specs
- [`docs/ideation/domain-names-as-bridge.md`](docs/ideation/domain-names-as-bridge.md) — why domain names are the prerequisite

## License

MIT — see [LICENSE](LICENSE).
