# Rides Project in Golang

Welcome to the Rides sub-project of the Concurrency Playground! 🎢 This project explores a scenario involving Democrats and Republicans waiting to ride together. The goal is to implement a solution that allows specific conditions for riders to join.

## Problem Statement

This project is inspired by the Uber Ride Problem. [Check out the problem statement here.](https://www.educative.io/courses/java-multithreading-for-senior-engineering-interviews/uber-ride-problem)

Imagine at the end of a political conference, Republicans and Democrats are trying to leave the venue and ordering Uber rides at the same time. To avoid conflicts, each ride can have either all Democrats or Republicans or two Democrats and two Republicans.

All other combinations can result in a fist-fight.

Your task as the Uber developer is to model the ride requestors as threads. Once an acceptable combination of riders is possible, threads are allowed to proceed to ride.

Each thread invokes the method `seated()` when selected by the system for the next ride. When all the threads are seated, any one of the four threads can invoke the method `drive()` to inform the driver to start the ride.

## Implementation

The `rides` package now provides an interface and an implementation for the Taxi project, defining behaviors for Democrats and Republicans.

### Taxi Interface

```go
type Taxi interface {
    Ride(id int, riderType RiderType)
}
```

### Example Usage

```go
package main

import (
	"sync"
	"github.com/vasusheoran/concurrency/rides"
)

func main() {
	var wg sync.WaitGroup
	taxiInstance := rides.NewRides(&wg)

	// Add your code to simulate Democrats and Republicans riding together

	wg.Wait()
}
```

### Ride Functions

The project now includes the `Ride` method, which takes a rider's ID and type (Democrat or Republic) and simulates the rider's behavior.

### Seated and Drive Functions

Additionally, the `Seated` and `Drive` functions have been introduced to customize the print messages for when a rider is seated and when the car is driven.

### How It Works

The project still utilizes `sync.Mutex`, `sync.Cond`, and `sync.WaitGroup` to synchronize and control the flow of Democrats and Republicans.

## Contributing

Feel free to contribute, report issues, or suggest improvements. Your insights and ideas are highly appreciated.

## License

This project is licensed under the [MIT License](LICENSE).

---