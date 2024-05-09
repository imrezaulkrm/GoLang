package main

import "fmt"

func main() {

	var dnum int

	fmt.Print("Enter Your Decimal Numeber: ")
	fmt.Scanf("%v", &dnum)

	fmt.Printf("This is Binary Number= %b\nThis is Octal Number= %o\nThis is Hexa-decimal Number= %x\n ", dnum, dnum, dnum)

}
