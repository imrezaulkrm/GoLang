package main

import "fmt"

func main() {
	// your code goes here
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var x, y, total float32
		fmt.Scan(&x, &y)
		total = (x * 4) + y
		fmt.Println(total)
	}
}
