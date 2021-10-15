// equalbinarytree.go
// find if two binary trees are equivalent, using concurrency

package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Walker(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}
func Same(t1, t2 *tree.Tree) bool {
	// p := fmt.Println

	c1 := make(chan int)
	c2 := make(chan int)

	go Walker(t1, c1)
	go Walker(t2, c2)

	for x := range c1 {
		y := <-c2

		if x != y {
			return false
		}

		// p("x", x)
		// p("y", y)
	}

	return true

}

func main() {
	p := fmt.Println
	// pf := fmt.Printf

	t1 := tree.New(1)
	t2 := tree.New(1)

	p(t1)
	p(t2)

	p(Same(t1, t2))
}
