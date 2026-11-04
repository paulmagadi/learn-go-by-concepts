package main

import "fmt"

func main() {
	favLanguage := "Go"
	years := 2
	morningCoder := true
	confidenceLevel := 0.8

	fmt.Printf("My favorite programming language is %s\n", favLanguage)
	fmt.Printf("I have been coding for %d years\n", years)
	fmt.Printf("Am I a morning coder? %t\n", morningCoder)
	fmt.Printf("My confidence level in coding is %.1f\n", confidenceLevel)
}