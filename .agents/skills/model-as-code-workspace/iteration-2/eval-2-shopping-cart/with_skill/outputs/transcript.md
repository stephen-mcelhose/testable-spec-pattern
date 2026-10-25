# Shopping Cart Model Transcript

## 1. Subjects Enumerated

| Kind | Name | Rationale |
|---|---|---|
| **Actor** | `Customer` | Initiates the checkout flow. |
| **Object** | `Cart` | Container for items being purchased. |
| **Object** | `Item` | Entity added to the cart. |
| **Object** | `PaymentInfo` | Credentials/method used for payment. |
| **Record** | `PaymentReceipt` | Proof of payment processing. |
| **Outcome** | `OrderConfirmation` | The final result of a successful checkout. |

## 2. Top-Level Operations

| Operation | Inputs | Outputs | Description |
|---|---|---|---|
| `AddToCart` | `Customer`, `Cart`, `Item` | `Cart`, `error` | Adds an item to a cart. |
| `Checkout` | `Customer`, `Cart`, `PaymentInfo` | `OrderConfirmation`, `error` | Orchestrates the checkout process. |
| `ProcessPayment` | `Customer`, `Cart`, `PaymentInfo` | `PaymentReceipt`, `error` | Leaf: Handles payment transaction. |
| `ConfirmOrder` | `Customer`, `Cart`, `PaymentReceipt` | `OrderConfirmation`, `error` | Leaf: Finalizes the order. |

## 3. Undefined

The following stayed as `Undefined` (via interfaces):
- The internal structure of `Customer`, `Item`, `Cart`, `PaymentInfo`, `PaymentReceipt`, and `OrderConfirmation`.
- These are treated as opaque domain entities until their internal fields are required for business logic orchestration.

## 4. Non-Obvious Decisions

- **Cart Immutability (Sketch):** `AddToCart` returns a `Cart` and an `error`, following a functional/immutable style which is common in domain modeling to avoid side-effect confusion, although the implementation is currently a leaf.
- **PaymentReceipt as a Bridge:** Introduced `PaymentReceipt` to connect `ProcessPayment` and `ConfirmOrder`. This makes the data flow visible: you can't confirm an order without proof of payment.
- **Checkout as Composition:** `Checkout` is explicitly modeled as a composition of `ProcessPayment` and `ConfirmOrder` to show the workflow requested in the prompt.
