package main

import "fmt"

func Task(sequence []string) map[string]struct{} {
    set := make(map[string]struct{})
    for _, item := range sequence {
        set[item] = struct{}{}
    }
    return set
}

func main() {
    sequence := []string{"cat", "cat", "dog", "cat", "tree"}
    set := Task(sequence)
    for key := range set {
        fmt.Println(key)
    }
}
