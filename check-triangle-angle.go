package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Check whether triangle is valid or not if angles are given")
	fmt.Print("Input first angle: ")
	fmt.Scanf("%v\n", &a)
	fmt.Print("Input second angle: ")
	fmt.Scanf("%v\n", &b)
	fmt.Print("Input third angle: ")
	fmt.Scanf("%v\n", &c)

	if a > 0 && b > 0 && c > 0 && 180 == a+b+c {
		fmt.Println("The triangle is valid")
	} else {
		fmt.Println("The triangle is invalid")
	}

}
