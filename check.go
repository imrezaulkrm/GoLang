package main

import "fmt"

func main() {
	// arithmetic operator -> +, -, *, /, %
	var num1, num2 int
	fmt.Printf("num1= ")
	fmt.Scan(&num1)
	fmt.Printf("num2= ")
	fmt.Scan(&num2)

	result := num1 + num2
	fmt.Printf("%v+%v=%v\n", num1, num2, result)

	result = num1 - num2
	fmt.Printf("%v-%v=%v\n", num1, num2, result)

	result = num1 * num2
	fmt.Printf("%v*%v=%v\n", num1, num2, result)

	var result2 = float32(num1) / float32(num2)
	fmt.Printf("%v/%v=%v\n", num1, num2, result2)

	result = num1 % num2
	fmt.Printf("%v%%%v=%v\n", num1, num2, result)
}
