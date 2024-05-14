package main

import "fmt"

func main() {
	var alphabet string
	fmt.Println("Check vowel or consonant")
	fmt.Print("Write Your Alphabet: ")
	fmt.Scanf("%v", &alphabet)

	switch alphabet {
	case "a", "e", "i", "o", "u":
		fmt.Printf("%v is vowels\n", alphabet)
	case "b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z":
		fmt.Printf("%v is consonants\n", alphabet)
	default:
		fmt.Printf("%v is Not Alphabet\n", alphabet)
	}
}

/*
func main() {
    var alphabet string
    fmt.Print("Write Your Alphabet: ")
    fmt.Scanf("%v", &alphabet)

    if alphabet == "a" || alphabet == "e" || alphabet == "i" || alphabet == "o" || alphabet == "u" {
        fmt.Printf("%v is vowels\n", alphabet)
    } else if alphabet >= "b" && alphabet <= "z" {
        fmt.Printf("%v is consonants\n", alphabet)
    } else {
        fmt.Printf("%v is Not Alphabet\n", alphabet)
    }
}
*/
