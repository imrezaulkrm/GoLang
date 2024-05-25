package main

import "fmt"

func main() {
	// your code goes here
	var w, x, y, z, n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {

		fmt.Scan(&w, &x, &y, &z)
		emptyCount := 0
		if w == 1 {
			emptyCount++
		}
		if x == 1 {
			emptyCount++
		}
		if y == 1 {
			emptyCount++
		}
		if z == 1 {
			emptyCount++
		}

		if emptyCount >= 1 {
			fmt.Println("OUT")
		} else {
			fmt.Println("IN")
		}
	}

}
