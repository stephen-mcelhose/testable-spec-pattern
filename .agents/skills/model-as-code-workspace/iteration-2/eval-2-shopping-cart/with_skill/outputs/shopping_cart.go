package shoppingcart

import "errors"

// Undefined is the explicit placeholder for anything not yet modeled.
type Undefined = any

// ErrUnimplemented is returned by every leaf. Never panic in a model.
var ErrUnimplemented = errors.New("unimplemented")

// Subjects

type Customer interface{ Undefined }
type Item interface{ Undefined }
type Cart interface{ Undefined }
type PaymentInfo interface{ Undefined }
type PaymentReceipt interface{ Undefined }
type OrderConfirmation interface{ Undefined }

// Operations

// AddToCart adds an item to the customer's cart.
func AddToCart(customer Customer, cart Cart, item Item) (Cart, error) {
	return nil, ErrUnimplemented
}

// Checkout initiates the checkout process for a cart.
func Checkout(customer Customer, cart Cart, payment PaymentInfo) (OrderConfirmation, error) {
	// Composed operation:
	// 1. Process payment
	// 2. Confirm order
	receipt, err := ProcessPayment(customer, cart, payment)
	if err != nil {
		return nil, err
	}
	return ConfirmOrder(customer, cart, receipt)
}

// ProcessPayment handles the payment for the items in the cart.
func ProcessPayment(customer Customer, cart Cart, payment PaymentInfo) (PaymentReceipt, error) {
	// Leaf: Payment processing logic is an external detail.
	return nil, ErrUnimplemented
}

// ConfirmOrder finalizes the order and provides confirmation.
func ConfirmOrder(customer Customer, cart Cart, receipt PaymentReceipt) (OrderConfirmation, error) {
	// Leaf: Order persistence and confirmation generation.
	return nil, ErrUnimplemented
}
