package main

import (
    "fmt"
    "time"
)

func producer(ch chan<- int, duration time.Duration) {
    timer := time.NewTimer(duration)
    for i := 0; ; i++ {
        select {
        case ch <- i:
            fmt.Printf("Отправлено: %d\n", i)
            time.Sleep(500 * time.Millisecond)
        case <-timer.C:
            fmt.Println("Время истекло, отправка завершена.")
            close(ch)
            return
        }
    }
}

func consumer(ch <-chan int) {
    for value := range ch {
        fmt.Printf("Прочитано: %d\n", value)
    }
}

func main() {
    ch := make(chan int)
    duration := 5 * time.Second
    go producer(ch, duration)
    go consumer(ch)
    time.Sleep(duration + 2*time.Second)
    fmt.Println("Программа завершена.")
}
