package main

import "fmt"

func main() {
	var year int
	fmt.Print("Enter Check Leepyear Year: ")
	fmt.Scanln(&year)
	if year%4 == 0 {
		fmt.Printf("%v is a Leep Year\n", year)
	} else {
		fmt.Printf("%v is a Normal Year\n", year)
	}
}
