package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// Create struct using literal
	p1 := Person{
		Name: "Alice",
		Age:  21,
	}

	// Create struct without specifying field names
	p2 := Person{"Bob", 25}

	// Create an empty struct and assign values later
	var p3 Person
	p3.Name = "Charlie"
	p3.Age = 30

	// Create a pointer to a struct
	p4 := &Person{Name: "David", Age: 40}
	p4.Age = 41 // update through pointer

	fmt.Println("Person 1:", p1)
	fmt.Println("Person 2:", p2)
	fmt.Println("Person 3:", p3)
	fmt.Println("Person 4:", *p4)

	// Access individual fields
	fmt.Println("Alice's age is", p1.Age)
}
