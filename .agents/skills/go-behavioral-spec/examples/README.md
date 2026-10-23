# go-behavioral-spec examples

Canonical examples for the behavioral-spec discipline. Copy **structure and naming**, not the PetStore domain.

**One module:** `go.mod` at this directory. Packages underneath share types without `replace`.

| # | Path | Pattern | Package |
|---|------|---------|---------|
| 01 | [01-pure-function/](./01-pure-function/) (code in [`pet/`](./pet/)) | Pure function | `…/examples/pet` |
| 02 | `02-injected-dependency/` | Injected leaf (functional + method) | Planned |
| 03 | `03-petstore-api/` | Resource-oriented CRUD | Planned |

```bash
cd .agents/skills/go-behavioral-spec/examples
go test ./... -count=1
```

See `.agents/plans/stephen-workflow-skills/` (incl. `lessons-learned.md`). Follow `../references/discipline.md` when extending.
