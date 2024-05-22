package main

import "fmt"

func main() {
	var amount, tAmount, surcharge, ebill float32
	fmt.Println("Calculate electricity bill")
	fmt.Print("Enter total units consumed: ")
	fmt.Scanf("%v", &amount)

	if amount <= 50 {
		tAmount = amount * 0.50
	} else if amount <= 150 {
		tAmount = (0.50 * 50) + ((amount - 50) * 0.75)
	} else if amount <= 250 {
		tAmount = (0.50 * 50) + (0.75 * 100) + ((amount - 150) * 1.20)
	} else if amount > 250 {
		tAmount = (0.50 * 50) + (0.75 * 100) + (1.20 * 100) + ((amount - 250) * 1.50)
	}
	surcharge = (20.0 / 100) * tAmount
	ebill = tAmount + surcharge
	fmt.Printf("Electric Bill is %.2f Taka.\n", ebill)
}
