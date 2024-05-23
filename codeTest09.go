package main

import "fmt"

func main() {
	// your code goes here
	var x, n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {

		fmt.Scan(&x)
		if x <= 100 {
			fmt.Printf("%v\n", x)
		} else if 100 < x && x <= 1000 {
			x -= 25
			fmt.Printf("%v\n", x)
		} else if 1000 < x && x <= 5000 {
			x -= 100
			fmt.Printf("%v\n", x)
		} else if x > 5000 {
			x -= 500
			fmt.Printf("%v\n", x)
		}
	}
}
