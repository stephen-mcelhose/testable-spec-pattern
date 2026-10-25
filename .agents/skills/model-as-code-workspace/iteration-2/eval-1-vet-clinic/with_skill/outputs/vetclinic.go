package vetclinic

import "errors"

// Primitives

// Undefined is the explicit placeholder for anything not yet modeled.
type Undefined = any

// ErrUnimplemented is returned by every leaf. Never panic in a model.
var ErrUnimplemented = errors.New("unimplemented")

// Subjects

// Actors
type PetOwner interface{ Undefined }
type Vet interface{ Undefined }

// Implied Actors (Surfaced as questions)
// - Is the Pet a separate entity from the PetOwner? (Assuming yes)
type Pet interface{ Undefined }

// Objects
type Appointment interface{ Undefined }

// Records
type VisitRecord interface{ Undefined }

// Outcomes
type BillingRecord interface{ Undefined }

// Top-level operations

// BookAppointment allows a pet owner to book an appointment for their pet.
func BookAppointment(owner PetOwner, pet Pet) (Appointment, error) {
	// Composed: Assigning a vet is part of booking in this system.
	vet, err := AssignVet(pet)
	if err != nil {
		return nil, err
	}

	return ConfirmAppointment(owner, pet, vet)
}

// RecordVisit documents the details of a completed appointment.
func RecordVisit(appointment Appointment) (VisitRecord, error) {
	// Composed: recording a visit produces a billing record in this system.
	record, err := CreateVisitRecord(appointment)
	if err != nil {
		return nil, err
	}

	_, err = ProduceBillingRecord(record)
	if err != nil {
		return nil, err
	}

	return record, nil
}

// Relations

// AssignVet assigns a veterinarian to a pet's case.
func AssignVet(pet Pet) (Vet, error) {
	// Leaf: Vet assignment logic not specified.
	return nil, ErrUnimplemented
}

// ConfirmAppointment finalizes the appointment booking.
func ConfirmAppointment(owner PetOwner, pet Pet, vet Vet) (Appointment, error) {
	// Leaf: Confirmation process not specified.
	return nil, ErrUnimplemented
}

// CreateVisitRecord generates the record of the visit.
func CreateVisitRecord(appointment Appointment) (VisitRecord, error) {
	// Leaf: Visit record details not specified.
	return nil, ErrUnimplemented
}

// ProduceBillingRecord generates the bill for the visit.
func ProduceBillingRecord(visit VisitRecord) (BillingRecord, error) {
	// Leaf: Billing logic not specified.
	return nil, ErrUnimplemented
}
