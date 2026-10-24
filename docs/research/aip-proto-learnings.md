---
type: spike
title: "AIP Proto Learnings — iter-1 Petstore"
description: Concrete AIP rule discoveries from the iter-1 Petstore proto revision — the delta between what naive proto authoring produces and what api-linter + AIP require.
resource: https://google.aip.dev/
tags: [aip, proto, resource-oriented-design, api-linter, iter-1]
timestamp: 2026-10-24T18:48:20Z
---

# AIP Proto Learnings — iter-1 Petstore

## Goal

Document the concrete AIP rule discoveries made when revising the Petstore proto from iter-0 (baseline) to iter-1 (0 api-linter violations, rubric 22/22). These are the gaps between naive proto authoring and what the AIP rules actually require.

## Approach

Ran api-linter against iter-0, iteratively fixed violations, and cross-checked against the AIP specification. Violations were categorised into "wrong assertion in our eval", "gap in our eval", and "genuine skill gap". See the iter-1 SCORECARD and the `resource-oriented-design` skill evals.

## Findings

### `name` field behavior: IDENTIFIER, not OUTPUT_ONLY

The `name` field on a resource message must be annotated `IDENTIFIER`, not `OUTPUT_ONLY`. Both communicate "server-assigned, not user-writable", but `IDENTIFIER` is the semantically correct choice for resource name fields and is what api-linter enforces. `OUTPUT_ONLY` would also technically pass api-linter's field-behavior checks but is the wrong semantic annotation.

**Rule:** `(google.api.field_behavior) = IDENTIFIER` on `name` — MUST (AIP-203).

### Self-referential `resource_reference` creates a mutable cycle

Adding `(google.api.resource_reference)` to the `name` field *on the resource message itself* (e.g., `Pet.name` referencing `petstore.example.com/Pet`) causes api-linter rule `0121::no-mutable-cycles` to fire. The resource_reference annotation on the resource's own name is redundant (the resource annotation already declares the type) and creates a mutable cycle in the type graph.

**Rule:** Do NOT add self-referential `resource_reference` to resource `name` fields — MUST NOT (AIP-121). Request message `name` fields (GetPetRequest.name, DeletePetRequest.name) still SHOULD have `resource_reference`.

### State over Status, nested inside the resource (AIP-216)

Top-level enums named `PetStatus` or `OrderStatus` are wrong in two ways:
1. AIP-216 prefers `State` over `Status` as the name.
2. The enum should be nested *inside* the resource message (`Pet.State`) rather than declared at the top level.

Nesting makes the namespace clear, avoids collisions between resources, and matches the AIP reference examples. api-linter enforces the nesting requirement.

**Rule:** `enum State { STATE_UNSPECIFIED = 0; ... }` nested inside the resource message — SHOULD (AIP-216).

### Enum zero value must be `STATE_UNSPECIFIED` (AIP-126)

Every proto enum must have a zero value named `{EnumName}_UNSPECIFIED`. For a nested `State` enum, this is `STATE_UNSPECIFIED = 0`. This value must also have a doc comment.

**Rule:** `STATE_UNSPECIFIED = 0` — MUST (AIP-126). All enum values including unspecified must have doc comments — MUST (AIP-192).

### Timestamp field naming: `_time` suffix (AIP-142)

Timestamp fields must end in `_time` (e.g., `create_time`, `ship_time`). Names like `create_date`, `ship_date`, or `created_at` violate AIP-142.

**Rule:** Timestamp fields end in `_time` — MUST (AIP-142).

### Human name fields: `given_name` / `family_name` (AIP-148)

Fields for human names must use `given_name` and `family_name`, not `first_name` and `last_name`. This is a hard AIP-148 requirement that api-linter enforces.

**Rule:** Use `given_name` / `family_name` — MUST (AIP-148).

### Cross-resource fields: no `_name` suffix (AIP-122)

Fields referencing another resource should be named after the resource type in snake_case *without* a `_name` suffix. For example, a field referencing a `Pet` resource on an `Order` message should be named `pet`, not `pet_name`.

