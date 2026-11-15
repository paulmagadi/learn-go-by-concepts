package main

import "fmt"

func userInfo() (int, string) {
	return 34, "Jane"
}

func main(){
	_, name := userInfo()
	fmt.Println(name)

	age, _ := userInfo()
	fmt.Println(age)

	// additional example
	ages := []int{21, 23, 24, 25}

	for _, age := range ages {
		fmt.Println(age)
	}

	for i, _ := range ages {
		fmt.Println(i)
	}
}