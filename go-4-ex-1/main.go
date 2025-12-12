package main

import "fmt"

// TODO: implement the function computeGrade

func main() {
	// TODO: call the function computeGrade
	fmt.Println(computerGrade(17.5, 28.0)) //4.125
	fmt.Println(computerGrade(35, 35))     //6
	fmt.Println(computerGrade(0, 20))      //1
}

func computerGrade(gotPoints float64, maxPoints float64) float64 {
	if gotPoints < 0 || gotPoints > maxPoints || maxPoints <= 0 {
		fmt.Println("Ungueltige Eingaben")
		return -1 // Rückgabewert für ungültige Eingabe
	}
	return gotPoints/maxPoints*5 + 1
}
