# AIP Eval Coverage Gap Audit

**Method:** `curl` → raw markdown from `aip-dev/google.aip.dev` → grep normative keywords →
manual triage (proto-observable vs behavioral) → compare against `evals.json`.

**Why curl, not WebFetch:** WebFetch is an LLM summarizer. It compresses and drops requirements.
Raw markdown is the ground truth. This file is the audit trail for that pipeline.

**Source:** `https://raw.githubusercontent.com/aip-dev/google.aip.dev/master/aip/general/{aip}.md`
**Extracted:** 2026-10-24
**Raw normative lines file:** `normative-raw.md`

---

## Triage key

- ✅ Covered — assertion exists in `evals.json`
- ❌ Gap — proto-observable, no assertion
- ⬜ Behavioral — runtime/server behavior, not verifiable from proto alone
- 🔶 MAY — optional permission, not a constraint

---

## AIP-0122 Resource Names

| Raw normative line | Observable? | Status |
| --- | --- | --- |
| All resource names **must** be unique within that API | No (runtime) | ⬜ |
| Resource name components **should** usually alternate between collection/ID segments | Yes | ✅ |
| Resource names **must** use `/` to separate segments | Yes | ✅ |
| Non-terminal segments **must not** contain `/` | Yes | ✅ |
| Terminal segment **should not** contain `/` | Yes | ✅ (implied) |
| Resource names **should** only use DNS-safe characters | Yes | ✅ (ID field guidance) |
| Resource IDs **should not** use upper-case letters | Yes | ✅ |
| Resource names **should not** use URL-escaping chars | Yes | ✅ (implied) |
| Unicode **must** be stored in Normalization Form C | No (storage) | ⬜ |
| Resources **must** expose a `name` field | Yes | ✅ |
| Separate resource ID field **must** apply `OUTPUT_ONLY` | Yes | ✅ |
| `uid` field **must** apply `OUTPUT_ONLY` | Yes | **❌ GAP** |
| Resources **must not** expose tuples, self-links, or other forms of resource identification | Yes | **❌ GAP** |
| All ID fields **should** be strings | Yes | ✅ |
| Collection identifiers **must** be plural | Yes | ✅ |
| Collection identifiers **must** be concise American English | Yes | ✅ |
| Collection identifiers **must** be in camelCase | Yes | ✅ |
| Collection identifiers **must** begin with lowercase, ASCII only | Yes | ✅ |
| Collection segments **must not** "coin" words (e.g. "infos") | Yes | ✅ (implied by plural rule) |
| Collection identifiers **must** be unique within any given single resource name | Yes | **❌ GAP** |
| Nested collections: API **must** follow omitted-prefix format consistently | Yes | ✅ (implied) |
| User-specified IDs: API **must** document allowed formats | No (docs) | ⬜ |
| User-specified IDs **should** conform to RFC-1034 | No (runtime) | ⬜ |
| Characters outside ASCII **should not** be permitted in user-specified IDs | No (runtime) | ⬜ |
| Aliases: all data **must** use canonical resource name | No (runtime) | ⬜ |
| Resource `name` field **must** be of type `string` and **must** be called `name` | Yes | ✅ |
| Message **should** include `google.api.resource` annotation | Yes | ✅ |
| Action request first field **must** be type `string`, **must** be called `name` | Yes | ✅ |
| Action request `name` field **should** be annotated with `google.api.resource_reference` | Yes | ✅ |
| Fields **must not** be called `name` except for the resource name | Yes | ✅ |
| Collection request first field **should** be `string parent` | Yes | ✅ |
| `parent` field **should** be annotated with `google.api.resource_reference` | Yes | ✅ |
| If multiple parent types, `parent` **should** use `child_type` on `resource_reference` | Yes | ✅ (implied by resource_reference annotation) |
| Fields **should not** be called `parent` except for parent collection | Yes | ✅ |
| Field representing another resource **should** be type `string` (not message type) | Yes | **❌ GAP** |
| Field name **should** be equivalent to corresponding message's name in snake_case | Yes | **❌ GAP** |
| Field names **should not** use `_name` suffix unless ambiguous | Yes | **❌ GAP** |
| Fields representing another resource **should** provide `google.api.resource_reference` | Yes | ✅ |
| If ID only: field **should** use `_id` suffix | Yes | ✅ (implied) |
| Field **should not** be of type `message` using the message that implements the resource | Yes | **❌ GAP** (duplicate of string type gap above — same rule) |

**AIP-122 gaps: 5** (uid OUTPUT_ONLY, no tuples/self-links, collection ID uniqueness, field snake_case naming, no _name suffix, not message type)

---

