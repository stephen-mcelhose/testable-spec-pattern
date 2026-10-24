---
name: go-project-scaffold
description: Scaffold a new Go project with idiomatic structure and layout. Use when starting a new Go module, service, or library — or when restructuring an existing project to follow Go conventions. Trigger on "scaffold a Go project", "set up a new Go module", "lay out this Go service", "restructure this Go project".
license: MIT
allowed-tools: Read Glob Bash(mkdir:*) Bash(go mod init:*) Bash(go mod tidy:*) Bash(go build:*) Bash(go test:*) Write Edit
---

# Go Project Scaffold

Scaffold or restructure a Go module with idiomatic layout. This skill encodes opinionated decisions grounded in official Go guidance and community practice. Apply the layout fully — don't apply it piecemeal.

---

## Step 1 — Identify the project type

| Type        | Description                                                                                     |
| ----------- | ----------------------------------------------------------------------------------------------- |
| **Service** | Self-contained application (API server, worker, CLI). Has one or more binaries. Not designed for external import. |
| **Library** | A package intended to be imported by other modules. No binary entry point.                      |
| **Mixed**   | Exports a public API *and* ships binaries. Uncommon — split into separate modules when in doubt. |

---

## Step 2 — Verify the module path

The `go.mod` module path MUST be the full VCS path:

```
github.com/{org}/{repo}
```

For a sub-module nested in a monorepo, include the subdirectory:

```
github.com/{org}/{repo}/{subdir}
```

Never use a short vanity path. Never omit the VCS host.

---

## Step 3 — Apply the layout

### Service project

A service is not imported externally. All Go code goes inside `internal/`; binaries live in `cmd/`. Treat the directory containing `go.mod` as `/` — do not repeat the project name inside `internal/`.

```
/                           ← directory containing go.mod
  go.mod
  internal/
    {domain}/               ← one directory per domain concept (e.g. pet/, order/)
      {domain}.go           # domain types; package {domain}
      {domain}_test.go
    service/                ← business logic
      service.go            # package service
      service_test.go
    store/                  ← persistence interface + implementation
      store.go
      store_test.go
  cmd/
    {appname}/              ← one directory per binary; name = deployed binary name
      main.go               # package main — wire internals, call serve/run
```

**Rules:**
- Package name = last path component of the directory: `internal/pet/` → `package pet`.
- `internal/` is compiler-enforced: the Go toolchain prevents any external module from importing it.
- `cmd/{appname}/main.go` wires internals together. Keep it minimal — no business logic in `main`.
- Name `cmd/{appname}` after the deployed binary, not a generic name like `server` or `app`.
- **Interfaces belong with the consumer, not the implementor.** The `Store` interface used by `service/` is declared in `internal/service/`, not in `internal/store/`. `internal/store/` holds the concrete implementation only. See [Go Code Review Comments — Interfaces](https://go.dev/wiki/CodeReviewComments#interfaces).

---

### Library project

```
/                           ← directory containing go.mod
  go.mod
  {module}.go               # root package; name = last component of module path
  {module}_test.go
  {subpkg}/                 # exported sub-package
    {subpkg}.go             # package {subpkg}
    {subpkg}_test.go
  internal/                 # private helpers
    {helper}/
      {helper}.go
```

**Rule:** Do not create a top-level `pkg/` directory. Use direct top-level packages for exported code and `internal/` for private helpers. See the Reasoning section.

---

## Step 4 — Test placement

Test files live **alongside** the source they test — not in a separate top-level `test/` directory.

| Test type         | Placement                        | Package declaration |
| ----------------- | -------------------------------- | ------------------- |
| Unit / spec       | Same directory as the source     | See note below      |
| Shared helpers    | Same directory, `{name}_test.go` | `package {pkg}_test` |
| Integration / E2E | `test/` or `testdata/` at root   | Separate build tag  |

**Package declaration for tests** is governed by the `go-behavioral-spec` skill — defer to it for the `package {pkg}` vs `package {pkg}_test` decision and for how test cases are structured and named.

---

## Step 5 — Package naming

| Rule                                     | Example                                         |
| ---------------------------------------- | ----------------------------------------------- |
| Match the directory name                 | `internal/auth/` → `package auth`               |
| Root package = last component of module path | `module github.com/org/petstore` → `package petstore` |
| One package per directory                | All `.go` files in a directory share one name   |
| No stutter                               | `petstore.Pet`, not `petstore.PetPet`           |
| No generic names                         | Avoid `util`, `common`, `helpers` — name by what the package does |
| **Prefer descriptive over abbreviated**  | `authentication` over `auth`; `service` over `svc`; `petstore` over `ps` |

---

## Step 6 — What NOT to create

| Directory  | Why not                                                               |
| ---------- | --------------------------------------------------------------------- |
| `/src`     | Go workspace artifact. Produces redundant, ugly import paths.         |
| `/pkg`     | Not recommended by the Go team. Prefer top-level packages (library) or `internal/` (service). |
| `/vendor`  | Generated by `go mod vendor`. Commit only for offline-build services; never for libraries. |

---

## Step 7 — Verify

```bash
go build ./...
go test ./...
```

All non-skip tests must pass. Fix every compile error and import path mismatch before delivering.

---

## Reasoning for opinionated choices

**`internal/` for all service code, not a flat root.**
The official Go team explicitly recommends this for server projects: it prevents accidental public API surface, gives compiler-enforced privacy, and lets you refactor internal packages freely. ([go.dev/doc/modules/layout](https://go.dev/doc/modules/layout))

**No `pkg/`.**
The `golang-standards/project-layout` repo popularised `pkg/`, but Russ Cox opened [golang-standards/project-layout#117](https://github.com/golang-standards/project-layout/issues/117) stating the layout is *not* standard and the `/pkg` pattern does not represent how Go packages actually work. The repo's own README now carries a bold disclaimer: *"This is NOT an official standard defined by the core Go dev team."* For a new project, `internal/` gives compiler enforcement and top-level packages give public API — `pkg/` adds indirection without unique benefit.

**`cmd/{binary-name}/`, not `cmd/server/`.**
Naming the subdirectory after the deployed binary (e.g., `cmd/petstore/`) rather than a generic name makes multi-binary repos self-documenting and keeps `go install` paths meaningful.

**Group by domain concept inside `internal/`, not by technical layer.**
`internal/pet/` and `internal/order/` outlast `internal/handlers/` and `internal/repositories/`. Domain names align with the ubiquitous language of the codebase; layer names leak implementation choices that change.

**Do not repeat the project name inside `internal/`.**
When the project root *is* the module root (the directory containing `go.mod`), nesting code under `internal/petstore/` is redundant — you are already inside `petstore`. Name sub-packages by domain concept instead.

**Interfaces belong with the consumer, not the implementor.**
The Go Code Review Comments state: *"Go interfaces generally belong in the package that uses values of the interface type, not the package that implements those values."* ([go.dev/wiki/CodeReviewComments#interfaces](https://go.dev/wiki/CodeReviewComments#interfaces)). Defining a `Store` interface in `internal/service/` (the consumer) rather than `internal/store/` (the implementor) means you can evolve the implementation freely, test with a hand-stub without touching `internal/store/`, and keep the interface minimal (you only express the methods you actually call).
