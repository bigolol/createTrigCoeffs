package createTrigCoeffs

import (
	"math"

	"github.com/bigolol/numericIntegration"
)

type TrigCoeffs struct {
	Offset int32     `json:"offset"`
	Ais    []float64 `json:"ais"`
	Bis    []float64 `json:"bis"`
}

//CreateCoeffs returns a struct containing all calculated coeffs, beginning at offset. If offset is 0, b0 will be 0
func CreateCoeffs(f func(float64) float64, amt int32, offset int32) TrigCoeffs {
	coeffs := TrigCoeffs{Offset: offset, Ais: make([]float64, amt), Bis: make([]float64, amt)}
	for i := 0; i < int(amt); i++ {
		ai, bi := createCoeffPair(f, int32(int(offset)+i))
		coeffs.Ais[i] = ai
		coeffs.Bis[i] = bi
	}
	return coeffs
}

//returns first ai, then bi
func createCoeffPair(f func(float64) float64, offset int32) (float64, float64) {
	if offset == 0 {
		return (1 / (2 * math.Pi)) * numericIntegration.Integrate(f, -math.Pi, math.Pi), 0
	}
	aiIntFunc := func(x float64) float64 {
		return f(x) * math.Cos(float64(offset)*x)
	}
	ai := (1 / math.Pi) * numericIntegration.Integrate(aiIntFunc, -math.Pi, math.Pi)
	biIntFunc := func(x float64) float64 {
		return f(x) * math.Sin(float64(offset)*x)
	}
	bi := (1 / math.Pi) * numericIntegration.Integrate(biIntFunc, -math.Pi, math.Pi)
	return ai, bi
}
