package main

import "fmt"

type Student struct {
	ID     int
	Name   string
	Age    int
	GPA    float64
	Active bool
}

func main() {
	studentMap := map[int]Student{
		1: {ID: 1, Name: "Alice", Age: 20, GPA: 3.8, Active: true},
		2: {ID: 5, Name: "Jane", Age: 21, GPA: 3.9, Active: true},
		3: {ID: 2, Name: "Mills", Age: 19, GPA: 3.7, Active: false},
	}

	fmt.Println(studentMap)
	fmt.Println(studentMap[2])
	fmt.Println(studentMap[0].Name)

	for _, s := range studentMap {
		fmt.Println(s)
	}

	for i, s := range studentMap {
		fmt.Println("Index", i, "is", s.Name)
	}
}