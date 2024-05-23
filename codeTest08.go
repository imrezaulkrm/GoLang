package main

import "fmt"

func main() {
	// your code goes here
	var x, y, z, n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {

		fmt.Scan(&x, &y, &z)
		emptyCount := 0
		if x == 0 {
			emptyCount++
		}
		if y == 0 {
			emptyCount++
		}
		if z == 0 {
			emptyCount++
		}

		if emptyCount >= 2 {
			fmt.Println("Water filling time")
		} else {
			fmt.Println("Not Now")
		}
	}

}
