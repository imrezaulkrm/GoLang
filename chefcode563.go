package main

import "fmt"

func main() {
	// your code goes here
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var p float64
		fmt.Scan(&p)
		p *= 1000
		p /= 100
		fmt.Printf("%d\n", int(p))
	}
}
