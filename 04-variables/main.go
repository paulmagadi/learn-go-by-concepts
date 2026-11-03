package main

import "fmt"

func main() {
	city := "Nairobi"
	year := 2024
	price := 97.50

	msg := fmt.Sprintf("In %s in the year %d the price of petrol was ksh %.2f", city, year, price)
	fmt.Println(msg)

	fmt.Println("In", city, "in the year", year, "the price of petrol was ksh", price)
}
