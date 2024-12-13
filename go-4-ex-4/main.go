package main

import "fmt"

func convertCelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9.0 / 5.0) + 32
}

func convertFahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5.0 / 9.0
}

func main() {
	fmt.Println(convertCelsiusToFahrenheit(0.0))
	fmt.Println(convertCelsiusToFahrenheit(-40.0))

	fmt.Println(convertFahrenheitToCelsius(77.0))
	fmt.Println(convertFahrenheitToCelsius(-40.0))
}
