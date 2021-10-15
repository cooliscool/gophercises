// purpose
// to divide and conquer a summing task
package sub2

import (
	"fmt"
)

func sum(s []int, c chan int) {

	sum := 0
	for _, item := range s {
		sum += item
	}

	c <- sum
}

func main() {
	p := fmt.Println

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	c := make(chan int) // make is used for maps, slices and chans

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c

	p(x + y)
}
