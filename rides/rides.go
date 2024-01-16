package rides

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Taxi interface {
	Ride(id int, riderType RiderType)
}

type taxi struct {
	m            sync.Mutex
	republicSema int
	demoSema     int

	republicWaiting sync.Mutex // republics waiting
	republicCond    sync.Cond

	demoWaiting sync.Mutex //demo waiting
	demoCond    sync.Cond
	wg          *sync.WaitGroup
}

func (t *taxi) Ride(id int, riderType RiderType) {
	if riderType == Democrat {
		go t.democrat(id, seatedFn, driveFn)
	} else {
		go t.republic(id, seatedFn, driveFn)
	}
}

func (t *taxi) democrat(id int, seated Seated, drive Drive) {
	defer t.wg.Done()
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

	t.m.Lock()
	so := sync.Once{}
	defer so.Do(t.m.Unlock)

	t.demoWaiting.Lock()
	defer t.demoWaiting.Unlock()

	isRider := false

	t.demoSema++
	if t.demoSema == 4 {
		// Release 3 other waiting democrats
		t.demoCond.Signal()
		t.demoCond.Signal()
		t.demoCond.Signal()
		isRider = true
		t.demoSema = 0
	} else if t.demoSema >= 2 && t.republicSema >= 2 {
		// Release 2 rep and 1 demo
		t.demoCond.Signal()
		t.republicCond.Signal()
		t.republicCond.Signal()
		isRider = true
		t.demoSema -= 2
		t.republicSema -= 2
	} else {
		// wait for riders
		so.Do(t.m.Unlock) // unlock current waiting mutex for taxi struct
		t.demoCond.Wait()
	}

	seated(id, Democrat)

	if isRider {
		drive(id, Democrat)
	}

}

func (t *taxi) republic(id int, seated Seated, drive Drive) {
	defer t.wg.Done()

	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	t.m.Lock()
	so := sync.Once{}
	defer so.Do(t.m.Unlock)

	t.republicWaiting.Lock()
	defer t.republicWaiting.Unlock()

	isRider := false

	t.republicSema++
	if t.republicSema == 4 {
		// Release 3 other waiting democrats
		t.republicCond.Signal()
		t.republicCond.Signal()
		t.republicCond.Signal()
		isRider = true
		t.republicSema = 0
	} else if t.demoSema >= 2 && t.republicSema >= 2 {
		// Release 2 rep and 1 demo
		t.demoCond.Signal()
		t.demoCond.Signal()
		t.republicCond.Signal()
		isRider = true
		t.demoSema -= 2
		t.republicSema -= 2
	} else {
		// TODO wait for riders
		so.Do(t.m.Unlock) // unlock current waiting mutex for taxi struct
		t.republicCond.Wait()
	}

	seated(id, Republic)

	if isRider {
		drive(id, Democrat)
	}
}

func NewRides(wg *sync.WaitGroup) Taxi {
	p := &taxi{
		m:               sync.Mutex{},
		republicWaiting: sync.Mutex{},
		demoWaiting:     sync.Mutex{},
		republicSema:    0,
		demoSema:        0,
		wg:              wg,
	}
	p.republicCond = sync.Cond{L: &p.republicWaiting}
	p.demoCond = sync.Cond{L: &p.demoWaiting}
	return p
}

type RiderType string

const (
	Democrat RiderType = "democrat"
	Republic RiderType = "republic"
)

type Seated func(ID int, riderType RiderType)

type Drive func(ID int, riderType RiderType)

func seatedFn(ID int, riderType RiderType) {
	fmt.Printf("%s rider %d is seated.\n", riderType, ID)
}

func driveFn(ID int, riderType RiderType) {
	fmt.Printf("Drive car signal from %s with id: %d\n", riderType, ID)
}
