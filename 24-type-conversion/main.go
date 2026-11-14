package main

import "fmt"

func main() {
	var a int = 10
	var b int = 3
	div := a / b
	fmt.Println("Dividing two intergers will produces an int:", div)

	c := float64(a)
	d := float64(b)
	fmt.Printf("Float of a is %f and b is %f\n", c, d)
	
	fmt.Println("The operation 10/3 produces decimal points:", c/d)
}

