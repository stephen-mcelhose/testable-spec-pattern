---
type: proposal
title: "Modeling the Workflow with Code"
description: A code-style sketch of the software development lifecycle this repo is building toward — expressed as typed functions with Undefined placeholders for types not yet modeled.
tags: [workflow, lifecycle, modeling, ideation]
timestamp: 2026-10-24T22:45:00Z
---

# Modeling the Workflow with Code

**Status:** sketch — types are intentionally underspecified

The entire workflow this repo is building can be expressed as a single function:

```go
func ExecuteSoftwareDevelopmentWorkflow(seed Seed) (CompletedProject, error)
```

This is not aspirational — it is definitional. Every skill, every convention, every decision in this repo exists to make one or more of the sub-functions below easier to execute correctly.

---

## The type system

```go
// Undefined is a placeholder for any type we have not yet modeled.
// Use it at the leaves rather than inventing a premature data model.
type Undefined = any

// Idea is the raw starting material — a product idea, a user story,
// a feature request, a domain description. Not yet structured.
type Idea = Undefined

// CompletedProject is the full output of the workflow.
type CompletedProject struct {
    Requirements  Requirements
    Contracts     ProtoContracts
    Repository    Repository
    Specification Specification
}

// Phase output types — intentionally underspecified for now.
type Requirements  = Undefined
type ProtoContracts = Undefined
type Repository    = Undefined
type Specification = Undefined
```

---

## The lifecycle

```go
// ExecuteSoftwareDevelopmentWorkflow is the top-level function.
// It composes the four phases in sequence.
func ExecuteSoftwareDevelopmentWorkflow(idea Idea) (CompletedProject, error)

// GatherRequirements takes an idea and produces structured requirements —
// use cases, domain type sketches, Gherkin scenarios.
func GatherRequirements(idea Idea) (Requirements, error)

// BuildContracts takes requirements and produces machine-readable contracts —
// proto definitions, OpenAPI specs, ROD resource definitions.
func BuildContracts(requirements Requirements) (ProtoContracts, error)

// Implement takes requirements and contracts and produces a repository —
// scaffolded Go project, wired services, generated proto types.
// This function has significant internal structure (many steps).
func Implement(requirements Requirements, contracts ProtoContracts) (Repository, error)

// RunRunnableSpecification takes a specification and runs it.
// Returns an error if any requirement is violated.
// A passing run is the definition of a completed project.
func RunRunnableSpecification(specification Specification) error
```

---

## Open questions

- Where does `Specification` come from? It is not yet an explicit output of any phase. Most likely it is derived from the `Repository` (the spec lives in the test suite) or it is a separate artifact produced during `Implement`. To be resolved when `Specification` is no longer `Undefined`.
- `Implement` hides significant internal structure. The sub-functions inside it (scaffold, generate, wire, test) will eventually be worth naming explicitly.
- `Idea` is the most underspecified type. It could be a freeform string, a structured document, or an existing codebase. The `GatherRequirements` phase has to handle all of those inputs.
