package main

import "fmt"

func main() {
	var alphabet string
	fmt.Println("Check whether a character is Uppercase or Lowercase")
	fmt.Print("Write Your Alphabet: ")
	fmt.Scanf("%v", &alphabet)

	if alphabet >= "a" && alphabet <= "z" {
		fmt.Printf("%v is Lowercase Alphabet\n", alphabet)
	} else if alphabet >= "A" && alphabet <= "Z" {
		fmt.Printf("%v is Uppercase Alphabet\n", alphabet)
	} else {
		fmt.Printf("%v is Not Alphabet\n", alphabet)
	}
}
