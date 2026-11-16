package main

import (
	"33-creating-packages/messages"
	"fmt"
)

func main() {
	msg := messages.WelcomeMessage("Alice")
	fmt.Println(msg)

	fmt.Println(messages.Greetings("Jane"))
}
