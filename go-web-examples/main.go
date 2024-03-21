package main

import (
	"fmt"
)

func dotProduct[F ~float32 | ~float64](v1, v2 []F) F {
	var s F
	for i, x := range v1 {
		y := v2[i]
		s += x * y
	}
	return s
}

func dotProduct1[F float32 | float64](v1, v2 []F) F {
	var s F
	for i, x := range v1 {
		y := v2[i]
		s += x * y
	}
	return s
}

type floats float32

func main() {
	// example.HelloWorld()

	// example.HelloServer()

	// example.Templates()

	// example.Websockets()
	v1 := []floats{1, 2, 3}
	v2 := []floats{0, 1, 0}
	fmt.Println(dotProduct(v1, v2))
	// fmt.Println(dotProduct1(v1, v2))
}
