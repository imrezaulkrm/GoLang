package main

import "fmt"

func main() {
	// your code goes here
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var x, y, z, students float32
		fmt.Scan(&x, &y, &z)
		students = x * y
		if students*2 > 50 {
			fmt.Printf("YES\n")
		} else {
			fmt.Printf("NO\n")
		}
	}

}
