package pet

// Pet is the resource entity managed by PetService.
type Pet struct {
	Name        string // pets/{pet} — output only / Update identifier; ignored on Create (AIP-133)
	DisplayName string
	Etag        string // server-computed; optional freshness token on Update/Delete (AIP-154)
}

// CreatePetRequest is AIP-133 shaped: user-specified pet_id + resource body.
// Top-level collection (no parent). Pet.Name on the body is ignored on create.
type CreatePetRequest struct {
	PetID string // becomes pets/{pet_id}
	Pet   Pet
}

// UpdatePetRequest is AIP-134 shaped: the resource carries its own name (+ optional etag).
// update_mask and allow_missing are out of scope for this example.
type UpdatePetRequest struct {
	Pet Pet
}

// ListPetsRequest is AIP-132 shaped for a top-level collection.
// filter, order_by, show_deleted, and parent are out of scope.
type ListPetsRequest struct {
	// PageSize is the maximum number of pets to return. Defaults to 50; capped at 100.
	// Negative values are invalid.
	PageSize int32
	// PageToken from a previous ListPets response retrieves the next page.
	PageToken string
}

// ListPetsResponse is AIP-132 shaped.
type ListPetsResponse struct {
	Pets          []Pet
	NextPageToken string // empty means no further pages
}
