package main

import "fmt"

// Schülerstruktur mit Vor- und Nachnamen
type Student struct {
	FirstName string
	LastName  string
}

// Klassenstruktur mit einer Liste von Schülern
type Class struct {
	Name     string
	Students []Student
}

func main() {
	// Definition von zwei Klassen mit jeweils drei Schülern
	class1 := Class{
		Name: "Klasse 1A",
		Students: []Student{
			{"Anna", "Müller"},
			{"Peter", "Schmidt"},
			{"Lisa", "Klein"},
		},
	}

	class2 := Class{
		Name: "Klasse 2B",
		Students: []Student{
			{"Max", "Mustermann"},
			{"Sophie", "Meier"},
			{"Jonas", "Lehmann"},
		},
	}

	// Module, die von Klassen besucht werden
	modules := map[int][]Class{
		104: {class1, class2},
		210: {class1},
		346: {class2},
	}

	// Ausgabe der Daten
	fmt.Println("Klassenverwaltung:")
	fmt.Println(class1)
	fmt.Println(class2)
	fmt.Println("Module und besuchte Klassen:")
	fmt.Println(modules)
}