## AIP-0123 Resource Types

| Raw normative line | Observable? | Status |
| --- | --- | --- |
| APIs **must** define a resource type for each resource | Yes | ✅ |
| Type name **must** start uppercase, alphanumeric, singular, PascalCase | Yes | ✅ |
| APIs **should** annotate using `google.api.resource` | Yes | ✅ |
| Patterns **must** correspond to resource name | Yes | ✅ |
| Pattern variables **must** use snake_case, **must not** use `_id` suffix | Yes | ✅ |
| Pattern variables **must** conform to `[a-z][_a-z0-9]*[a-z0-9]` | Yes | ✅ |
| Pattern variables **must** be unique within pattern | Yes | ✅ |
| Resources with multiple patterns **must** preserve ordering; new at end; **must not** remove/reorder | Yes | ✅ |
| `singular` **must** be lower camelCase | Yes | ✅ |
| Pattern variables **must** be singular form of resource type | Yes | ✅ |
| `plural` **must** be lower camelCase plural | Yes | ✅ |
| Pattern collection identifiers **must** match plural | Yes | ✅ |
| Multiple patterns **must** be mutually unique | Yes | ✅ |
| **Must not** define mutually non-unique patterns | Yes | ✅ |

**AIP-123 gaps: 0** ✅

---

## AIP-0131 Get

| Raw normative line | Observable? | Status |
| --- | --- | --- |
| APIs **must** provide a get method | No (design) | ⬜ |
| RPC name **must** begin with `Get` | Yes | ✅ |
| RPC name **should** be singular form of resource | Yes | ✅ |
| Request message **must** match RPC name + `Request` suffix | Yes | ✅ |
| Response message **must** be the resource itself | Yes | ✅ |
| Response **should** include fully-populated resource | No (runtime) | ⬜ |
| HTTP verb **must** be `GET` | Yes | ✅ |
| URI **should** contain variable for `name` field | Yes | ✅ |
| Field **should** be called `name` | Yes | ✅ |
| `name` field **should** be only variable in URI path | Yes | ✅ |
| Remaining params **should** map to URI query params | Yes | ✅ |
| **Must not** have `body` key in annotation | Yes | ✅ |
| **Should** be exactly one `method_signature` annotation | Yes | ✅ |
| Resource name field **must** be included | Yes | ✅ |
| Field **should** be called `name`, **should** be annotated as required | Yes | ✅ |
| Field **must** identify resource type via `resource_reference` | Yes | ✅ |
| Comment **should** document resource pattern | Yes | ✅ |
| Request **must not** contain other required fields | Yes | ✅ |
| Request **should not** contain other optional fields | Yes | **❌ GAP** |

**AIP-131 gaps: 1** (SHOULD NOT contain other optional fields)

---

## AIP-0132 List

| Raw normative line | Observable? | Status |
| --- | --- | --- |
| APIs **must** provide List method (unless singleton) | No (design) | ⬜ |
| RPC name **must** begin with `List` | Yes | ✅ |
| RPC name **should** be plural form | Yes | ✅ |
| Request and response messages **must** match RPC name + suffix | Yes | ✅ |
| HTTP verb **must** be `GET` | Yes | ✅ |
| Collection **should** map to URI path | Yes | ✅ |
| `parent` **should** be only variable in URI path | Yes | ✅ |
| Remaining params **should** map to URI query params | Yes | ✅ |
| Collection identifier **must** be a literal string | Yes | ✅ |
| `body` key **must** be omitted | Yes | ✅ |
| `parent` field **must** be included (non-top-level) | Yes | ✅ |
| `parent` **should** be called `parent` | Yes | ✅ |
| `parent` **should** be annotated as required | Yes | ✅ |
| `parent` **must** identify resource type via `resource_reference` | Yes | ✅ |
| `page_size` and `page_token` **must** be included | Yes | **❌ GAP** (evals say SHOULD via AIP-158; AIP-132 says MUST) |
| `page_token` **must** be included on all list request messages | Yes | **❌ GAP** (same) |
| Request **must not** contain other required fields | Yes | ✅ |
| Request **should not** contain other optional fields (except per AIP) | Yes | **❌ GAP** |
| Response **must** include one repeated field for the resources | Yes | ✅ |
| Response **should not** include any other repeated fields | Yes | **❌ GAP** |
| Response **should** include fully-populated resources | No (runtime) | ⬜ |
| `next_page_token` **must** be included on all list response messages | Yes | **❌ GAP** (evals say SHOULD via AIP-158; AIP-132 says MUST) |
| `next_page_token` **must** be set if subsequent pages; **must not** be set if final page | No (runtime) | ⬜ |
| `order_by` **should** be `string` if sorting provided | Yes | 🔶 MAY (sorting is optional) |
| `filter` **should** be `string` if filtering provided | Yes | ✅ |

