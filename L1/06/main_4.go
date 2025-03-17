package main

import (
    "fmt"
    "time"
)

func Task4(done *bool) { // Флаг завершения
    for {
        if *done {
            fmt.Println("Горутина завершена")
            return
        }
        fmt.Println("Горутина работает...")
        time.Sleep(1 * time.Second)
    }
}

func main() {
    done := false
    go Task4(&done)
    time.Sleep(3 * time.Second)
    done = true
    time.Sleep(1 * time.Second)
}
