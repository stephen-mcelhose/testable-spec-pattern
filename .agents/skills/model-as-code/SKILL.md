---
name: model-as-code
description: Model a domain using Go types as the modeling language. Use when someone wants to sketch a domain, turn an idea into types and functions, or make the shape of a problem visible before implementing it. Trigger on "model this as code", "sketch this in Go", "what are the types here", "turn this idea into types", or any request to explore a problem space as compilable Go.
license: MIT
allowed-tools: Read Glob Grep Write Edit Bash(go build:*) Bash(go vet:*)
metadata:
  version: "1.0.0"
---

# model-as-code

`ModelAsCode(idea Idea) (CodeArtifact, error)`

Turn an idea into a compilable Go package that makes the shape of the domain visible. The model is not a prototype — it is a sketch. It names what is known, marks what is unknown, and shows where real behaviour will be plugged in later.

---

## The primitives

```go
// Undefined is the explicit placeholder for anything not yet modeled.
type Undefined = any

// ErrUnimplemented is returned by every leaf. Never panic in a model.
var ErrUnimplemented = errors.New("unimplemented")
```

Every type and function in the model is built from these two.

---

## Steps

- `ExtractSubjects(idea)` → for each subject: `type <Subject> interface{ Undefined }`
- `ExtractRelationships(idea, subjects)` → for each relation: `func <Verb><Subject>(subjectA, subjectB?) (<Result>, error)`
- Composed functions call sub-functions explicitly in the body
- Leaf functions (side effects, unknowns) return zero + `ErrUnimplemented`
- Delegation points become a Strategy type — doc comment names the backing skill
- `go build ./...` must pass before delivering
