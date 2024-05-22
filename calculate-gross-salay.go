package main

import "fmt"

func main() {
	var salary, hra, da, gross float32
	fmt.Println("Enter basic salary and calculate gross salary of an employee")
	fmt.Print("Input basic salary of an employee: ")
	fmt.Scan(&salary)

	if salary <= 10000 {
		hra = (20.0 / 100) * salary
		da = (80.0 / 100) * salary
		gross = salary + hra + da
		fmt.Printf("If Your Salary is %v then Your Gross Salary : %.2f\n", salary, gross)
	} else if 10001 <= salary && salary <= 20000 {
		hra = (25.0 / 100) * salary
		da = (90.0 / 100) * salary
		gross = salary + hra + da
		fmt.Printf("If Your Salary is %v then Your Gross Salary : %.2f\n", salary, gross)
	} else if salary >= 20001 {
		hra = (30.0 / 100) * salary
		da = (95.0 / 100) * salary
		gross = salary + hra + da
		fmt.Printf("If Your Salary is %v then Your Gross Salary : %.2f\n", salary, gross)
	}
}
