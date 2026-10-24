# Petstore Proto Score Card — Iter 1

Scored against `.agents/skills/resource-oriented-design/evals/evals.json`
and `evals/rubric.md`. Domain adapted from Library API assertions to
Petstore (Pet/Order/User — all top-level resources, no parent hierarchy).

**api-linter result: 0 violations** (68 suppressed by disabling language-specific
rules for java, csharp, php, ruby — irrelevant to a Go-only target).

Changes from iter-0 baseline:
- `pet_name` → `pet` (AIP-122: no `_name` suffix on cross-resource fields)
- `PetStatus` → `Pet.State`, `OrderStatus` → `Order.State` (AIP-216: nested, prefer State)
- `name` on resources → `IDENTIFIER` (not `OUTPUT_ONLY`)
- `name` on resources: removed self-referential `resource_reference` (causes mutable cycle)
- All fields annotated with `field_behavior`
- All enum values documented; `STATE_UNSPECIFIED = 0` per AIP-126
- `method_signature` for List: omitted (top-level)
- `method_signature` for Create: includes `{resource}_id`
- `Pet.state` → `OPTIONAL` (user-settable; `0216::state-field-output-only` suppressed inline)
- `Order.complete` → `OPTIONAL` (user-settable)
- `first_name` / `last_name` → `given_name` / `family_name` (AIP-148)
- `ship_date` → `ship_time` (AIP-142)
- `ListUsers` added (AIP-121)
- File layout: service before messages (AIP-191)

---

## Eval results (evals.json assertions)

### AIP-122 Resource Names

| Assertion | Pass? | Notes |
| --------- | ----- | ----- |
| `name` field type `string`, called `name` | ✅ | All three resources |
| `name` is first field | ✅ | Field number 1 on all resources |
| Collection identifiers plural (`pets`, `orders`, `users`) | ✅ | |
| Collection identifiers lowercase, ASCII, camelCase | ✅ | Single-word, trivially correct |
| Segments alternate collection/ID: `pets/{pet}` | ✅ | |
| `google.api.resource` annotation present | ✅ | All three resources |
| Action request first field is `string name` | ✅ | Get, Delete requests |
| `name` field annotated `resource_reference` | ✅ | All Get/Delete requests |
| `uid` if present must be OUTPUT_ONLY | N/A | No uid field |
| Resource MUST NOT expose tuples/self-links | ✅ | Only `name` string identifier |
| Collection IDs unique within pattern | ✅ | Single-segment patterns |
| Cross-resource field named after message (snake_case) | ✅ | `pet` on Order |
| Cross-resource field SHOULD NOT use `_name` suffix | ✅ | Fixed: `pet` not `pet_name` |
| Cross-resource field SHOULD NOT be message type | ✅ | `pet` is string |

**AIP-122: 14/14** ✅

---

### AIP-123 Resource Types

| Assertion | Pass? | Notes |
| --------- | ----- | ----- |
| `type` in `google.api.resource` follows `{Service}/{Type}` | ✅ | `petstore.example.com/Pet` |
| `pattern` uses collection/ID alternation | ✅ | `pets/{pet}` |
| `singular` and `plural` set | ✅ | |

**AIP-123: 3/3** ✅

---

### AIP-131 Get

| Assertion | Pass? | Notes |
| --------- | ----- | ----- |
| Request message: `{Rpc}Request` suffix | ✅ | |
| Response is the resource itself | ✅ | Returns `Pet`, `Order`, `User` |
| HTTP verb `GET` | ✅ | |
| No `body` key | ✅ | |
| Resource name field included | ✅ | |
| `resource_reference` annotation | ✅ | |
| No other required fields | ✅ | |
| RPC name suffix is singular | ✅ | |
| `name` field called `name`, annotated REQUIRED | ✅ | |
| `name` only variable in URI path | ✅ | |
| One `method_signature` annotation with `"name"` | ✅ | |
| Comment documents resource pattern | ✅ | `// Format: pets/{pet}` |
| SHOULD NOT contain other optional fields | ✅ | No extras |

**AIP-131: 13/13** ✅

---

### AIP-132 List + AIP-158 + AIP-160

