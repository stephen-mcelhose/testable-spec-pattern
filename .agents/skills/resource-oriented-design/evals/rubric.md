---
title: "Eval rubric: resource-oriented-design skill"
status: draft
created: 2025
---

# Eval rubric: `resource-oriented-design` skill

Scoring rubric for evaluating proto output produced (or guided) by the `resource-oriented-design` skill. Applied per-artifact (a `.proto` file set covering a domain).

> **Location:** `.agents/skills/resource-oriented-design/evals/` per the [agentskills standard](https://agentskills.io/skill-creation/evaluating-skills) (evals live inside the skill directory). Confirmed by [ADR-006](../../../../docs/decisions/006-skill-eval-location.md).

**Score:** 0 / 1 / 2 per dimension. Aggregate = sum. Max = 22.

| Score | Meaning                                        |
| ----- | ---------------------------------------------- |
| 0     | Missing or wrong                               |
| 1     | Present but incomplete or inconsistent         |
| 2     | Correct, complete, consistent                  |

---

## Dimensions

### 1. Standard compliance *(user-identified)*
*The design follows AIP guidelines: resource names as strings, standard method signatures, PATCH+FieldMask for update, AIP-158 pagination on List, AIP-136 for custom methods.*

Indicators:
- Name fields are strings, not integers
- UpdateFoo takes the resource + FieldMask (not a full replacement PUT shape)
- ListFoo has `page_size`, `page_token`, `next_page_token`
- Custom methods use the `:verb` URL pattern and are reflected in the RPC name
- No invented patterns that contradict an existing AIP

---

### 2. Annotation completeness *(user-identified)*
*Every message and field that should carry an annotation does.*

Indicators:
- `google.api.resource` on every top-level resource message (`type` + `pattern`)
- `google.api.field_behavior` on every field (REQUIRED, OUTPUT_ONLY, OPTIONAL, IMMUTABLE as appropriate)
- `google.api.resource_reference` on cross-resource reference fields
- `google.api.http` bindings on every RPC (or explicitly waived with reason)

---

### 3. Annotation correctness
*The annotations say the right thing, not just present. A proto can have annotations everywhere but say wrong things — a distinct failure mode from completeness.*

Indicators:
- `OUTPUT_ONLY` on server-assigned fields (name, create_time, update_time, etag)
- `REQUIRED` only on fields that are truly required on every call (not just Create)
- `IMMUTABLE` on fields that cannot change after creation
- `resource_reference.type` points to the correct resource type string
- FieldMask used correctly on Update (not on Get/Delete)

---

### 4. Complete domain implementation *(user-identified)*
*All resources and operations from the source domain are represented.*

Indicators:
- Every resource in the domain has a message
- Every standard operation the domain needs has an RPC
- Embedded types (not standalone resources) are modelled as nested messages, not separate resources with their own CRUD
- Fields on resource messages cover the domain's data model (no glaring omissions)

---

### 5. Domain names *(user-identified)*
*Resource types, collection names, method names, and field names reflect the domain language — not generic tech terms.*

Indicators:
- Collection names are plural, domain-meaningful nouns (`pets`, not `items`)
- Resource type strings follow `{org}/{Service}/{Resource}` convention
- RPC names are `VerbNoun` in domain terms (`UploadPhoto`, not `AddFile`)
- Field names use domain vocabulary (`status`, `category`, `photo_urls` — not `state`, `group`, `image_list`)
- Enum value names carry the type prefix (`PET_STATUS_AVAILABLE`, not just `AVAILABLE`)

---

### 6. Reasonableness *(user-identified)*
*The design is proportionate to the domain — not over-specified or under-specified.*

Indicators:
- Deviations from the source spec are intentional and justified (e.g. AIP over OpenAPI PUT)
- Scope is respected: out-of-scope operations are absent
- No gold-plating: no extra resources, methods, or fields invented beyond the domain
- Custom methods are truly custom — not standard methods in disguise

---

### 7. Proto idioms
*The proto files are idiomatic proto3, independent of AIP.*

Indicators:
- `syntax = "proto3"` declared; file, package, and Go package options set correctly
- Field numbers are stable and gap-free from 1; reserved ranges used if fields were removed
- Enums have an `UNSPECIFIED = 0` value
- Well-known types used where appropriate: `google.protobuf.Timestamp`, `google.protobuf.FieldMask`, `google.protobuf.Empty`
- Import paths are correct and minimal (no unused imports)

---

### 8. AIP deviation explicitness
*Where the source spec and AIP conflict, the proto reflects a deliberate choice, not an accident.*

Indicators:
- Integer IDs → string name fields: done consistently, not mixed
- PUT → PATCH+FieldMask: done consistently across all resources
- Multiple filter endpoints → single List with `filter` string: done consistently
- Remaining deviations from a relevant AIP have a comment explaining why

---

### 9. Method orthogonality
*Methods are independent and non-overlapping. No method does more than one thing.*

Indicators:
- Get does not mutate
- Create does not also return a list
- Custom methods are not standard methods renamed
- No upsert RPCs hiding Create+Update logic unless explicitly AIP-compliant

---

### 10. Forward compatibility
*The design can evolve without breaking clients.*

Indicators:
- `optional` used on fields that may be absent (proto3 explicit presence)
- Response messages (not primitives) returned from all RPCs — allows adding fields later
- `page_token` is opaque (no documented internal structure)
- Enums are open to new values (no exhaustive handling implied in proto)

---

### 11. File organisation
*The proto files are laid out in a way a maintainer can navigate.*

Indicators:
- One service per file, or a coherent grouping with a clear rationale
- Messages co-located with the service that owns them
- Shared types in a separate `common.proto` if used across services
- Directory structure matches `api/proto/{service}/{version}/`

---

## Scoring summary template

```
Artifact: examples/petstore/api/proto/petstore/v1/

| #  | Dimension                  | Score | Notes |
|----|----------------------------|-------|-------|
|  1 | Standard compliance        |  / 2  |       |
|  2 | Annotation completeness    |  / 2  |       |
|  3 | Annotation correctness     |  / 2  |       |
|  4 | Complete domain impl       |  / 2  |       |
|  5 | Domain names               |  / 2  |       |
|  6 | Reasonableness             |  / 2  |       |
|  7 | Proto idioms               |  / 2  |       |
|  8 | AIP deviation explicitness |  / 2  |       |
|  9 | Method orthogonality       |  / 2  |       |
| 10 | Forward compatibility      |  / 2  |       |
| 11 | File organisation          |  / 2  |       |

Total: / 22
```

---

## Eval loop

1. Apply the `resource-oriented-design` skill to a domain
2. Score the output against this rubric
3. For each 0 or 1: identify whether the failure is in the skill prompt, the model, or the source material
4. Adjust the skill and re-run
5. Track scores across runs to measure improvement
