// bufferedChannel.go
// channel wont emit until buffer overflows ? Nop!
package sub

import (
	"fmt"
	"time"
)

func writeChan(c chan int, v int) {
	time.Sleep(10 * time.Millisecond)
	c <- v
}

func main() {
	p := fmt.Println
	pf := fmt.Printf

	c := make(chan int, 2) // means only two can be put into this buffer at a time.
	// if you need to put more, you have to clear the buffer.

	// buffered also means that channel will not get blocked for a read/write to happen.
	go writeChan(c, 1)
	go writeChan(c, 2)
	go writeChan(c, 3)
	go writeChan(c, 4)

	for {
		select {
		case x := <-c:
			p(x)
		default:
			pf("yo")
			time.Sleep(100 * time.Millisecond)
		}

	}

}
