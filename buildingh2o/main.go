package buildingh2o

import (
	"fmt"
	"sync"
)

type H2O interface {
	ReleaseOxygen()
	ReleaseHydrogen()
}

type h2o struct {
	h         int // 10
	o         int // 5
	hConsumed int
	oConsumed int
	isOxygen  bool
	m         sync.Mutex
	c         sync.Cond
	wg        *sync.WaitGroup
}

// HHO
// HOH
// OHH

func (h *h2o) ReleaseHydrogen() {
	h.m.Lock()
	defer h.m.Unlock()

	for {
		if h.hConsumed < 2 {
			break
		} else {
			h.c.Wait()
		}
		//fmt.Println("", h.hConsumed, h.isOxygen)
	}
	fmt.Print("H")
	h.hConsumed += 1

	if h.oConsumed == 1 && h.hConsumed == 2 {
		h.oConsumed = 0
		h.hConsumed = 0
	}

	h.c.Broadcast()
	h.wg.Done()
}

func (h *h2o) ReleaseOxygen() {
	h.m.Lock()
	defer h.m.Unlock()

	for {
		if h.oConsumed < 1 {
			break
		} else {
			h.c.Wait()
		}
	}

	fmt.Print("O")
	h.oConsumed++

	if h.oConsumed == 1 && h.hConsumed == 2 {
		h.oConsumed = 0
		h.hConsumed = 0
	}

	h.c.Broadcast()
	h.wg.Done()
}

func NewH2O(wg *sync.WaitGroup) H2O {
	h := &h2o{
		h:  0,
		o:  0,
		m:  sync.Mutex{},
		wg: wg,
	}

	h.c = sync.Cond{L: &h.m}
	return h
}
