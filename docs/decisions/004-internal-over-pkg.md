---
type: decision
title: "ADR-004: Use internal/ not pkg/ for Go packages"
description: Explains why the go-project-scaffold skill prescribes internal/ and avoids pkg/.
tags: [adr, go, layout, scaffold]
timestamp: 2026-10-25T00:00:00Z
---

# ADR-004: Use `internal/` not `pkg/` for Go packages

## Status

Accepted

## Context

While building the `go-project-scaffold` skill, we needed a position on the `internal/` vs `pkg/` debate — a long-running community dispute. The `golang-standards/project-layout` repo popularised `pkg/` as a conventional directory for exported Go packages, leading many projects to adopt it.

The Go core team and a vocal part of the community dispute this convention. The debate is documented in:

- **[golang-standards/project-layout#117](https://github.com/golang-standards/project-layout/issues/117)** — Russ Cox (Go team lead) comments that `golang-standards/project-layout` is not an official standard and that `pkg/` in particular has no basis in the Go toolchain or official guidance.
- **[Go blog: Organizing a Go module](https://go.dev/blog/organizing-go-module)** — the official guidance uses `internal/` for packages that should not be imported by other modules, and makes no mention of `pkg/`.
- **[Go spec: internal packages](https://pkg.go.dev/cmd/go#hdr-Internal_Directories)** — `internal/` is a toolchain-enforced visibility boundary; `pkg/` is not.

## Decision

The `go-project-scaffold` skill prescribes `internal/` and does **not** recommend `pkg/`.

Concretely, for a service project:

```
internal/
  {domain}/       ← domain types, errors, value objects
  service/        ← service struct, Store interface, business logic
cmd/
  {appname}/      ← main package; wires up and starts the server
```

The `examples/petstore/` example follows this layout exactly.

## Rationale

| Criterion           | `internal/`                                    | `pkg/`                                      |
| ------------------- | ---------------------------------------------- | ------------------------------------------- |
| Toolchain support   | Enforced by `go` tool — cross-module imports blocked | No toolchain meaning; purely conventional   |
| Official guidance   | Referenced in Go blog and spec                 | Not mentioned in any official Go doc        |
| Encapsulation       | Prevents accidental leakage across module boundaries | No protection; anything can import `pkg/`   |
| Community consensus | Endorsed by Go team                            | Disputed — see issue #117                   |

For example code that demonstrates a service, `internal/` also more accurately models production Go code: a real service's domain packages are not exported to the world, they are internal implementation detail.

## Consequences

- The `go-project-scaffold` skill will not suggest or mention `pkg/`.
- Agents following the skill will put all non-binary packages under `internal/`.
- If a future example needs to export packages for consumption by other modules (e.g. a shared client library), a new decision should address that case — the default remains `internal/`.
