package main

import "fmt"

// Strukturen
type Student struct {
	FirstName string
	LastName  string
}

type Class struct {
	Name     string
	Students []Student
}

type Module struct {
	Number  int
	Classes []Class
}

func main() {
	// Beispiel-Students
	classAStudents := []Student{
		{"Max", "Muster"},
		{"Anna", "Meier"},
		{"Luca", "Schmid"},
	}
	classBStudents := []Student{
		{"Sophie", "Keller"},
		{"Tim", "Fischer"},
		{"Laura", "Huber"},
	}

	// Klassen
	classA := Class{"1A", classAStudents}
	classB := Class{"1B", classBStudents}

	// Module
	modules := []Module{
		{346, []Class{classA, classB}},
		{347, []Class{classA}},
		{348, []Class{classB}},
	}

	// Mit KI
	fmt.Println("Klassenverwaltung:")
	for _, m := range modules {
		fmt.Printf("Module %d besucht von Klassen:\n", m.Number)
		for _, c := range m.Classes {
			fmt.Printf("  %s: ", c.Name)
			for _, s := range c.Students {
				fmt.Printf("%s %s, ", s.FirstName, s.LastName)
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
