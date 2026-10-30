package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, number := range numbers {
			double(number, c)
		}
	}()
	for i := range c {
		fmt.Println(i)
	}
	wg.Wait()
	fmt.Println("donee")
}

func double(a int, c chan int) {
	c <- a * a
}
