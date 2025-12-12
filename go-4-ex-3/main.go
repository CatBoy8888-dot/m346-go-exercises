package main

import (
	"fmt"
	"math"
	"strconv"
)

// TODO: implement the function computeQuadraticFormula

func main() {
	equations := [][]float64{
		{3, 4, 1},
		{2, 4, 2},
		{3, 4, 2},
	}
	var output [][]string

	for _, eq := range equations {
		a, b, c := eq[0], eq[1], eq[2]
		D := computeQuadraticFormula(a, b, c)
		row := []string{D, strconv.FormatFloat(a, 'f', -1, 64), strconv.FormatFloat(b, 'f', -1, 64), strconv.FormatFloat(c, 'f', -1, 64)}
		output = append(output, row)
	}

	for _, row := range output {
		fmt.Println(row)
	}

}

func computeQuadraticFormula(a, b, c float64) string {
	discriminant := math.Pow(b, 2) - 4*a*c

	if discriminant > 0 {
		return "+"
	} else if discriminant == 0 {
		return "0"
	} else {
		return "-"
	}
}
