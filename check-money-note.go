package main

import "fmt"

func main() {
	var money, note1000, note500, note200, note100, note50, note20, note10, note5, note2, note1, totalNote int
	fmt.Println("Count total number of notes in given amount")
	fmt.Print("Enter amount of money: ")
	fmt.Scanf("%v\n", &money)

	if money >= 1000 {
		note1000 = money / 1000
		money = money % 1000
	}
	if money >= 500 {
		note500 = money / 500
		money = money % 500
	}
	if money >= 200 {
		note200 = money / 200
		money = money % 200
	}
	if money >= 100 {
		note100 = money / 100
		money = money % 100
	}
	if money >= 50 {
		note50 = money / 50
		money = money % 50
	}
	if money >= 20 {
		note20 = money / 20
		money = money % 20
	}
	if money >= 10 {
		note10 = money / 10
		money = money % 10
	}
	if money >= 5 {
		note5 = money / 5
		money = money % 5
	}
	if money >= 2 {
		note2 = money / 2
		money = money % 2
	}
	if money >= 1 {
		note1 = money / 1
		money = money % 1
	}
	totalNote = note1000 + note500 + note200 + note100 + note50 + note20 + note10 + note5 + note2 + note1
	fmt.Printf("Total number of notes: %v\n 1000: %v\n 500: %v\n 200: %v\n 100: %v\n 50: %v\n 20: %v\n 10: %v\n 5: %v\n 2: %v\n 1: %v\n", totalNote, note1000, note500, note200, note100, note50, note20, note10, note5, note2, note1)
}
