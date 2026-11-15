package main

import "fmt"

func minMax(a, b int) (int, int, bool) {
	if a == b {
		return  a, b, false
	} 
	
	if a < b {
		return a, b, true
	}

	return a, b, true
}

func main(){
	min, max, ok := minMax(2, 3)

	if !ok {
        fmt.Println("The numbers are equal")
        return
    }

	fmt.Println(min)
	fmt.Println(max)
	fmt.Println(ok)

}