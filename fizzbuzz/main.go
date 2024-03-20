package fizzbuzz

import (
	"fmt"
	"sync"
)

type FizzBuzz interface {
	Fizz()
	Buzz()
	FizzBuzz()
	Number()
	Data() []string
}

type fizzBuzzCond struct {
	n    int
	i    int
	data []string
	m    sync.Mutex
	c    sync.Cond
	wg   *sync.WaitGroup
}

func (f *fizzBuzzCond) Data() []string {
	return f.data
}

func (f *fizzBuzzCond) Fizz() {
	f.m.Lock()
	defer f.m.Unlock()

	for f.i%3 != 0 || (f.i%3 == 0 && f.i%5 == 0) {
		f.c.Wait()
	}

	f.data = append(f.data, "Fizz")
	fmt.Print("Fizz", ",")
	f.i++
	f.c.Broadcast()
	f.wg.Done()
}

func (f *fizzBuzzCond) Buzz() {
	f.m.Lock()
	defer f.m.Unlock()

	for f.i%5 != 0 || (f.i%3 == 0 && f.i%5 == 0) {
		f.c.Wait()
	}

	fmt.Print("Buzz", ",")
	f.data = append(f.data, "Buzz")
	f.i++
	f.c.Broadcast()
	f.wg.Done()
}

func (f *fizzBuzzCond) FizzBuzz() {
	f.m.Lock()
	defer f.m.Unlock()

	for {
		if f.i%3 == 0 && f.i%5 == 0 {
			break
		} else {
			f.c.Wait()
		}
	}

	fmt.Print("FizzBuzz")
	f.data = append(f.data, "FizzBuzz")
	f.i++
	f.c.Broadcast()
	f.wg.Done()
}

func (f *fizzBuzzCond) Number() {
	f.m.Lock()
	defer f.m.Unlock()

	for f.i%3 == 0 || f.i%5 == 0 {
		f.c.Wait()
	}

	fmt.Print(f.i, ",")
	f.data = append(f.data, fmt.Sprint(f.i))
	f.i++
	f.c.Broadcast()
	f.wg.Done()
}

func NewFizzBuzz(n int, wg *sync.WaitGroup) FizzBuzz {
	f := &fizzBuzzCond{
		n:    n,
		i:    1,
		data: make([]string, 0),
		m:    sync.Mutex{},
		wg:   wg,
	}
	f.c = sync.Cond{L: &f.m}
	return f
}
