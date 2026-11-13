package main

import "fmt"

func main() {
	result := double(4)
	fmt.Println(result)

	triple(5)

	fmt.Println(add(3, 7))

	name := name()
	fmt.Println(name)

	greet("Alice")

	fullName := fullName("John", "Mark")
	fmt.Println(fullName)
}

func double(n int) int {
	return n * 2
}

func triple(n int) {
	fmt.Println(n * 3)
}

func add(a int, b int) int {
	return a + b
}

func name() string {
	return "John Doe"
}

func greet(name string) {
	fmt.Println("Hello, " + name + "!")
}

func fullName(firstName string, lastName string) string {
	return firstName + " " + lastName
}