**AIP-132 gaps: 5** (page_size/page_token MUST not SHOULD, page_token on all requests, SHOULD NOT other optional fields, SHOULD NOT other repeated fields in response, next_page_token MUST)

---

## AIP-0133 Create

| Raw normative line | Observable? | Status |
| --- | --- | --- |
| APIs **should** generally provide create | No (design) | ⬜ |
| RPC name **must** begin with `Create` | Yes | ✅ |
| RPC name **should** be singular form | Yes | ✅ |
| Request message **must** match RPC name + `Request` suffix | Yes | ✅ |
| Response **must** be the resource; if LRO **must** be `google.longrunning.Operation` | Yes (LRO branch) | **❌ GAP** |
| Response **should** include fully-populated resource | No (runtime) | ⬜ |
| HTTP verb **must** be `POST` | Yes | ✅ |
| Collection **should** map to URI path | Yes | ✅ |
| `parent` **should** be only variable in URI path | Yes | ✅ |
| Collection identifier **must** be literal string | Yes | ✅ |
| **Must** be a `body` key mapping to resource field | Yes | ✅ |
| Remaining fields **should** map to URI query params | Yes | ✅ |
| `parent` field **must** be included (non-top-level) | Yes | ✅ |
| `parent` **should** be called `parent`, **should** be annotated required | Yes | ✅ |
| `parent` **must** identify resource type via `resource_reference` | Yes | ✅ |
| `{resource}_id` **must** be included for management plane | Yes | ✅ |
| `{resource}_id` **should** be included for data plane | Yes | ✅ |
| Resource field **must** be included and **must** map to POST body | Yes | ✅ |
| Request **must not** contain other required fields | Yes | ✅ |
| Request **should not** contain other optional fields | Yes | ✅ |
| LRO: `response_type` **must** be set to the resource; `response_type` and `metadata_type` **must** be specified | Yes (LRO branch) | **❌ GAP** |
| `{resource}_id` **must** exist on request message, not the resource | Yes | ✅ |
| `name` field on resource **must** be ignored on creation | Yes | ✅ |
| ID format **should** follow AIP-122 guidance | Yes | ✅ (implied) |

**AIP-133 gaps: 2** (LRO return type, LRO operation_info annotation)

---

## AIP-0134 Update

| Raw normative line | Observable? | Status |
| --- | --- | --- |
| APIs **should** generally provide update | No (design) | ⬜ |
| RPC name **must** begin with `Update` | Yes | ✅ |
| RPC name **should** be singular form | Yes | ✅ |
| Request message **must** match RPC name + `Request` suffix | Yes | ✅ |
| Response **must** be the resource; if LRO **must** be `google.longrunning.Operation` | Yes (LRO branch) | **❌ GAP** |
| Response **should** include fully-populated resource | No (runtime) | ⬜ |
| Method **should** support partial update, HTTP verb **should** be `PATCH` | Yes | ✅ |
| HTTP verb **may** be `PUT` (strongly discouraged) | Yes | ✅ (SHOULD NOT PUT) |
| Resource `name` field **should** map to URI path | Yes | ✅ |
| `{resource}.name` **should** be only variable in URI path | Yes | ✅ |
| **Must** be `body` key mapping to resource field | Yes | ✅ |
| Remaining fields **should** map to URI query params | Yes | ✅ |
| One `method_signature` annotation **should** exist | Yes | ✅ |
| Request **must** contain a field for the resource | Yes | ✅ |
| Resource field **must** map to PATCH body | Yes | ✅ |
| Resource field **should** be annotated required | Yes | ✅ |
| `name` field **must** be included in resource; **should** be called `name` | Yes | ✅ |
| `name` field **must** identify resource type via `resource_reference` | Yes | ✅ |
| If partial update: field mask **must** be included; **must** be `google.protobuf.FieldMask`; **must** be called `update_mask` | Yes | ✅ |
| `update_mask` **must** be optional | Yes | ✅ |
| If empty `update_mask` (`paths: []`), **should** treat as no-op | No (runtime) | ⬜ |
| Update masks **must** support `*` (full replacement) | No (runtime) | ⬜ |
| Request **must not** contain other required fields | Yes | ✅ |
| Request **should not** contain other optional fields | Yes | ✅ |
| Update methods **should not** trigger side effects | No (design) | ⬜ |
| State fields **must not** be directly writable | Yes | ✅ |
| LRO: `response_type` and `metadata_type` **must** be specified | Yes (LRO branch) | **❌ GAP** |
| `allow_missing` bool field **may** be exposed | Yes | 🔶 MAY |
| User **must** have update permissions even with `allow_missing` | No (runtime) | ⬜ |
| `etag` field **may** be required or optional | Yes | 🔶 MAY |
| etag: request **must** fail with ABORTED if mismatch | No (runtime) | ⬜ |

