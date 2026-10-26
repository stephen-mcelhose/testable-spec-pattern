---
type: specification
title: Proto Contracts
description: Normative RFC 2119 behavioral specification for proto contracts following Resource Oriented Design — the requirements a proto file must satisfy to be a valid wire contract.
tags: [proto, contracts, aip, rod, rfc-2119, specification]
resource: https://google.aip.dev/121
timestamp: 2026-10-25T21:24:53Z
---

# Proto Contracts

## 1. Introduction & Scope

This specification defines the behavioral requirements for proto contracts in this workflow. A proto contract is a `.proto` file defining gRPC services and their resources, produced during the [Contracts & Domain Modeling](./index.md) phase.

**In scope:** resource definitions, service method signatures, field annotations, file layout, naming, constraints, and linter compliance.

**Out of scope:** implementation details, generated code, transport configuration, and deployment topology.

The canonical design philosophy is Resource Oriented Design ([AIP-121](https://google.aip.dev/121)). This specification is a **summarization** of that standard for use in this workflow. Where any statement here conflicts with or is silent on a point covered by the AIP index, the AIP index is authoritative. Always defer to [google.aip.dev](https://google.aip.dev/) as the primary reference. Contracts MUST be enforced with annotations, not comments.

## 2. Terminology

The key words "MUST", "MUST NOT", "REQUIRED", "SHOULD", "SHOULD NOT", "RECOMMENDED", "MAY", and "OPTIONAL" in this document are to be interpreted as described in [RFC 2119](https://datatracker.ietf.org/doc/html/rfc2119).

**Resource** — a named domain entity exposed by the API, acted on by standard methods.  
**Standard method** — one of: List, Get, Create, Update, Delete.  
**Custom method** — an RPC that does not map cleanly to any standard method.  
**Field behavior** — a `google.api.field_behavior` annotation declaring how a field is used.

## 3. Behavioral Requirements

### 3.1 Design Order

1. A proto contract MUST be designed in this sequence: resources → relationships and hierarchy → resource schema → methods.
2. Methods MUST NOT be designed before resources and their schemas are established.

### 3.2 Resources

3. Every API MUST expose its domain entities as named resources ([AIP-121](https://google.aip.dev/121)).
4. Resource messages MUST carry a `(google.api.resource)` annotation declaring `type`, `pattern`, `singular`, and `plural`.
5. Every resource SHOULD support a List method. If List is genuinely inapplicable, the omission MUST be suppressed with an inline linter comment.
6. Resource names MUST follow the `{collection}/{id}` pattern ([AIP-122](https://google.aip.dev/122)).

### 3.3 Standard Methods

7. Services MUST prefer standard methods (List, Get, Create, Update, Delete) over custom methods.
8. Custom methods MUST only be used for operations that do not cleanly map to any standard verb.
9. Method signatures MUST conform to the following rules ([AIP-131](https://google.aip.dev/131)–[AIP-135](https://google.aip.dev/135)):
   - Top-level List (no parent): omit `method_signature` annotation entirely
   - List with parent: `"parent"`
   - Get / Delete: `"name"`
   - Create: `"{resource},{resource}_id"` — `{resource}_id` MUST be included even when optional
   - Update: `"{resource},update_mask"`
   - Custom method: all REQUIRED fields, comma-separated

### 3.4 Field Behavior

10. Every field MUST carry at least one `(google.api.field_behavior)` annotation.
11. The `name` field on a resource MUST use `IDENTIFIER`. It MUST NOT use `OUTPUT_ONLY`.
12. A `resource_reference` annotation MUST NOT be placed on a resource's own `name` field. It MUST only appear on request message fields that reference the resource.
13. Server-assigned fields (e.g. `create_time`, `etag`, managed state) MUST use `OUTPUT_ONLY`.
14. Fields the caller must always supply MUST use `REQUIRED`.
15. Fields the caller may omit (e.g. `page_size`, `page_token`, `filter`, `update_mask`, `{resource}_id`) MUST use `OPTIONAL`.
16. Fields set at creation and never changed (e.g. `username`) MUST use `IMMUTABLE`.
17. `update_mask` MUST be `OPTIONAL`.

### 3.5 Naming

18. Cross-resource reference fields MUST be named after the referenced resource type in `snake_case`, with no `_name` suffix (e.g. `pet`, never `pet_name`).
19. Human name fields MUST use `given_name` and `family_name`, not `first_name` or `last_name` ([AIP-148](https://google.aip.dev/148)).
20. Timestamp fields MUST end in `_time` (e.g. `ship_time`, not `ship_date`) ([AIP-142](https://google.aip.dev/142)).
21. Lifecycle enums MUST use `State` as the suffix, not `Status` (e.g. `PetState`, not `PetStatus`) ([AIP-216](https://google.aip.dev/216)).

### 3.6 File Layout

22. Declarations in a `.proto` file MUST appear in this order ([AIP-191](https://google.aip.dev/191)):
    1. `syntax`, `package`, `option`, `import`
    2. Service definition
    3. Request / response messages (in RPC order)
    4. Resource messages
    5. Enums

### 3.7 Constraints

23. Field constraints MUST be enforced with annotations, not comments.
24. Field constraints MUST use `buf.validate` (protovalidate), not ad-hoc comment conventions.

### 3.8 Enums

25. Every enum value — including `UNSPECIFIED = 0` — MUST have a doc comment.

### 3.9 Linter Compliance

26. Every proto contract MUST pass `api-linter` with no violations before it is considered complete.
27. Violations MUST be diagnosed and fixed at root cause. Blanket suppression of linter rules is NOT RECOMMENDED.

## 4. Data Models / Interfaces

### Resource annotation shape

```proto
option (google.api.resource) = {
  type: "{Service}/{ResourceType}"
  pattern: "{collections}/{id}"
  singular: "{resource}"
  plural: "{resources}"
};
```

### Field behavior annotation shape

```proto
string name = 1 [(google.api.field_behavior) = IDENTIFIER];
google.protobuf.Timestamp create_time = 2 [(google.api.field_behavior) = OUTPUT_ONLY];
string display_name = 3 [(google.api.field_behavior) = REQUIRED];
string description = 4 [(google.api.field_behavior) = OPTIONAL];
```

### Constraint annotation shape

```proto
repeated CreateBookRequest requests = 1 [
  (google.api.field_behavior) = REQUIRED,
  (buf.validate.field).repeated.max_items = 1000
];
```
