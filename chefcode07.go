package main

import "fmt"

func main() {
	// your code goes here
	var n, x, y, beg int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {

		fmt.Scan(&x, &y)
		beg = x * y
		beg /= 100
		fmt.Println(beg)
	}

}
