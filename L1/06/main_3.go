package main

import (
    "fmt"
    "sync"
    "time"
)

func Task3(wg *sync.WaitGroup) { // WaitGroup
    defer wg.Done()
    for i := 0; i < 3; i++ {
        fmt.Println("Горутина работает...")
        time.Sleep(1 * time.Second)
    }
}

func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go Task3(&wg)
    wg.Wait()
    fmt.Println("Горутина завершена")
}
