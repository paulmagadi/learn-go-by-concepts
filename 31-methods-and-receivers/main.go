package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() int {
	return (r.Width + r.Height) * 2
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println("The area is:", rect.Area())
	fmt.Printf("The perimeter is %dcm.\n", rect.Perimeter())
}