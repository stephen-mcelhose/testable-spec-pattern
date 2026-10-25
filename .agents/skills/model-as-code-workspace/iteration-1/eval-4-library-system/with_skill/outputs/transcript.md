# Transcript: Library System Model

## Top-Level Operations
The following top-level operations were named first:
- `CheckoutBook(patron Patron, book Book) (Loan, error)`
- `ReturnBook(loan Loan) error`
- `ReserveBook(patron Patron, book Book) (Reservation, error)`

## Subjects Identified
The following subjects were extracted from the prompt:
- `Book`
- `Patron`
- `Loan` (distinct from Book, representing the checkout record)
- `Reservation`
- `Fine`

## Undefined Types
All subjects stayed as `interface{ Undefined }` because their internal fields were not required to define the relationships or the top-level flow. This keeps the model focused on the "what" rather than the "how".

## Non-Obvious Decisions
- **`IsOverdue` Predicate**: Per the skill instructions, this returns a `bool` rather than an error or a complex type.
- **Composition of `ReturnBook`**: The return process explicitly checks if the loan is overdue and triggers fine calculation before marking the book as returned.
- **Separation of Loan and Book**: `Loan` is a separate type that tracks the relationship between a `Patron` and a `Book`, ensuring the `Book` type remains a simple representation of the item itself.
- **No Strategy Types**: No external skills were named in the prompt, so no Strategy types were introduced.
- **Error Handling**: `ErrUnimplemented` is used for all leaf functions to signal where implementation is required.
