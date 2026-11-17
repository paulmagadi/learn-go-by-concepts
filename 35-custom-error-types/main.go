package main

import "fmt"

type ValidationError struct {
	ErrorField string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("Validation failed for field: %s", e.ErrorField)
}

func validateName(name string) error {
	if name == "" {
		return ValidationError{ErrorField: "name"}
	}
	return nil
}



//

type NotFoundError struct {
    Resource string
}

func (e NotFoundError) Error() string {
    return fmt.Sprintf("%s not found", e.Resource)
}

func findUser(id int) error {
    return NotFoundError{Resource: "user"}
}

func main() {
	err := validateName("John")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Name is valid")


	err2 := findUser(7)
    if err2 != nil {
        fmt.Println("Error:", err2)
        return
    }
    fmt.Println("User found")
}