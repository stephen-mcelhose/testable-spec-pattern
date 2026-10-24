# AIP Normative Language — Raw Extract
Extracted 2026-10-24 via curl from aip-dev/google.aip.dev raw markdown.


## AIP-0122

All resource names defined by an API **must** be unique within that API. (See
- Resource name components **should** usually alternate between collection
- Resource names **must** use the `/` character to separate individual segments
  - Non-terminal segments of a resource name **must not** contain a `/`
  - The terminal segment of a resource name **should not** contain a `/`
- Resource names **should** only use characters available in DNS names, as
  - Additionally, resource IDs **should not** use upper-case letters.
  - If additional characters are necessary, resource names **should not** use
  - If Unicode characters can not be avoided, resource names **must** be stored
- Resources **must** expose a `name` field that contains its resource name.
  - Resources **may** provide the resource ID as a separate field (e.g.
    `book_id`). This field **must** apply the
  - Resources **may** expose a separate, system-generated unique ID field
    [(`uid`)](./0148.md#uid). This field **must** apply the
  - Resources **must not** expose tuples, self-links, or other forms of
  - All ID fields **should** be strings.
The collection identifier segments in a resource name **must** be the plural
- Collection identifiers **must** be concise American English terms.
- Collection identifiers **must** be in `camelCase`.
- Collection identifiers **must** begin with a lower-cased letter and contain
- Collection identifiers **must** be plural.
    is correct. Collection segments **must not** "coin" words by adding "s" in
- Within any given single resource name, collection identifiers **must** be
collection's name **may** omit the prefix. For example, given a collection of
An API **should** use the less-redundant form:
**Note:** APIs wishing to do this **must** follow this format consistently
- If resource IDs are user-specified, the API **must** document allowed
  formats. User-specified resource IDs **should** conform to [RFC-1034][];
  - Additionally, user-specified resource IDs **should** restrict letters to
  - Characters outside of ASCII **should not** be permitted; however, if
    Unicode characters are necessary, APIs **must** follow guidance in
- If resource IDs are not user-settable, the API **should** document the
APIs **may** provide programmatic aliases for common lookup patterns. However,
all data returned from the API **must** use the canonical resource name.
arbitrary API. In this situation, the service **should** use the _full resource
**Note:** The full resource name **should not** be used for cross-API
When defining a resource, the first field **should** be the resource name,
which **must** be of type `string` and **must** be called `name` for the
resource name. The message **should** include a `google.api.resource`
**should** be the resource name, which **must** be of type `string` and
**must** be called `name` for the resource name. The field **should** also be
**Note:** Fields **must not** be called `name` except for this purpose. For
of the request message **should** be of type `string` and **should** be called
`parent` for the resource name of the collection. The `parent` field **should**
If there is more than one possible parent type, the `parent` field **should**
**Note:** Fields **should not** be called `parent` except for this purpose. For
When a field represents another resource, the field **should** be of type
**should** be equivalent to the corresponding message's name in snake case.
- Field names **may** include a leading adjective if appropriate (such as
- Field names **should not** use the `_name` suffix unless the field would be
- Fields representing another resource **should** provide the
strictly necessary, the field **should** use an `_id` suffix (e.g. `shelf_id`).
The field **should not** be of type `message` using the `message` that

## AIP-0123

APIs **must** define a resource type for each resource in the API, according to
the following pattern: `{Service Name}/{Type}`. The type name **must**:
APIs **should** annotate the resource types for each resource in the API using
- Patterns **must** correspond to the [resource name][aip-122].
- Pattern variables (the segments within braces) **must** use `snake_case`, and
  **must not** use an `_id` suffix.
- Pattern variables **must** conform to the format `[a-z][_a-z0-9]*[a-z0-9]`.
- Pattern variables **must** be unique within any given pattern. (e.g.
- Resources with [multiple patterns][multi-pattern-resources] **must**
  preserve ordering: new patterns **must** be added at the end of the list, and
  existing patterns **must not** be removed or re-ordered, as this breaks client
- Singular **must** be the lower camel case of the type.
  - Pattern variables **must** be the singular form of the resource type e.g.
- Plural **must** be the lower camel case plural of the singular.
  - Pattern collection identifier segments **must** match the plural of the
When multiple patterns are defined within a resource, these patterns **must**
Therefore the following two patterns **must not** be defined within the same

## AIP-0131

APIs **must** provide a get method for resources. The purpose of the get method
- The RPC's name **must** begin with the word `Get`. The remainder of the RPC
  name **should** be the singular form of the resource's message name.
- The request message **must** match the RPC name, with a `Request` suffix.
- The response message **must** be the resource itself. (There is no
  - The response **should** usually include the fully-populated resource unless
- The HTTP verb **must** be `GET`.
- The URI **should** contain a single variable field corresponding to the
  - This field **should** be called `name`.
  - The URI **should** have a variable corresponding to this field.
  - The `name` field **should** be the only variable in the URI path. All
    remaining parameters **should** map to URI query parameters.
- There **must not** be a `body` key in the `google.api.http` annotation.
- There **should** be exactly one `google.api.method_signature` annotation,
- A resource name field **must** be included. It **should** be called `name`.
  - The field **should** be [annotated as required][aip-203].
  - The field **must** identify the [resource type][aip-123] that it
- The comment for the `name` field **should** document the resource pattern.
- The request message **must not** contain any other required fields, and
  **should not** contain other optional fields except those described in

## AIP-0132

APIs **must** provide a `List` method for resources unless the resource is a
- The RPC's name **must** begin with the word `List`. The remainder of the RPC
  name **should** be the plural form of the resource being listed.
- The request and response messages **must** match the RPC name, with
- The HTTP verb **must** be `GET`.
- The collection whose resources are being listed **should** map to the URI
  - The collection's parent resource **should** be called `parent`, and
    **should** be the only variable in the URI path. All remaining parameters
    **should** map to URI query parameters.
  - The collection identifier (`books` in the above example) **must** be a
- The `body` key in the `google.api.http` annotation **must** be omitted.
- If the resource being listed is not a top-level resource, there **should**
  **should** be either no `google.api.method_signature` annotation, or exactly
- A `parent` field **must** be included unless the resource being listed is a
  top-level resource. It **should** be called `parent`.
  - The field **should** be [annotated as required][aip-203].
  - The field **must** identify the [resource type][aip-123] of the resource
- The `page_size` and `page_token` fields, which support pagination, **must**
  - The comment above the `page_size` field **should** document the maximum
    to `0`). If preferred, the API **may** state that the server will use a
    sensible default. This default **may** change over time.
    **should** coerce the value to the maximum allowed.
  - If a user provides a negative or other invalid value, the API **must** send
- The `page_token` field **must** be included on all list request messages.
- The request message **may** include fields for common design patterns
- The request message **must not** contain any other required fields, and
  **should not** contain other optional fields except those described in this
**Note:** List methods **should** return the same results for any user that has
- The response message **must** include one repeated field corresponding to the
  resources being returned, and **should not** include any other repeated
  - The response **should** usually include fully-populated resources unless
- The `next_page_token` field, which supports pagination, **must** be included
  on all list response messages. It **must** be set if there are subsequent
  pages, and **must not** be set if the response represents the final page. For
- The message **may** include a `int32 total_size` (or `int64 total_size`)
  - The value **may** be an estimate (the field **should** clearly document
  - If filtering is used, the `total_size` field **should** reflect the size of
`List` methods **may** allow clients to specify sorting order; if they do, the
request message **should** contain a `string order_by` field.
- Values **should** be a comma separated list of fields. For example:
- The resulting list order **should** be based on the field type's natural
  lexicographically, etc. However, APIs **may** choose to use a different
  ordering; if so, it **must** be documented in the `order_by` definition.
List methods **may** allow clients to specify filters; if they do, the request
message **should** contain a `string filter` field. Filtering is described in
APIs that do this **should not** include deleted resources by default in list
requests. APIs with soft deletion of a resource **should** include a

## AIP-0133

APIs **should** generally provide a create method for resources unless it is
- The RPC's name **must** begin with the word `Create`. The remainder of the
  RPC name **should** be the singular form of the resource being created.
- The request message **must** match the RPC name, with a `Request` suffix.
- The response message **must** be the resource itself. There is no
  - The response **should** include the fully-populated resource, and **must**
    message **must** be a `google.longrunning.Operation` which resolves to the
- The HTTP verb **must** be `POST`.
- The collection where the resource is being added **should** map to the URI
  - The collection's parent resource **should** be called `parent`, and
    **should** be the only variable in the URI path.
  - The collection identifier (`books` in the above example) **must** be
- There **must** be a `body` key in the `google.api.http` annotation, and it
  **must** map to the resource field in the request message.
  - All remaining fields **should** map to URI query parameters.
- There **should** be exactly one `google.api.method_signature` annotation,
  [strong consistency][]: the completion of a create operation **must** mean
- A `parent` field **must** be included unless the resource being created is a
  top-level resource. It **should** be called `parent`.
  - The field **should** be [annotated as required][aip-203].
  - The field **must** identify the [resource type][aip-123] of the resource
- A `{resource}_id` field **must** be included for management plane resources,
  and **should** be included for data plane resources.
- The resource field **must** be included and **must** map to the POST body.
- The request message **must not** contain any other required fields and
  **should not** contain other optional fields except those described in this
regular API request. In this situation, the API **should** use a long-running
- The response type **must** be set to the resource (what the return type would
- Both the `response_type` and `metadata_type` fields **must** be specified.
**Important:** Declarative-friendly resources (AIP-128) **should** use
long-running operations. The service **may** return an LRO that is already set
An API **must** allow a user to specify the ID component of a resource (the last
On the [data plane][], an API **should** allow a user to specify the ID.
**may** allow the `{resource}_id` field have the [field_behavior][] `OPTIONAL`,
- The `{resource}_id` field **must** exist on the request message, not the
  - The field **may** be required or optional. If it is required, it **should**
- The `name` field on the resource **must** be ignored.
- There **should** be exactly one `google.api.method_signature` annotation on
- The documentation **should** explain what the acceptable format is, and the
  format **should** follow the guidance for resource name formatting in
  duplicate resource name, the service **must** error with `ALREADY_EXISTS`.
    duplicate resource, the service **must** error with `PERMISSION_DENIED`

## AIP-0134

APIs **should** generally provide an update method for resources unless it is
- The RPC's name **must** begin with the word `Update`. The remainder of the
  RPC name **should** be the singular form of the resource's message name.
- The request message **must** match the RPC name, with a `Request` suffix.
- The response message **must** be the resource itself. (There is no
  - The response **should** include the fully-populated resource, and **must**
    message **must** be a `google.longrunning.Operation` which resolves to the
- The method **should** support partial resource update, and the HTTP verb
  **should** be `PATCH`.
    HTTP verb **may** be `PUT`. However, this is strongly discouraged because
- The resource's `name` field **should** map to the URI path.
  - The `{resource}.name` field **should** be the only variable in the URI
- There **must** be a `body` key in the `google.api.http` annotation, and it
  **must** map to the resource field in the request message.
  - All remaining fields **should** map to URI query parameters.
- There **should** be exactly one `google.api.method_signature` annotation, with
  [strong consistency][]: the completion of an update operation **must** mean
- The request message **must** contain a field for the resource.
  - The field **must** map to the `PATCH` body.
  - The field **should** be [annotated as required][aip-203].
  - A `name` field **must** be included in the resource message. It **should**
  - The field **must** identify the [resource type][aip-123] of the resource
- If partial resource update is supported, a field mask **must** be included.
  It **must** be of type `google.protobuf.FieldMask`, and it **must** be
  - The field **must** be optional, and the service **must** treat an omitted
    `paths: []`), the service **should** treat the request as a no-op and
  - Update masks **must** support a special value `*`, meaning full replacement
- The request message **must not** contain any other required fields, and
  **should not** contain other optional fields except those described in this
Update methods **should not** trigger other side effects. Instead, side effects
**should** be triggered by custom methods.
In particular, this entails that [state fields][] **must not** be directly
regular API request. In this situation, the API **should** use a long-running
- The response type **must** be set to the resource (what the return type would
- Both the `response_type` and `metadata_type` fields **must** be specified.
**Note:** Declarative-friendly resources ([AIP-128][]) **should** use long-running
If the service uses client-assigned resource names, `Update` methods **may**
The user **must** have the update permissions to call `Update` even with
creating resources using the update method, IAM conditions **should** be used.
In this situation, the resource **should** contain a `string etag` field:
The `etag` field **may** be either [required][] or [optional][]. If it is set,
then the request **must** succeed if and only if the provided etag matches the
server-computed value, and **must** fail with an `ABORTED` error otherwise. The
In this situation, an API **may** return back only the fields that were
updated, and omit the rest, and **should** document this behavior if they do
does not exist, the service **must** error with `NOT_FOUND` (HTTP 404) unless
- **2024-03-14**: Make `update_mask` optional field_behaviour guidance a **must**.
- **2023-07-17**: Make `update_mask` name guidance a **must**.

## AIP-0135

APIs **should** generally provide a delete method for resources unless it is
- The RPC's name **must** begin with the word `Delete`. The remainder of the
  RPC name **should** be the singular form of the resource's message name.
- The request message **must** match the RPC name, with a `Request` suffix.
- The response message **should** be `google.protobuf.Empty`.
    **should** be the resource itself.
    message **must** be a `google.longrunning.Operation` which resolves to the
- The HTTP verb **must** be `DELETE`.
- The request message field receiving the resource name **should** map to the
  - This field **should** be called `name`.
  - The `name` field **should** be the only variable in the URI path. All
    remaining parameters **should** map to URI query parameters.
- There **must not** be a `body` key in the `google.api.http` annotation.
- There **should** be exactly one `google.api.method_signature` annotation,
  with a value of `"name"`. If an etag or force field are used, they **may** be
  [strong consistency][]: the completion of a delete operation **must** mean
- The API **must** fail with a `FAILED_PRECONDITION` error if child resources
  - If the only child resource type is a [Singleton][aip-156], deletion **must**
The Delete method **should** succeed if and only if a resource was present and
was successfully deleted. If the resource did not exist, the method **should**
- A `name` field **must** be included. It **should** be called `name`.
  - The field **should** be [annotated as required][aip-203].
  - The field **must** identify the [resource type][aip-123] that it
- The comment for the field **should** document the resource pattern.
- The request message **must not** contain any other required fields, and
  **should not** contain other optional fields except those described in this
regular API request. In this situation, the API **should** use a long-running
- The response type **must** be set to the appropriate return type if the RPC
- Both the `response_type` and `metadata_type` fields **must** be specified
**should** provide a `bool force` field on the request, which the user sets to
The API **must** fail with a `FAILED_PRECONDITION` error if the `force` field
the delete request **may** accept the etag (as either required or optional):
request **must** fail with a `ABORTED` error code.
**Note:** Declarative-friendly resources (AIP-128) **must** provide the `etag`
If the service uses client-assigned resource names, `Delete` methods **may**
**Note:** Declarative-friendly resources (AIP-128) **should** expose the
whether or not it exists, the service **must** error with `PERMISSION_DENIED`
(HTTP 403). Permission **must** be checked prior to checking if the resource
exist, the service **must** error with `NOT_FOUND` (HTTP 404) unless

## AIP-0136

Custom methods **should** only be used for functionality that can not be easily
- The name of the method **should** be a verb followed by a noun.
  - The name **must not** contain prepositions ("for", "with", etc.).
  - The verb in the name **should not** contain any of the standard method verbs ([Get][],
  - The name **must not** include the term `Async`. Instead, if the intention is
    **may** be used for this purpose. For example, to create a long-running book creation
- The HTTP method **must** be `GET` or `POST`:
  - `GET` **must** be used for methods retrieving data or resource state.
    - `POST` **may** be used for data retrieval methods if the request payload could
  - `POST` **must** be used if the method has side effects or mutates resources
- The HTTP URI **must** use a `:` character followed by the custom verb
  (`:archive` in the above example), and the verb in the URI **must** match the
  - If word separation is required, `camelCase` **must** be used.
- The `body` clause in the `google.api.http` annotation **should** be `"*"`.
- Custom methods **should** take a request message matching the RPC
- Custom methods **should** return a response message matching the RPC
  - When operating on a specific resource, a custom method **may** return the
Custom methods **must** operate on a resource if the API can be modeled
- The parameter for the resource's name **must** be called `name`, and
**may** operate on a collection instead:
- If the collection's resource has a parent, that resource **must** be called
- The collection key (`books` in the above example) **must** be literal.
  example), the field name in the request message **should** be the name of the
  scope resource. If word separators are necessary, `snake_case` **must** be
- The URI **should** place both the verb and noun after the `:` separator
- Stateless methods **must** use `POST` if they involve billing.
Declarative-friendly resources usually **should not** employ custom methods

## AIP-0158

RPCs returning collections of data **must** provide pagination _at the outset_,
- Request messages for collections **should** define an `int32 page_size`
  - The `page_size` field **must not** be required.
    chooses an appropriate default, which the API **should** document. The API
    **must not** return an error.
    API, the API **should** coerce down to the maximum permitted page size.
  - If the user specifies a negative value for `page_size`, the API **must**
  - The API **may** return fewer results than the number requested (including
- Request messages for collections **should** define a `string page_token`
  - The `page_token` field **must not** be required.
    service **must** honor the new page size.
    any arguments are different, the API **should** send an `INVALID_ARGUMENT`
- The response **must not** be a streaming response.
- Response messages for collections **should** define a
  - The field containing pagination results **should** be the first field in
    the message and have a field number of `1`. It **should** be a repeated
    **must** be empty. This is the _only_ way to communicate
    determine in time), the API **must** provide a `next_page_token`.
- Response messages for collections **may** provide an `int32 total_size`
  - This total **may** be an estimate (but the API **should** explicitly
The request definition for a paginated operation **may** define an `int32 skip`
The `skip` value **must** refer to the number of individual resources to skip,
querying a massive data set, the response **must** be `200 OK` with an empty
collection, the response **must not** include a `next_page_token`.
Page tokens provided by APIs **must** be opaque (but URL-safe) strings, and
**must not** be user-parseable. This is because if users are able to
contain sensitive data, an API **may** obfuscate the page token by defining an
Page tokens **must** be limited to providing an indication of where to continue
the pagination process only. They **must not** provide any form of
authorization to the underlying resources, and authorization **must** be
**may** expire page tokens a reasonable time after they have been sent, in
fields, they **must** be _actually implemented_ with a non-infinite default

## AIP-0160

APIs **may** provide filtering to users on `List` methods (or similar methods
**should** follow the common specification for filters discussed here. The
When employing filtering, a request message **should** have exactly one
Literals appearing alone (with no specified field) **should** usually be
However, a service **may** choose to only consider certain fields; if so, it
**must** document which fields it considers. A service **may** include new
fields over time, but **should** do so judiciously and consider impact on
Filtering implementations **should** provide the binary operators:
examples **should** encourage the use of explicit parentheses to avoid
confusion, but **should not** require explicit parentheses.
Filtering implementations **should** provide the unary operators `NOT` and `-`.
These are used interchangeably, and a service that supports negation **must**
Filtering implementations **should** provide the binary comparison operators
duration fields (but **should not** provide them for booleans or enums).
**Note:** Unlike in most programming languages, field names **must** appear on
Additionally, when comparing strings for equality, services **should** support
Filtering implementations **should** provide the `.` operator, which indicates
Traversal **must** be written using the field names from the resource. If a
service wishes to support "implicit fields" of some kind, they **must** do so
through well-documented [functions](#functions). A service **may** specify a
the service **should** return an error with `INVALID_ARGUMENT`. A service
**may** permit traversal to undefined keys on maps and structs, and **should**
in the chain is not set on the entry being evaluated, the entry **should** be
**Important:** The `.` operator **must not** be used to traverse through a
Filtering implementations **must** provide the `:` operator, which means "has".
  implementations **may** choose to support automatic conversion between camel
API-specific extensions. An API **may** define a function using the
`call(arg...)` syntax, and **must** document any specific functions it
A service **may** specify further structure or limitations for filter queries,
Further structure or limitations **must** be clearly documented, and
**must not** violate requirements set forth in this document.
the API **should** error with `INVALID_ARGUMENT`. Wherever validation is relaxed
for `filter`, the API **must** document the difference.
- Fields referenced in the `filter` **must** exist on the filtered schema
- Field values provided in the `filter` **must** align to the type of the field
  **must** be a valid value in the set
- Field values for standardized types e.g. `Timestamp` **must** conform to the
