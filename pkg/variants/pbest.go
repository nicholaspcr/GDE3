package variants

import (
	"errors"
	"math"
	"math/rand"

	"github.com/nicholaspcr/gde3/pkg/models"
)

// pbest
type pbest struct{}

func Pbest() models.Variant {
	return &pbest{}
}

func (p *pbest) Name() string {
	return "pbest"
}

func (p *pbest) Mutate(
	elems, rankZero models.Population,
	params models.VariantParams,
) (models.Vector, error) {

	ind := make([]int, 3)
	ind[0] = params.CurrPos

	err := generateIndices(1, len(elems), ind)
	if err != nil {
		return models.Vector{}, errors.New(
			"insufficient size for the population, must me equal or greater than 5",
		)
	}

	indexLimit := int(math.Ceil(float64(len(rankZero)) * params.P))
	bestIndex := rand.Int() % indexLimit

	arr := make([]float64, params.DIM)

	r1, r2 := elems[ind[1]], elems[ind[2]]
	curr := elems[params.CurrPos]
	best := rankZero[bestIndex]

	for i := 0; i < params.DIM; i++ {
		arr[i] = curr.X[i] + params.F*(best.X[i]-curr.X[i]) + params.F*(r1.X[i]-r2.X[i])
	}

	ret := models.Vector{
		X: arr,
	}
	return ret, nil
}
