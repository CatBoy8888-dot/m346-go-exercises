package main

import "fmt"

type FullName struct {
	// TODO: add fields
	FirstName string
	LastName  string
}

// TODO: declare a structure for birth date

type BirthDay struct {
	DayOfBirth   byte
	MonthOfBirth byte
	YearOfBirth  int16
}

type Profile struct {
	// TODO: embed full name and birth date information

	FullName
	BirthDay
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		FullName: FullName{
			FirstName: "Max",
			LastName:  "Mustermann",
		},
		BirthDay: BirthDay{
			DayOfBirth:   2,
			MonthOfBirth: 4,
			YearOfBirth:  2000,
		},
		NumberOfSiblings: 0,        // TODO: adjust
		ZodiacSign:       '\u2652', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister

	me.NumberOfSiblings = 1
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
