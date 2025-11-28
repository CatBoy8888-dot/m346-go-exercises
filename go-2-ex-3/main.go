package main

import "fmt"

func main() {
	// TODO: create a map called "modules"

	modules := map[uint]string{
		104: "Keine Ahnung",
		117: "Netzwerk",
		346: "Cloud",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// TODO: delete one
	// TODO: add one
	// TODO: replace one
	fmt.Println(modules)
}
