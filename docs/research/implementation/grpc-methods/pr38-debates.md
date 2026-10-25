---
type: research
title: "PR #38 architecture debates — P3 responses"
description: Responses to the three deferred (CLARIFICATION_NEEDED) comments from PR #38 review — doc format, proto in the service layer, and use-cases vs business-rules naming.
tags: [grpc, clean-architecture, ddd, documentation, petstore]
timestamp: 2026-10-25T21:30:00Z
---

# PR #38 — Priority 3 debate responses

Three comments from the PR review needed research before deciding. This doc
collects the position on each, with verifiable links, so the decision is
durable rather than conversational.

---

## F — How-to vs RFC 2119 spec for `process.md`

**Your comment:** "Change this to RFC2119 language. Hmm… actually we may revisit, I think the how-to actually makes sense, but only if we write real how-to (with verifiable assertions). Lets talk about it."

### What the formats actually mean

[Diátaxis](https://diataxis.fr/how-to-guides/) (the canonical doc taxonomy used by Django, NumPy, and the Linux kernel docs project) draws the line clearly:

| Format    | Written from      | Goal                          | Completeness         |
|-----------|-------------------|-------------------------------|----------------------|
| How-to    | User's perspective | Get a specific thing done     | Omit the unnecessary |
| Reference | Machinery's perspective | Describe every input/output | Seek completeness |
| Specification (RFC 2119) | System's perspective | Define required behaviour | Normative, exhaustive |

A how-to makes a **verifiable contract**: *"If you execute steps A, B, C under conditions X, you will get result Y."* If a competent reader follows the steps and does not get the promised result, the contract is broken. That is the test.

An RFC 2119 spec defines **what a system must/shall/may do** — it is normative and describes the system, not the reader's workflow.

### Position

`process.md` is correctly a **how-to** — it guides *you* through filling out an RPC, not through specifying what the runtime must do. The existing RFC 2119 MUSTs in the doc are not wrong in a how-to; they function as
conditional imperatives: *"If you are writing an adapter method, you MUST
validate shape first."* That is Diátaxis-legal: it guides the reader's
judgment, not just their keystrokes.

**What is missing** is verifiability at each step. The how-to contract is currently implicit. To make it real:

- Step 1 (shape) needs a "you can verify this by…" — e.g. a failing test when `protovalidate.Validate` is removed.
- Step 2 (business rules) needs a concrete "if display_name is empty after trim, the service returns `InvalidArgument`" assertion you can run.
- Step 3 (Store) needs: "run `go test ./...` and confirm no panics on nil domain fields."

**Recommendation:** Keep the how-to format. Harden each step by adding a "verify" bullet that a reader can actually execute. The RFC 2119 MUST seeds already in the doc are the right starting point for a future behavioral spec (`go-behavioral-spec` skill); they do not need to replace the how-to now.

---

## G — Should proto types flow into the service layer? Do we need a separate server/grpc layer?

**Your comments:**
- "This is a questionable choice. Explore DDD/Clean/Onion. Should we have moved to domain objects by this point? Proto feels like external interface contract. Pros/Cons it with me."
- "Do we need server/grpc-interface layer? PROS/CONS. Again take inspiration from DDD/Clean Architecture for 'debate'."

### What the architecture literature says

Three independent sources agree on the same point:

**Clean Architecture — Robert C. Martin** ([blog post](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)):

> "Source code dependencies can only point inwards. No code in an inner circle can know anything at all about anything declared in an outer circle. Outer data formats must not be passed into inner circles."

Generated proto types (`petstorev1.CreatePetRequest`) are outer-circle data formats — they are artefacts of `buf generate` and the gRPC framework. The service (`PetService`) is inner circle. The Dependency Rule prohibits proto types in `PetService`.

**Hexagonal Architecture — Alistair Cockburn** ([original paper](https://alistair.cockburn.us/hexagonal-architecture/)):

> "The application is 'blissfully ignorant' of the concrete technologies interacting with it."

gRPC is a technology. The application core (service) must not know it exists.

**Three Dots Tech — Combining DDD, CQRS, and Clean Architecture in Go** ([article](https://threedots.tech/post/introducing-clean-architecture/)):

> "Proto-specific structures are kept entirely inside the adapter. The mapping and conversion from protobuf types to domain-friendly native Go types happen inside the adapter — invisible to the application service."

All three reach the same conclusion: proto at the transport adapter, domain types inside.

### Pros/Cons: proto in the service layer

| | Proto in service | Domain types in service (current) |
|---|---|---|
| **Boilerplate** | Less — no mapping step | More — explicit mapping in adapter |
| **Testability** | Must construct proto messages in service tests | Service tests use plain Go values |
| **Coupling** | Proto schema changes ripple into business logic | Proto changes are contained at the adapter |
| **Second transport** | Adding HTTP/JSON requires extracting the layer you skipped | HTTP adapter maps to the same service interface |
| **Spec alignment** | Business rules expressed in proto vocabulary | Business rules expressed in domain vocabulary |
| **Architecture pattern** | Operation Script / Anemic service facade | Domain Facade (DDD standard per [Fowler](https://martinfowler.com/eaaCatalog/serviceLayer.html)) |

**Verdict:** Proto in the service layer is an Anemic Domain Model anti-pattern. The current design — proto types stay in `internal/transport/grpc`, service speaks `pet.Pet` — is correct. This is not a trade-off; it is the intended direction of the Dependency Rule.

### Pros/Cons: separate server/grpc layer vs collapsed

| | Collapsed (`PetService` embeds `UnimplementedPetServiceServer`) | Separate (current) |
|---|---|---|
| **Setup cost** | One struct, one file | Two structs, explicit wiring |
| **Shape validation** | Lives next to business rules — mixed concern | `validateShape` isolated in transport |
| **Business rule tests** | Require proto request construction | Plain Go args, no proto dependency |
| **Code review** | All concerns in one method — hard to audit what is transport vs application | Clear boundary, each layer reviewable independently |
| **Future transports** | Forced extraction when a second transport appears | Zero refactor cost |
| **Clean Architecture** | Collapses "Interface Adapters" and "Use Cases" rings | Preserves the Dependency Rule |

**Verdict:** The separate layer is worth the wiring cost. The petstore example is explicitly teaching the layering; collapsing it would undermine that goal. For a service with no business rules and a single transport forever, collapsing is acceptable — that caveat is now documented in `process.md`.

---

## H — "Business rules" vs "Use cases" for the application layer

**Your comment:** "Should we call this use cases? I'm torn, mostly I find myself needing to validate here, but really it is the application of the business use case. Maybe we should be aiming for a tighter association."

### What the architecture literature says

Clean Architecture ([Uncle Bob](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)) names the rings:

1. **Entities** — *Enterprise* business rules (reusable across apps, e.g. `pet.NormalizePetName`)
2. **Use Cases** — *Application-specific* business rules (e.g. display_name required on Create for *this* app)
3. Interface Adapters
4. Frameworks & Drivers

The word "business rules" in Clean Architecture maps to **Entities**, not Use Cases. What `PetService` is doing — coordinating input normalisation, enforcing application-specific constraints, calling the Store — is exactly what Uncle Bob calls a **Use Case**.

Martin Fowler's [Service Layer](https://martinfowler.com/eaaCatalog/serviceLayer.html) reaches the same conclusion through DDD:

> "Every public method on an Application Service corresponds directly to a specific user action or application use case."

`CreatePet`, `GetPet`, etc. are use cases. `validateCreatePetBusinessRules` is the use case's enforcement logic — it is *applying the use case*, not validating in a generic sense.

### Why you feel the tension

The validation that happens in `validate*BusinessRules` is genuinely dual-natured:

- **Normalization** (`NormalizePetName`, `strings.TrimSpace`) — arguably a transport concern that slipped in
- **Application rule** (`display_name required after trim`) — a true use-case constraint

The `validate` prefix emphasises the technical act (checking); the `BusinessRules` suffix tries to name the policy layer. Neither quite captures "this is the enforcement step of the CreatePet use case."

### Position

Your instinct is architecturally grounded. The section header in `process.md` ("Business rules — application") uses Clean Architecture vocabulary incorrectly: by Uncle Bob's definition, what lives there is **use-case logic**, not entity-level business rules.

**Recommended rename for the section header:**

> ### 2. Use-case logic — application-specific rules

This aligns with Clean Architecture vocabulary, matches Fowler's Service Layer framing, and does not require renaming `validate*BusinessRules` in code now (that debate is tracked separately under item A).

The functions themselves can still be called `validate*BusinessRules` for now — that is a code-naming choice, not a documentation one. The doc should use the correct architectural vocabulary even if the code does not yet.

**Action:** Update `process.md` section heading and overview bullet from "Business rules" → "Use-case logic" in a follow-up commit.

---

## Summary of positions

| Item | Position | Action |
|------|----------|--------|
| **F** | Keep how-to. Add "verify" bullet per step to make the contract explicit. | Follow-up: harden each step with a verifiable assertion |
| **G (proto)** | Proto in service violates the Dependency Rule. Current design is correct. | No change — document stands |
| **G (layers)** | Separate server/service layer is worth the wiring cost. Collapse is documented as non-preferred. | Already addressed in process.md (item D) |
| **H** | "Use-case logic" is the architecturally correct term. Rename doc section, not code yet. | Follow-up: rename section heading in process.md |

## References

- [The Clean Architecture — Robert C. Martin](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Hexagonal Architecture — Alistair Cockburn](https://alistair.cockburn.us/hexagonal-architecture/)
- [Service Layer — Martin Fowler](https://martinfowler.com/eaaCatalog/serviceLayer.html)
- [Introducing Clean Architecture in Go — Three Dots Tech](https://threedots.tech/post/introducing-clean-architecture/)
- [Diátaxis: How-to guides](https://diataxis.fr/how-to-guides/)
- [RFC 2119](https://www.rfc-editor.org/rfc/rfc2119)
