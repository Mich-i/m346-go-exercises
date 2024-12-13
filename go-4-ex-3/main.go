package main

import (
	"fmt"
	"math"
)

func computeQuadraticFormula(a float64, b float64, c float64) []float64 {
	discriminant := math.Pow(b, 2) - 4*a*c

	if discriminant > 0 {
		x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		x2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		return []float64{x1, x2}

	} else if discriminant == 0 {
		x := -b / (2 * a)
		return []float64{x}

	} else {
		return []float64{}
	}
}

func main() {
	fmt.Println(computeQuadraticFormula(3, 4, 1))
	fmt.Println(computeQuadraticFormula(2, 4, 2))
	fmt.Println(computeQuadraticFormula(3, 4, 2))
}
