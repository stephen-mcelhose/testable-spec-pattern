# PetService — Behavioural Specification

## 1. Introduction & Scope

This document specifies the external behaviour of the application Pet service: business-rule validation, interaction with persistence, and the errors and resources a caller observes.

**In scope:** Create, Get, Update, Delete, and List operations on Pet; name and display rules; pagination defaults and coercion; freshness (etag) handling; error codes returned to the caller.

**Out of scope:** gRPC or HTTP transport, request-message shape validation at the wire boundary (see [PetServer specification](../../transport/grpc/specification.md)), and concrete storage technology.

## 2. Terminology

The key words "MUST", "MUST NOT", "REQUIRED", "SHOULD", "SHOULD NOT", "RECOMMENDED", "MAY", and "OPTIONAL" in this document are to be interpreted as described in [RFC 2119](https://datatracker.ietf.org/doc/html/rfc2119).

| Term | Meaning |
| ---- | ------- |
| **Caller** | The component invoking PetService (typically the transport adapter). |
| **Store** | The persistence port that creates, reads, updates, deletes, and lists Pets. |
| **Pet** | The domain pet resource (see §4). |
| **Resource name** | The canonical identifier of a Pet in the form `pets/{pet}`. |
| **Pet id** | The final segment of the resource name; used on create to form the resource name. |

## 3. Behavioral Requirements

### 3.1 Common error handling

1. When the Store returns a domain error the service recognizes (`INVALID_ARGUMENT`, `NOT_FOUND`, `ALREADY_EXISTS`, `ABORTED`, or `INTERNAL` as produced by the Store), PetService MUST return that error to the Caller unchanged in kind.
2. When the Store returns any other failure, PetService MUST return an `INTERNAL` error to the Caller.
3. When a business rule fails, PetService MUST return an `INVALID_ARGUMENT` error and MUST NOT call the Store for that operation.

### 3.2 CreatePet

4. The Caller MUST supply a pet id and a display name.
5. PetService MUST reject a missing or ill-formed pet id with `INVALID_ARGUMENT`.
6. A well-formed pet id MUST match `[a-z0-9-]{1,63}` with no leading or trailing hyphen; the resulting resource name MUST be `pets/{pet_id}` after trimming surrounding whitespace on the id.
7. PetService MUST reject a display name that is empty after trimming surrounding whitespace with `INVALID_ARGUMENT`.
8. On success, PetService MUST ask the Store to create a Pet whose name is the canonical resource name derived from the pet id and whose display name is the trimmed display name.
9. PetService MUST ignore any caller-supplied resource name on the create body; the pet id alone MUST determine the resource name.
10. When the Store reports that the resource already exists, PetService MUST return `ALREADY_EXISTS`.
11. On success, PetService MUST return the Pet as stored (including any server-assigned fields such as etag).

### 3.3 GetPet

12. The Caller MUST supply a resource name.
13. PetService MUST reject a resource name that is not a canonical `pets/{pet}` id (after trimming surrounding whitespace) with `INVALID_ARGUMENT`.
14. PetService MUST request the Pet from the Store using the canonical resource name.
15. When the Store reports the resource is missing, PetService MUST return `NOT_FOUND`.
16. On success, PetService MUST return the stored Pet.

### 3.4 UpdatePet

17. The Caller MUST supply a resource name and a display name; the Caller MUST supply an etag.
18. PetService MUST reject a resource name that is not a canonical `pets/{pet}` id (after trimming) with `INVALID_ARGUMENT`.
19. PetService MUST reject a display name that is empty after trimming with `INVALID_ARGUMENT`.
20. PetService MUST ask the Store to update the Pet identified by the canonical name with the trimmed display name, passing the trimmed etag (empty when omitted).
21. When the Store reports the resource is missing, PetService MUST return `NOT_FOUND`.
22. When an etag is provided and the Store reports a freshness mismatch, PetService MUST return `ABORTED`.
23. When etag is omitted (empty after trim), PetService MUST still attempt the update; the Store MUST treat an empty etag as no freshness precondition.
24. On success, PetService MUST return the Pet as stored.

### 3.5 DeletePet

25. The Caller MUST supply a resource name; the Caller MAY supply an etag.
26. PetService MUST reject a resource name that is not a canonical `pets/{pet}` id (after trimming) with `INVALID_ARGUMENT`.
27. PetService MUST ask the Store to delete the Pet at the canonical name, passing the trimmed etag (empty when omitted).
28. When the Store reports the resource is missing, PetService MUST return `NOT_FOUND`.
29. When an etag is provided and the Store reports a freshness mismatch, PetService MUST return `ABORTED`.
30. When etag is omitted, PetService MUST still attempt the delete; the Store MUST treat an empty etag as no freshness precondition.
31. On success, PetService MUST complete without returning a Pet body.

### 3.6 ListPets

32. The Caller MAY supply a page size and a page token.
33. When page size is unspecified or zero, PetService MUST use a default page size of 50.
34. When page size is greater than 100, PetService MUST coerce it to 100 and MUST NOT fail the request solely for that reason.
35. Negative page size is outside this service’s contract; the transport boundary MUST reject it before delegation (see [PetServer specification](../../transport/grpc/specification.md)).
36. PetService MUST request a page from the Store using the effective page size and the supplied page token.
37. On success, PetService MUST return the page of Pets and a next-page token (empty when no further pages exist).
38. An empty collection MUST be returned as success with an empty Pet list, not as `NOT_FOUND`.

## 4. Data Models / Interfaces

### 4.1 Pet

| Field | Description |
| ----- | ----------- |
| `name` | Resource name `pets/{pet}` — identifies the resource. |
| `display_name` | Human-readable name; REQUIRED non-empty after trim on create and update. |
| `etag` | OPTIONAL freshness token; server-assigned on write; Caller MAY send it on update/delete. |

### 4.2 Store port

The Store MUST support create, get, update (with optional etag precondition), delete (with optional etag precondition), and list (page size + page token → pets + next-page token). Observed error kinds from the Store that PetService preserves are listed in §3.1.

### 4.3 Error codes

| Code | When observed by the Caller |
| ---- | --------------------------- |
| `INVALID_ARGUMENT` | Business-rule failure (name/id/display). |
| `NOT_FOUND` | Get/Update/Delete when the Pet does not exist. |
| `ALREADY_EXISTS` | Create when the Pet already exists. |
| `ABORTED` | Update/Delete when a provided etag does not match. |
| `INTERNAL` | Non-domain Store failure. |
