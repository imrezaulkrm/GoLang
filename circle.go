package main

import "fmt"

func main() {
	var r, result float32
	fmt.Print("Enter Your Circle Redius: ")
	fmt.Scanf("%v\n", &r)

	result = 3.1416 * r * r
	fmt.Printf("This is your area of Circle: %.3f\n", result)
}
