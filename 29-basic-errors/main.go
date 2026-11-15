package main

import (
	"errors"
	"fmt"
	"math"
)

func squareRoot(n int) (float64, error) {
	if n < 0 {
		return 0, errors.New("number cannot be negative")
	}
	return math.Sqrt(float64(n)), nil
}


func main(){
	ans, error := squareRoot(25)

	if error != nil {
		fmt.Println("Error:", error)
	} else {
		fmt.Println("The square root of", 25, "is", ans)
	}

	ans, error = squareRoot(-25)

	if error != nil {
		fmt.Println("Error:", error)
	} else {
		fmt.Println("The square root of", -25, "is", ans)
	}
}