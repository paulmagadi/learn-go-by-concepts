package main

import "fmt"

type Book struct {
	Title string
	Author string
	Year int
}

func main() {
	book1 := Book{
		Title: "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Year: 1925, 
	}

	fmt.Println("Book Title:", book1.Title)
	fmt.Println("Book Author:", book1.Author)
	fmt.Println("Book Year:", book1.Year)
}