package main

import "fmt"

func main(){

status := "success" // You can change this value to test different cases
	// status := "error" 
	// status := "loading" 
	switch status {
	case "success":
		fmt.Println("Operation completed successfully.")
	case "error":
		fmt.Println("Something went wrong.")
	case "loading":
		fmt.Println("Please wait...")
	default:
		fmt.Println("Unknown status.")
	}

	// Additional switch statement to demonstrate multiple cases
	day := "Monday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday.")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend.")
	default:
		fmt.Println("Unknown day.")
	}


	// Switch statement with fallthrough
	// This (grade := "B"), will print both "Good job!" and "You can do better."
	grade := "B" // You can change this value to test different cases
	switch grade {
	case "A":
		fmt.Println("Excellent!")
	case "B":
		fmt.Println("Good job!")
		fallthrough
	case "C":
		fmt.Println("You can do better.")
	default:
		fmt.Println("Invalid grade.")
	}

}