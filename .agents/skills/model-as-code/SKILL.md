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

## The process

### 0. Specify only what you know

Model only what is stated in the idea or directly derivable from it. If you derive something non-obvious, document it in a comment. If something is needed to model a function but not given, ask rather than invent. This is design — it is expected to be incomplete and revisited.

`Undefined` is not a failure. It is the model being honest.

### 1. Name the top-level operations

Name what the system does as function signatures — inputs, outputs, errors. Name all types as you go; prefer `type <DomainName> interface{ Undefined }` as the initial form. Do not fill in the internals of a type until they are necessary to define something else — internals are implementation details until proven otherwise.

```go
func BookAppointment(request AppointmentRequest) (Appointment, error)
func RecordVisit(appointment Appointment) (VisitRecord, error)
```

Filling in details before you understand the system as a whole is the trap this step avoids.

### 2. Decompose into subjects and relations

```
subjects  := ExtractSubjects(idea)
relations := ExtractRelationships(idea, subjects)

for each subject  → type <Subject> interface{ Undefined }
for each relation → func <Verb><Subject>(subjectA, subjectB?) (<Result>, error)
  if the function composes sub-steps  → call them explicitly in the body
  if the function is a leaf           → return zero, ErrUnimplemented
  if the function delegates to a skill → declare a Strategy type instead
```

### 3. Iterate

Expand each non-leaf function the same way: express it as a composition, then decompose further. Stop when every function is either a composition of named sub-functions or a leaf returning `ErrUnimplemented`.

---

## Subjects → types

Each subject becomes a named interface embedding `Undefined`:

```go
// Patient is a person receiving veterinary care.
type Patient interface{ Undefined }
```

Promote `Undefined` to a struct only when the fields matter to the model — not speculatively.

---

## Relations → functions

Each relation becomes a named function. The name is the claim:

```go
// ScheduleAppointment books a patient into a time slot.
func ScheduleAppointment(patient Patient, slot TimeSlot) (Appointment, error)
```

**Composed:** call sub-functions in the body — make the structure visible.

**Leaf:** return zero + `ErrUnimplemented`. Add a `// Leaf:` comment explaining why.

**Strategy seam:** only introduce a Strategy type when the step is explicitly backed by a named external skill. If you don't know which skill backs it, don't add it.

```go
// SchedulingStrategy is backed by the scheduling-engine skill.
type SchedulingStrategy func(patient Patient, clinic Clinic) (Appointment, error)
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
