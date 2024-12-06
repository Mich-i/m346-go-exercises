package main

import "fmt"

func main() {
	// Start der Fibonacci-Folge
	var fibs = []int{1, 1, 0, 0, 0}

	// Berechnung der nächsten drei Werte
	fibs[2] = fibs[0] + fibs[1]
	fibs[3] = fibs[1] + fibs[2]
	fibs[4] = fibs[2] + fibs[3]

	// Hinzufügen von vier weiteren Werten
	fibs = append(fibs, fibs[3]+fibs[4]) // 8
	fibs = append(fibs, fibs[4]+fibs[5]) // 13
	fibs = append(fibs, fibs[5]+fibs[6]) // 21
	fibs = append(fibs, fibs[6]+fibs[7]) // 34

	fmt.Println(fibs) // Erwartete Ausgabe: [1 1 2 3 5 8 13 21 34]
}
