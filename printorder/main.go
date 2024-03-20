package main

import (
	"fmt"
	"sync"
)

var f, s chan bool
var wg *sync.WaitGroup

func first() {
	defer wg.Done()
	fmt.Print("first")
	f <- true
}

func second() {
	defer wg.Done()
	<-f
	fmt.Print("second")
	s <- true
}

func third() {
	defer wg.Done()
	<-s
	fmt.Print("third")
}

func main() {
	wg = &sync.WaitGroup{}
	f = make(chan bool)
	s = make(chan bool)

	wg.Add(3)

	go first()
	go third()
	go second()

	wg.Wait()

}
