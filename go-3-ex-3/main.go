package main

import "fmt"

const (
	Lower = 1
	Upper = 30
)

func main() {
	for i := 1; i <= 30; i++ {
		if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%5 == 0 && i%3 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println(i)
		}
	}
}
