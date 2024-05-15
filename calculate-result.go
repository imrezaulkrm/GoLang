package main

import "fmt"

func main() {
	var ban, math, phy, chemis, bio, eng, percentage float32
	fmt.Println("Enter student marks and find percentage and grade")
	fmt.Print("Input marks of Bangla subjects: ")
	fmt.Scanf("%v\n", &ban)
	fmt.Print("Input marks of English subjects: ")
	fmt.Scanf("%v\n", &eng)
	fmt.Print("Input marks of Math subjects: ")
	fmt.Scanf("%v\n", &math)
	fmt.Print("Input marks of Physics subjects: ")
	fmt.Scanf("%v\n", &phy)
	fmt.Print("Input marks of Chemistry subjects: ")
	fmt.Scanf("%v\n", &chemis)
	fmt.Print("Input marks of Biology subjects: ")
	fmt.Scanf("%v\n", &bio)

	percentage = (ban + eng + math + phy + chemis + bio) / 6
	if ban > 33 && eng > 33 && math > 33 && phy > 33 && chemis > 33 && bio > 33 {
		if percentage >= 90 {
			fmt.Printf("\nYou Are Pass and Your Total Result Persentage is %.2f%%\nYour Grade Golden A+\n", percentage)
		} else if percentage >= 80 {
			fmt.Printf("\nYou Are Pass and Your Total Result Persentage is %.2f%%\nYour Grade A+\n", percentage)
		} else if percentage >= 70 {
			fmt.Printf("\nYou Are Pass and Your Total Result Persentage is %.2f%%\nYour Grade A\n", percentage)
		} else if percentage >= 60 {
			fmt.Printf("\nYou Are Pass and Your Total Result Persentage is %.2f%%\nYour Grade A-\n", percentage)
		} else if percentage >= 50 {
			fmt.Printf("\nYou Are Pass and Your Total Result Persentage is %.2f%%\nYour Grade B\n", percentage)
		} else if percentage >= 40 {
			fmt.Printf("\nYou Are Pass and Your Total Result Persentage is %.2f%%\nYour Grade C\n", percentage)
		} else if percentage >= 33 {
			fmt.Printf("\nYou Are Pass and Your Total Result Persentage is %.2f%%\nYour Grade D\n", percentage)
		}
	} else {
		fmt.Printf("\nYour are Fail\n")

	}

}
