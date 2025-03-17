package main

import (
    "fmt"
    "time"
)

func Task1(stopChan <-chan bool) { // Канал для сигнализации
    for {
        select {
        case <-stopChan:
            fmt.Println("Горутина завершена")
            return
        default:
            fmt.Println("Горутина работает...")
            time.Sleep(1 * time.Second)
        }
    }
}

func main() {
    stopChan := make(chan bool)
    go Task1(stopChan)
    time.Sleep(3 * time.Second)
    stopChan <- true
}
