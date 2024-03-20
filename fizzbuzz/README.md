## FizzBuzz - Concurrent Printing

This challenge implements the classic FizzBuzz problem using concurrent programming in Go. It showcases how to print "Fizz" for multiples of 3, "Buzz" for multiples of 5, and "FizzBuzz" for multiples of both 3 and 5, following the sequence from 1 to a specified number.

## Functionality

The `fizzbuzz` package defines a `FizzBuzz` interface with methods to print the appropriate output based on divisibility conditions. Internally, a `fizzBuzzCond` struct implements the interface and utilizes synchronization primitives (`sync.Mutex`, `sync.Cond`, and `sync.WaitGroup`) to ensure coordinated execution of concurrent goroutines handling "Fizz", "Buzz", "FizzBuzz", and number printing.

## Usage

1. Import the `fizzbuzz` package in your Go program.
2. Create an instance of `FizzBuzz` using the `NewFizzBuzzCond` function, specifying the maximum number and a wait group reference.
3. Launch separate Go routines for each printing functionality (`Fizz()`, `Buzz()`, `FizzBuzz()`, and multiple calls to `Number()`) to achieve concurrent execution.
4. Call the `Data()` method on the `FizzBuzz` interface to retrieve the final printed sequence as a slice of strings.

## Example

```go
package main

import (
    "fmt"
    "sync"
    "github.com/vasusheoran/concurrency/fizzbuzz" // Assuming the fizzbuzz.go file is in the same directory
)

func main() {
    n := 15  // Set the maximum number
    var wg sync.WaitGroup
    wg.Add(n) // Add the number of expected goroutines (n)

    fb := fizzbuzz.NewFizzBuzz(n, &wg)

    go fb.Fizz()
    go fb.Buzz()
    go fb.FizzBuzz()
    for i := 1; i <= n; i++ {
        go fb.Number()
    }

    wg.Wait() // Wait for all goroutines to finish
    fmt.Println("\nResult:", fb.Data()) // Print the final output
}
```