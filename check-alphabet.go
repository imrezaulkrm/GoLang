package main

import "fmt"

func main() {
	var alphabet string
	fmt.Print("Write Your Alphabet: ")
	fmt.Scanf("%v", &alphabet)

	switch alphabet {
	case "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z":
		fmt.Printf("%v is Alphabet\n", alphabet)
	default:
		fmt.Printf("%v is Not Alphabet\n", alphabet)
	}
}
