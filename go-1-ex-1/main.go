package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	var firstName string = "Max"
	var lastName string = "Schaad"
	var dayOfBirth uint8 = 14
	var monthOfBirth uint8 = 2
	var yearOfBirth uint16 = 2009
	var numberOfSiblings uint8 = 0
	var heightInMeters float32 = 1.8
	var zodiacSign rune = '\u2652'

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
