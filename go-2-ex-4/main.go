package main

import "fmt"

type Student struct {
	FirstName string
	LastName  string
}

type Class struct {
	Name     string
	Students []Student
}

func main() {
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

	modules := map[int][]Class{
		104: {class1, class2},
		210: {class1},
		346: {class2},
	}

	fmt.Println("Klassenverwaltung:")
	fmt.Println(class1)
	fmt.Println(class2)
	fmt.Println("Module und besuchte Klassen:")
	fmt.Println(modules)
}
