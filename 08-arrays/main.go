package main

import "fmt"

func main() {
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	n := numbers[2] // Accessing the third element of the array
	fmt.Println("The third element of this array is:", n)

	fmt.Println("Array elements:")
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	grades := [3]float64{85.5, 90.0, 78.5}
	fmt.Println("Grades:", grades)
	fmt.Println("First grade:", grades[0])
	fmt.Println("Length of grades array:", len(grades))
}