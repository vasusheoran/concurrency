package main

import (
	"fmt"
	"sync"
)

var wg *sync.WaitGroup
var n int

var ch chan bool

func foo() {
	defer wg.Done()
	for i := 0; i < n; i++ {
		fmt.Println("foo")
		ch <- true // allow 1 process to enter cs
		<-ch       // wait till allowed
	}
}

// q ---

func bar() {
	defer wg.Done()
	for i := 0; i < n; i++ {
		<-ch // wait till allowed
		fmt.Println("bar")
		ch <- true // allow 1 process to enter cs
	}
}

func main() {

	wg = &sync.WaitGroup{}
	ch = make(chan bool)

	n = 5

	wg.Add(2)

	go foo()
	go bar()

	wg.Wait()
}
