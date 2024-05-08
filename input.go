package main

import "fmt"

func main() {
	var fristName, lastName string
	var roll int
	var gpa float32

	fmt.Print("Enter Your Name: ")
	fmt.Scanf("%v %v", &fristName, &lastName)

	fmt.Print("Enter Your Roll: ")
	fmt.Scanln(&roll)

	fmt.Print("Enter Your CGPA: ")
	fmt.Scan(&gpa)

	fmt.Printf("My name is %v %v \nmy roll is %v \nmy CGPA is %v \n", fristName, lastName, roll, gpa)
}
