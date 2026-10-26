# AIP Coverage Audit

## Summary
Total normative requirements (proto-observable): 106
Covered by assertions: 90
**Gaps: 16**

## Per-AIP breakdown

### AIP-122 Resource Names
Proto-observable requirements: 25 / Covered: 19

**Gaps:**
- [MUST NOT] Resources must not expose tuples, self-links, or other forms of resource identification → Observable in resource message fields.
- [MUST] If a resource exposes a system-generated unique ID field (`uid`), this field must apply the `OUTPUT_ONLY` field behavior annotation → Observable in proto field annotations.
- [SHOULD] The field name equivalent to the corresponding referenced message's name in `snake_case` → Observable in field naming.
- [SHOULD NOT] Field names should not use the `_name` suffix unless the field would be ambiguous → Observable in field naming.
- [SHOULD NOT] The field should not be of type message using the message that implements the resource (with exceptions) → Observable in field types.
- [MUST] Collection identifiers must be unique within any given single resource name path pattern → Observable in resource pattern string.

### AIP-123 Resource Types
Proto-observable requirements: 14 / Covered: 14

**Gaps:**
- 0 Gaps.

### AIP-131 Get Methods
Proto-observable requirements: 10 / Covered: 10

**Gaps:**
- 0 Gaps.

### AIP-132 List Methods
Proto-observable requirements: 15 / Covered: 15

**Gaps:**
- 0 Gaps.

### AIP-133 Create Methods
Proto-observable requirements: 12 / Covered: 10

**Gaps:**
- [MUST] If a resource takes longer to create than reasonable, the API should use a long-running operation; response must be `google.longrunning.Operation` → Observable in RPC return type.
- [MUST] LRO metadata: `response_type` and `metadata_type` must be specified in `operation_info` → Observable in RPC annotations.

### AIP-134 Update Methods
Proto-observable requirements: 15 / Covered: 12

**Gaps:**
- [MUST] If update is long-running, response must be `google.longrunning.Operation` → Observable in RPC return type.
- [MUST] LRO metadata: `response_type` and `metadata_type` must be specified in `operation_info` → Observable in RPC annotations.
- [MAY] Request message may expose a `bool allow_missing` field for create-on-update → Observable in request message fields.

### AIP-135 Delete Methods
Proto-observable requirements: 11 / Covered: 9

**Gaps:**
- [MUST] If delete is long-running, response must be `google.longrunning.Operation` → Observable in RPC return type.
- [MUST] LRO metadata: `response_type` and `metadata_type` must be specified in `operation_info` → Observable in RPC annotations.

### AIP-136 Custom Methods
Proto-observable requirements: 12 / Covered: 9

**Gaps:**
- [MAY] Suffix `LongRunning` may be used for long-running RPCs → Observable in RPC name.
- [MUST] Stateless methods must use `POST` if they involve billing → Observable in HTTP annotation.
- [MUST] Custom methods must operate on a resource if the API can be modeled as such → Structural design requirement.

### AIP-158 Pagination
Proto-observable requirements: 6 / Covered: 6

**Gaps:**
- 0 Gaps.

### AIP-160 Filtering
Proto-observable requirements: 4 / Covered: 4

**Gaps:**
- 0 Gaps.
