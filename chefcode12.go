package main

import "fmt"

func main() {
	// your code goes here
	var x, y, n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {

		fmt.Scan(&x, &y)
		if x <= y*2 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
