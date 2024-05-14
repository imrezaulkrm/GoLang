package main

import "fmt"

func main() {
	var week int
	fmt.Println("Enter week number and print day of week")
	fmt.Print("Enter week day number: ")
	fmt.Scanf("%v\n", &week)

	switch week {
	case 1:
		fmt.Println("Saturday")
	case 2:
		fmt.Println("Sunday")
	case 3:
		fmt.Println("Monday")
	case 4:
		fmt.Println("Tuesday")
	case 5:
		fmt.Println("Wednesday")
	case 6:
		fmt.Println("Thursday")
	case 7:
		fmt.Println("Friday")
	default:
		fmt.Println("Invalid Input! Please enter week number between 1-7.")
	}

}
