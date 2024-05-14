package main

import "fmt"

func main() {
	var month int
	fmt.Println("Find number of days in month")
	fmt.Print("Enter month number: ")
	fmt.Scanf("%v\n", &month)

	if month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12 {
		fmt.Println("This Month is 31 days")
	} else if month == 2 {
		fmt.Println("This Month is 28 days or 29 days")
	} else if month >= 4 && month <= 11 {
		fmt.Println("This Month is 30 days")
	} else {
		fmt.Println("Invalid Input! Please enter month number between 1-12.")
	}
}

/*func main() {
	var week int
	fmt.Println("Find number of days in month")
	fmt.Print("Enter week day number: ")
	fmt.Scanf("%v\n", &week)

	switch week {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("Saturday")
	case 2:
		fmt.Println("Sunday")
	case 4, 6, 9, 11:
		fmt.Println("Monday")
	default:
		fmt.Println("Invalid Input! Please enter month number between 1-12.")
	}

}
*/
