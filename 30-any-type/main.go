package main

import "fmt"

func logValue(label string, v any) {

	fmt.Println(label+":", v)
	fmt.Println("Type:", fmt.Sprintf("%T", v))
}

func main(){
	logValue("Name", "John")
	logValue("Age", 16)
	logValue("Price", 9.99)

	data := map[string]any{
		"name" : "Alice",
		"age" : 15,
		"price" : 9.99,
	}

	fmt.Println("name Type:", fmt.Sprintf("%T", data["name"]))
	fmt.Println("age Type:", fmt.Sprintf("%T", data["age"]))
	fmt.Println("price Type:", fmt.Sprintf("%T", data["price"]))
}