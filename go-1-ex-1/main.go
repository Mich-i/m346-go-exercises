package main

import "fmt"

func main() {
	firstName := "Michael"
	lastName := "Stutz"
	dayOfBirth := 24
	monthOfBirth := 6
	yearOfBirth := 2007
	numberOfSiblings := 1
	heightInMeters := 1.73
	zodiacSign := 'K'

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}