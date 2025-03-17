package main

import (
    "fmt"
    "math"
)

func Task(t []float64) map[int][]float64 {
    groups := make(map[int][]float64)
    for _, temp := range t {
        key := int(math.Floor(temp/10) * 10)
        groups[key] = append(groups[key], temp)
    }
    return groups
}

func main() {
    temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
    groups := Task(temperatures)
    for key, temps := range groups {
        fmt.Printf("%d: %v\n", key, temps)
    }
}
