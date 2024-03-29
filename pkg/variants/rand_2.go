package variants

import (
	"errors"

	"github.com/nicholaspcr/gde3/pkg/models"
)

// rand2 a + F(b-c) + F(d-e)
type rand2 struct{}

func Rand2() models.Variant {
	return &rand2{}
}

func (r *rand2) Name() string {
	return "rand2"
}

func (r *rand2) Mutate(
	elems, rankZero models.Population,
	p models.VariantParams,
) (models.Vector, error) {

	// generating random indices different from current pos
	ind := make([]int, 6)
	ind[0] = p.CurrPos
	err := generateIndices(1, len(elems), ind)
	if err != nil {
		return models.Vector{}, errors.New(
			"insufficient size for the population, must me equal or greater than 4",
		)
	}

	arr := make([]float64, p.DIM)

	i1, i2, i3, i4, i5 := ind[1], ind[2], ind[3], ind[4], ind[5]
	r1, r2, r3, r4, r5 := elems[i1], elems[i2], elems[i3], elems[i4], elems[i5]

	for i := 0; i < p.DIM; i++ {
		arr[i] = r1.X[i] + p.F*(r2.X[i]-r3.X[i]) + p.F*(r4.X[i]-r5.X[i])
	}
	ret := models.Vector{
		X: arr,
	}
	return ret, nil
}
