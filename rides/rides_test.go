package rides

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestConcurrentRides(t *testing.T) {
	var wg sync.WaitGroup
	taxiInstance := NewRides(&wg)

	// Define a function to simulate multiple concurrent riders
	simulateConcurrentRides := func(id int, riderType RiderType) {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		taxiInstance.Ride(id, riderType)
	}

	// Simulate multiple concurrent Democrats and Republicans taking rides
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go simulateConcurrentRides(i, Democrat)
	}

	for i := 5; i <= 8; i++ {
		wg.Add(1)
		go simulateConcurrentRides(i, Republic)
	}

	// Wait for the taxi to finish processing
	wg.Wait()
}
