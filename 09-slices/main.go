package main

import "fmt"

func main() {
	// Create a slice of strings
	fruits := []string{"apple", "banana", "cherry", "date"}

	// Print the entire fruits slice
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

	// An uninitialized slice equals to nil and has length 0
	var vegetables []string
	fmt.Println("Vegetables slice:", vegetables)
	fmt.Println("Length of vegetables slice:", len(vegetables))

	// To create a slice with non-zero length, use the builtin make function
	// The first argument is the type of the slice, the second is the length, and the third is the capacity
	vegetables = make([]string, 3)
	fmt.Println("Vegetables slice after make:", vegetables)
	fmt.Println("Length of vegetables slice after make:", len(vegetables))

	// Assigning values to the vegetables slice
	vegetables[0] = "carrot"
	vegetables[1] = "broccoli"
	vegetables[2] = "spinach"
	fmt.Println("Vegetables slice after assigning values:", vegetables)

	// Appending elements to the vegetables slice
	vegetables = append(vegetables, "pepper", "onion")
	fmt.Println("Vegetables slice after appending:", vegetables)
	fmt.Println("Length of vegetables slice after appending:", len(vegetables))

	// Copying slices
	// Create a new slice to copy the vegetables slice
	copiedVegetables := make([]string, len(vegetables))
	copy(copiedVegetables, vegetables)
	fmt.Println("Copied Vegetables slice:", copiedVegetables)


	// Slicing a slice
	// Slicing the fruits slice to get a new slice containing the first three fruits
	firstThreeFruits := fruits[:3]
	fmt.Println("First three fruits:", firstThreeFruits)

	// Slicing the fruits slice to get a new slice containing the middle fruits
	middleFruits := fruits[1:4]
	fmt.Println("Middle fruits (from index 1 to 3):", middleFruits)

	// Slicing the vegetables slice to get a new slice containing the last two vegetables
	lastTwoVegetables := vegetables[len(vegetables)-2:]
	fmt.Println("Last two vegetables:", lastTwoVegetables)

	toTheEndFruits := fruits[2:]
	fmt.Println("Fruits from index 2 to the end:", toTheEndFruits)

	// Modifying the sliced slices will affect the original slices since they share the same underlying array
	firstThreeFruits[0] = "avocado"
	fmt.Println("First three fruits after modification:", firstThreeFruits)
	fmt.Println("Original fruits slice:", fruits)

	// Modifying the original fruits slice will also affect the sliced slices
	fruits[3] = "dragonfruit"
	fmt.Println("Fruits slice after modification:", fruits)
	fmt.Println("Middle fruits slice after original modification:", middleFruits)


	// Demonstrating the capacity of slices
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice
	fmt.Println("Capacity of fruits slice:", cap(fruits))
	fmt.Println("Capacity of first three fruits slice:", cap(firstThreeFruits))
	fmt.Println("Capacity of middle fruits slice:", cap(middleFruits))
	fmt.Println("Capacity of last two vegetables slice:", cap(lastTwoVegetables))


	// Creating a slice with a specified capacity using the make function
	// The first argument is the type of the slice, the second is the length, and the third is the capacity
	customSlice := make([]int, 5, 10)
	fmt.Println("Custom slice:", customSlice)
	fmt.Println("Length of custom slice:", len(customSlice))
	fmt.Println("Capacity of custom slice:", cap(customSlice))


	// Demonstrating the difference between length and capacity when appending elements
	// Appending elements to the custom slice
	customSlice = append(customSlice, 1, 2, 3, 4, 5)
	fmt.Println("Custom slice after appending:", customSlice)
	fmt.Println("Length of custom slice after appending:", len(customSlice))
	fmt.Println("Capacity of custom slice after appending:", cap(customSlice))

	// Appending more elements to the custom slice to exceed its initial capacity
	customSlice = append(customSlice, 6, 7, 8, 9, 10)
	fmt.Println("Custom slice after exceeding initial capacity:", customSlice)
	fmt.Println("Length of custom slice after exceeding initial capacity:", len(customSlice))
	fmt.Println("Capacity of custom slice after exceeding initial capacity:", cap(customSlice))

	// Multi-dimensional slices (slices of slices)
	// Creating a 2D slice to represent a matrix
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Matrix:", matrix)


	// Accessing elements in the 2D slice
	fmt.Println("Element at row 1, column 2:", matrix[1][2]) // This will print 6


	// looping through the 2D slice to print all elements
	fmt.Println("Matrix elements:")
	for i := range matrix {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("Element at row %d, column %d: %d\n", i, j, matrix[i][j])
		}
	}

}