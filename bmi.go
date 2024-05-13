package main

import "fmt"

func main() {
	var height, weight, weight1, bmi, bmi1, m, need float32
	bmi1 = 22.0

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

	weight1 = bmi1 * m
	need = weight1 - weight

	if weight1 < weight {
		fmt.Printf("You need to lose weight : %.2fKG\nYou may choose to follow any of these diet plans, but it's important to consult with a reputable dietitian beforehand.\n\n1. Vegan diet\n2. Low carb diet\n3. Mediterranean diet\n4. Atkins diet\n5. Keto diet\n", need)
	} else if weight1 > weight {
		fmt.Printf("You need to gain weight : %.2fKG\nYou may choose to follow any of these diet plans, but it's important to consult with a reputable dietitian beforehand.\n\n1. Bulking diet\n2. High-Calorie Diet\n3. Mass-Gaining Diet\n4. Hypercaloric Diet\n5. Calorie Surplus\n", need)
	}
}
