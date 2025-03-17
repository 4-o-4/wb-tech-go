package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

type AtomicCounter struct {
    value int64
}

func (c *AtomicCounter) Increment() {
    atomic.AddInt64(&c.value, 1)
}

func (c *AtomicCounter) Value() int64 {
    return atomic.LoadInt64(&c.value)
}

func main() {
    counter := &AtomicCounter{}
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.Increment()
        }()
    }
    wg.Wait()
    fmt.Printf("Итоговое значение: %d\n", counter.Value())
}
