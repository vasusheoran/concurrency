package diningphilosopher

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func pickLeftFork(pID int) {
	fmt.Printf("Left fork picked by %d\n", pID)
}

func pickRightFork(pID int) {
	fmt.Printf("Right fork picked by %d\n", pID)
}

func putLeftFork(pID int) {
	fmt.Printf("Left fork put by %d\n", pID)
}

func putRightFork(pID int) {
	fmt.Printf("Right fork put by %d\n", pID)
}

func eat(pID int) {
	fmt.Printf("Philosopher %d is eating\n", pID)
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(200))) // Eating takes a random amount of time
	fmt.Printf("Philosopher %d is done eating\n", pID)
}

const (
	attemptToEat = 2
)

func TestNewDiningPhilosopherWithSignal(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(5 * attemptToEat)
	ph := NewDiningPhilosopherWithSignal(wg)

	go func() {
		for i := 0; i < attemptToEat; i++ {
			go ph.WantsToEat(0, pickLeftFork, pickRightFork, putLeftFork, putRightFork, eat)
			go ph.WantsToEat(1, pickLeftFork, pickRightFork, putLeftFork, putRightFork, eat)
			go ph.WantsToEat(2, pickLeftFork, pickRightFork, putLeftFork, putRightFork, eat)
			go ph.WantsToEat(3, pickLeftFork, pickRightFork, putLeftFork, putRightFork, eat)
			go ph.WantsToEat(4, pickLeftFork, pickRightFork, putLeftFork, putRightFork, eat)
		}
	}()

	wg.Wait()

	fmt.Printf("\n\n%#v", ph.GetEatingOrder())
}

func TestNewDiningPhilosophersWithPanicHandling(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(5 * attemptToEat)
	ph := NewDiningPhilosophersWithMutex(wg)

	go func() {
		for i := 0; i < attemptToEat; i++ {
			go ph.WantsToEat(0, pickLeftFork, pickRightFork, putLeftFork, putRightFork, eat)
			go ph.WantsToEat(1, pickLeftFork, pickRightFork, putLeftFork, putRightFork, eat)
			go ph.WantsToEat(2, pickLeftFork, pickRightFork, putLeftFork, putRightFork, eat)
			go ph.WantsToEat(3, pickLeftFork, pickRightFork, putLeftFork, putRightFork, eat)
			go ph.WantsToEat(4, pickLeftFork, pickRightFork, putLeftFork, putRightFork, eat)
		}
	}()

	wg.Wait()

	fmt.Printf("\n\n%#v", ph.GetEatingOrder())
}

func TestNewDiningPhilosophers_NoWaitBeforeAttempts(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(5 * attemptToEat)
	ph := NewDiningPhilosopherWithSignal(wg)

	go func() {
		for i := 0; i < attemptToEat; i++ {
			go ph.WantsToEat(0, pickLeftFork, pickRightFork, putLeftFork, putRightFork, eat)
			go ph.WantsToEat(1, pickLeftFork, pickRightFork, putLeftFork, putRightFork, eat)
			go ph.WantsToEat(2, pickLeftFork, pickRightFork, putLeftFork, putRightFork, eat)
			go ph.WantsToEat(3, pickLeftFork, pickRightFork, putLeftFork, putRightFork, eat)
			go ph.WantsToEat(4, pickLeftFork, pickRightFork, putLeftFork, putRightFork, eat)
		}
	}()

	wg.Wait()

	fmt.Printf("\n\n%#v", ph.GetEatingOrder())

}

func TestNewDiningPhilosophers_WaitBeforeAttempts(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(5 * attemptToEat)
	ph := NewDiningPhilosopherWithSignal(wg)

	go func() {
		for i := 0; i < attemptToEat; i++ {
			ph.WantsToEat(0, pickLeftFork, pickRightFork, putLeftFork, putRightFork, eat)
		}
	}()

	go func() {
		for i := 0; i < attemptToEat; i++ {
			ph.WantsToEat(1, pickLeftFork, pickRightFork, putLeftFork, putRightFork, eat)
		}
	}()

	go func() {
		for i := 0; i < attemptToEat; i++ {
			ph.WantsToEat(2, pickLeftFork, pickRightFork, putLeftFork, putRightFork, eat)
		}
	}()

	go func() {
		for i := 0; i < attemptToEat; i++ {
			ph.WantsToEat(3, pickLeftFork, pickRightFork, putLeftFork, putRightFork, eat)
		}
	}()

	go func() {
		for i := 0; i < attemptToEat; i++ {
			ph.WantsToEat(4, pickLeftFork, pickRightFork, putLeftFork, putRightFork, eat)
		}
	}()

	wg.Wait()

	fmt.Printf("\n\n%#v", ph.GetEatingOrder())

}
