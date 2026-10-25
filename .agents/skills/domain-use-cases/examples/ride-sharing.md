# Ride Sharing — Domain Use-Cases

Worked example for the `domain-use-cases` skill. Domain concepts discovered through a client interview. No implementation consulted.

---

## Use-Cases

---

### Request a Ride

```
Use Case: Request a Ride
Actor:    Rider
Goal:     A trip is created and a driver is being matched

Main flow:
  1. Rider enters a pickup location and destination
  2. System estimates the fare including any surge multiplier
  3. Rider confirms the request
  4. Trip is created in requested state
  5. Matching begins — nearby online drivers are considered
```

---

### Accept a Trip Offer

```
Use Case: Accept a Trip Offer
Actor:    Driver
Goal:     Driver is assigned to the trip and heads to pickup

Main flow:
  1. Driver receives a trip offer with rider location and estimated fare
  2. Driver accepts the offer within the response window
  3. Trip state changes to accepted
  4. Driver navigates to the pickup location
  5. Trip state changes to en route
```

---

### Complete a Trip

```
Use Case: Complete a Trip
Actor:    Driver
Goal:     Trip is completed, fare is settled, and both parties can rate each other

Main flow:
  1. Rider enters the vehicle — trip state changes to in progress
  2. Driver navigates to the destination
  3. Driver marks the trip complete at the destination
  4. Actual fare is calculated from the real route
  5. Rider's payment method is charged
  6. Trip state changes to completed
  7. Rider rates the driver; driver rates the rider
```

---

### Cancel a Trip

```
Use Case: Cancel a Trip
Actor:    Rider or Driver
Goal:     Trip is cancelled before the ride begins

Main flow:
  1. Rider or Driver cancels before the trip enters in progress state
  2. Trip state changes to cancelled
  3. Cancellation is recorded against the cancelling party's account
```

---

### Go Online

```
Use Case: Go Online
Actor:    Driver
Goal:     Driver is available to receive trip offers

Main flow:
  1. Driver opens the app and starts a shift
  2. System checks that the driver account is active and a vehicle is approved
  3. Driver's availability changes to online
  4. Driver begins receiving trip offers
```

---

### Register a Vehicle

```
Use Case: Register a Vehicle
Actor:    Driver
Goal:     Vehicle is submitted for approval and can be activated once approved

Main flow:
  1. Driver submits vehicle details — make, model, year, plate number
  2. Vehicle record is created in pending state
  3. Platform reviews the vehicle submission
  4. Vehicle state changes to approved or rejected
```

---

### Receive a Weekly Payout

```
Use Case: Receive a Weekly Payout
Actor:    Driver
Goal:     Driver receives earnings for all completed trips in the payout period

Main flow:
  1. Payout period closes
  2. System calculates total earnings across all completed trips in the period
  3. Payout record is created in pending state
  4. Payout is processed to the driver's bank account
  5. Payout state changes to paid
```

---

## Extracted Concepts

---

### Actors

---

```markdown
---
type: concept
title: "Rider"
description: A person who requests trips to travel between locations.
tags: [actor]
timestamp: 2026-10-25T03:14:00Z
---

## Rider

**Kind:** Actor

**Definition:** A person who requests trips and pays the fare upon completion.
```

---

```markdown
---
type: concept
title: "Driver"
description: A person with an approved vehicle who accepts trip offers and earns a fare for completing them.
tags: [actor]
timestamp: 2026-10-25T03:14:00Z
---

## Driver

**Kind:** Actor

**Definition:** A person with an approved vehicle who accepts trip offers, transports riders, and earns a payout for completed trips.
```

---

### Domain Objects

---

```markdown
---
type: concept
title: "Trip"
description: A single journey from a pickup location to a destination, from the moment a rider requests until the fare is settled.
tags: [entity, aggregate-root]
timestamp: 2026-10-25T03:14:00Z
---

## Trip

**Kind:** Domain Object

**Definition:** A single journey between a pickup location and a destination, owned by a Rider and fulfilled by a Driver.

**Lifecycle states:** requested, matched, accepted, en route, in progress, completed, cancelled
```

---

```markdown
---
type: concept
title: "Vehicle"
description: A driver's registered car that must be approved before it can be used to accept trips.
tags: [entity]
timestamp: 2026-10-25T03:14:00Z
---

## Vehicle

**Kind:** Domain Object

**Definition:** A car registered by a Driver that must pass approval before it can be used on the platform.

**Lifecycle states:** pending, approved, rejected
```

---

