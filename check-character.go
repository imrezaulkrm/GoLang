package main

import "fmt"

func main() {
	var character string
	fmt.Println("Check whether a character is alphabet, digit or special character")
	fmt.Print("Enter Character: ")
	fmt.Scanf("%v\n", &character)

	if character >= "a" && character <= "z" {
		fmt.Printf("%v is a Alphabet\n", character)
	} else if character >= "A" && character <= "Z" {
		fmt.Printf("%v is a Alphabet\n", character)
	} else if character >= "0" && character <= "9" {
		fmt.Printf("%v is a Digit\n", character)
	} else {
		fmt.Printf("%v is a Special Character\n", character)
	}
}
