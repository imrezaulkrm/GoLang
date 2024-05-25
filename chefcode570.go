package main

import "fmt"

func main() {
	// your code goes here
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		if x < y || x == y {
			fmt.Println(0)
		} else if x > y {
			x -= y
			if x < 4 {
				fmt.Println(1)
			} else if x%4 != 0 {
				x /= 4
				x += 1
				fmt.Println(x)
			} else if x%4 == 0 {
				x /= 4
				fmt.Println(x)
			}
		}
	}
}
