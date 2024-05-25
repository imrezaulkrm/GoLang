package main

import "fmt"

func main() {
	// your code goes here
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		if x == y {
			fmt.Println("SAME")
		} else if x < y {
			fmt.Println("BIKE")
		} else if x > y {
			fmt.Println("CAR")
		}
	}
}
