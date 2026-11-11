package main

import "fmt"

func main() {
	countries := map[string]string{
		"US": "United States",
		"DE": "Germany",
		"FR": "France",
	}
	fmt.Println("Original map:", countries)

	// Adding a new key-value pair to the map
	countries["JP"] = "Japan"
	fmt.Println("After adding Japan:", countries)

	// Deleting a key-value pair from the map
	delete(countries, "DE")
	fmt.Println("After deleting Germany:", countries)

	// Updating a value in the map
	countries["US"] = "USA"
	fmt.Println("After updating US:", countries)

	// Checking if a key exists in the map
	if country, exists := countries["FR"]; exists {
		fmt.Println("Country found:", country)
	} else {
		fmt.Println("Country not found")
	}

	if country, exists := countries["KE"]; exists {
		fmt.Println("Country found:", country)
	} else {
		fmt.Println("Country not found")
	}

	// Iterating over the map
	for code, country := range countries {
		fmt.Printf("Code: %s, Country: %s\n", code, country)
	}

	for code := range countries {
		fmt.Println("Country code:", code)
	}

	for _, country := range countries {
		fmt.Println("Country name:", country)
	}

	// Getting the length of the map
	fmt.Println("Number of countries in the map:", len(countries))

	// Clearing the map by reinitializing it
	countries = make(map[string]string)
	fmt.Println("Map after clearing:", countries)

	// Demonstrating that the map is empty
	if len(countries) == 0 {
		fmt.Println("The map is now empty.")
	} else {
		fmt.Println("The map is not empty.")
	}

	// Demonstrating that maps are reference types
	originalMap := map[string]string{
		"IN": "India",
		"CN": "China",
	}
	fmt.Println("Original map:", originalMap)

	// Creating a new reference to the original map
	newReference := originalMap
	newReference["IN"] = "Bharat"
	fmt.Println("After modifying new reference:", newReference)
	fmt.Println("Original map after modification:", originalMap)

	// Creating empty maps using make
	emptyMap1 := make(map[string]int)
	emptyMap2 := make(map[string]int, 10)

	fmt.Println("Empty map 1:", emptyMap1)
	fmt.Println("Empty map 2 with initial capacity of 10:", emptyMap2)

	// Adding elements to the empty maps
	emptyMap1["A"] = 1
	emptyMap2["B"] = 2

	fmt.Println("Empty map 1 after adding an element:", emptyMap1)
	fmt.Println("Empty map 2 after adding an element:", emptyMap2)

	// Demonstrating that maps can have different types for keys and values
	mixedMap := map[int]string{
		1: "One",
		2: "Two",
	}
	fmt.Println("Mixed map:", mixedMap)

	// Clearing keys/values from the mixed map using builtin clear function

	clear(mixedMap)
	fmt.Println("Mixed map after clearing:", mixedMap)

}