```markdown
---
type: concept
title: "Payout"
description: A weekly settlement of a driver's earnings across all completed trips in the payout period.
tags: [entity]
timestamp: 2026-10-25T03:14:00Z
---

## Payout

**Kind:** Domain Object

**Definition:** A weekly record of a Driver's earnings for all completed trips in the period, settled to their bank account.

**Lifecycle states:** pending, processing, paid, failed
```

---

```markdown
---
type: concept
title: "DriverAvailability"
description: A driver's current willingness to receive trip offers, independent of their account status.
tags: [value-object]
timestamp: 2026-10-25T03:14:00Z
---

## DriverAvailability

**Kind:** Domain Object

**Definition:** Whether a Driver is currently online and able to receive trip offers. Independent of account status — a driver can be account-active but offline.

**Lifecycle states:** online, offline
```

---

### Domain Events

---

```markdown
---
type: concept
title: "TripRequested"
description: A rider has confirmed a ride request and a trip has been created awaiting driver matching.
tags: [domain-event]
timestamp: 2026-10-25T03:14:00Z
---

## TripRequested

**Kind:** Domain Event

**Definition:** A Rider has confirmed a ride request — a Trip has been created and driver matching has begun.
```

---

```markdown
---
type: concept
title: "TripAccepted"
description: A driver has accepted a trip offer and is heading to the pickup location.
tags: [domain-event]
timestamp: 2026-10-25T03:14:00Z
---

## TripAccepted

**Kind:** Domain Event

**Definition:** A Driver has accepted the trip offer and the Trip moves to en route state.
```

---

```markdown
---
type: concept
title: "TripCompleted"
description: The driver has delivered the rider to the destination and the fare has been settled.
tags: [domain-event]
timestamp: 2026-10-25T03:14:00Z
---

## TripCompleted

**Kind:** Domain Event

**Definition:** The Driver has delivered the Rider to the destination, the fare has been settled, and ratings are open.
```

---

```markdown
---
type: concept
title: "TripCancelled"
description: A trip was cancelled by either the rider or driver before the ride began.
tags: [domain-event]
timestamp: 2026-10-25T03:14:00Z
---

## TripCancelled

**Kind:** Domain Event

**Definition:** A Trip was cancelled by either the Rider or Driver before entering in progress state.
```

---

```markdown
---
type: concept
title: "VehicleApproved"
description: A driver's registered vehicle has passed platform review and can be used to accept trips.
tags: [domain-event]
timestamp: 2026-10-25T03:14:00Z
---

## VehicleApproved

**Kind:** Domain Event

**Definition:** A Driver's Vehicle has passed platform review and is now eligible for use on the platform.
```

---

```markdown
---
type: concept
title: "PayoutSettled"
description: A driver's weekly earnings have been successfully transferred to their bank account.
tags: [domain-event]
timestamp: 2026-10-25T03:14:00Z
---

## PayoutSettled

**Kind:** Domain Event

**Definition:** A Driver's weekly Payout has been successfully processed and transferred to their bank account.
```

---

### Actions

---

```markdown
---
type: concept
title: "RequestTrip"
description: A rider submits a ride request from a pickup location to a destination.
tags: [action]
timestamp: 2026-10-25T03:14:00Z
---

## RequestTrip

**Kind:** Action

**Definition:** A Rider submits a ride request specifying pickup and destination, triggering TripRequested and driver matching.
```

---

```markdown
---
type: concept
title: "AcceptTripOffer"
description: A driver accepts a trip offer within the response window, committing to collect the rider.
tags: [action]
timestamp: 2026-10-25T03:14:00Z
---

## AcceptTripOffer

**Kind:** Action

**Definition:** A Driver accepts a trip offer within the response window, triggering TripAccepted and beginning navigation to the pickup.
```

---

```markdown
---
type: concept
title: "CompleteTrip"
description: A driver marks a trip complete at the destination, triggering fare settlement and rating.
tags: [action]
timestamp: 2026-10-25T03:14:00Z
---

## CompleteTrip

**Kind:** Action

**Definition:** A Driver marks a Trip complete at the destination, triggering fare calculation, payment charge, and TripCompleted.
```

---

```markdown
---
type: concept
title: "CancelTrip"
description: A rider or driver cancels a trip before the ride begins.
tags: [action]
timestamp: 2026-10-25T03:14:00Z
---

## CancelTrip

**Kind:** Action

**Definition:** A Rider or Driver cancels a Trip before it enters in progress state, triggering TripCancelled and recording the cancellation.
```

---

```markdown
---
type: concept
title: "GoOnline"
description: A driver starts a shift and signals readiness to receive trip offers.
tags: [action]
timestamp: 2026-10-25T03:14:00Z
---

## GoOnline

**Kind:** Action

**Definition:** A Driver starts a shift, changing DriverAvailability to online and enabling trip offer delivery.
```
