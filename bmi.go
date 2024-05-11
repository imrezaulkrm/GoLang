package main

import "fmt"

func main() {
	var height, weight, bmi, m float32

	fmt.Print("Enter Your Height in inc : ")
	fmt.Scanln(&height)
	fmt.Print("Enter Your Weight in KG : ")
	fmt.Scanln(&weight)

	m = height * 0.0254
	m *= m
	bmi = weight / m

	if bmi <= 18.4 {
		fmt.Printf("You are Underweight\nYour BMI is : %.2f\n", bmi)
	} else if 18.5 <= bmi && bmi <= 24.9 {
		fmt.Printf("You are Normal\nYour BMI is : %.2f\n", bmi)
	} else if 25.0 <= bmi && bmi <= 39.9 {
		fmt.Printf("You are Overweight\nYour BMI is : %.2f\n", bmi)
	} else {
		fmt.Printf("You are Obese\nYour BMI is : %.2f\n", bmi)
	}

}
