## Dining Philosophers in Golang

This project explores the classic Dining Philosophers problem using Go concurrency features. It implements two solutions to demonstrate different approaches to synchronization:

* `diningPhilosopherWithSignal`: This approach utilizes mutexes and condition variables to ensure philosophers acquire forks in the correct order (left then right) and avoid deadlocks.
* `diningPhilosopherWithMutex`: This solution leverages mutexes. It demonstrates an alternative approach to manage concurrent access to shared resources.

## Functionality

Both solutions define a `DiningPhilosophers` interface with two methods:

* `WantsToEat(pID int, pickLeftFork PickLeftFork, pickRightFork PickRightFork, putLeftFork PutLeftFork, putRightFork PutRightFork, eat Eat)`: This function simulates a philosopher attempting to eat. It takes arguments for functions representing actions like picking up forks, eating, and putting forks down. These functions are for demonstration purposes and can be implemented to simulate specific behavior.
* `GetEatingOrder() []int`: This function retrieves the order in which philosophers were able to eat.

## Usage

1. Import the `diningphilosopher` package in your Go program.
2. Create an instance of `DiningPhilosophers` using either `NewDiningPhilosopherWithSignal` or `NewDiningPhilosophersWithMutex`, depending on the desired synchronization approach.
3. Call the `WantsToEat` method concurrently in separate Go routines to simulate multiple philosophers.
4. Implement the provided function callbacks (`pickLeftFork`, `pickRightFork`, `putLeftFork`, `putRightFork`, `eat`) to define specific behavior for each philosopher's actions.
5. After running all philosophers, call `GetEatingOrder` to retrieve the order in which they were able to eat.

## Example Usage (Illustrative - Implement callbacks)

```go
package main

import (
    "fmt"
    "sync"
	"time"
    "github.com/vasusheoran/concurrency/diningphilosopher"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(5 * 5) // 5 philosophers

    // Choose your implementation (WithSignal or WithMutex)
    philosophers := diningphilosopher.NewDiningPhilosopherWithSignal(&wg)

    for i := 0; i < 5; i++ {
        go func(pID int) {
            philosophers.WantsToEat(pID, func(id int) {
                fmt.Println("Philosopher", id, "picked up left fork")
            }, func(id int) {
                fmt.Println("Philosopher", id, "picked up right fork")
            }, func(id int) {
                fmt.Println("Philosopher", id, "put down left fork")
            }, func(id int) {
                fmt.Println("Philosopher", id, "put down right fork")
            }, func(id int) {
                fmt.Println("Philosopher", id, "is eating")
                // Simulate eating time
                time.Sleep(time.Millisecond * 100)
            })
            wg.Done()
        }(i)
    }

    wg.Wait()

    order := philosophers.GetEatingOrder()
    fmt.Println("Eating order:", order)
}
```
