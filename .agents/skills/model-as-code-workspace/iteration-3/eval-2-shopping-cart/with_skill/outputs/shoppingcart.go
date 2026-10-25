package shoppingcart

import "errors"

// Undefined is the explicit placeholder for anything not yet modeled.
type Undefined = any

// ErrUnimplemented is returned by every leaf. Never panic in a model.
var ErrUnimplemented = errors.New("unimplemented")

// Subjects

type Customer interface{ Undefined }
type Item     interface{ Undefined }
type Cart     interface{ Undefined }
type Payment  interface{ Undefined }
type Order     interface{ Undefined }

// Top-level operations

// AddItem adds an item to the shopping cart.
func AddItem(customer Customer, cart Cart, item Item) (Cart, error) {
	return addItem(customer, cart, item)
}

// Checkout initiates the checkout process for a cart.
func Checkout(customer Customer, cart Cart) (Payment, error) {
	return checkout(customer, cart)
}

// ProcessPayment processes the payment for an order.
func ProcessPayment(payment Payment) (Order, error) {
	return processPayment(payment)
}

// ConfirmOrder confirms the final order.
func ConfirmOrder(order Order) (Order, error) {
	// Note: Confirmation is the outcome, modeling it as a state transition or a final check.
	return confirmOrder(order)
}

// Decompositions

func addItem(customer Customer, cart Cart, item Item) (Cart, error) {
	// Leaf: logic for adding an item to a cart.
	return nil, ErrUnimplemented
}

func checkout(customer Customer, cart Cart) (Payment, error) {
	// Leaf: logic for starting checkout and generating a payment request.
	return nil, ErrUnimplemented
}

func processPayment(payment Payment) (Order, error) {
	// Leaf: logic for interacting with a payment provider.
	return nil, ErrUnimplemented
}

func confirmOrder(order Order) (Order, error) {
	// Leaf: logic for final order confirmation.
	return nil, ErrUnimplemented
}
