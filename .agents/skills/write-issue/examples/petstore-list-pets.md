# Example: Petstore — List Pets endpoint

This is a worked example of the `write-issue` skill output.

**Topic given:** Petstore CRUD application using Resource Oriented Design (ROD) — List Pets.  
**Scope:** One resource endpoint. Additional CRUD operations (Create, Get, Update, Delete) are out of scope for this issue and would each be separate issues.

---

## What

A `GET /pets` endpoint that returns a paginated list of pets from the petstore collection, conforming to the Resource Oriented Design (ROD) conventions.

## Why

The petstore API has no list endpoint. Callers cannot discover existing pets without knowing their IDs in advance, which makes the API unusable for any browse or search flow. ROD requires that every collection resource exposes a standard list operation before individual resource operations are built on top of it.

## Use-case

A frontend developer building the petstore UI calls `GET /pets` with an optional `page_token` query parameter. The server returns the first page of pets as a JSON array under a `pets` key, along with a `next_page_token` for cursor-based pagination. If the collection is empty, the response is `200 OK` with an empty `pets` array — not a `404`. The developer does not need to know pet IDs in advance to populate the initial view.

## Acceptance criteria

- [ ] `GET /pets` returns `200 OK` with a JSON body containing a `pets` array
- [ ] Each pet object in the response includes at minimum: `name` (string), `id` (string), `status` (one of `available`, `pending`, `sold`)
- [ ] The response includes a `next_page_token` field; the value is an opaque string when more results exist, and omitted (or `null`) on the last page
- [ ] A `page_size` query parameter (integer, default `20`, max `100`) controls the number of results returned
- [ ] A `page_token` query parameter, when provided, resumes from the cursor position it encodes
- [ ] An empty collection returns `200 OK` with `{"pets": []}`, not `404`
- [ ] The endpoint is documented in the OpenAPI spec under the `pets` collection resource

## Out of scope

- Filtering or searching pets by field values (separate issue)
- Authentication and authorisation on the endpoint (separate issue — list is read-only and public in the initial spec)
- The Create, Get by ID, Update, and Delete pet endpoints (each a separate issue)
- Client SDK generation from the OpenAPI spec (downstream of the spec being complete)
