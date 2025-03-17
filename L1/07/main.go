package main

import (
    "fmt"
    "sync"
)

func main() {
    data := make(map[string]int)
    var mu sync.Mutex
    var wg sync.WaitGroup
    for i := 0; i < 50; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            mu.Lock()
            data[fmt.Sprintf("key %d", i)] = i
            mu.Unlock()
        }(i)
    }
    wg.Wait()
    for key, value := range data {
        fmt.Printf("%s: %d\n", key, value)
    }
}
