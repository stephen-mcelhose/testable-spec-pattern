# Petstore Proto Score Card

Scored against `.agents/skills/resource-oriented-design/evals/evals.json`
and `evals/rubric.md`. Domain adapted from Library API assertions to
Petstore (Pet/Order/User — all top-level resources, no parent hierarchy).

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
| Cross-resource field named after message (snake_case) | ✅ | `pet_name` on Order |
| Cross-resource field SHOULD NOT use `_name` suffix | ❌ | `pet_name` violates this — should be `pet` |
| Cross-resource field SHOULD NOT be message type | ✅ | `pet_name` is string |

**AIP-122: 13/14** — one failure: `pet_name` should be `pet` per SHOULD NOT `_name` suffix rule.

---

### AIP-123 Resource Types

| Assertion | Pass? | Notes |
| --------- | ----- | ----- |
| `google.api.resource` with `type` field | ✅ | All three |
| `type` follows `{Service}/{Type}` pattern | ✅ | `petstore.example.com/Pet` etc. |
| Type name starts uppercase, PascalCase, singular, alphanumeric | ✅ | `Pet`, `Order`, `User` |
| `pattern` field present | ✅ | |
| Pattern variables use snake_case, no `_id` suffix | ✅ | `{pet}`, `{order}`, `{user}` |
| Pattern variables singular form of type | ✅ | |
| `singular` and `plural` set | ✅ | |
| `singular` lower camelCase, `plural` lower camelCase plural | ✅ | |

**AIP-123: 8/8** ✅

---

### AIP-131 Get

| Assertion | Pass? | Notes |
| --------- | ----- | ----- |
| RPC begins with `Get` | ✅ | GetPet, GetOrder, GetUser |
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

**AIP-131: 14/14** ✅

---

### AIP-132 List + AIP-158 + AIP-160

| Assertion | Pass? | Notes |
| --------- | ----- | ----- |
| RPC begins with `List` | ✅ | ListPets, ListOrders |
| Request: `ListXxxRequest`, Response: `ListXxxResponse` | ✅ | |
| HTTP verb `GET` | ✅ | |
| No `body` key | ✅ | |
| Collection identifier is literal string in URI | ✅ | `/v1/pets` |
| `parent` field present (non-top-level only) | N/A | Top-level resources — no parent required |
| RPC name suffix is plural | ✅ | `ListPets`, `ListOrders` |
| `method_signature` annotation exists | ⚠️ | Present as `""` — should be omitted for top-level list |
| Results field is first, field number 1 | ✅ | `repeated Pet pets = 1` |
| `int32 page_size` field | ✅ | |
| `string page_token` field | ✅ | |
| `page_size` NOT required | ✅ | |
| `page_token` NOT required | ✅ | |
| `string next_page_token` in response | ✅ | |
| Response SHOULD NOT have other repeated fields | ✅ | Only the resource list |
| RPC MUST NOT use streaming | ✅ | No `stream` keyword |
| `string filter` field | ✅ | |
| `filter` NOT required | ✅ | |

**AIP-132/158/160: 16/17** — one warning: empty `method_signature = ""` is unusual for top-level List; should likely be omitted.

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
| One `method_signature` | ✅ | |
| `name` field on resource is OUTPUT_ONLY | ✅ | |
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
| `name` has `resource_reference` | ❌ | `name` on Pet/User has no `resource_reference` annotation — only field_behavior |
| `update_mask` of type `FieldMask`, called `update_mask` | ✅ | |
| `update_mask` NOT required | ✅ | |
| No other required fields | ✅ | |
| HTTP verb NOT `PUT` | ✅ | Uses PATCH |
| HTTP verb `PATCH` | ✅ | |
| `{resource}.name` maps to URI path | ✅ | `{pet.name=pets/*}` |
| `{resource}.name` only URI variable | ✅ | |
| One `method_signature` with `"{resource},update_mask"` | ✅ | |
| Resource field annotated REQUIRED | ✅ | |
| `name` called `name` | ✅ | |
| State fields NOT directly writable | ❌ | `status` on Pet is OUTPUT_ONLY which is too restrictive — but more importantly, if status IS writable it should NOT be OUTPUT_ONLY. Annotation is inconsistent with domain intent. |
| SHOULD NOT contain other optional fields | ✅ | |

**AIP-134: 17/19** — two failures: missing `resource_reference` on resource `name` field; `status` annotation ambiguous (OUTPUT_ONLY on a user-settable field).

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

**AIP-136: 12/12** ✅

---

## Rubric score

| # | Dimension | Score | Notes |
|---|---|---|---|
| 1 | Standard compliance | 2/2 | All methods AIP-shaped; custom method correct |
| 2 | Annotation completeness | 2/2 | resource, field_behavior, resource_reference, http all present |
| 3 | Annotation correctness | 1/2 | `status`/`complete` marked OUTPUT_ONLY but are user-settable; `name` on resource missing `resource_reference` |
| 4 | Complete domain impl | 2/2 | All three resources; UploadPhoto; Category/Tag as embedded types; Login/Logout absent |
| 5 | Domain names | 1/2 | `pet_name` violates SHOULD NOT `_name` suffix; should be `pet` with `resource_reference` |
| 6 | Reasonableness | 2/2 | No gold-plating; deviations justified (PATCH over PUT, filter over findByStatus) |
| 7 | Proto idioms | 2/2 | `syntax = "proto3"`, package options, UNSPECIFIED = 0, well-known types, no unused imports |
| 8 | AIP deviation explicitness | 2/2 | Integer IDs → string names; PUT → PATCH; findByStatus/findByTags → filter |
| 9 | Method orthogonality | 2/2 | No method does two things; UploadPhoto is genuinely custom |
| 10 | Forward compatibility | 2/2 | All RPCs return messages; optional fields; no page_token structure implied |
| 11 | File organisation | 2/2 | One service per file; messages co-located; `api/proto/petstore/v1/` structure |

**Total: 20/22**

---

## Failures attributed to skill gaps

| Failure | What the skill says | What it should say |
| ------- | ------------------- | ------------------ |
| `pet_name` instead of `pet` | Nothing about cross-resource field naming | Add: when referencing another resource, name the field after the resource type in snake_case — NOT `{resource}_name` |
| `status` marked OUTPUT_ONLY incorrectly | Nothing about when to use field_behavior annotations | Add: OUTPUT_ONLY for server-assigned fields only (name, create_time, etag); user-settable fields are not OUTPUT_ONLY |
| `name` on resource missing `resource_reference` | Nothing about `resource_reference` | Add: the `name` field on the resource itself should have `(google.api.resource_reference)` pointing to its own type |
| Empty `method_signature = ""` for top-level List | Nothing about method_signature | Add: top-level List (no parent) should omit method_signature or use `""` only if no required fields |

**4 skill gaps identified. Score would be 22/22 with those gaps closed.**