| Assertion | Pass? | Notes |
| --------- | ----- | ----- |
| RPC begins with `List` | ✅ | ListPets, ListOrders, ListUsers |
| Request: `ListXxxRequest`, Response: `ListXxxResponse` | ✅ | |
| HTTP verb `GET` | ✅ | |
| No `body` key | ✅ | |
| Collection identifier is literal string in URI | ✅ | `/v1/pets` |
| `parent` field present (non-top-level only) | N/A | Top-level resources — no parent required |
| RPC name suffix is plural | ✅ | `ListPets`, `ListOrders`, `ListUsers` |
| `method_signature` omitted for top-level list | ✅ | No annotation present |
| Results field is first, field number 1 | ✅ | `repeated Pet pets = 1` |
| `int32 page_size` field | ✅ | |
| `string page_token` field | ✅ | |
| `page_size` NOT required | ✅ | OPTIONAL |
| `page_token` NOT required | ✅ | OPTIONAL |
| `string next_page_token` in response | ✅ | |
| Response SHOULD NOT have other repeated fields | ✅ | Only the resource list |
| RPC MUST NOT use streaming | ✅ | No `stream` keyword |
| `string filter` field | ✅ | |
| `filter` NOT required | ✅ | OPTIONAL |

**AIP-132/158/160: 17/17** ✅

---

### AIP-133 Create

| Assertion | Pass? | Notes |
| --------- | ----- | ----- |
| RPC begins with `Create` | ✅ | |
| Request: `CreateXxxRequest` | ✅ | |
| Response is resource itself | ✅ | |
| HTTP verb `POST` | ✅ | |
| `body` maps to resource field | ✅ | `body: "pet"` |
| Collection identifier literal in URI | ✅ | `/v1/pets` |
| `parent` field (non-top-level only) | N/A | |
| Resource field present and maps to POST body | ✅ | |
| `{resource}_id` on request, NOT resource | ✅ | `pet_id` on CreatePetRequest |
| No other required fields beyond resource + id | ✅ | |
| RPC name suffix singular | ✅ | |
| One `method_signature` with `"{resource},{resource}_id"` | ✅ | `"pet,pet_id"` |
| `name` field on resource is IDENTIFIER | ✅ | Fixed from OUTPUT_ONLY |
| SHOULD NOT contain other optional fields | ✅ | |

**AIP-133: 13/13** ✅

---

### AIP-134 Update

| Assertion | Pass? | Notes |
| --------- | ----- | ----- |
| RPC begins with `Update` | ✅ | |
| Request: `UpdateXxxRequest` | ✅ | |
| Response is resource itself | ✅ | |
| `body` maps to resource field | ✅ | |
| Resource field present in request | ✅ | |
| `name` field in resource | ✅ | |
| `name` is IDENTIFIER (no resource_reference cycle) | ✅ | Self-referential resource_reference removed |
| `update_mask` of type `FieldMask`, called `update_mask` | ✅ | |
| `update_mask` is OPTIONAL | ✅ | Fixed |
| No other required fields | ✅ | |
| HTTP verb NOT `PUT` | ✅ | Uses PATCH |
| HTTP verb `PATCH` | ✅ | |
| `{resource}.name` maps to URI path | ✅ | `{pet.name=pets/*}` |
| `{resource}.name` only URI variable | ✅ | |
| One `method_signature` with `"{resource},update_mask"` | ✅ | |
| Resource field annotated REQUIRED | ✅ | |
| `name` called `name` | ✅ | |
| State fields correctly annotated | ✅ | `Pet.state` OPTIONAL (user-settable); `Order.state` OUTPUT_ONLY (server-managed) |
| SHOULD NOT contain other optional fields | ✅ | |

**AIP-134: 19/19** ✅

---

### AIP-135 Delete

| Assertion | Pass? | Notes |
| --------- | ----- | ----- |
| RPC begins with `Delete` | ✅ | |
| Request: `DeleteXxxRequest` | ✅ | |
| HTTP verb `DELETE` | ✅ | |
| No `body` key | ✅ | |
| `name` field, type string | ✅ | |
| `resource_reference` on `name` | ✅ | |
| No other required fields | ✅ | |
| RPC name suffix singular | ✅ | |
| Response is `google.protobuf.Empty` | ✅ | |
| `name` called `name`, annotated REQUIRED | ✅ | |
| `name` maps to URI path, only variable | ✅ | |
| One `method_signature` with `"name"` | ✅ | |
| SHOULD NOT contain other optional fields | ✅ | |

