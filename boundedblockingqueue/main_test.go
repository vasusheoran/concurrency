package boundedblockingqueue

import (
	"fmt"
	"testing"
	"time"
)

func Test_chan(t *testing.T) {
	w := make(chan bool)
	ch := make(chan int, 2)

	go func() {
		ch <- 10
		fmt.Println("added 10")
		ch <- 10
		fmt.Println("added 10")
		ch <- 20
		fmt.Println(<-ch)
		fmt.Println("added 20")
		time.Sleep(5 * time.Second)
	}()

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		ch <- 30
		w <- true
	}()

	<-w
}

func TestQueue(t *testing.T) {
	q := New(2)

	ch := make(chan bool)

	go func() {
		q.Enqueue(10)
		fmt.Println("added 10")
		q.Enqueue(10)
		fmt.Println("added 10")
		q.Enqueue(20)
		fmt.Println("added 20")
	}()

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println(q.Dequeue())
		fmt.Println(q.Dequeue())
		fmt.Println(q.Dequeue())
		ch <- true
	}()

	<-ch
}
