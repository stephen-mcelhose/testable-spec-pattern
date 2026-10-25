# Glossary

Reference for terms used in the domain-use-cases skill.

---

## Actor

A named person or system that initiates actions in the domain. Actors are specific roles, not generic system labels.

The test: an actor is the *who* in a use-case — the one pursuing a goal. If you cannot write "Actor: X" at the top of a use-case and have X do something meaningful in the main flow, it is not an actor.

External systems can be actors when they actively initiate something (`PaymentGateway` triggering a `PaymentFailed` event). Passive consumers — systems that only receive or forward messages without pursuing a goal — are not actors; they are infrastructure.

✅ `Customer`, `StoreStaff`, `Vet`, `Driver`
❌ `User`, `Person`, `Client` — too generic to name a domain role
❌ `OrderService`, `UserRepository` — these coordinate or store; they do not pursue goals
❌ `MessageQueue`, `NotificationService` — passive consumers, not actors

---

## Domain Object

A named thing in the domain that is either tracked over time (entity) or defined purely by its attributes (value object). The key question: does the domain care about *which one* it is, or only *what it contains*?

✅ `Pet`, `Order`, `Appointment`, `Money`, `Address`
❌ `CreateOrderRequest`, `GetPetResponse` — wire shapes produced when designing contracts, not domain objects
❌ `OrderService`, `PetRepository` — these coordinate or store; they are not domain things with meaning of their own

Two subtypes — use `entity` or `value-object` as the tag when you can tell them apart:

**Entity** — the domain cares *which one it is*. The same entity persists through state changes: an `Order` placed → approved → delivered is still the same `Order`. Has identity. Has lifecycle states.

**Value object** — the domain only cares *what it contains*. Immutable; defined entirely by its attributes. Two `Money` objects both representing £10 are the same thing — there is no "which one". No identity, no lifecycle. Examples: `Money`, `Address`, `DateRange`.

`Category` and `Tag` feel like value objects in many domains, but become entities if the business needs to rename, merge, or deactivate them. When uncertain, use `entity` and refine later.

---

## Domain Event

Something that happened in the domain. Named in past tense. An immutable fact — it cannot be undone, only compensated.

✅ `OrderPlaced`, `PetAdopted`, `OrderDelivered`
❌ `PlaceOrder`, `AdoptPet` — these are actions, not events
❌ `OrderCreated` — prefer business language: `OrderPlaced` not `OrderCreated`; `PetAdopted` not `PetSold`

---

## Action

Something an actor does that causes a domain event. Named in imperative form. An action is the intent; the event is the result.

✅ `PlaceOrder`, `AdoptPet`, `ApproveApplication`, `CancelBooking`
❌ `OrderPlaced` — that is a domain event, not an action
❌ `CreateOrder` — prefer domain language: `PlaceOrder` not `CreateOrder`
❌ `UpdatePetDetails` — too close to CRUD; prefer a name that says *why*: `CorrectPetRecord`, `AddPetPhoto`

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
