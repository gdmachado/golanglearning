package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := x
    dz := 1.0
    for dz > 1e-10 {
        z0 := z
        z = z-(z*z-x)/(2*z)
        dz = z0 - z
    }
    return z
}

func main() {
    for i := 0; i < argv[0]; i++ {
        fmt.Printf("Sqrt(%v): %v %v\n", i, Sqrt(float64(i)), math.Sqrt(float64(i)))
    }
}
