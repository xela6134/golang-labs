package main

import "fmt"

func main() {
	// Arrays: var arr = [3]int{1, 2, 3}
	// Fixed-length collection of elements of the same type
	// The length is part of its type, [5]int and [10]int are different types
	// Stored directly as a value (copied when assigned)

	// Fixed size array
	var arr [5]int
	arr[0] = 10
	arr[1] = 20
	fmt.Println("Array:", arr)
	fmt.Println("Length:", len(arr))

	var arrr [5]string
	arrr[0] = "asd"
	arrr[1] = "qwe"
	fmt.Println("Array:", arrr)
	fmt.Println("Length:", len(arrr))
	fmt.Println("")

	// Array literal
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array literal:", arr2)

	// Auto-length array
	arr3 := [...]string{"a", "b", "c"}
	fmt.Println("Auto-length array:", arr3)

	// Slices: var slice = []int{1, 2, 3}
	// Dynamic, flexible view (or window) into an underlying array
	// Changeable length
	// Descriptor that contains
	// 	- a pointer to the underlying array
	// 	- a length (len)
	// 	- a capacity (cap)
	slice := []int{10, 20, 30}
	fmt.Println("\nSlice:", slice)
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))

	// Slicing an array
	full := [5]int{1, 2, 3, 4, 5}
	sub := full[1:4] // elements at index 1, 2, 3
	fmt.Println("\nFull array:", full)
	fmt.Println("Sliced portion (1:4):", sub)

	// Modifying the slice modifies the original array
	sub[0] = 999
	fmt.Println("\nAfter modifying slice:")
	fmt.Println("Sub:", sub)
	fmt.Println("Full:", full)

	// Appending to slices
	nums := []int{1, 2, 3}
	nums = append(nums, 4, 5)
	fmt.Println("\nAfter append:", nums)
	fmt.Println("New length:", len(nums))
	fmt.Println("New capacity:", cap(nums))

	// Append another slice using ...
	more := []int{6, 7, 8}
	nums = append(nums, more...)
	fmt.Println("\nAfter another append:", nums)
	fmt.Println("New new length:", len(nums))
	fmt.Println("New new capacity:", cap(nums))

	// Copying slices (does not share memory)
	src := []int{10, 20, 30}
	dest := make([]int, len(src))
	copy(dest, src)
	dest[0] = 999
	fmt.Println("\nSource:", src)
	fmt.Println("Copied (modified) dest:", dest)

	// Using make() to control capacity
	made := make([]int, 3, 10)
	fmt.Println("\nMade slice:", made)
	fmt.Println("Length:", len(made), "| Capacity:", cap(made))

	// Slicing slices
	letters := []string{"a", "b", "c", "d", "e"}
	part1 := letters[:3]  // first 3 elements
	part2 := letters[2:]  // from index 2 to end
	part3 := letters[1:4] // from index 1 to 3
	fmt.Println("\nOriginal letters:", letters)
	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
	fmt.Println("part3:", part3)
	fmt.Println("")

	// Capacity expansion demonstration
	dynamic := []int{}
	for i := 1; i <= 8; i++ {
		dynamic = append(dynamic, i)
		fmt.Printf("After appending %d: len=%d cap=%d\n", i, len(dynamic), cap(dynamic))
	}
}