**AIP-134 gaps: 2** (LRO return type, LRO operation_info annotation)

---

## AIP-0135 Delete

| Raw normative line | Observable? | Status |
| --- | --- | --- |
| APIs **should** generally provide delete | No (design) | ⬜ |
| RPC name **must** begin with `Delete` | Yes | ✅ |
| RPC name **should** be singular form | Yes | ✅ |
| Request message **must** match RPC name + `Request` suffix | Yes | ✅ |
| Response **should** be `google.protobuf.Empty`; if soft-delete **should** be the resource | Yes | ✅ |
| If LRO: response **must** be `google.longrunning.Operation` | Yes (LRO branch) | **❌ GAP** |
| HTTP verb **must** be `DELETE` | Yes | ✅ |
| `name` field **should** map to URI path | Yes | ✅ |
| `name` field **should** be called `name` | Yes | ✅ |
| `name` field **should** be only variable in URI path | Yes | ✅ |
| Remaining params **should** map to URI query params | Yes | ✅ |
| **Must not** be `body` key in annotation | Yes | ✅ |
| One `method_signature` annotation **should** exist | Yes | ✅ |
| `name` field **must** be included | Yes | ✅ |
| `name` field **should** be annotated as required | Yes | ✅ |
| `name` field **must** identify resource type via `resource_reference` | Yes | ✅ |
| Comment **should** document resource pattern | Yes | ✅ |
| Request **must not** contain other required fields | Yes | ✅ |
| Request **should not** contain other optional fields | Yes | ✅ |
| LRO: `response_type` and `metadata_type` **must** be specified | Yes (LRO branch) | **❌ GAP** |
| `bool force` field **should** be provided if child resources possible | Yes | 🔶 MAY (design choice) |
| etag: request **must** fail with ABORTED if mismatch | No (runtime) | ⬜ |
| Declarative-friendly: **must** provide `etag` field | Yes | 🔶 MAY (AIP-128 specific) |
| `allow_missing` bool **may** be exposed | Yes | 🔶 MAY |
| PERMISSION_DENIED check: **must** error before resource check | No (runtime) | ⬜ |
| NOT_FOUND: **must** error if resource not found | No (runtime) | ⬜ |

**AIP-135 gaps: 2** (LRO return type, LRO operation_info annotation)

---

## AIP-0136 Custom Methods

| Raw normative line | Observable? | Status |
| --- | --- | --- |
| Custom methods **should** only be used when standard methods can't | No (design) | ⬜ |
| Method name **should** be verb + noun | Yes | ✅ |
| Name **must not** contain prepositions | Yes | ✅ |
| Verb **should not** be a standard method verb | Yes | ✅ |
| Name **must not** include `Async` | Yes | ✅ |
| HTTP method **must** be `GET` or `POST` | Yes | **❌ GAP** (we only cover the POST/mutating case; GET for retrieval not asserted) |
| `GET` **must** be used for methods retrieving data or resource state | Yes | **❌ GAP** |
| `POST` **must** be used if method has side effects or mutates | Yes | ✅ |
| URI **must** use `:` + custom verb | Yes | ✅ |
| Verb in URI **must** match verb in RPC name | Yes | ✅ |
| If word separation in URI verb: **must** use camelCase | Yes | ✅ (implied) |
| `body` **should** be `"*"` | Yes | ✅ |
| Custom methods **should** take `Request`-suffixed message | Yes | ✅ |
| Custom methods **should** return `Response`-suffixed message | Yes | ✅ |
| Custom methods **must** operate on a resource if the API can be modeled as such | No (design) | ⬜ |
| Resource name param **must** be called `name`; **must** be only variable in URI path | Yes | ✅ |
| Collection-based: parent resource **must** be called `parent`; **must** be only variable | Yes | **❌ GAP** (we only cover resource-based; collection-based custom methods not covered) |
| Collection key **must** be literal | Yes | **❌ GAP** |
| Stateless scope field **should** be named after scope resource; **must** use snake_case | Yes | **❌ GAP** |
| Stateless methods **must** use `POST` if billing involved | Yes | **❌ GAP** |
| Declarative-friendly **should not** employ custom methods | No (design) | ⬜ |

**AIP-136 gaps: 6** (GET for retrieval, collection-based parent=`parent`, collection key literal, stateless snake_case, stateless billing POST, overall GET/POST constraint only partially covered)

