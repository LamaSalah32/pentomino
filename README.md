# Pentomino DLX Solver in Go

This project implements a **Pentomino tiling puzzle solver** in Go, using Donald Knuth’s **Dancing Links (DLX)** [paper](https://www.ocf.berkeley.edu/~jchu/publicportal/sudoku/0011047.pdf) algorithm to solve the Exact Cover problem.

## Example

```go
package main

import (
    pentomino "github.com/lamasalah32/pentomino-tiling"
)

func main() {
    width := 5
    height := 5
    pieces := []string{"L", "P", "U", "F", "W"}

    pentomino.Solve(width, height, pieces)
}
```

### Output
```
[U U F L L]
[U F F F L]
[U U W F L]
[P P W W L]
[P P P W W]
```

---
