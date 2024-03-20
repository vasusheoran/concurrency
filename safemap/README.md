## Go Package for Managing Visited Keys with Thread Safety

This package provides a `SafeMap` interface that acts as a concurrent-safe set for tracking visited keys. It offers functionalities to check if a key exists (`Exist`) and add new keys (`Add`).

## Functionality

`SafeMap` utilizes a `sync.RWMutex` to ensure safe access to an underlying `map` from multiple Go routines simultaneously.

* **`data`:** An internal `map[interface{}]interface{}` stores key-value pairs. In practice, the values are typically set to a dummy value like `true`.
* **`mutex`:** A `sync.RWMutex` object provides read-write locking mechanisms for thread-safe access to the `data` map.

* **`Add(key interface{})`:**
    - Acquires a write lock on the `mutex`.
    - Adds a new key-value pair (key, true) to the `data` map.
    - Releases the write lock.
* **`Exist(key interface{}) bool`:**
    - Acquires a read lock on the `mutex`.
    - Checks if the provided `key` exists within the `data` map.
    - Releases the read lock.
    - Returns `true` if the key exists, `false` otherwise.

## Usage

1. Import the `safemap` package in your Go program.
2. Create a new instance of `SafeMap` using the `NewSafeMap` function.
3. Utilize the `Add(key interface{})` method to add keys to the set.
4. Employ the `Exist(key interface{}) bool` method to verify if a key exists in the set.

## Example

```go
package main

import (
  "fmt"
  "github.com/vasusheoran/concurrency"
)

func main() {
  visited := safemap.NewSafeMap()

  visited.Add("key1")
  visited.Add("key2")

  if visited.Exist("key1") {
    fmt.Println("key1 exists")
  }

  if !visited.Exist("key3") {
    fmt.Println("key3 does not exist")
  }
}
```