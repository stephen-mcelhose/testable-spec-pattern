---
type: how-to
title: "gRPC method implementation process"
description: Three-step fill-out for each generated RPC — proto validation, business rules, then storage — distilled from wiring PetService to petstorev1 types.
tags: [grpc, proto, implementation, petstore, aip]
timestamp: 2026-10-25T18:21:19Z
---

# gRPC method implementation process

## Overview

When a service already has generated request/response types (`protoc-gen-go` / `protoc-gen-go-grpc`), filling out each RPC method is a straight three-step pipeline. Do not invent parallel DTOs. Accept the generated `*FooRequest`, return the generated response type, and keep domain helpers (name normalization, typed errors) in a small package beside the service.

Worked example: `examples/petstore` `PetService` over `internal/gen/petstore/v1` (`petstorev1`).

## The three steps

Every method body follows the same order. First step is **always** proto validation.

### 1. Proto / shape validation — must be a valid `<Type>Request`

Reject anything that is not a well-formed request message before business logic runs:

- `req == nil`
- REQUIRED message fields missing (e.g. `CreatePetRequest.pet`, `UpdatePetRequest.pet`)
- REQUIRED identifiers empty or the wrong resource-name shape (often via a shared normalizer such as `NormalizePetName`)

Return `InvalidArgument` (or the project's typed stand-in). Do not call storage.

When protovalidate / `buf.validate` constraints exist on the proto, this step is where they run. Until then, mirror REQUIRED field behavior and name format in code.

### 2. Business rules — listed rules, or room left

Rules that are true for this product but are not (yet) encoded as proto constraints:

- Required display fields (`display_name` non-empty after trim) even when the proto marks the field OPTIONAL
- Defaults and coercion (List `page_size`: default 50, coerce above max 100 — AIP-132)
- Optional freshness tokens (etag on Update/Delete — AIP-154); mismatch is a storage concern that surfaces as `Aborted`
- Explicitly **out of scope** items stay as comments / test headers (system-generated ids, `update_mask`, `filter`, `allow_missing`, …) — leave room, do not implement quietly

If a business rule later becomes a proto constraint, move the check up into step 1.

### 3. Implementation — usually delegate to storage

Build the minimal resource the store needs (ignore create-body `name`; use `pet_id` for the resource name — AIP-133). Call `Store`. Map leaf errors:

- Known domain codes pass through (`AlreadyExists`, `NotFound`, `Aborted`, …)
- Anything else → `Internal`

Return the stored resource (or `google.protobuf.Empty` for Delete). Do not re-derive server fields the store already set (etag, etc.).

## Skeleton

```go
func (s *PetService) CreatePet(ctx context.Context, req *petstorev1.CreatePetRequest) (*petstorev1.Pet, error) {
	// 1. Proto / shape — valid CreatePetRequest
	if req == nil || req.GetPet() == nil {
		return nil, invalidArgument("…")
	}

	// 2. Business rules
	name, err := nameFromID(req.GetPetId())
	if err != nil {
		return nil, err
	}
	display := strings.TrimSpace(req.GetPet().GetDisplayName())
	if display == "" {
		return nil, invalidArgument("display_name is required")
	}

	// 3. Storage
	created, err := s.Store.Create(ctx, &petstorev1.Pet{Name: name, DisplayName: display})
	if err != nil {
		return nil, mapStoreError(err)
	}
	return created, nil
}
```

## Wiring checklist

1. `buf.gen.yaml` emits Go under `internal/gen/` (`paths=source_relative`).
2. Commit generated `*.pb.go` / `*_grpc.pb.go`.
3. Service embeds `UnimplementedPetServiceServer` and satisfies `PetServiceServer` (unimplemented RPCs stay on the embed until their issue).
4. `Store` speaks generated resource types (`*petstorev1.Pet`), not hand-written twins.
5. Behavioral specs take `*petstorev1.<Method>Request` and assert KnownError codes — same MUST language as before.

## Non-goals (this process)

- gRPC server registration / transport wiring
- Mapping `KnownError` → `google.rpc.Status` codes at the wire boundary
- Implementing every RPC on the service in one pass (UploadPhoto can wait)

## References

- Petstore implementation: `examples/petstore/internal/service/service.go`
- Generated types: `examples/petstore/internal/gen/petstore/v1/`
- [AIP Proto Learnings — iter-1 Petstore](../../aip-proto-learnings.md)
- [Contract-First as the Foundation](../../../ideation/contract-first.md)
- [03 — PetStore resource-oriented API](../../../plans/stephen-workflow-skills/03-petstore-api.md)
