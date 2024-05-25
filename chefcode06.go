package main

import "fmt"

func main() {
	// your code goes here
	var n, weapen int
	fmt.Scan(&n)
	evenCount := 0
	oddCount := 0
	for i := 0; i < n; i++ {

		fmt.Scan(&weapen)
		if weapen%2 == 0 {
			evenCount++
		} else if weapen%2 != 0 {
			oddCount++
		}
	}
	if evenCount > oddCount {
		fmt.Println("READY FOR BATTLE")
	} else {
		fmt.Println("NOT READY")
	}

}