**Rule:** Cross-resource reference fields SHOULD NOT use a `_name` suffix — SHOULD NOT (AIP-122).

### `method_signature` for top-level List: omit entirely (AIP-132)

For top-level resources (no parent), the List RPC should omit `method_signature` entirely — there is no `parent` parameter to include. Providing an empty string or incorrect signature causes api-linter violations.

For child resources, `method_signature = "parent"` is correct.

**Rule:** Top-level List omits `method_signature`; child List uses `"parent"` — SHOULD (AIP-132).

### `method_signature` for Create: include `{resource}_id` (AIP-133)

The Create RPC's `method_signature` should include the resource ID field: `"parent,pet,pet_id"` for a child resource, or `"pet,pet_id"` for a top-level resource. Just `"parent,pet"` is insufficient.

**Rule:** Create `method_signature` = `"parent,{resource},{resource}_id"` (or without parent for top-level) — SHOULD (AIP-133).

### State fields: `OPTIONAL` when user-settable; suppress api-linter if needed (AIP-216)

AIP-216 says state fields should be `OUTPUT_ONLY`, implying the server manages state transitions. However, for APIs where the client directly sets state (e.g., a simple pet store where `state` is just an enum the user provides), this conflicts with the design intent. The correct approach is:
- If state is server-managed → `OUTPUT_ONLY`
- If state is user-settable → `OPTIONAL`, with api-linter rule `0216::state-field-output-only` suppressed inline using `(google.api.field_info)` or a buf lint ignore comment

**Rule:** Disable `0216::state-field-output-only` selectively when state is genuinely user-settable — don't fight the linter with the wrong annotation.

### File layout: service before messages (AIP-191)

The service definition should appear before message definitions in the `.proto` file. This is a SHOULD per AIP-191 and is enforced by api-linter.

**Rule:** Service block at top of file, then messages — SHOULD (AIP-191).

### Every field needs `field_behavior` (AIP-203)

All fields on resource messages and request messages must have at least one `(google.api.field_behavior)` annotation. Missing annotations cause api-linter violations.

Summary of correct annotations:
| Field type | Correct annotation |
|------------|--------------------|
| Resource name (`name`) | `IDENTIFIER` |
| Server-assigned (timestamps, etags) | `OUTPUT_ONLY` |
| Required request field | `REQUIRED` |
| Optional request/resource field | `OPTIONAL` |
| Client-settable state | `OPTIONAL` |

## Recommendation

These learnings have been incorporated into:
- The `resource-oriented-design` skill (`SKILL.md`) — added explicit rules for all of the above
- The eval suite (`evals/evals.json`) — new eval #9 covers AIP-203/216/142/148/191 together; fixed wrong assertions in evals #5 and #6
- The Petstore proto (`v1/`) — iter-1 achieves 0 api-linter violations

For future proto work: run api-linter early and often. The violations are highly informative — treat them as a checklist rather than an obstacle.

## References

- [AIP-121](https://google.aip.dev/121) — Resource-oriented design
- [AIP-122](https://google.aip.dev/122) — Resource names
- [AIP-126](https://google.aip.dev/126) — Enum default values
- [AIP-132](https://google.aip.dev/132) — Standard List
- [AIP-133](https://google.aip.dev/133) — Standard Create
- [AIP-142](https://google.aip.dev/142) — Timestamp fields
- [AIP-148](https://google.aip.dev/148) — Human names
- [AIP-191](https://google.aip.dev/191) — Proto file structure
- [AIP-192](https://google.aip.dev/192) — Documentation comments
- [AIP-203](https://google.aip.dev/203) — Field behaviors
- [AIP-216](https://google.aip.dev/216) — Lifecycle states
- [Iter-1 SCORECARD](../../examples/petstore/api/proto/petstore/v1/SCORECARD-iter1.md)
- [resource-oriented-design skill evals](../../.agents/skills/resource-oriented-design/evals/evals.json)
