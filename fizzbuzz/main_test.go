package fizzbuzz

import (
	"fmt"
	"sync"
	"testing"
)

func TestNewFizzBuzzCond(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(15)
	f := NewFizzBuzz(15, wg)

	go f.Fizz()
	go f.Fizz()
	go f.Fizz()
	go f.Fizz()

	go f.Buzz()
	go f.Buzz()

	go f.FizzBuzz()

	go f.Number()
	go f.Number()
	go f.Number()
	go f.Number()
	go f.Number()
	go f.Number()
	go f.Number()
	go f.Number()

	wg.Wait()

	fmt.Printf("%#v\n", f.Data())
}
