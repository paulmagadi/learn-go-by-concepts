package main

import "fmt"

func main(){

	var score int = 78

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else if score >= 60 {
		fmt.Println("Grade: D")
	} else if score >= 50 {
		fmt.Println("Grade: E")
	} else {
		fmt.Println("Grade: F")
	}

}