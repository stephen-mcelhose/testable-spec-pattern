package functional

import (
	"context"
	"errors"
	"fmt"

	pet "example.com/go-behavioral-spec/examples/01-pure-function"
)

// Pet is the entity returned by a PetGetter leaf (may be loaded via I/O in production).
type Pet struct {
	Name        string
	DisplayName string
}

// PetView is the pure projection BuildPetView returns to callers.
type PetView struct {
	Name        string
	DisplayName string
}

// PetGetter loads a Pet by resource name. Production implementations may perform I/O;
// tests inject a stub (or a counterfeiter fake when call-count/arg capture is needed).
type PetGetter interface {
	Get(ctx context.Context, name string) (Pet, error)
}

// BuildPetView normalizes the pet name, loads the Pet via getter, and returns a view.
// Dependency is passed as a function argument (functional injection).
func BuildPetView(ctx context.Context, name string, getter PetGetter) (PetView, error) {
	canonical, err := pet.NormalizePetName(name)
	if err != nil {
		return PetView{}, err
	}

	p, err := getter.Get(ctx, canonical)
	if err != nil {
		var nf *pet.NotFoundError
		if errors.As(err, &nf) {
			return PetView{}, err
		}
		return PetView{}, &pet.InternalError{Message: fmt.Sprintf("get pet: %v", err)}
	}

	return PetView{
		Name:        p.Name,
		DisplayName: p.DisplayName,
	}, nil
}
