package main

import "fmt"

func main() {
	var num int
	fmt.Println("Check whether a number is divisible by 5 and 11 or not")
	fmt.Print("Enter Number: ")
	fmt.Scanln(&num)

	if num%5 == 0 && num%11 == 0 {
		fmt.Printf("%v number is divisible by 5 and 11\n", num)
	} else if num%5 == 0 && num%11 != 0 {
		fmt.Printf("%v number is just divisible by 5\n", num)
	} else if num%5 != 0 && num%11 == 0 {
		fmt.Printf("%v number is just divisible by 11\n", num)
	} else {
		fmt.Printf("%v number is not divisible by 5 or 11\n", num)
	}
}
