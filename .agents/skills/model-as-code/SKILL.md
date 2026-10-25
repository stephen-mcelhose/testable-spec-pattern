---
name: model-as-code
description: Model a domain using Go types as the modeling language. Use when someone wants to sketch a domain, turn an idea into types and functions, or make the shape of a problem visible before implementing it. Trigger on "model this as code", "sketch this in Go", "what are the types here", "turn this idea into types", or any request to explore a problem space as compilable Go.
license: MIT
allowed-tools: Read Glob Grep Write Edit Bash(go build:*) Bash(go vet:*)
metadata:
  version: "1.2.0"
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

## The process

### 0. Specify only what you know

Model only what is stated in the idea or directly derivable from it. If you derive something non-obvious, document it in a comment. If something is needed to model a function but not given, ask rather than invent. This is design — it is expected to be incomplete and revisited.

`Undefined` is not a failure. It is the model being honest.

### 1. Enumerate subjects

Before writing any function signature, list every entity in the idea. Entities come in four kinds:

- **Actors** — who initiates or performs actions (e.g., the person who books, pays, or requests something)
- **Objects** — what is acted on (e.g., the thing being booked, moved, or changed)
- **Records** — what gets created or persisted as a result (e.g., a log, a receipt, an audit entry)
- **Outcomes** — the final deliverable of a process (e.g., a confirmation, a report, a certificate)

The grammatical actor of a sentence is still a domain entity. *"A driver parks a vehicle"* produces two subjects: `Driver` and `Vehicle`. Don't let an actor dissolve into a function parameter without first declaring it as a type.

If domain knowledge implies an entity the idea doesn't name, surface it as a question rather than inventing or ignoring it.

Each subject becomes a named interface embedding `Undefined` before any function is written:

```go
type Driver  interface{ Undefined }
type Vehicle interface{ Undefined }
type Spot    interface{ Undefined }
```

### 2. Name the top-level operations

With all subjects declared, name what the system does as function signatures — inputs, outputs, errors. Prefer `type <DomainName> interface{ Undefined }` for any new type you encounter. Do not fill in the internals of a type until they are necessary to define something else — internals are implementation details until proven otherwise.

```go
func ParkVehicle(driver Driver, vehicle Vehicle) (ParkingTicket, error)
func ReleaseVehicle(ticket ParkingTicket) (Receipt, error)
```

Filling in details before you understand the system as a whole is the trap this step avoids.

### 3. Decompose into relations and iterate

```
relations := ExtractRelationships(idea, subjects)

for each relation → func <Verb><Subject>(subjectA, subjectB?) (<Result>, error)
  if the function composes sub-steps  → call them explicitly in the body
  if the function is a leaf           → return zero, ErrUnimplemented
  if the function delegates to a skill → declare a Strategy type instead
```

Expand each non-leaf function the same way: express it as a composition, then decompose further. Stop when every function is either a composition of named sub-functions or a leaf returning `ErrUnimplemented`.

---

## Subjects → types

Promote `Undefined` to a struct only when the fields matter to the model — not speculatively.

---

## Relations → functions

Names — of types and functions — come from the domain, not from implementation tools or technology. A type is `GeneratedCode`, not `ProtoGeneration`. A function is `GenerateFromContracts`, not `GenerateGoFromProto`. If the name contains a library, a protocol, or a framework, it is probably wrong.

Each relation becomes a named function. The name is the claim:

```go
// AssignSpot allocates a parking spot to a vehicle.
func AssignSpot(vehicle Vehicle, zone Zone) (Spot, error)
```

**Composed:** call sub-functions in the body — make the structure visible.

**Leaf:** return zero + `ErrUnimplemented`. Add a `// Leaf:` comment explaining why.

**Strategy seam:** only introduce a Strategy type when the step is explicitly backed by a named external skill. If you don't know which skill backs it, don't add it.

```go
// SpotSelectionStrategy is backed by the space-optimizer skill.
type SpotSelectionStrategy func(vehicle Vehicle, zone Zone) (Spot, error)
```

**Predicates** return `bool`, not `error` and not a wrapper type:

```go
func IsVerifiedProject(project BuiltProject) bool
```

**Comments:** let each symbol speak for itself. Banner-style section dividers (`// ------`) are noise — omit them.

---

## Verify

```bash
go build ./...
grep -r panic .
```

The model must compile and must not contain any `panic`. Fix both before delivering.
