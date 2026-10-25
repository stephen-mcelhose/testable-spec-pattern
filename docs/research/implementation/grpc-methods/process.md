---
type: how-to
title: "gRPC method implementation process"
description: Layered fill-out for each generated RPC — transport protovalidate + map, application business rules, domain Store — distilled from wiring PetService.
tags: [grpc, proto, implementation, petstore, aip, protovalidate, clean-architecture]
timestamp: 2026-10-25T19:05:48Z
---

# gRPC method implementation process

## Overview

Generated request/response types are the **external contract**, not the application model. Filling out each RPC is a layered pipeline:

1. **Transport** — `protovalidate.Validate` shape; map proto ↔ domain
2. **Application** — `validate*BusinessRules`; call Store
3. **Leaf** — persistence speaking **domain** types

Do not invent parallel DTOs for the wire. Do not leak generated types into Store or domain packages.

Worked example: `examples/petstore` — `internal/transport/grpc.PetServer` → `internal/service.PetService` → `Store` on `pet.Pet`.

This page is a how-to. The RFC 2119 MUSTS below are the seed for a future behavioral-spec / skill checklist.

## Layers

This layering follows [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) (Robert C. Martin): dependencies point inward only. The transport ring (`PetServer`) depends on the application ring (`PetService`), which depends on the domain (`pet.Pet`, `Store`). Nothing in an inner ring imports from an outer ring. `proto`/`gRPC` types are outer-ring concerns — they stay at the transport boundary.

```
PetServer (PetServiceServer)     ← generated server interface lives here
  validateShape / protovalidate
  map proto fields → domain args
       ↓
PetService                       ← no proto / gRPC imports
  validate*BusinessRules
  Store (pet.Pet)
       ↓
map domain → proto response
```

## The three steps (where they live)

### 1. Proto / shape validation — MUST be a valid `<Type>Request` (transport)

Encode shape in the `.proto` with `buf.validate`. At the top of every **adapter** method:

- A method **MUST** validate request shape with `protovalidate.Validate` (via `validateShape[T proto.Message]`).
- Shape violations **MUST** map to `InvalidArgument` (project typed stand-in OK).
- The application service **MUST NOT** be called on shape failure.

### 2. Business rules — application

Encapsulate in `validate<Method>BusinessRules(...)` on domain/primitive args. A no-op that returns success is the easy default when there are no extra rules.

Examples: required `display_name` after trim; `NormalizePetName`; List `page_size` default/coerce. Negative `page_size` is shape (`int32.gte = 0`), not business.

### 3. Implementation — application → Store on domain types

Build the minimal **domain** resource. Call `Store`. Map leaf errors (`AlreadyExists`, `NotFound`, `Aborted`, … pass through; else `Internal`). Adapter maps the domain result back to proto.

## Skeleton

```go
// transport/grpc — owns shape + mapping
func (s *PetServer) CreatePet(ctx context.Context, req *petstorev1.CreatePetRequest) (*petstorev1.Pet, error) {
	if err := validateShape(req); err != nil {
		return nil, err
	}
	created, err := s.Pets.CreatePet(ctx, req.GetPetId(), req.GetPet().GetDisplayName())
	if err != nil {
		return nil, err
	}
	return petToProto(created), nil
}

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

// service — owns business rules + Store (domain pet.Pet)
func (s *PetService) CreatePet(ctx context.Context, petID, displayName string) (pet.Pet, error) {
	name, display, err := validateCreatePetBusinessRules(petID, displayName)
	if err != nil {
		return pet.Pet{}, err
	}
	return s.Store.Create(ctx, pet.Pet{Name: name, DisplayName: display})
}
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
```

After editing constraints: `buf generate`, commit regenerated Go under `internal/gen/`.

## Boundary choices (template defaults)

These are the defaults this example teaches. Thin “proto all the way down” is a drill, not the production shape.

### Store on domain objects (not proto)

**Default:** `Store` speaks `pet.Pet`. Proto `petstorev1.Pet` is a wire DTO mapped in `internal/transport/grpc`.

**Why:** generated types are an external interface. Putting them in Store makes the core depend on the contract and teaches the wrong dependency direction.

### Transport adapter vs embedding the generated server on the app service

**Default:** `PetServer` embeds `UnimplementedPetServiceServer` and satisfies `PetServiceServer`. `PetService` has **no** generated server types.

**Why:** shape checks, proto mapping, and (later) `KnownError` → `google.rpc.Status` are transport concerns. Application rules and persistence stay free of gRPC.

**Still deferred (wiring, not layering):** listener / `RegisterPetServiceServer`; real status-code mapping.

### Alternative: collapsed transport + service (non-preferred)

A common shortcut embeds `UnimplementedPetServiceServer` directly on the application service, so one struct handles shape validation, business rules, and persistence calls.

```go
// collapsed: NOT the default taught here
type PetService struct {
    petstorev1.UnimplementedPetServiceServer
    Store Store
}

func (s *PetService) CreatePet(ctx context.Context, req *petstorev1.CreatePetRequest) (*petstorev1.Pet, error) {
    // shape + business rules + store call — all here
}
```

**Why it appears often:** less boilerplate, faster to wire up for small services.

**Risks relative to Clean Architecture:**
- Proto/gRPC types enter the application layer, coupling business logic to the wire contract. Schema changes in the `.proto` ripple into the service struct.
- `protovalidate` and mapping code sit alongside business rules, making each method a mixed concern. Unit-testing business rules requires constructing proto messages.
- There is no stable domain interface for the service — consumers (tests, future transports) must speak proto.
- Adding a second transport (HTTP/JSON gateway, CLI) forces extraction of the layer you skipped.

Use this shape only when the service is genuinely a thin adapter with no business rules and you are certain a second transport will never appear.

Future skills: parent `wire-external-interface` (shape → rules → port → leaf); child `wire-grpc-service` (generated `*Server` adapter, protovalidate, domain map, status map).

## Wiring checklist

1. `buf.gen.yaml` emits Go under `internal/gen/` (`paths=source_relative`).
2. Request messages carry `buf.validate`; **transport** calls `protovalidate.Validate`.
3. Commit generated `*.pb.go` / `*_grpc.pb.go`.
4. Adapter embeds `UnimplementedPetServiceServer` and satisfies `PetServiceServer`.
5. `Store` speaks domain resource types (`pet.Pet`).
6. Split tests by layer: `transport/grpc` owns shape + mapping to `PetAPI`; `service` owns business rules + Store. Do not assert business-rule failures in transport tests.

## Non-goals (this process)

- gRPC server registration / listen loop
- Mapping `KnownError` → `google.rpc.Status` at the wire boundary
- Implementing every RPC in one pass (UploadPhoto can wait)

## References

- Adapter: `examples/petstore/internal/transport/grpc/`
- Application: `examples/petstore/internal/service/service.go`
- Domain: `examples/petstore/internal/pet/`
- Generated types: `examples/petstore/internal/gen/petstore/v1/`
- [protovalidate](https://github.com/bufbuild/protovalidate)
- [AIP Proto Learnings — iter-1 Petstore](../../aip-proto-learnings.md)
- [Contract-First as the Foundation](../../../ideation/contract-first.md)
- [ADR-004: Use internal/ not pkg/](../../../decisions/004-internal-over-pkg.md)
