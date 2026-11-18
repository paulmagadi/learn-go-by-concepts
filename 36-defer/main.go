package main

import "fmt"

func saveReport() {
	fmt.Println("Saving report")
	defer fmt.Println("Closing resources")
	defer fmt.Println("Cleaning temporary data")
	fmt.Println("Writing data")
}

func deferHalfNum() {
	for i := 0; i <= 7; i++ {
		if i > 3 {
			defer fmt.Println(i)
		} else {
			fmt.Println(i)
		}
	}
}


func reverseLetters() {
	for i := 'A'; i <= 'E'; i++ {
		defer fmt.Println(string(i))
	}
}

func main() {
	saveReport()
	deferHalfNum()
	reverseLetters()
}
