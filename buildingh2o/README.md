# BuildingH2O - Concurrent Molecule Creation

This project explores concurrent programming in Go by simulating the creation of water molecules (H2O).

## Functionality

`buildingh2o` defines an `H2O` interface with methods to release hydrogen (`ReleaseHydrogen`) and oxygen (`ReleaseOxygen`) atoms. Internally, it uses a mutex (`sync.Mutex`), condition variable (`sync.Cond`), and wait group (`sync.WaitGroup`) to synchronize access and ensure correct molecule formation (HHO, HOH, OHH).

## Usage

1. Import the `buildingh2o` package in your Go program.
2. Create an instance of `H2O` using the `NewH2O` function, passing a wait group.
3. Call the `ReleaseHydrogen` and `ReleaseOxygen` methods concurrently in separate Go routines to simulate atom release.

## Example

```go
package main

import (
    "fmt"
    "sync"
    "github.com/vasusheoran/concurrency/buildingh2o"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(9) // Number of H and O atoms (2H + 1O) * 3 molecules

    h2o := buildingh2o.NewH2O(&wg)
    
    for i := 0; i < 3; i++ {
        go h2o.ReleaseHydrogen()
        go h2o.ReleaseHydrogen()
        go h2o.ReleaseOxygen()
    }

    wg.Wait()
    fmt.Println("\nMolecules formed!")
}
```
