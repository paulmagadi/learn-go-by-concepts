package main

import "fmt"

func main(){
	const favNumber int = 7
	const favQuote string = "Keep the fire in you burning"

	fmt.Println("My favourite number is:", favNumber)
	fmt.Printf("My favourite quote is: %q\n", favQuote)
}