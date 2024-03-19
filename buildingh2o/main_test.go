package buildingh2o

import (
	"sync"
	"testing"
)

func TestNewH2O(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(15)
	h := NewH2O(wg)

	str := "OOOOOHHHHHHHHHH"
	//fmt.Print(len(str))

	for i := 0; i < len(str); i++ {
		ch := string(str[i])
		if ch == "O" {
			//fmt.Println("Releasing O")
			go h.ReleaseOxygen()
		} else {
			//fmt.Println("Releasing H")
			go h.ReleaseHydrogen()
		}
	}

	wg.Wait()
}
