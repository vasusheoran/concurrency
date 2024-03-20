package readerwriter

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

func TestNewReaderWriter(t *testing.T) {

	var rCount, wCount int
	rc := make(chan int)
	wc := make(chan int)

	wg := &sync.WaitGroup{}
	rw := NewReaderWriter(5, wg, rc, wc)
	readers := 10
	writers := 2

	go func() {
		for {
			select {
			case n := <-rc:
				rCount += n
			case n := <-wc:
				wCount += n
			}

			fmt.Printf("%s%s\n", strings.Repeat("R", rCount), strings.Repeat("W", wCount))
		}
	}()

	wg.Add(readers + writers)
	for i := 0; i < readers; i++ {
		go fmt.Println(rw.Read())
	}

	for i := 0; i < writers; i++ {
		go rw.Write(i)
	}

	wg.Wait()
	//close(rc)
	//close(wc)
}
