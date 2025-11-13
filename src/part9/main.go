package main

import "fmt"

func main() {
	// new(T): Allocates memory for type T and returns *T, no internal initialisation
	// make(T, ...): Creates data structure and initialises internal values
	ages := make(map[string]int)

	// Add values
	ages["Alex"] = 99
	ages["Bob"] = 12
	ages["Charlie"] = 30

	// Access values
	fmt.Println("Alex age:", ages["Alex"])

	// Check if a key exists
	age, exists := ages["David"]
	fmt.Printf("age: %v, exists: %v\n", age, exists)
	if exists {
		fmt.Println("David's age:", age)
	} else {
		fmt.Println("David not found")
	}

	// Update a value
	ages["Bob"] = 26

	// Delete a key
	delete(ages, "Charlie")

	// Iterate over map
	fmt.Println("\nAll ages:")
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}
}
