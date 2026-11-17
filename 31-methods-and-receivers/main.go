package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

// Methods can be defined for either pointer or value receiver types. 
// Here’s an example of a value receiver.
func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() int {
	return (r.Width + r.Height) * 2
}

// This area method has a receiver type of *Rectangle.
func (r *Rectangle) Area1() int {
    return r.Width * r.Height
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println("The area is:", rect.Area())
	fmt.Printf("The perimeter is %dcm.\n", rect.Perimeter())


	// Go automatically handles conversion between values and pointers for method calls. 
	// You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.
	rp := &rect
    fmt.Println("area: ", rp.Area1())
    fmt.Println("perim:", rp.Perimeter())
}



