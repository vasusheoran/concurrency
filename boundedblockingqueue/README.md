# Bounded Blocking Queue in Golang

A simple implementation of a bounded blocking queue in Golang. This package provides a thread-safe, bounded blocking queue that supports enqueue and dequeue operations, making it suitable for concurrent programming scenarios.

## Usage

### Importing the Package

```go
import "github.com/vasusheoran/concurrency/boundedblockingqueue"
```

### Creating a Bounded Blocking Queue

```go
capacity := 10 // Set your desired capacity
bbq := boundedblockingqueue.New(capacity)
```

### Enqueue an Item

```go
item := "YourItem"
bbq.Enqueue(item)
```

### Dequeue an Item

```go
dequeuedItem := bbq.Dequeue()
```

## How It Works

This implementation utilizes the `sync.Mutex` and `sync.Cond` primitives from the Golang standard library to achieve thread safety. The queue has a specified capacity, and when it reaches full capacity, the enqueue operation will block until space becomes available. Similarly, the dequeue operation will block if the queue is empty.

## Contributing

Feel free to contribute, report issues, or suggest improvements. Your feedback is highly valued!