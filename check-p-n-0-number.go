package main

import "fmt"

func main() {
	var num float32
	fmt.Println("Check whether a number is positive, negative or zero")
	fmt.Print("Enter Number: ")
	fmt.Scanln(&num)
	if num > 0 {
		fmt.Printf("%v Positive Number\n", num)
	} else if num < 0 {
		fmt.Printf("%v Negative Number\n", num)
	} else {
		fmt.Printf("This Number is Zero\n")
	}

}
