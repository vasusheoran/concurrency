package ratelimiter

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNewRateLimiter(t *testing.T) {
	wg := &sync.WaitGroup{}

	rl := NewRateLimiter(wg, 5, 1)

	fn := func(ID int) {
		fmt.Printf("Executing task - %d\n", ID)
	}

	for i := 0; i < 8; i++ {
		go rl.RateLimiter(fn, i)
	}

	for i := 9; i < 10; i++ {
		go rl.RateLimiter(fn, i)
		time.Sleep(time.Duration(2) * time.Second)
	}

	wg.Wait()

	time.Sleep(1 * time.Second)
}
