package main

import "fmt"

func main() {
	var fever, cough, headache, tiredness, muscle, vomiting, diarrhea, sneezing, pain, swelling string
	fmt.Print("You Have Fever: ")
	fmt.Scanln(&fever)
	fmt.Print("You Have Cough: ")
	fmt.Scanln(&cough)
	fmt.Print("You Have Headache: ")
	fmt.Scanln(&headache)
	fmt.Print("You Have Tiredness: ")
	fmt.Scanln(&tiredness)
	fmt.Print("You Have Muscle: ")
	fmt.Scanln(&muscle)
	fmt.Print("You Have Vomiting: ")
	fmt.Scanln(&vomiting)
	fmt.Print("You Have Diarrhea: ")
	fmt.Scanln(&diarrhea)
	fmt.Print("You Have Sneezing: ")
	fmt.Scanln(&sneezing)
	fmt.Print("You Have Pain: ")
	fmt.Scanln(&pain)
	fmt.Print("You Have Swelling: ")
	fmt.Scanln(&swelling)

	if fever == "y" || cough == "y" || headache == "y" || tiredness == "y" || muscle == "y" || vomiting == "y" || diarrhea == "y" {
		fmt.Println("You Have Flu")
	} else if cough == "y" || headache == "y" || tiredness == "y" || sneezing == "y" {
		fmt.Println("You Have Cold")
	} else if cough == "y" || pain == "y" || swelling == "y" || fever == "y" {
		fmt.Println("You Have Sore throat")
	} else {
		fmt.Println("You are fit")
	}
}
