---
type: how-to
title: "gRPC method implementation process"
description: Three-step fill-out for each generated RPC — protovalidate shape, business rules, then storage — distilled from wiring PetService to petstorev1 types.
tags: [grpc, proto, implementation, petstore, aip, protovalidate]
timestamp: 2026-10-25T18:56:31Z
---

# gRPC method implementation process

## Overview

When a service already has generated request/response types (`protoc-gen-go` / `protoc-gen-go-grpc`), filling out each RPC method is a straight three-step pipeline. Do not invent parallel DTOs. Accept the generated `*FooRequest`, return the generated response type, and keep domain helpers (name normalization, typed errors) in a small package beside the service.

Worked example: `examples/petstore` `PetService` over `internal/gen/petstore/v1` (`petstorev1`).

This page is a how-to. The RFC 2119 MUSTS below are the seed for a future behavioral-spec / skill checklist — we are not splitting a second competing RFC doc here.

## The three steps

Every method body follows the same order. First step is **always** proto shape validation via `protovalidate.Validate`.

### 1. Proto / shape validation — MUST be a valid `<Type>Request`

Encode shape in the `.proto` with `buf.validate` annotations (`required`, `string.min_len`, `int32.gte`, …).

- A method **MUST** validate request shape with `protovalidate.Validate` (constraints from `buf.validate` on the request message). Use the shared `validateShape[T proto.Message]` helper (nil / typed-nil guard, then `protovalidate.Validate`).
- Shape violations **MUST** map to `InvalidArgument` (project typed stand-in OK).
- Storage **MUST NOT** be called on shape failure.

Do **not** hand-roll REQUIRED-field ifs when a `buf.validate` constraint already covers them. Shape lives in the proto; Go runs the stereotypical check.

```go
if err := validateShape(req); err != nil {
	return nil, err
}
```

### 2. Business rules — listed rules, or room left

Encapsulate in `validate<Method>BusinessRules(...)`. A no-op that returns success (and any coerced values) is the easy default when there are no extra rules.

Rules that are true for this product but are not (yet) encoded as `buf.validate` constraints:

- Required display fields (`display_name` non-empty after trim) even when the proto marks the field OPTIONAL
- Canonical resource-name format after trim (`NormalizePetName`) — more than `min_len`
- Defaults and coercion (List `page_size`: default 50, coerce above max 100 — AIP-132). Negative `page_size` is shape (`int32.gte = 0`), not business.
- Optional freshness tokens (etag on Update/Delete — AIP-154); mismatch is a storage concern that surfaces as `Aborted`
- Explicitly **out of scope** items stay as comments / test headers (system-generated ids, `update_mask`, `filter`, `allow_missing`, …) — leave room, do not implement quietly

If a business rule later becomes a `buf.validate` constraint, move the check up into step 1 (delete the Go branch).

### 3. Implementation — usually delegate to storage

Build the minimal resource the store needs (ignore create-body `name`; use `pet_id` for the resource name — AIP-133). Call `Store`. Map leaf errors:

- Known domain codes pass through (`AlreadyExists`, `NotFound`, `Aborted`, …)
- Anything else → `Internal`

Return the stored resource (or `google.protobuf.Empty` for Delete). Do not re-derive server fields the store already set (etag, etc.).

## Skeleton

```go
func (s *PetService) CreatePet(ctx context.Context, req *petstorev1.CreatePetRequest) (*petstorev1.Pet, error) {
	// 1. Proto / shape — stereotypical check: protovalidate.Validate
	if err := validateShape(req); err != nil {
		return nil, err
	}

	// 2. Business rules
	name, display, err := validateCreatePetBusinessRules(req)
	if err != nil {
		return nil, err
	}

	// 3. Storage
	created, err := s.Store.Create(ctx, &petstorev1.Pet{Name: name, DisplayName: display})
	if err != nil {
		return nil, mapStoreError(err)
	}
	return created, nil
}

// Stereotypical proto shape engine: protovalidate.Validate over buf.validate.
func validateShape[T proto.Message](msg T) error {
	var zero T
	if any(msg) == any(zero) || isNilProto(msg) {
		return invalidArgument("request is required")
	}
	if err := protovalidate.Validate(msg); err != nil {
		return invalidArgument(err.Error())
	}
	return nil
}

// No-op default when a method has no extra rules:
// func validateFooBusinessRules(req *FooRequest) error { return nil }
```

## Proto side (shape contract)

```proto
import "buf/validate/validate.proto";

message CreatePetRequest {
  Pet pet = 1 [
    (google.api.field_behavior) = REQUIRED,
    (buf.validate.field).required = true
  ];
  string pet_id = 2 [(google.api.field_behavior) = OPTIONAL];
}

message ListPetsRequest {
  int32 page_size = 1 [
    (google.api.field_behavior) = OPTIONAL,
    (buf.validate.field).int32.gte = 0
  ];
}
```

After editing constraints: `buf generate`, commit regenerated Go under `internal/gen/`.

## Boundary choice

Defaults for this thin-wire example; flip when the criteria below are met.

### Store on proto types vs domain objects

**Default here:** `Store` speaks `*petstorev1.Pet`. The teaching target is the method fill-out ritual (shape → rules → storage). A parallel domain `Pet` is ceremony until domain language diverges.

**Flip when:** domain terms diverge from proto; multiple transports share one model; rich invariants must not depend on generated code. Then map `petstorev1.Pet` ↔ `domain.Pet` at the boundary; Store speaks domain.

### Service embeds generated server vs transport adapter

**Default here:** `PetService` embeds `UnimplementedPetServiceServer` and satisfies `PetServiceServer`. That *is* wiring to the external contract. Typed `KnownError` stands in for `google.rpc.Status`.

**Flip when:** mapping KnownError → gRPC codes/details; interceptors; multi-service registration. Then a thin `internal/transport/grpc` (or similar) adapter owns status mapping; domain/application service stays free of generated server types.

Future skills: parent `wire-external-interface` (shape → rules → port → leaf); child `wire-grpc-service` (generated `*Server`, protovalidate, `validate*BusinessRules`, status map, optional domain map).

## Wiring checklist

1. `buf.gen.yaml` emits Go under `internal/gen/` (`paths=source_relative`).
2. Request messages carry `buf.validate` for shape; service calls `protovalidate.Validate` via `validateShape`.
3. Commit generated `*.pb.go` / `*_grpc.pb.go`.
4. Service embeds `UnimplementedPetServiceServer` and satisfies `PetServiceServer` (unimplemented RPCs stay on the embed until their issue).
5. `Store` speaks generated resource types (`*petstorev1.Pet`), not hand-written twins — until Boundary choice says otherwise.
6. Behavioral specs take `*petstorev1.<Method>Request` and assert KnownError codes — same MUST language as before.

## Non-goals (this process)

- gRPC server registration / transport wiring
- Mapping `KnownError` → `google.rpc.Status` codes at the wire boundary
- Implementing every RPC on the service in one pass (UploadPhoto can wait)
- Introducing a domain package solely to satisfy Clean/Onion layering

## References

- Petstore implementation: `examples/petstore/internal/service/service.go`
- Generated types: `examples/petstore/internal/gen/petstore/v1/`
- [protovalidate](https://github.com/bufbuild/protovalidate)
- [AIP Proto Learnings — iter-1 Petstore](../../aip-proto-learnings.md)
- [Contract-First as the Foundation](../../../ideation/contract-first.md)
- [03 — PetStore resource-oriented API](../../../plans/stephen-workflow-skills/03-petstore-api.md)
