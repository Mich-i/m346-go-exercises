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

	// Entfernen eines Moduls
	delete(modules, 117)

	// Hinzufügen eines neuen Moduls
	modules[210] = "Webentwicklung"

	// Ersetzen eines Moduls
	modules[346] = "DevOps-Konzepte"

	fmt.Println(modules)
}
