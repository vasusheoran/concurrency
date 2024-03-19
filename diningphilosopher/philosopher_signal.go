package diningphilosopher

import "sync"

type fork struct {
	m       sync.Mutex
	c       sync.Cond
	isTaken bool
}

type diningPhilosopherWithSignal struct {
	forks []fork
	data  []int
	wg    *sync.WaitGroup
}

func (d diningPhilosopherWithSignal) WantsToEat(pID int, pickLeftFork PickLeftFork, pickRightFork PickRightFork, putLeftFork PutLeftFork, putRightFork PutRightFork, eat Eat) {
	defer d.wg.Done()
	left := pID
	right := (pID + 1) % 5

	d.forks[left].m.Lock()
	defer d.forks[left].m.Unlock()

	for {
		//pickLeftFork(pID)
		if d.forks[right].m.TryLock() {
			break
		}
		//putLeftFork(pID)
		d.forks[left].c.Wait()
	}

	defer d.forks[right].m.Unlock()
	pickLeftFork(pID)
	pickRightFork(pID)

	eat(pID)
	d.data = append(d.data, pID)

	putRightFork(pID)
	putLeftFork(pID)

	d.forks[right].c.Broadcast()
	d.forks[left].c.Broadcast()
}

func (d diningPhilosopherWithSignal) GetEatingOrder() []int {
	return d.data
}

func NewDiningPhilosopherWithSignal(wg *sync.WaitGroup) DiningPhilosophers {
	d := &diningPhilosopherWithSignal{
		forks: make([]fork, 5),
		wg:    wg,
	}

	for i := 0; i < 5; i++ {
		d.forks[i] = fork{m: sync.Mutex{}, isTaken: false}
		d.forks[i].c = sync.Cond{L: &d.forks[i].m}
	}

	return d
}
