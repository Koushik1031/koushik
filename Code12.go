package main

import (
    "fmt"
    "math"
)

func main() {
    x := 3
    y := 4
    f := math.Sqrt(float64(x*x + y*y))
    z := uint(f)
    fmt.Printf("%d %d %d\n", x, y, z)
}
