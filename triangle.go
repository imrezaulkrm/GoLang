package main

import "fmt"

func main() {
	var base, height, result float32
	fmt.Print("Enter Your Triangle Base: ")
	fmt.Scanln(&base)

	fmt.Print("Enter Your Triangle Height: ")
	fmt.Scanln(&height)

	result = 0.5 * base * height

	fmt.Printf("This is the are of Triangle:%.3f\n", result)

}
