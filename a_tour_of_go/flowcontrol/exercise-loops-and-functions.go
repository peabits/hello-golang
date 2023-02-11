package main

import (
	"fmt"
)

func Abs(x float64) float64 {
	if x >= 0 {
		return x
	}
	return -x
}

func Sqrt(x float64) float64 {
	z := 1.
	for Abs(z*z-x) > 1e-5 {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
