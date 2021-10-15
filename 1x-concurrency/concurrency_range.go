package main

import (
	"fmt"
)

func emit(c chan int) {
	c <- 1
	c <- 2
	c <- 3
	close(c)
}

func main() {
	p := fmt.Println
	// pf := fmt.Printf

	c1 := make(chan int)
	c2 := make(chan int)

	go emit(c1)
	go emit(c2)

	for i := range c1 {
		j := <-c2
		p(i, j)
	}
}
