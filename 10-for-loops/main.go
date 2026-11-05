package main

import "fmt"

func main() {
	i := 0
	for i <= 5 {
		fmt.Println("Index:", i)
		i++
	}

	for i := 1; i < 4; i++ {
		fmt.Println("Line:", i)
	}

	names := []string{"Alice", "Bob", "Charlie"}
	for i := range len(names){
		fmt.Println("Name:", names[i])
	}

	for i, name := range names {
		fmt.Println("Index:", i, "Name:", name)
	}

	for _, name := range names {
		fmt.Println("Name:", name)
	}

	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println("Number:", num)
	}

	// Infinite loop example (commented out to prevent execution)
	// for {
	// 	fmt.Println("This is an infinite loop")
	// }

	for {
		fmt.Println("This loop will run forever unless broken")
		break // Break the loop to prevent infinite execution
	}


	// Loop with continue statement
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Println("Odd Number:", i)
	}


	// Nested loop example
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 2; j++ {
			fmt.Printf("Outer loop: %d, Inner loop: %d\n", i, j)
		}
	}


	// Example1: PrintCombinations of two slices
	colors := []string{"Red", "Green", "Blue"}
	shapes := []string{"Circle", "Square"}

	for _, color := range colors {
		for _, shape := range shapes {
			fmt.Printf("Combination: %s %s\n", color, shape)
		}
	}

	// Example2: PrintCombinations of three slices
	sizes := []string{"Small", "Medium", "Large"}
	for _, color := range colors {
		for _, shape := range shapes {
			for _, size := range sizes {
				fmt.Printf("Combination: %s %s %s\n", color, shape, size)
			}
		}
	}

	// Example3: PrintComb
	// Write a program that prints, in ascending order and on a single line: all unique combinations of three different digits so that, the first digit is lower than the second, and the second is lower than the third.

	// These combinations are separated by a comma and a space.

	for i := 0; i <= 7; i++ {
		for j := i + 1; j <= 8; j++ {
			for k := j + 1; k <= 9; k++ {
				fmt.Printf("%d%d%d", i, j, k)
				if i != 7 || j != 8 || k != 9 {
					fmt.Print(", ")
				}
			}
		}
	}
	
}