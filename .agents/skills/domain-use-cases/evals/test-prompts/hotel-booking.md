# Eval: Hotel Booking

**Prompt to run against the skill:**

> We're designing a hotel booking system. Before we touch any schema or API, I want to map out the domain. Can you write use-cases covering the guest experience and the front desk side, then extract the domain concepts?

**What good output looks like:**

- Actors include at minimum: Guest, FrontDesk (or Receptionist), Housekeeper (may emerge from check-out flow)
- Domain objects include: Reservation, Room, Stay (the in-progress occupancy — distinct from Reservation)
- Domain events include: ReservationMade, CheckedIn, CheckedOut, RoomAssigned, ReservationCancelled
- Lifecycle states: `available`, `occupied`, `cleaning` for Room; `pending`, `confirmed`, `cancelled` for Reservation; `active`, `completed` for Stay
- No HTTP or database language in main flow steps
- The Reservation/Stay distinction is surfaced: a Reservation is a future commitment; a Stay is the active occupancy — two separate domain objects
- Concept documents produced with OKF frontmatter and correct kind tags

**Known hard parts:**

- Reservation vs Stay: the skill should ask whether there's a separate concept for the period when the guest is actually in the room vs. when they've just booked
- Room assignment timing: is a room assigned at reservation or at check-in? This changes the domain model significantly
- Cancellation policy: cancellation may have its own lifecycle (grace period, penalty) — does it warrant its own concept?
