package main

import "fmt"

func main() {
	var c, s, p, l int
	fmt.Println("Calculate profit or loss")
	fmt.Print("Input cost price: ")
	fmt.Scanf("%v\n", &c)
	fmt.Print("Input selling price: ")
	fmt.Scanf("%v\n", &s)

	if c < s {
		p = s - c
		fmt.Printf("Profit: %v TAKA\n", p)
	} else if c > s {
		l = c - s
		fmt.Printf("Loss: %v TAKA\n", l)
	}
}
