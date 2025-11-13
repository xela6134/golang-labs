package main

import (
	"fmt"
	"project8/mathutil"
)

func main() {
	sum := mathutil.Add(3, 4)
	product := mathutil.Multiply(2, 5)

	fmt.Println("Sum:", sum)
	fmt.Println("Product:", product)
}
