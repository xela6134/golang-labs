package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumber(id int, wg *sync.WaitGroup) {
	defer wg.Done() // mark this goroutine as finished when function returns

	for i := 0; i < 5; i++ {
		fmt.Printf("Goroutine %d → %d\n", id, i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup

	// We will start 3 goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)              // tell WaitGroup we are adding one goroutine
		go printNumber(i, &wg) // start the goroutine
	}

	fmt.Println("Main goroutine is waiting")
	wg.Wait() // wait for all goroutines to finish
	fmt.Println("All goroutines finished")
}
