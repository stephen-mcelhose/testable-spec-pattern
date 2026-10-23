# go-behavioral-spec examples

Canonical examples for the behavioral-spec discipline. Copy **structure and naming**, not the PetStore domain.

**One module:** `go.mod` here. Numbered folders hold examples; `package` names may not match the directory (`pet` under `01-pure-function`).

| Need | Open |
|------|------|
| Pure function | [`01-pure-function/`](./01-pure-function/) |
| Injected dep — **function arg** | [`02-injected-dependency/functional/`](./02-injected-dependency/functional/) |
| Injected dep — **struct method** | [`02-injected-dependency/method/`](./02-injected-dependency/method/) |
| Resource-oriented CRUD | [`03-petstore-api/`](./03-petstore-api/) — open `service_<method>_test.go` |

```bash
cd .agents/skills/go-behavioral-spec/examples
go test ./... -count=1
```

See `.agents/plans/stephen-workflow-skills/` (incl. `lessons-learned.md`). Follow `../references/discipline.md` when extending.
