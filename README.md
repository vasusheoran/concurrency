# Concurrency Playground in Golang

Welcome to the Concurrency Playground repository! 🚀 This collection of Golang projects explores various aspects of concurrent programming, offering hands-on examples and implementations.

## Projects

1. **[Bounded Blocking Queue](boundedblockingqueue/README.md)**
   - A thread-safe, bounded blocking queue implementation for managing tasks in a concurrent environment. (**Difficulty:** Medium)
   - [Leetcode Challenge](https://leetcode.com/problems/design-bounded-blocking-queue/)
2. **[Rides](rides/README.md)**
   - Simulates a scenario where Democrats and Republicans order Uber rides, considering specific conditions to prevent conflicts. (**Difficulty:** Hard)
3. **[BuildingH2O](buildingh2o/README.md)**
   - Demonstrates concurrent molecule creation using Go routines and synchronization primitives, simulating the formation of water molecules (H2O). (**Difficulty:** Medium)
   - [Leetcode Challenge](https://leetcode.com/problems/building-h2o/)
4. **[Dining Philosophers](diningphilosopher/README.md)**
   - Implements the classic Dining Philosophers problem using Go concurrency features, exploring synchronization techniques for resource sharing. (**Difficulty:** Medium)
   - [Leetcode Challenge](https://leetcode.com/problems/the-dining-philosophers/) 
5. **[FizzBuzz](fizzbuzz/README.md)**
   - Implements the classic FizzBuzz problem concurrently, printing numbers divisible by 3 as "Fizz", divisible by 5 as "Buzz", and divisible by both as "FizzBuzz". (**Difficulty:** Easy)
   - [Leetcode Challenge](https://leetcode.com/problems/fizz-buzz-multithreaded/)
6. **[FooBar](foobar/README.md)**
   - Implements a concurrent solution to the FooBar problem, where two goroutines take turns printing "Foo" and "Bar" in a specific order. (**Difficulty:** Easy)
   - [Leetcode Challenge](https://leetcode.com/problems/print-foobar-alternately/)
7. **[PrintOrder](printorder/README.md)**
   - Demonstrates synchronization techniques using goroutines and channels to ensure a specific order of printing messages from multiple concurrent routines.(**Difficulty:** Medium)
   - [Leetcode Challenge](https://leetcode.com/problems/print-in-order/)
8. **[RateLimiter](ratelimiter/README.md)**
   - Provides a thread-safe implementation for rate limiting functions based on leaky bucket algorithm, allowing control over the number of allowed calls within a defined time window. (**Difficulty:** Easy)
9. **[ReaderWriter](readerwriter/README.md)**
   - Offers a package for managing concurrent read and write access to shared data using read-write locks and optional channels for synchronization. (**Difficulty:** Medium)
10. **[SafeMap](safemap/README.md)**
    - Implements a thread-safe map structure that allows concurrent access and modification of key-value pairs without data races. (**Difficulty:** Medium)
11. **[ScheduleTasks](scheduletasks/README.md)**
    - Demonstrates scheduling tasks for execution at specific times or intervals using concurrency primitives like timers or goroutines. (**Difficulty:** Medium)

## Usage

### Import concurrency module

```bash
go get -u github.com/vasusheoran/concurrency@latest
```

### Explore Sub-Projects

Navigate to individual sub-projects to explore their README files and code.

## Contributing

Feel free to contribute, report issues, or suggest improvements. Your insights and ideas are highly appreciated.

## License

This repository is licensed under the [MIT License](LICENSE).