---

## AIP-0158 Pagination

| Raw normative line | Observable? | Status |
| --- | --- | --- |
| RPCs returning collections **must** provide pagination at the outset | No (design) | ⬜ |
| `int32 page_size` **should** be defined | Yes | ✅ |
| `page_size` **must not** be required | Yes | ✅ |
| API **should** document default page size | No (docs) | ⬜ |
| API **should** coerce large page_size to max | No (runtime) | ⬜ |
| Negative page_size: API **must** send error | No (runtime) | ⬜ |
| `string page_token` **should** be defined | Yes | ✅ |
| `page_token` **must not** be required | Yes | ✅ |
| API **must** honor new page_size on subsequent pages | No (runtime) | ⬜ |
| Response **must not** be a streaming response | Yes | **❌ GAP** (observable: RPC return type must not be `stream`) |
| `string next_page_token` **should** be defined | Yes | ✅ |
| Results field **should** be first field at field number 1 | Yes | ✅ |
| Results field **should** be repeated | Yes | ✅ |
| `next_page_token` **must** be empty at end of collection | No (runtime) | ⬜ |
| `next_page_token` **must** be provided if not at end | No (runtime) | ⬜ |
| `int32 skip` field **may** be defined | Yes | 🔶 MAY |
| Page tokens **must** be opaque URL-safe strings | No (runtime) | ⬜ |
| Page tokens **must not** be user-parseable | No (runtime) | ⬜ |
| Page tokens **must** be limited to pagination continuation only | No (runtime) | ⬜ |
| Page tokens **must not** provide authorization | No (runtime) | ⬜ |
| Pagination fields **must** actually be implemented | No (runtime) | ⬜ |

**AIP-158 gaps: 1** (response must not be streaming — observable as RPC return type)

---

## AIP-0160 Filtering

All proto-observable requirements (the single `string filter` field) are covered.
The remaining normative lines concern filter-language evaluation semantics (runtime/behavioral).

**AIP-160 gaps: 0** ✅

---

## Summary

| AIP  | Proto-observable | Covered | Gaps |
| ---- | ---------------- | ------- | ---- |
| 122  | 25               | 20      | **5** |
| 123  | 14               | 14      | 0 |
| 131  | 16               | 15      | **1** |
| 132  | 18               | 13      | **5** |
| 133  | 13               | 11      | **2** |
| 134  | 15               | 13      | **2** |
| 135  | 14               | 12      | **2** |
| 136  | 14               | 8       | **6** |
| 158  | 8                | 7       | **1** |
| 160  | 2                | 2       | 0 |
| **Total** | **139**     | **115** | **24** |

---

## Gaps to add to evals.json

### AIP-122
1. `uid` field, if present, **must** be annotated `OUTPUT_ONLY`
2. Resource **must not** expose tuples, self-links, or other resource identification forms
3. Collection identifiers **must** be unique within any given single resource name pattern
4. Field referencing another resource **should** be named after the message in snake_case
5. Field referencing another resource **should not** use `_name` suffix unless ambiguous
6. Field referencing another resource **should not** be of the message type that implements the resource (use `string`)

### AIP-131
7. Request **should not** contain other optional fields except those described in AIP-131 or another AIP

### AIP-132
8. `page_size` field **must** be included (not just SHOULD — AIP-132 is stronger than AIP-158 on this)
9. `page_token` field **must** be included on all list request messages
10. `next_page_token` field **must** be included on all list response messages
11. Request **should not** contain other optional fields (except per AIP)
12. Response **should not** include any other repeated fields beyond the resources field

### AIP-133
13. If LRO: response **must** be `google.longrunning.Operation` resolving to the resource
14. If LRO: `response_type` and `metadata_type` **must** be specified in `operation_info`

### AIP-134
15. If LRO: response **must** be `google.longrunning.Operation` resolving to the resource
16. If LRO: `response_type` and `metadata_type` **must** be specified in `operation_info`

### AIP-135
17. If LRO: response **must** be `google.longrunning.Operation`
18. If LRO: `response_type` and `metadata_type` **must** be specified in `operation_info`

### AIP-136
19. HTTP method **must** be `GET` for methods that only retrieve data or resource state
20. Collection-based custom method: parent resource param **must** be called `parent`; **must** be only variable in URI path
21. Collection-based custom method: collection key **must** be a literal string (not a variable)
22. Stateless method scope field: word separators **must** use snake_case in request message field name
23. Stateless methods **must** use `POST` if they involve billing

### AIP-158
24. RPC **must not** use streaming response — observable: return type is not `stream Foo`
