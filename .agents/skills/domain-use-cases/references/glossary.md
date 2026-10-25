# Glossary

Reference for terms used in the domain-use-cases skill.

---

## Actor

A named person or system that initiates or receives domain events. Actors are specific roles, not generic system labels.

✅ `Customer`, `StoreStaff`, `Vet`, `PetOwner`
❌ `User`, `Person`, `Client` — too generic to name a domain role
❌ `OrderService`, `UserRepository` — these are implementation concepts, not actors

---

## Domain Object

A thing that holds its own state and has a lifecycle you can track. It exists independently and can be identified over time.

✅ `Pet`, `Order`, `Appointment`
❌ `CreateOrderRequest`, `GetPetResponse` — wire shapes produced when designing contracts, not domain objects
❌ `OrderService`, `PetRepository` — these coordinate or store; they do not have a lifecycle of their own

Two subtypes — use `entity` or `value-object` as the tag when you can tell them apart:

**Entity** — has identity; the same concept persists through state changes (`Order` placed → approved → delivered is still the same `Order`).

**Value object** — has no identity; defined entirely by its attributes; immutable (`Category`, `Tag`, `Money`). Two `Category` objects with the same name are the same thing.

When uncertain, use `entity` and refine later.

---

## Domain Event

Something that happened in the domain. Named in past tense. An immutable fact — it cannot be undone, only compensated.

✅ `OrderPlaced`, `PetAdopted`, `OrderDelivered`
❌ `PlaceOrder`, `AdoptPet` — these are actions, not events
❌ `OrderCreated` — prefer business language: `OrderPlaced` not `OrderCreated`; `PetAdopted` not `PetSold`

---

## Action

Something an actor does that causes a domain event. Named in imperative form. An action is the intent; the event is the result.

✅ `PlaceOrder`, `AdoptPet`, `UpdatePetDetails`
❌ `OrderPlaced` — that is a domain event, not an action
❌ `CreateOrder` — prefer domain language: `PlaceOrder` not `CreateOrder`

---

## Lifecycle State

A named state a domain object can be in at a point in time. States are named in domain language, not technical shorthand.

✅ `Pet`: `available`, `pending`, `adopted`
✅ `Order`: `placed`, `approved`, `delivered`
❌ `ACTIVE`, `INACTIVE`, `DONE` — generic technical labels that carry no domain meaning

Lifecycle states are the basis for scenario names in behavioral specs.

---

## Ubiquitous Language

The shared vocabulary of the domain — used consistently by everyone: business stakeholders, developers, and tests. A concept document is a definition in the ubiquitous language. If a word means different things to different people on the team, it is not yet part of the ubiquitous language — clarify it first.
