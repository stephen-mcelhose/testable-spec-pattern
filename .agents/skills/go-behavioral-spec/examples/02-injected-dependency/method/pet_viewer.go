package method

import (
	"context"

	"example.com/go-behavioral-spec/examples/02-injected-dependency/functional"
)

// PetViewer is the method-form wiring: PetGetter is injected on the struct.
type PetViewer struct {
	Getter functional.PetGetter
}

// BuildPetView is the same contract as functional.BuildPetView; only construction differs.
func (v *PetViewer) BuildPetView(ctx context.Context, name string) (functional.PetView, error) {
	return functional.BuildPetView(ctx, name, v.Getter)
}
