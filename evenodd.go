package main

import "fmt"

func main() {
	var num int
	fmt.Println("Check whether a number is even or odd")
	fmt.Print("Enter Number: ")
	fmt.Scanf("%v", &num)

	if num%2 == 0 {
		fmt.Printf("%v number is Even\n", num)
	} else {
		fmt.Printf("%v number is Odd\n", num)
	}
}
