package main

import "fmt"

type Book struct{
	title string
	year int
}

func main(){
	var age int
	var name string
	var isLegal bool
	var price []float64
	var b Book

	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(isLegal)
	fmt.Println(price)
	fmt.Println(b)
}