---
name: resource-oriented-design
description: Guide API design using Resource Oriented Design (ROD). Use when designing a new API, reviewing a proto file, naming resources or methods, or deciding whether to use a standard vs custom method. Trigger on "design this API", "how should I name this resource", "ROD review", "AIP compliant", "resource-oriented design".
license: MIT
compatibility: Requires buf (https://buf.build/docs/installation) and api-linter (github.com/googleapis/api-linter) — both must be on PATH for the verify step.
allowed-tools: Read Bash(buf:*) Bash(api-linter:*)
---

# Resource Oriented Design

Design APIs as a hierarchy of named **resources** acted on by a small, stable set of standard methods over a stateless protocol. Follow [AIP-121](https://google.aip.dev/121) as the canonical reference.

## Design order

Work through these in sequence before writing any method signatures:

1. What **resources** (nouns) does the API expose?
2. What **relationships and hierarchy** exist between them?
3. What is the **schema** of each resource?
4. What **methods** does each resource need — and can standard methods cover them?

## Standard methods

Prefer standard methods. Custom methods are for operations that don't cleanly map to any standard verb — not a general escape hatch.

| Method | HTTP   | Request body | Response body |
| ------ | ------ | ------------ | ------------- |
| List   | GET    | None         | Resource list |
| Get    | GET    | None         | Resource      |
| Create | POST   | Resource     | Resource      |
| Update | PATCH  | Resource     | Resource      |
| Delete | DELETE | None         | None          |

## Key AIPs

| AIP                                   | What it covers                                     |
| ------------------------------------- | -------------------------------------------------- |
| [AIP-121](https://google.aip.dev/121) | Resource-oriented design — foundational philosophy |
| [AIP-122](https://google.aip.dev/122) | Resource names — `{collection}/{id}` patterns      |
| [AIP-123](https://google.aip.dev/123) | Resource types — `{Service}/{Type}` annotation     |
| [AIP-131](https://google.aip.dev/131) | Standard Get                                       |
| [AIP-132](https://google.aip.dev/132) | Standard List                                      |
| [AIP-133](https://google.aip.dev/133) | Standard Create                                    |
| [AIP-134](https://google.aip.dev/134) | Standard Update                                    |
| [AIP-135](https://google.aip.dev/135) | Standard Delete                                    |
| [AIP-154](https://google.aip.dev/154) | Resource freshness validation (ETags)              |
| [AIP-158](https://google.aip.dev/158) | Pagination                                         |

For AIPs not listed above, consult the full index at **https://google.aip.dev/**.

## Compliance directive

Every proto you produce MUST pass `api-linter` clean. After writing, run:

```bash
# from the module root (where buf.yaml lives)
buf build -o /tmp/api.pb
# from the proto root (e.g. api/proto/) — run from here so package paths resolve
api-linter --descriptor-set-in=/tmp/api.pb \
  --disable-rule core::0191::java-package \
  --disable-rule core::0191::java-multiple-files \
  --disable-rule core::0191::java-outer-classname \
  --disable-rule core::0191::csharp-namespace \
  --disable-rule core::0191::php-namespace \
  --disable-rule core::0191::ruby-package \
  <proto-files>
```

Read every violation. Reason about the root cause. Fix. Re-run. Repeat until clean.

## Annotation quick-reference

| Annotation                      | Import                          | Purpose                                                                         |
| ------------------------------- | ------------------------------- | ------------------------------------------------------------------------------- |
| `(google.api.resource)`         | `google/api/resource.proto`     | Marks a message as a named resource; set `type`, `pattern`, `singular`, `plural` |
| `(google.api.resource_reference)` | `google/api/resource.proto`   | Declares a string field holds another resource's name; set `type`               |
| `(google.api.field_behavior)`   | `google/api/field_behavior.proto` | Documents how a field behaves — see rules below                               |
| `(google.api.http)`             | `google/api/annotations.proto`  | Maps an RPC to an HTTP verb + URL                                               |
| `(google.api.method_signature)` | `google/api/client.proto`       | Positional signature — **import `google/api/client.proto`** when used           |

## Field behavior rules

Every field MUST carry at least one `field_behavior` annotation. Pick by role:

| Behavior     | Use for                                                                                               |
| ------------ | ----------------------------------------------------------------------------------------------------- |
| `IDENTIFIER` | The resource `name` field only — server-assigned, uniquely identifies the resource                    |
| `OUTPUT_ONLY` | Other server-assigned fields: `create_time`, `etag`, server-managed state transitions               |
| `REQUIRED`   | Caller must always supply                                                                             |
| `OPTIONAL`   | Caller may omit — `page_size`, `page_token`, `filter`, `update_mask`, `{resource}_id`, most body fields |
| `IMMUTABLE`  | Set at creation, never changed (e.g. `username`)                                                      |

- The `name` field on a resource MUST use `IDENTIFIER` (not `OUTPUT_ONLY`). Do **not** add `resource_reference` to the resource's own `name` field — that creates a self-referential cycle. Add `resource_reference` on request message fields that reference the resource (Get, Delete, Update requests), not on the resource itself.
- `update_mask` MUST be `OPTIONAL`.

## Naming rules

- **Cross-resource fields**: name the field after the resource type in snake_case, no `_name` suffix. A field holding a `Pet` name is `pet`, never `pet_name`.
- **Human names**: use `given_name` / `family_name`, not `first_name` / `last_name` (AIP-148).
- **Timestamps**: field name must end in `_time` — `ship_time` not `ship_date` (AIP-142).
- **Lifecycle enums**: prefer `State` over `Status` — `PetState` not `PetStatus` (AIP-216).

## File layout (AIP-191)

Declarations in a `.proto` file MUST appear in this order:

1. `syntax`, `package`, `option`, `import`
2. Service definition
3. Request / response messages (in RPC order)
4. Resource messages
5. Enums

## Constraints

Contracts MUST be enforced with annotations, not comments. A comment saying "max 1000 items" is aspirational; an annotation is a contract the runtime can verify.

Use `buf.validate` (protovalidate) for field constraints:

```proto
import "buf/validate/validate.proto";

repeated CreateBookRequest requests = 1 [
  (google.api.field_behavior) = REQUIRED,
  (buf.validate.field).repeated.max_items = 1000
];
```

Add `buf.build/bufbuild/protovalidate` to `buf.yaml` deps and run `buf dep update` when introducing this import.

Common constraint annotations:

| Constraint                              | Example value |
| --------------------------------------- | ------------- |
| `(buf.validate.field).repeated.max_items` | `1000`      |
| `(buf.validate.field).repeated.min_items` | `1`         |
| `(buf.validate.field).string.max_len`   | `256`         |
| `(buf.validate.field).int32.gte`        | `0`           |
| `(buf.validate.field).int32.lte`        | `100`         |

## Comments

Every enum value — including `UNSPECIFIED = 0` — MUST have a doc comment.

## Method signature rules

| RPC                           | `method_signature`                          |
| ----------------------------- | ------------------------------------------- |
| Top-level List (no parent)    | Omit the annotation entirely                |
| List with parent              | `"parent"`                                  |
| Get / Delete                  | `"name"`                                    |
| Create                        | `"{resource},{resource}_id"` (always include `{resource}_id` even when optional) |
| Update                        | `"{resource},update_mask"`                  |
| Custom method                 | All REQUIRED fields, comma-separated        |

## Resource completeness

All resources SHOULD support List (AIP-121). If List is genuinely inapplicable, suppress the rule with an inline comment on the service:

```proto
// (-- api-linter: core::0121::resource-must-support-list=disabled --)
service FooService {
```
