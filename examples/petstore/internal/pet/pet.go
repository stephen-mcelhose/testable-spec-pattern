package pet

// Pet is the domain resource managed by PetService.
// Proto petstorev1.Pet is the wire DTO; map at the transport boundary.
type Pet struct {
	Name        string // pets/{pet}
	DisplayName string
	Etag        string // server-computed; optional freshness on Update/Delete (AIP-154)
}
