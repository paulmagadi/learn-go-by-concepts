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
	students := []Student{
		{ID: 1, Name: "Alice", Age: 20, GPA: 3.8, Active: true},
		{ID: 2, Name: "Bob", Age: 22, GPA: 3.4, Active: false},
	}

	fmt.Println(students)
	fmt.Println(students[1])
	fmt.Println(students[1].Name)

	for _, student := range students {
		fmt.Println(student.Name)
	}
}