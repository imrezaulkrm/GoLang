package main

import "fmt"

func main() {
	// your code goes here
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var x int
		fmt.Scan(&x)

		if x+3 <= 10 {
			fmt.Printf("YES\n")
		} else if x+3 > 10 {
			fmt.Printf("NO\n")
		}
	}

}
