package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Check whether triangle is equilateral, scalene or isosceles")
	fmt.Print("Input first sides: ")
	fmt.Scanf("%v\n", &a)
	fmt.Print("Input second sides: ")
	fmt.Scanf("%v\n", &b)
	fmt.Print("Input third sides: ")
	fmt.Scanf("%v\n", &c)

	if a > 0 && b > 0 && c > 0 && a == b && b == c && c == a {
		fmt.Println("The triangle is Equilateral")
	} else if a > 0 && b > 0 && c > 0 && a == b || a == c || b == c {
		fmt.Println("The triangle is Scalene")
	} else {
		fmt.Println("The triangle is Isosceles")
	}

}
