package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Printf("Hi %v\n", i)
	}

	for i := range 3 {
		fmt.Printf("Hi %v\n", i)
	}

	names := []string{"Alex Amazon", "Natalie Netflix", "George Google"}

	for index, name := range names {
		fmt.Printf("%v: %v\n", index, name)
		split := strings.Fields(name)
		firstName := split[0]
		lastName := split[1]
		fmt.Printf("Firstname: %v, Lastname: %v\n", firstName, lastName)
	}
}
