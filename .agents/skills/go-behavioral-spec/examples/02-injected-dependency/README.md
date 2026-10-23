# 02 — Injected dependency

Path-based lookup for `stephen-workflow` / agents:

| Wiring | Open |
|--------|------|
| Function arg injection | [`functional/`](./functional/) |
| Struct field + method | [`method/`](./method/) |

Same MUST contracts; method wraps `functional.BuildPetView`.

## MUST parity list

1. `MUST call Get with the canonical pet name and return a PetView from the Get result`
2. `MUST return InvalidArgument without calling Get when the pet name is invalid`
3. `MUST return NotFound when Get reports not found`
4. `MUST return Internal when Get fails for a non-not-found reason`

## Verify

```bash
cd .agents/skills/go-behavioral-spec/examples
go test ./02-injected-dependency/... -count=1
```
