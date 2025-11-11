package main

import "fmt"

// Basic function with no parameters and no return
func sayHello() {
	fmt.Println("Hello from Go!")
}

// Function with parameters
func add(a int, b int) int {
	return a + b
}

// Function with multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Named return values
func rectangleStats(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	// return area, perimeter
	return // implicit return of named values
}

// Variadic function (accepts any number of arguments)
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Function that takes another function as a parameter
func applyTwice(f func(int) int, x int) int {
	return f(f(x))
}

func main() {
	// Basic call
	sayHello()

	// Parameters and return
	result := add(3, 4)
	fmt.Println("3 + 4 =", result)

	// Multiple returns
	quotient, err := divide(10, 2)
	fmt.Printf("\nQuotient: %v, Error: %v\n", quotient, err)

	// Dividing by zero
	quotient, err = divide(5, 0)
	fmt.Printf("Quotient: %v, Error: %v\n", quotient, err)

	// Named return values
	area, perimeter := rectangleStats(5, 3)
	fmt.Printf("Rectangle area: %v, perimeter: %v\n", area, perimeter)

	// Variadic function
	total := sum(1, 2, 3, 4, 5)
	fmt.Println("\nSum of numbers:", total)

	// Passing function as argument
	double := func(x int) int { return x * 2 }
	fmt.Println("\nApplyTwice(double, 5):", applyTwice(double, 5))
}
