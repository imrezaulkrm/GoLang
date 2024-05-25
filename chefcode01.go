package main

import "fmt"

func main() {
	// your code goes here
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var a, b, c float32
		fmt.Scan(&a, &b, &c)

		if (a + b) >= (2 * c) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

}
