package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Check whether triangle is valid or not if sides are given")
	fmt.Print("Input first sides: ")
	fmt.Scanf("%v\n", &a)
	fmt.Print("Input second sides: ")
	fmt.Scanf("%v\n", &b)
	fmt.Print("Input third sides: ")
	fmt.Scanf("%v\n", &c)

	if a > 0 && b > 0 && c > 0 && a+b > c && b+c > a && c+a > b {
		fmt.Println("The triangle is valid")
	} else {
		fmt.Println("The triangle is invalid")
	}

}
