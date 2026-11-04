package main

import "fmt"

func main() {
	var numbers [5]int
	fmt.Println("Numbers array:", numbers) // This will print the default values of the array elements (0 for int) [0 0 0 0 0]

	// Assigning values to the array elements
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	n := numbers[2] // Accessing the third element of the array
	fmt.Println("The third element of this array is:", n)

	// Printing the entire array
	fmt.Println("Numbers array:", numbers)

	// Looping through the array
	fmt.Println("Array elements:")
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i]) // This will print each element of the array on a new line
	}

	fmt.Println(numbers[2:4]) // This will print the elements from index 2 to 3 (30, 40)
	fmt.Println(numbers[:3])  // This will print the first three elements (10, 20, 30)
	fmt.Println(numbers[3:])  // This will print the elements from index 3 to the end (40, 50)


	// Array of float64 to store grades
	grades := [3]float64{85.5, 90.0, 78.5}
	fmt.Println("Grades:", grades)
	fmt.Println("First grade:", grades[0])
	fmt.Println("Length of grades array:", len(grades))

	// Array of strings to store book titles
	var books [3]string
	books[0] = "The Great Gatsby"
	books[1] = "To Kill a Mockingbird"
	books[2] = "1984"
	fmt.Println("Books:", books)

	books[1] = "Pride and Prejudice" // Modifying the second element
	fmt.Println("Updated Books:", books)

	// Declaring and initializing an array without specifying the size 
	names := [...]string{"Alice", "Bob", "Charlie", "David"}
	fmt.Println("Names:", len(names))

	b := [...]int{100, 3: 400, 500}
	fmt.Println("Array b:", b) // This will print [100 0 0 400 500], where the 3rd index is set to 400 and the rest are default values (0 for int)


	// Multi-dimensional array (2D array) to store a matrix
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("Matrix:", matrix)
	fmt.Println("Element at row 1, column 2:", matrix[1][2]) // Accessing the element at row 1, column 2 (which is 6)

	// Looping through the 2D array to print all elements
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("Element at [%d][%d]: %d\n", i, j, matrix[i][j])
		}
	}

	// Multi-dimensional array (3D array) to store a cube of integers 
	var multiDimArray [2][3][4]int

	// Assigning values to the 3D array
	for i := range 2 {
		for j := range 3 {
			for k := range 4 {
				multiDimArray[i][j][k] = i + j + k
			}
		}
	}

	fmt.Println("Multi-dimensional array:", multiDimArray)

}