package main

import "fmt"

func main() {
	// your code goes here
	var a, b, x, y, n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {

		fmt.Scan(&a, &b, &x, &y)
		if a*b <= x*y {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
