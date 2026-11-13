package main

import "fmt"

func main() {
	var scores map[string]int
    var names []string
    var value *float64

	if scores == nil {
		fmt.Println("Score is required.")
	}

	if names == nil {
		fmt.Println("Names are required.")
	}

	if value == nil {
		fmt.Println("Value cannot be nil.")
	}

	fmt.Println(scores == nil)
	fmt.Println(names == nil)
	fmt.Println(value == nil)

}