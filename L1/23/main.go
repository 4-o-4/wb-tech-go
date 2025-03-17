package main

import "fmt"

func Task(slice []int, i int) []int {
    if i < 0 || i >= len(slice) {
        return slice
    }
    return append(slice[:i], slice[i+1:]...)
}

func main() {
    slice := []int{1, 2, 3, 4, 5}
    slice = Task(slice, 1)
    fmt.Println(slice) // [1 3 4 5]
}
