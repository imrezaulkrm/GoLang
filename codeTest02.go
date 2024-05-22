package main

import "fmt"

func main() {
	// your code goes here
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var n, x, s, total int
		fmt.Scan(&n, &x)

		if n > 6 && n%6 != 0 {
			s = n / 6
			s += 1
			total = s * x
			fmt.Printf("%v\n", total)
		} else if n%6 == 0 {
			s = n / 6
			total = s * x
			fmt.Printf("%v\n", total)
		} else {
			fmt.Printf("%v\n", x)
		}
	}

}
