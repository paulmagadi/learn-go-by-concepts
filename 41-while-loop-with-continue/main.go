package main

import "fmt"

func main() {
	i := 1

	for i <= 6 {
		if i%2 != 0 {
			i++
			continue
		}
		fmt.Println("Even", i)
		i++
	}
}