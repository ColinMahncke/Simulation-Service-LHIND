package functions

import (
	"math"
)

//the function to display the rise and fall (minimum and maximum)
func f(x float64) float64 {
	a := 2.4
	f := (10 * math.Pow(a, 3)) / (math.Pow(x-12, 2) + 15*math.Pow(a, 2))
	return f
}

func Bell(x float64) float64 {
	return f(x) - f(0.0)
}
