# Petstore — Domain Use-Cases

Worked example for the `domain-use-cases` skill. Domain concepts discovered through a client interview; use-cases written from that interview without reference to any implementation.

---

## Use-Cases

---

### Browse Available Pets

```
Use Case: Browse Available Pets
Actor:    Customer
Goal:     Customer sees a list of pets that are available for adoption

Main flow:
  1. Customer requests the list of available pets
  2. Customer optionally filters by category or tag
  3. Store returns matching pets in available state
```

---

### Adopt a Pet

```
Use Case: Adopt a Pet
Actor:    Customer
Goal:     An order is placed and the pet is reserved for that customer

Main flow:
  1. Customer selects an available pet
  2. Customer places an adoption order for the pet
  3. Pet state changes to pending
  4. Order is confirmed as placed
```

---

### Approve an Adoption Order

```
Use Case: Approve an Adoption Order
Actor:    StoreStaff
Goal:     Order is approved and handoff is arranged

Main flow:
  1. StoreStaff reviews a placed order
  2. StoreStaff approves the order
  3. Order state changes to approved
  4. Handoff (delivery or pickup) is arranged with the customer
```

---

### Complete a Delivery

```
Use Case: Complete a Delivery
Actor:    StoreStaff
Goal:     Pet is adopted and the order is marked delivered

Main flow:
  1. Pet is handed off to the customer
  2. StoreStaff marks the order as delivered
  3. Order state changes to delivered
  4. Pet state changes to adopted
```

---

### List a Pet for Adoption

```
Use Case: List a Pet for Adoption
Actor:    StoreStaff
Goal:     Pet is listed and visible to customers as available

Main flow:
  1. StoreStaff creates a pet record with name, category, and tags
  2. StoreStaff uploads one or more photos
  3. StoreStaff marks the pet as available
  4. Pet appears in the available listing
```

---

### Update Pet Details

```
Use Case: Update Pet Details
Actor:    StoreStaff
Goal:     Pet record reflects the latest accurate information

Main flow:
  1. StoreStaff selects a pet to update
  2. StoreStaff updates one or more fields (name, category, tags, photos)
  3. Pet record is updated
```

---

### Register a Customer Account

```
Use Case: Register a Customer Account
Actor:    Customer
Goal:     Customer has an account and can place adoption orders

Main flow:
  1. Customer provides a username and contact details
  2. Account is created for the customer
  3. Customer can log in and browse pets
```

---

### Onboard Store Staff

```
Use Case: Onboard Store Staff
Actor:    StoreStaff
Goal:     One or more staff accounts are created ready for use

Main flow:
  1. StoreStaff submits a list of new staff members with their details
  2. Accounts are created for each staff member
  3. Staff can log in and manage pets and orders
```

---

## Extracted Concepts

---

### Actors

---

```markdown
---
type: concept
title: "Customer"
description: A person who browses available pets and places adoption orders.
tags: [actor]
timestamp: 2026-10-25T03:14:00Z
---

## Customer

**Kind:** Actor

**Definition:** A person who browses available pets and places adoption orders.
```

---

```markdown
---
type: concept
title: "StoreStaff"
description: A member of the petstore team who manages pet listings and approves adoption orders.
tags: [actor]
timestamp: 2026-10-25T03:14:00Z
---

## StoreStaff

**Kind:** Actor

**Definition:** A member of the petstore team who manages pet listings, approves adoption orders, and completes deliveries.
```

---

### Domain Objects

---

```markdown
---
type: concept
title: "Pet"
description: An animal listed in the store with a name, category, and availability state.
tags: [entity, aggregate-root]
timestamp: 2026-10-25T03:14:00Z
---

## Pet

**Kind:** Domain Object

**Definition:** An animal listed in the store that progresses through availability states as it moves toward adoption.

**Lifecycle states:** available, pending, adopted
```

---

```markdown
---
type: concept
title: "Order"
description: A customer's request to adopt a specific pet, tracked through approval and delivery.
tags: [entity]
timestamp: 2026-10-25T03:14:00Z
---

## Order

**Kind:** Domain Object

**Definition:** A customer's request to adopt a specific pet, tracked from placement through approval and delivery.

**Lifecycle states:** placed, approved, delivered
```

---

