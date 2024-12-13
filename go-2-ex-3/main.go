package main

import "fmt"

func main() {
	// Definition der Module als Map
	modules := map[int]string{
		104: "Programmieren Grundlagen",
		117: "Systemadministration",
		346: "Cloud-Lösungen konzipieren und realisieren",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	delete(modules, 117)

	modules[210] = "Webentwicklung"

	modules[346] = "DevOps-Konzepte"

	fmt.Println(modules)
}
