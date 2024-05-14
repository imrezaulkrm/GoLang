package main

import "fmt"

func main() {
	var no1, no2, no3 float32
	fmt.Println("Find maximum between two numbers")
	fmt.Print("Enter Your First Number: ")
	fmt.Scanln(&no1)
	fmt.Print("Enter Your Second Number: ")
	fmt.Scanln(&no2)
	fmt.Print("Enter Your Third Number: ")
	fmt.Scanln(&no3)

	if no1 > no2 && no1 > no3 {
		fmt.Printf("%v is Maximum Number.\n", no1)
	} else if no2 > no1 && no2 > no3 {
		fmt.Printf("%v is Maximum Number.\n", no2)
	} else if no3 > no1 && no3 > no2 {
		fmt.Printf("%v is Maximum Number.\n", no3)
	} else {
		fmt.Printf("%v, %v and %v Both are equal Number.\n", no1, no2, no3)
	}

}
