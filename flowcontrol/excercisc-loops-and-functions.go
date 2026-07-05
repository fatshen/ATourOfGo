package main

import (
	"fmt"
	"math"
)

const Delta = 0.0001

func Sqrt(x float64) float64 {
	z := 1.0
	d := math.Abs(z*z - x)
	for d > Delta {
		z = (z*z - x) / (2 * z)
		d = math.Abs(z*z - x)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
