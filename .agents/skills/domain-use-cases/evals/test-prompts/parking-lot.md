# Eval: Parking Lot

**Prompt to run against the skill:**

> We're building a parking lot management system. I want to figure out the domain concepts before we write any code or design any API. Can you interview me to discover the actors, what objects have lifecycle, and what the key events are?

**What good output looks like:**

- Actors include at minimum: Driver (or parker), Attendant or System (if automated)
- Domain objects include: ParkingSpace (with availability states), Ticket or Session (tracking entry/exit), Vehicle
- Domain events include at minimum: VehicleEntered, VehicleExited, TicketIssued, PaymentCollected
- Lifecycle states are in domain language: `available`, `occupied` for space; `open`, `closed` or `active`, `settled` for a session
- No HTTP, endpoint, or database language in the main flow steps
- Concept documents produced for each extracted concept with OKF frontmatter and correct kind tags

**Known hard parts:**

- Differentiating a ParkingSpace (a physical slot) from a Ticket/Session (the occupancy event) — these are two separate domain objects with different lifecycles
- Determining whether payment is part of the session or a separate concept
- Clarifying whether "the system" is an actor or infrastructure
