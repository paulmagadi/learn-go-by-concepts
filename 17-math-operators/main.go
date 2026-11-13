package main

import "fmt"

func main() {
	bPrice := 50.0
	sPrice := 100.0
	taxRate := 0.08

	tax := sPrice * taxRate
	fmt.Println("Tax: " ,tax)

	total := sPrice + (sPrice * taxRate)
	fmt.Println("Total cost: ", total)

	profit := sPrice - bPrice
	fmt.Println("Profit: ", profit)

	percentageProfit := (profit / bPrice) * 100
	fmt.Println("Percentage profit: ", percentageProfit, "%")

	// modulus operator
	// returns the remainder
	fmt.Println(10%2, 10%3, 10%4, 10%5, 10%6)
}