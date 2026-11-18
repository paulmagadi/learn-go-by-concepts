package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	first := uuid.New()
	second := uuid.New()

	fmt.Println("First UUID:", first)
	fmt.Println("Second UUID:", second)
}
