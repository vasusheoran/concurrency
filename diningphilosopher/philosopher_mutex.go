package diningphilosopher

import (
	"math/rand"
	"sync"
	"time"
)

type PickLeftFork func(pID int)
type PickRightFork func(pID int)
type PutLeftFork func(pID int)
type PutRightFork func(pID int)
type Eat func(pID int)

type DiningPhilosophers interface {
	WantsToEat(pID int, pickLeftFork PickLeftFork, pickRightFork PickRightFork, putLeftFork PutLeftFork, putRightFork PutRightFork, eat Eat)
	GetEatingOrder() []int
}

type diningPhilosopherWithMutex struct {
	forks []sync.Mutex
	data  []int
	wg    *sync.WaitGroup
}

func (d *diningPhilosopherWithMutex) GetEatingOrder() []int {
	return d.data
}

func (d *diningPhilosopherWithMutex) canPick(left, right int) {
	d.forks[left].Lock()
	defer d.forks[left].Unlock()
}

func (d *diningPhilosopherWithMutex) WantsToEat(pID int, pickLeftFork PickLeftFork, pickRightFork PickRightFork, putLeftFork PutLeftFork, putRightFork PutRightFork, eat Eat) {
	defer d.wg.Done()
	left := pID
	right := (pID + 1) % 5

	isLeftUnlocked := true
	isRightUnlocked := true

	defer func() {

		if isLeftUnlocked == false {
			d.forks[left].Unlock()
			putLeftFork(pID)
		}
	}()

	defer func() {
		if isRightUnlocked == false {
			d.forks[right].Unlock()
			putRightFork(pID)
		}
	}()

	// wait to pick left fork
	for {
		isLeftUnlocked = false
		d.forks[left].Lock()
		pickLeftFork(pID)

		isRightUnlocked = false
		if d.forks[right].TryLock() {
			pickRightFork(pID)
			break
		}

		d.forks[left].Unlock()
		isLeftUnlocked = true
		putLeftFork(pID)

		// Wait for random time before picking fork again
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(200)))
	}

	// Philosopher has both left and right fork
	d.data = append(d.data, pID)
	eat(pID)

	putRightFork(pID)
	d.forks[right].Unlock()
	isRightUnlocked = true

	putLeftFork(pID)
	d.forks[left].Unlock()
	isLeftUnlocked = true
}

func NewDiningPhilosophersWithMutex(wg *sync.WaitGroup) DiningPhilosophers {
	return &diningPhilosopherWithMutex{
		forks: make([]sync.Mutex, 5),
		wg:    wg,
		data:  make([]int, 0),
	}
}
