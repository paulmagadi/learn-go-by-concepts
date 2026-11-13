package main

import "fmt"

func main() {
	fName := "John"
	favLanguage := "Go"

	fmt.Println("My name is " + fName + " and I love " + favLanguage + " programming.")
	fmt.Println("My name is", fName, "and my favorite programming language is", favLanguage, ".")

	fullName := fName + " Doe"
	fmt.Println("My full name is " + fullName + ".")

	// string concatenation with numbers
	age := 30
	fmt.Println("I am " + fmt.Sprint(age) + " years old.")
	fmt.Println("I am", age, "years old.")
}