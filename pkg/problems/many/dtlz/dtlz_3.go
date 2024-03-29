package dtlz

import (
	"errors"
	"math"

	"github.com/nicholaspcr/gde3/pkg/models"
)

type dtlz3 struct{}

func Dtlz3() models.Problem {
	return &dtlz3{}
}

func (v *dtlz3) Name() string {
	return "dtlz3"
}

func (v *dtlz3) Evaluate(e *models.Vector, M int) error {

	if len(e.X) <= M {
		return errors.New(
			"need to have an M lesser than the amount of variables",
		)
	}

	evalG := func(v []float64) float64 {
		g := 0.0
		for _, x := range v {
			g += math.Pow(x-0.5, 2) - math.Cos(20.0*math.Pi*(x-0.5))
		}
		k := float64(len(v))
		return 100.0 * (k + g)
	}

	g := evalG(e.X[M-1:])
	objs := make([]float64, M)

	for i := 0; i < M; i++ {
		prod := (1.0 + g)
		for j := 0; j < M-(i+1); j++ {
			prod *= math.Cos(e.X[j] * 0.5 * math.Pi)
		}
		if i != 0 {
			prod *= math.Sin(e.X[M-(i+1)] * 0.5 * math.Pi)
		}
		objs[i] = prod
	}

	// puts new objectives into the elem
	e.Objs = make([]float64, len(objs))
	copy(e.Objs, objs)
	return nil
}
