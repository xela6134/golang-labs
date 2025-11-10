package main

import "fmt"

func main() {
	var name = "Alex"
	age := 200
	const city = "Sydney"

	fmt.Println("Hello World")
	fmt.Println(name, "is", age, "years old")

	name = "Joe"

	fmt.Printf("%v lives in %v and is %v years old\n", name, city, age)
}
