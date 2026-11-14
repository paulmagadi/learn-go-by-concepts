package main

import "fmt"

func main() {
	name := "Notebook"
	price := 9.99

	fmt.Printf("%s costs $ %.2f \n", name, price)

	message := fmt.Sprintf("%s costs $ %.2f \n", name, price)

	fmt.Println(message)
}