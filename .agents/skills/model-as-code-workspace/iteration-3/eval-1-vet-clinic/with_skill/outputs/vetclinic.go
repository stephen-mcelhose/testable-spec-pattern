package vetclinic

import "errors"

// Undefined is the explicit placeholder for anything not yet modeled.
type Undefined = any

// ErrUnimplemented is returned by every leaf. Never panic in a model.
var ErrUnimplemented = errors.New("unimplemented")

// --- Subjects ---

// PetOwner is the actor who initiates a booking.
type PetOwner interface{ Undefined }

// Appointment is the object being booked.
type Appointment interface{ Undefined }

// Vet is the actor assigned to an appointment and who records the visit.
type Vet interface{ Undefined }

// Visit is the record of an appointment being conducted.
type Visit interface{ Undefined }

// BillingRecord is the outcome produced after a visit is recorded.
type BillingRecord interface{ Undefined }

// --- Top-level Operations ---

// BookAppointment initiates the booking process for a pet owner.
func BookAppointment(owner PetOwner) (Appointment, error) {
	// Composed: books the appointment and then assigns a vet.
	apt, err := CreateBooking(owner)
	if err != nil {
		return nil, err
	}

	vet, err := AssignVet(apt)
	if err != nil {
		return nil, err
	}

	return ConfirmAppointment(apt, vet)
}

// RecordVisit records the details of a visit by a vet for an appointment.
func RecordVisit(vet Vet, apt Appointment) (Visit, error) {
	// Composed: records the visit and then produces the billing record.
	visit, err := CreateVisitRecord(vet, apt)
	if err != nil {
		return nil, err
	}

	_, err = ProduceBillingRecord(visit)
	if err != nil {
		return nil, err
	}

	return visit, nil
}

// --- Relations ---

// CreateBooking creates the initial appointment record.
func CreateBooking(owner PetOwner) (Appointment, error) {
	// Leaf: implementation of booking creation is not yet specified.
	return nil, ErrUnimplemented
}

// AssignVet assigns a vet to an appointment.
func AssignVet(apt Appointment) (Vet, error) {
	// Leaf: assignment logic (availability, specialization) is not yet modeled.
	return nil, ErrUnimplemented
}

// ConfirmAppointment finalizes the appointment with an assigned vet.
func ConfirmAppointment(apt Appointment, vet Vet) (Appointment, error) {
	// Leaf: confirmation process not yet modeled.
	return nil, ErrUnimplemented
}

// CreateVisitRecord captures the details of the visit.
func CreateVisitRecord(vet Vet, apt Appointment) (Visit, error) {
	// Leaf: visit details capture is not yet modeled.
	return nil, ErrUnimplemented
}

// ProduceBillingRecord generates a billing record from a visit.
func ProduceBillingRecord(visit Visit) (BillingRecord, error) {
	// Leaf: billing logic and pricing not yet modeled.
	return nil, ErrUnimplemented
}
