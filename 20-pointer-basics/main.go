package main

import "fmt"

func main() {

	score := 100
	fmt.Println("Original score:", score)

	// a pointer to score variable
	ptr := &score

	// prints the memory address
	fmt.Println("The memory address of the score is:", ptr)

	ptrValue := *ptr
	fmt.Println("The value stored at the address is:", ptrValue)

	// Update the score using the dereferenced pointer
	*ptr = 150
	fmt.Println("New score:", score)
}