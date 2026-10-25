package vetclinic

import "errors"

// Undefined is the explicit placeholder for anything not yet modeled.
type Undefined = any

// ErrUnimplemented is returned by every leaf. Never panic in a model.
var ErrUnimplemented = errors.New("unimplemented")

// Subjects identified from the prompt.

// PetOwner is someone who owns a pet and interacts with the clinic.
type PetOwner interface{ Undefined }

// AppointmentRequest contains the details needed to request an appointment.
type AppointmentRequest interface{ Undefined }

// Appointment represents a scheduled time for a clinic visit.
type Appointment interface{ Undefined }

// Vet is a professional assigned to handle appointments.
type Vet interface{ Undefined }

// VisitRecord captures what happened during an appointment.
type VisitRecord interface{ Undefined }

// BillingRecord represents the financial bill produced from a visit.
type BillingRecord interface{ Undefined }

// Top-level operations

// BookAppointment initiates the booking process for a pet owner.
func BookAppointment(owner PetOwner, request AppointmentRequest) (Appointment, error) {
	appointment, err := createUnassignedAppointment(owner, request)
	if err != nil {
		return nil, err
	}
	return appointment, nil
}

// AssignVet assigns a specific vet to an existing appointment.
func AssignVet(appointment Appointment, vet Vet) (Appointment, error) {
	return assignVetToAppointment(appointment, vet)
}

// RecordVisit captures the details of the medical visit.
func RecordVisit(appointment Appointment) (VisitRecord, error) {
	return finalizeVisitRecord(appointment)
}

// ProduceBilling creates a bill based on the visit.
func ProduceBilling(visit VisitRecord) (BillingRecord, error) {
	return generateBilling(visit)
}

// Internal (composed) functions and leaves

func createUnassignedAppointment(owner PetOwner, request AppointmentRequest) (Appointment, error) {
	// Leaf: logic to initialize an appointment from the request.
	return nil, ErrUnimplemented
}

func assignVetToAppointment(appointment Appointment, vet Vet) (Appointment, error) {
	// Leaf: logic to attach a vet to an appointment.
	return nil, ErrUnimplemented
}

func finalizeVisitRecord(appointment Appointment) (VisitRecord, error) {
	// Leaf: logic to convert an appointment into a visit record.
	return nil, ErrUnimplemented
}

func generateBilling(visit VisitRecord) (BillingRecord, error) {
	// Leaf: logic to calculate and produce a billing record from a visit.
	return nil, ErrUnimplemented
}
