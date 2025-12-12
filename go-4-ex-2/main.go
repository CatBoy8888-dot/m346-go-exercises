package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt

// func main() {
// 	// TODO: call the function computeHypotenuse

// 	fmt.Println(computeHypotenuse(3, 4)) // 5
// 	fmt.Println(computeHypotenuse(5, 12)) // 13
// 	fmt.Println(computeHypotenuse(7, 24)) // 25
// }

type ShortSides struct {
	a float64
	b float64
}

func Hypotenuse(x ShortSides) float64 {
	return math.Sqrt(math.Pow(x.a, 2) + math.Pow(x.b, 2))
}
func main() {
	// TODO: call the function computeHypotenuse

	fmt.Println(computeHypotenuse(3, 4))  // 5
	fmt.Println(computeHypotenuse(5, 12)) // 13
	fmt.Println(computeHypotenuse(7, 24)) // 25
}

func computeHypotenuse(a float64, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}
