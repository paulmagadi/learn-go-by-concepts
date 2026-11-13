package main  

import "fmt"

func main() {

	loggedIn := true
	isStaff := false
	hasPermission := true

	if loggedIn && (isStaff || hasPermission) {
		fmt.Println("Action allowed")
	} else {
		fmt.Println("Action denied")
	}


	// Additional example with multiple conditions
	if loggedIn && isStaff && hasPermission {
		fmt.Println("Staff member with permission")
	} else if loggedIn && isStaff {
		fmt.Println("Staff member without permission")
	} else if loggedIn && hasPermission {
		fmt.Println("User with permission")
	}
}