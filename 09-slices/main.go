package main

import "fmt"

func main() {
	// Create a slice of strings
	fruits := []string{"apple", "banana", "cherry", "date"}
	fmt.Println("Fruits:", fruits)

	// Accessing elements in the slice
	fmt.Println("First fruit:", fruits[0])
	fmt.Println("Length of fruits slice:", len(fruits))

	// Modifying an element in the slice
	fruits[1] = "blueberry"
	fmt.Println("Updated Fruits:", fruits)

	// Appending elements to the slice
	fruits = append(fruits, "elderberry", "fig")
	fmt.Println("Fruits after appending:", fruits)
	fmt.Println("Length of fruits slice after appending:", len(fruits))


	// Fruits: [apple banana cherry date]
	// First fruit: apple
	// Length of fruits slice: 4
	// Updated Fruits: [apple blueberry cherry date]
	// Fruits after appending: [apple blueberry cherry date elderberry fig]
	// Length of fruits slice after appending: 6
}