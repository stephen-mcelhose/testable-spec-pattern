package librarysystem

import "errors"

// Undefined is the explicit placeholder for anything not yet modeled.
type Undefined = any

// ErrUnimplemented is returned by every leaf. Never panic in a model.
var ErrUnimplemented = errors.New("unimplemented")

// Subjects

// Book represents a physical item in the library.
type Book interface{ Undefined }

// Patron is a person who uses the library services.
type Patron interface{ Undefined }

// Loan represents the record of a patron checking out a book.
type Loan interface{ Undefined }

// Reservation represents a patron's request to hold a book.
type Reservation interface{ Undefined }

// Fine represents an overdue penalty.
type Fine interface{ Undefined }

// Top-level operations

// CheckoutBook records that a patron is taking a book.
func CheckoutBook(patron Patron, book Book) (Loan, error) {
	if !IsAvailable(book) {
		return nil, errors.New("book not available")
	}
	return RecordLoan(patron, book)
}

// ReturnBook records that a patron is returning a book.
func ReturnBook(loan Loan) error {
	if IsOverdue(loan) {
		_, err := CalculateOverdueFine(loan)
		if err != nil {
			return err
		}
	}
	return MarkReturned(loan)
}

// ReserveBook places a hold on a book for a patron.
func ReserveBook(patron Patron, book Book) (Reservation, error) {
	return CreateReservation(patron, book)
}

// Predicates

// IsAvailable checks if a book is currently in the library and not reserved.
func IsAvailable(book Book) bool {
	// Leaf: logic to check availability from inventory is not yet modeled.
	return false
}

// IsOverdue checks if a loan has exceeded its due date.
func IsOverdue(loan Loan) bool {
	// Leaf: date comparison logic is not yet modeled.
	return false
}

// Leaf functions

// RecordLoan creates the checkout record.
func RecordLoan(patron Patron, book Book) (Loan, error) {
	// Leaf: persistence of loan records is not yet modeled.
	return nil, ErrUnimplemented
}

// MarkReturned marks the loan as completed and the book as back in inventory.
func MarkReturned(loan Loan) error {
	// Leaf: updates to loan and book status are not yet modeled.
	return ErrUnimplemented
}

// CreateReservation records a patron's interest in a book.
func CreateReservation(patron Patron, book Book) (Reservation, error) {
	// Leaf: persistence of reservations is not yet modeled.
	return nil, ErrUnimplemented
}

// CalculateOverdueFine determines the fine amount for a late return.
func CalculateOverdueFine(loan Loan) (Fine, error) {
	// Leaf: fine calculation logic based on dates and rates is not yet modeled.
	return nil, ErrUnimplemented
}
