# testable-spec-pattern

A workflow for moving a software idea to a verified, runnable specification.

Enter at any phase. Each phase produces an artifact. Each artifact feeds the next.

---

## [Idea](./idea/index.md)

A raw concept, problem statement, or feature request. The workflow starts here.

**Produces:** an Idea — the input to Requirements Gathering.

---

## [Requirements Gathering](./requirements-gathering/index.md)

Converts an Idea into Requirements. Requirements are composable — any one of the following independently satisfies the phase:

| Artifact                                                           | Primary concern                        |
| ------------------------------------------------------------------ | -------------------------------------- |
| [Use Cases](./requirements-gathering/use-cases.md)                 | Actors, goals, interactions            |
| [Domain Types](./requirements-gathering/domain-types.md)           | Named entities and relationships       |
| [Gherkin Scenarios](./requirements-gathering/gherkin-scenarios.md) | Concrete Given/When/Then behaviours    |

**Produces:** Requirements — the input to Contracts & Domain Modeling, Specifications, and Planning.

---

## [Contracts & Domain Modeling](./contracts-and-domain-modeling/index.md)

Produces machine-readable contracts and a domain model from Requirements.

| Artifact                                                                  | What it is                              |
| ------------------------------------------------------------------------- | --------------------------------------- |
| [Contracts](./contracts-and-domain-modeling/contracts.md)                 | Proto / OpenAPI / schema — the shape    |
| [Domain Model](./contracts-and-domain-modeling/domain-model.md)           | Go types — the named entities in code   |

**Produces:** Contracts and a Domain Model — inputs to Specifications, Planning, and Implementation.

---

## [Specifications](./specifications/index.md)

Produces an RFC 2119 prose specification from Requirements and Contracts.

| Artifact                                          | What it is                                    |
| ------------------------------------------------- | --------------------------------------------- |
| [Specification](./specifications/specification.md) | Named behavioral requirements in prose        |

**Produces:** a Specification — the input to Planning, Implementation, and Testing.

---

## [Planning](./planning/index.md)

Produces an implementation plan from Specifications and Contracts.

| Artifact                              | What it is                                 |
| ------------------------------------- | ------------------------------------------ |
| [Plan](./planning/plan.md)            | Ordered implementation tasks with rationale |

**Produces:** a Plan — the input to Implementation.

---

## [Implementation](./implementation/index.md)

Produces working code from Contracts, Specifications, and a Plan.

| Artifact                                        | What it is                        |
| ----------------------------------------------- | --------------------------------- |
| [Implementation](./implementation/index.md)     | Scaffolded, wired, compilable Go  |

**Produces:** an Implementation — the input to Testing.

---

## [Testing](./testing/index.md)

Produces a Runnable Specification from the Implementation.

| Artifact                                                        | What it is                                    |
| --------------------------------------------------------------- | --------------------------------------------- |
| [Runnable Specification](./testing/runnable-specification.md)   | Go tests named as RFC 2119 requirements       |

**Produces:** a passing test run — the definition of done.

> A passing Runnable Specification means the Implementation satisfies its Contracts and Specifications.
