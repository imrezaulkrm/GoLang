package main

import "fmt"

func main() {
	// your code goes here
	var x, y, z, n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {

		fmt.Scan(&x, &y)
		if ((x * y) % 4) != 0 {
			z = (((x * y) / 4) + 1)
			fmt.Printf("%v\n", z)
		} else if ((x * y) % 4) == 0 {
			z = ((x * y) / 4)
			fmt.Printf("%v\n", z)
		}
	}
}
