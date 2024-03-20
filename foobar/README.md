## Go Concurrency with Channel Synchronization

This code demonstrates the use of channels for synchronization in concurrent Go routines. It implements a simple alternating printing of "foo" and "bar" messages `n` times each, ensuring they appear in the order "foo", "bar", "foo", "bar", and so on.

## Functionality

1. **Shared Variables:** The code uses global variables `wg` (a `sync.WaitGroup`) and `n` (an integer) to manage goroutine synchronization and the number of iterations.
2. **Channel:** A channel `ch` of type `bool` is created to control access to the critical section (CS) where printing occurs. It acts as a binary semaphore, allowing only one goroutine to enter the CS at a time.
3. **`foo` Routine:** This goroutine iterates `n` times:
    - Prints "foo".
    - Sends `true` to the channel (`ch <- true`), signaling that bar can enter the CS.
    - Waits by receiving from the channel (`<-ch`), ensuring "bar" has finished printing before proceeding.
4. **`bar` Routine:** This goroutine follows the same pattern as `foo`:
    - Waits to receive `true` from the channel before entering the CS.
    - Prints "bar".
    - Sends `true` back to the channel, allowing "foo" to proceed.
5. **`main` Function:**
    - Initializes `wg` and `ch`.
    - Sets `n` to the desired number of iterations.
    - Increments `wg` by 2 for the two goroutines.
    - Launches `foo` and `bar` as Go routines.
    - Waits for all goroutines to finish using `wg.Wait()`.

## Key Points

- This approach ensures that "foo" and "bar" are printed in an alternating sequence.
- Channels offer a flexible way to control access to shared resources between concurrent goroutines.

## Considerations

- This is a simple example. For more complex scenarios, consider using mutexes or other synchronization primitives alongside channels.
- Global variables should be used cautiously in Go due to potential concurrency issues. Consider refactoring to pass necessary data as arguments to functions.
