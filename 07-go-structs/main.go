package main

import "fmt"

type Book struct {
	Title string
	Author string
	Year int
	Price float64
}

func main() {
	book1 := Book{
		Title: "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Year: 1925, 
		Price: 10.99,
	}

	book2 := Book{
		Title: "To Kill a Mockingbird",
		Author: "Harper Lee",
		Year: 1960,
		Price: 12.99,
	}
	

	fmt.Println("Book Title:", book1.Title)
	fmt.Println("Book Author:", book1.Author)
	fmt.Println("Book Year:", book1.Year)
	fmt.Println("Book Price:", book1.Price)

	fmt.Println("Book Title:", book2.Title)
	fmt.Println("Book Author:", book2.Author)
	fmt.Println("Book Year:", book2.Year)
	fmt.Println("Book Price:", book2.Price)
}