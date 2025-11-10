package main

import "fmt"

func main() {
	// Basic if else
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is 5 or less")
	}

	// else if chain
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 75 {
		fmt.Println("Grade: B")
	} else if score >= 60 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// Inline variable declaration in if
	if n := 7; n%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}
	// Note: n is only accessible inside this if-block

	// Boolean values
	isRaining := false
	isWeekend := true

	if isRaining && isWeekend {
		fmt.Println("Stay home and relax!")
	} else if isRaining && !isWeekend {
		fmt.Println("Bring an umbrella to work!")
	} else if !isRaining && isWeekend {
		fmt.Println("Perfect day to go out!")
	} else {
		fmt.Println("Normal work day.")
	}

	// Comparison operators
	a, b := 5, 10
	fmt.Println("\nComparison results:")
	fmt.Println("a == b:", a == b)
	fmt.Println("a != b:", a != b)
	fmt.Println("a < b:", a < b)
	fmt.Println("a <= b:", a <= b)
	fmt.Println("a > b:", a > b)
	fmt.Println("a >= b:", a >= b)

	// Logical operators
	p := true
	q := false
	fmt.Println("\nLogical operations:")
	fmt.Println("p && q:", p && q) // AND
	fmt.Println("p || q:", p || q) // OR
	fmt.Println("!p:", !p)         // NOT
}
