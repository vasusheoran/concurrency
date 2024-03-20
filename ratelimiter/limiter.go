package ratelimiter

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter interface {
	RateLimiter(func(ID int), int)
}

type limiter struct {
	wg          *sync.WaitGroup
	m           *sync.Mutex
	attempts    int
	maxAttempts int
	duration    int
	nextReset   time.Time
}

func NewRateLimiter(wg *sync.WaitGroup, maxAttempts, duration int) RateLimiter {
	if wg == nil {
		wg = &sync.WaitGroup{}
	}

	return &limiter{
		wg:          wg,
		m:           &sync.Mutex{},
		attempts:    0,
		maxAttempts: maxAttempts,
		duration:    duration,
		nextReset:   time.Now(),
	}
}

func (rl *limiter) RateLimiter(fn func(int), ID int) {
	rl.m.Lock()
	defer rl.m.Unlock()

	// if current attempt is first then start timer
	if rl.attempts == 0 {
		rl.nextReset = time.Now().Add(time.Duration(rl.duration) * time.Second)
		go rl.watchTimer(time.NewTimer(time.Duration(rl.duration) * time.Second))
	}

	if rl.attempts >= rl.maxAttempts {
		fmt.Printf("Current attempts %d greater than max attempts %d. Please wait for %.0f seconds.\n",
			rl.attempts, rl.maxAttempts, rl.nextReset.Sub(time.Now()).Seconds())
		rl.attempts++
		return
	}

	// Execute task
	fn(ID)

	rl.attempts++
}

func (rl *limiter) watchTimer(timer *time.Timer) {
	defer rl.wg.Done()
	rl.wg.Add(1)

	select {
	case <-timer.C:
		rl.m.Lock()
		defer rl.m.Unlock()

		fmt.Println("Resetting attempts")

		rl.attempts = 0
	}
}
