---
name: go-behavioral-spec
description: Write Go tests that specify public interface behaviour using RFC 2119 language. Use when writing new tests, adding coverage, or specifying a function's contract. Trigger on phrases like "write tests for X", "spec out this function", "test the public interface of X", "write a behavioral spec".
license: MIT
allowed-tools: Read Glob Grep Lsp Bash(go test:*) Bash(go build:*) Write Edit
metadata:
  version: "1.0.0"
---

# Go Behavioral Spec

Write Go tests that read like a behavioral specification. Each `t.Run` name is a requirement statement. The suite documents what the public interface MUST do — not how it does it.

See `references/discipline.md` for the design principles behind this style.

---

## Before writing

1. **Read the source file completely.** Every exported function, every branch, every error condition. Never guess field names or enum values — wrong ones compile silently.
2. **Read existing `_test.go` files in the package.** Match their style. Append; never replace.
3. **Choose an ordering and document it** in a comment on the test function: spec-natural (happy path → skips → errors → gaps) or code-flow (mirrors the function top-to-bottom).

---

## Structure

One `Test<Thing>` function per exported function or logical unit. Use the external test package (`package foo_test`) by default. Switch to `package foo` only when testing unexported behaviour — add a comment explaining why.

Write a baseline helper that returns the smallest valid input. Every test starts from it and overrides only what it's testing:

```go
func validEC(eventType string) model.EventContext {
    return model.EventContext{
        CommitableOffset: kafka.CommitableOffset{Topic: "t", Offset: 1},
        Event:            &events.Event{Id: "ce-1", Type: eventType},
    }
}
```

---

## Prefer table-driven tests

Use a table when multiple inputs exercise the same code path. If a table row's assertion needs a conditional, those cases are different branches — use flat `t.Run`s instead.

```go
func TestInitializeEventContext(t *testing.T) {
    t.Parallel()
    // Order: handled event types → unhandled → parse errors

    cases := []struct {
        name           string
        // inputs
        msg            kafka.Message
        // expected
        wantErr        bool
        wantCommitOnly bool
        wantAnswerID   string
    }{
        {
            name:         "MUST parse a create event and populate Answer",
            msg:          msg(makeAnswerEvent(t, events.AnswerCreateEventType, "ans-1")),
            wantAnswerID: "ans-1",
        },
        {
            name:           "MUST set CommitOnlyReasonError for an unrecognised event type without error",
            msg:            msg(makeAnswerEvent(t, "some.other.type", "irrelevant")),
            wantCommitOnly: true,
        },
        {
            name:    "MUST return an error when Data is not valid JSON",
            msg:     msg(&events.Event{Type: events.AnswerCreateEventType, Data: []byte("not-json")}),
            wantErr: true,
        },
    }

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            t.Parallel()
            ec, err := model.InitializeEventContext()(t.Context(), tc.msg)
            if tc.wantErr {
                require.Error(t, err)
                return
            }
            require.NoError(t, err)
            assert.Equal(t, tc.wantCommitOnly, ec.IsCommitOnly())
            if tc.wantAnswerID != "" {
                assert.Equal(t, tc.wantAnswerID, ec.Answer.AnswerID)
            }
        })
    }
}
```

---

## RFC 2119 naming

Every `t.Run` name is a requirement statement. Always include the keyword.

| Keyword    | Use when                                            |
| ---------- | --------------------------------------------------- |
| `MUST`     | Required. This test verifies it.                    |
| `MUST NOT` | Required to be absent. This test verifies absence.  |
| `SHOULD`   | Required but not yet fully verifiable — use `t.Skip("OPEN: ...")` |
| `MAY`      | Optional or conditionally present.                  |

- Omit the function name as subject when `Test<Thing>` already makes it unambiguous.
- Include the subject when it is a specific field: `"SourceEventId MUST equal the CloudEvent id"`.
- For every optional field or conditional effect, write both directions: `MUST set X when present` and `MUST NOT set X when absent`.

---

## Known gaps

```go
t.Run("OPEN: CorrelationId MUST be included in annotations when present", func(t *testing.T) {
    t.Skip("OPEN: mapping not implemented; remove when converters.go handles CorrelationID")
})
```

---

## Stubs

Hand-roll stubs for simple dependency control. Use counterfeiter (`make generate-fakes`) when you need call-count verification or argument capture.

---

## Verify

```bash
go build ./...
go test -count=1 ./...
```

All non-skip tests must pass. Fix every compile error before delivering.

---

## Rules

- **One rule per test.** If the name needs "and", split it.
- **Presence AND absence.** For every conditional: test it set, test it absent.
- **`t.Parallel()` everywhere.** Top-level and every subtest.
- **testify only.** `require` for fail-fast (errors, nil guards). `assert` for field checks. Never raw `t.Error`/`t.Fatal`.
- **Early return after `require.Error`.** Prevent cascading failures.
- **Named skips, not silence.** Every known gap gets `t.Skip("OPEN: ...")`.
- **Never delete passing tests.** Append; never replace.
