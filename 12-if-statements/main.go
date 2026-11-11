package main

import "fmt"

func main() {
	score := 42

	if score > 49 {
		fmt.Println("Pass")
	} 

	if score < 50 {
		fmt.Println("Fail")
	}

	if score >= 50 {
		fmt.Println("Pass")
	}

	if score <= 49 {
		fmt.Println("Fail")
	}

	if score == 42 {
		fmt.Println("The answer to life, the universe, and everything")
	}

	if score >= 40 && score <= 50 {
		fmt.Println("Score is between 40 and 50")
	}

	if score < 0 || score > 100 {
		fmt.Println("Invalid score")
	}
}