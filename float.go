package main

import "fmt"

func main() {
	var a int
	var b int
	var z float32

	fmt.Print("Enter Your First Number: ")
	fmt.Scanf("%d\n", &a)
	fmt.Print("Enter Second Number: ")
	fmt.Scanf("%d\n", &b)
	z = float32(a) / float32(b)
	fmt.Printf("This Is Your Floating Number: %.20f\n", z)
	fmt.Printf("This is Binary Number: %b\n", int(z))
}