**AIP-135: 13/13** ✅

---

### AIP-136 Custom Methods

| Assertion | Pass? | Notes |
| --------- | ----- | ----- |
| HTTP method is `POST` (mutating) | ✅ | UploadPhoto mutates pet |
| URI uses `:` + custom verb | ✅ | `:uploadPhoto` |
| URI verb matches RPC name verb | ✅ | `UploadPhoto` → `uploadPhoto` |
| Resource name param called `name` | ✅ | |
| `name` only variable in URI path | ✅ | |
| RPC name NO prepositions | ✅ | |
| RPC name NO `Async` | ✅ | |
| Verb NOT a standard method verb | ✅ | `Upload` is not Get/List/Create/Update/Delete |
| RPC name is verb + noun | ✅ | `UploadPhoto` |
| `body` is `"*"` | ✅ | |
| Request matches RPC name + `Request` | ✅ | `UploadPhotoRequest` |
| Response matches RPC name + `Response` | ✅ | `UploadPhotoResponse` |
| `method_signature` includes all REQUIRED fields | ✅ | `"name,photo_data"` |

**AIP-136: 13/13** ✅

---

## Rubric score

| # | Dimension                    | Score | Notes |
|---|------------------------------|-------|-------|
| 1 | Standard compliance          | 2/2   | All methods AIP-shaped; custom method correct |
| 2 | Annotation completeness      | 2/2   | All fields annotated; resource, field_behavior, resource_reference, http all present |
| 3 | Annotation correctness       | 2/2   | `Pet.state` OPTIONAL (user-settable, inline lint suppression); `Order.state` OUTPUT_ONLY (server-managed); `name` IDENTIFIER |
| 4 | Complete domain impl         | 2/2   | All three resources; ListUsers added; UploadPhoto; Category/Tag embedded |
| 5 | Domain names                 | 2/2   | `pet` not `pet_name`; `given_name`/`family_name`; `ship_time`; `PetState`→`Pet.State` |
| 6 | Reasonableness               | 2/2   | No gold-plating; deviations justified |
| 7 | Proto idioms                 | 2/2   | `syntax = "proto3"`, UNSPECIFIED = 0, well-known types, nested enums per AIP-216 |
| 8 | AIP deviation explicitness   | 2/2   | Integer IDs → string names; PUT → PATCH; findByStatus/findByTags → filter |
| 9 | Method orthogonality         | 2/2   | No method does two things; UploadPhoto is genuinely custom |
| 10 | Forward compatibility        | 2/2   | All RPCs return messages; optional fields; no page_token structure implied |
| 11 | File organisation            | 2/2   | One service per file; service before messages (AIP-191); `api/proto/petstore/v1/` structure |

**Total: 22/22** ✅

---

## Skill additions that resolved iter-0 failures

| Iter-0 failure | Skill addition in iter-1 |
| -------------- | ------------------------- |
| `pet_name` instead of `pet` | Added cross-resource field naming rule: no `_name` suffix, use `resource_reference` |
| `status` annotation inconsistent | Added field behavior rules: OUTPUT_ONLY for server-assigned only; user-settable → OPTIONAL |
| `name` on resource had wrong annotation | Added: `name` on resource must be IDENTIFIER; do NOT add self-referential `resource_reference` |
| Empty `method_signature = ""` for top-level List | Added: top-level List omits `method_signature` entirely |
| `method_signature` for Create missing `{resource}_id` | Added: Create always includes `{resource},{resource}_id` |
| `PetStatus` / `OrderStatus` non-idiomatic | Added: prefer `State` over `Status`; nest enum inside resource as `State` |
| Enum values missing comments | Added: every enum value (including UNSPECIFIED = 0) MUST have a doc comment |
| `first_name` / `last_name` | Added: AIP-148 naming rule: `given_name` / `family_name` |
| `ship_date` instead of `ship_time` | Added: AIP-142 naming rule: field name must end in `_time` |
| File layout: messages before service | Added: AIP-191 file layout rule: service first |
| No `ListUsers` | Added: AIP-121 resource completeness: all resources SHOULD support List |
