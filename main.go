package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// Channels for streaming individual integers
	super := make(chan int)
	c := make(chan int)

	// Producer: Generates numbers
	go func() {
		genNumber(super)
	}()

	// Consumer 1: Squares the numbers
	go func() {
		for number := range super {
			square(number, c)
		}
		close(c) // Close the channel when all numbers are processed
	}()

	// Consumer 2: Prints the results
	for i := range c {
		fmt.Println(i)
	}

	fmt.Println("done")
}

// Renamed to 'square' to reflect what it actually does
func square(a int, c chan int) {
	c <- a * a
}

func genNumber(ch chan int) {
	for i := 0; i < 10; i++ {
		n := rand.IntN(101)
		ch <- n
	}
	close(ch) // Crucial: Close the channel so the 'range' loop in the receiver knows to stop
}