```markdown
---
type: concept
title: "Category"
description: A broad classification of a pet's type, such as dog, cat, or rabbit — fixed at listing time.
tags: [value-object]
timestamp: 2026-10-25T03:14:00Z
---

## Category

**Kind:** Domain Object

**Definition:** A broad classification of a pet's type — fixed at listing time and shared across many pets.
```

---

```markdown
---
type: concept
title: "Tag"
description: A descriptive label applied to a pet, such as a breed, characteristic, or care note.
tags: [value-object]
timestamp: 2026-10-25T03:14:00Z
---

## Tag

**Kind:** Domain Object

**Definition:** A descriptive label applied to a pet to support browsing and filtering. A pet may have many tags.
```

---

### Domain Events

---

```markdown
---
type: concept
title: "PetListedForAdoption"
description: A pet has been added to the store and marked as available for adoption.
tags: [domain-event]
timestamp: 2026-10-25T03:14:00Z
---

## PetListedForAdoption

**Kind:** Domain Event

**Definition:** A pet has been added to the store and marked as available for adoption by StoreStaff.
```

---

```markdown
---
type: concept
title: "OrderPlaced"
description: A customer has submitted an adoption order for an available pet.
tags: [domain-event]
timestamp: 2026-10-25T03:14:00Z
---

## OrderPlaced

**Kind:** Domain Event

**Definition:** A customer has submitted an adoption order for an available pet, reserving it as pending.
```

---

```markdown
---
type: concept
title: "OrderApproved"
description: StoreStaff have reviewed and approved a placed adoption order.
tags: [domain-event]
timestamp: 2026-10-25T03:14:00Z
---

## OrderApproved

**Kind:** Domain Event

**Definition:** StoreStaff have reviewed and approved a placed adoption order, and handoff is being arranged.
```

---

```markdown
---
type: concept
title: "OrderDelivered"
description: The pet has been handed off to the customer and the order is complete.
tags: [domain-event]
timestamp: 2026-10-25T03:14:00Z
---

## OrderDelivered

**Kind:** Domain Event

**Definition:** The pet has been handed off to the customer and the order is marked complete.
```

---

```markdown
---
type: concept
title: "PetAdopted"
description: A pet has been successfully adopted and is no longer available in the store.
tags: [domain-event]
timestamp: 2026-10-25T03:14:00Z
---

## PetAdopted

**Kind:** Domain Event

**Definition:** A pet has been successfully adopted and its state is permanently changed to adopted.
```

---

### Actions

---

```markdown
---
type: concept
title: "PlaceAdoptionOrder"
description: A customer submits an adoption order for an available pet.
tags: [action]
timestamp: 2026-10-25T03:14:00Z
---

## PlaceAdoptionOrder

**Kind:** Action

**Definition:** A customer submits an adoption order for an available pet, triggering OrderPlaced and moving the pet to pending.
```

---

```markdown
---
type: concept
title: "ApproveOrder"
description: StoreStaff approves a placed adoption order, triggering arrangement of the handoff.
tags: [action]
timestamp: 2026-10-25T03:14:00Z
---

## ApproveOrder

**Kind:** Action

**Definition:** StoreStaff approves a placed adoption order, triggering OrderApproved and arranging the handoff.
```

---

```markdown
---
type: concept
title: "CompleteDelivery"
description: StoreStaff marks an approved order as delivered after the pet is handed off to the customer.
tags: [action]
timestamp: 2026-10-25T03:14:00Z
---

## CompleteDelivery

**Kind:** Action

**Definition:** StoreStaff marks an approved order as delivered, triggering OrderDelivered and PetAdopted.
```

---

```markdown
---
type: concept
title: "ListPetForAdoption"
description: StoreStaff creates a pet record and marks it available so customers can browse and adopt it.
tags: [action]
timestamp: 2026-10-25T03:14:00Z
---

## ListPetForAdoption

**Kind:** Action

**Definition:** StoreStaff creates a pet record with name, category, tags, and photos, then marks it available.
```

---

```markdown
---
type: concept
title: "UploadPetPhoto"
description: StoreStaff attaches a photo to a pet listing.
tags: [action]
timestamp: 2026-10-25T03:14:00Z
---

## UploadPetPhoto

**Kind:** Action

**Definition:** StoreStaff attaches one or more photos to a pet listing to support customer browsing.
```
