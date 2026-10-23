package petstore

// Pet is the resource entity returned by PetService methods.
type Pet struct {
	Name        string // pets/{pet} — output / Update identifier; ignored on Create (AIP-133)
	DisplayName string
	Etag        string // server-computed; optional freshness on Update/Delete (AIP-154)
}

// CreatePetRequest is AIP-133 shaped: user-specified pet_id + resource body.
// Top-level collection (no parent). Pet.Name on the body is ignored.
type CreatePetRequest struct {
	PetID string // becomes pets/{pet_id}
	Pet   Pet
}

// UpdatePetRequest is AIP-134 shaped: resource carries name (+ optional etag).
// update_mask / allow_missing are out of scope for this teaching example.
type UpdatePetRequest struct {
	Pet Pet
}

// ListPetsRequest is AIP-132 shaped for a top-level collection (no parent).
// page_size / page_token follow AIP-158. filter, order_by, show_deleted are out of scope.
type ListPetsRequest struct {
	// PageSize is the maximum number of pets to return. The service may return fewer.
	// If 0/unspecified, defaults to 50. Values above 100 are coerced to 100.
	// Negative values are invalid.
	PageSize int32
	// PageToken from a previous ListPets response retrieves the next page.
	PageToken string
}

// ListPetsResponse is AIP-132 shaped: resources + next_page_token.
type ListPetsResponse struct {
	Pets          []Pet
	NextPageToken string // unset/empty means no subsequent pages
}
