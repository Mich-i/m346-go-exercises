package main

import "fmt"

func computerGrade(gotPoints float64, maxPoints float64) float64 {
	if maxPoints == 0 {
		return 1.0
	}

	note := 1.0 + 5.0*(gotPoints/maxPoints)

	if note > 6.0 {
		note = 6.0
	} else if note < 1.0 {
		note = 1.0
	}

	return note
}

func main() {
	fmt.Println(computerGrade(17.5, 28.0))
	fmt.Println(computerGrade(28.0, 28.0))
	fmt.Println(computerGrade(0.0, 28.0))
}
