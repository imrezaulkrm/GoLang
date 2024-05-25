package main

import "fmt"

func main() {
	// your code goes here
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var a float64
		fmt.Scan(&a)
		a *= 50
		a -= ((a / 100) * 70)
		fmt.Printf("%d\n", int(a))
	}
}
