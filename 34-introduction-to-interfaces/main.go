package main

import "fmt"

type Logger interface {
    Log(message string)
}

type ConsoleLogger struct{}

func (ConsoleLogger) Log(message string) {
    fmt.Println(message)
}

type PrefixedLogger struct {
    Prefix string
}

func (p PrefixedLogger) Log(message string) {
    fmt.Println(p.Prefix, message)
}

func runLog(l Logger, msg string) {
    l.Log(msg)
}

func main() {
    plain := ConsoleLogger{}
    info := PrefixedLogger{Prefix: "[INFO]"}

    runLog(plain, "Starting application")
    runLog(info, "User logged in")
}

// import (
// 	"fmt"
// 	"math"
// )

// type Shape interface {
// 	Area() float64
// }

// type Circle struct {
// 	Radius float64
// }

// func (c Circle) Area() float64 {
// 	return math.Pi * c.Radius * c.Radius
// }

// type Rectangle struct {
// 	Width  float64
// 	Height float64
// }

// func (r Rectangle) Area()  float64 {
// 	return r.Width * r.Height
// }

// func printArea(s Shape) {
// 	fmt.Println(s.Area())
// }

// func main() {
// 	c := Circle{Radius: 3}
// 	r := Rectangle{Width: 5, Height: 4}

// 	printArea(c)
// 	printArea(r)

// 	fmt.Println(r.Perimeter())
// }

// func (r Rectangle) Perimeter() float64 {
// 	return (r.Width + r.Height) * 2
// }
