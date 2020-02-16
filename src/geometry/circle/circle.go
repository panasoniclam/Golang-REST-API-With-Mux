package circle

import (
	"math"

	"calculate/calculate"
)

func Area(r float32) float32 {
	return math.Pi * calculate.Mul(r, r)
}
