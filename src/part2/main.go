package main

import "fmt"

func main() {
	var name = "Alex"
	age := 200
	const city string = "Sydney"
	cost := 3.1

	fmt.Println("Hello World")
	fmt.Println(name, "is", age, "years old")

	name = "Joe"

	fmt.Printf("%v lives in %v and is %v years old\n", name, city, age)
	fmt.Printf("%T, %T, %T, %T\n", name, age, city, cost)
}
