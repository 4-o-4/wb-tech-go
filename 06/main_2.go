package main

import (
    "context"
    "fmt"
    "time"
)

func Task2(ctx context.Context) { // Контекст
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Горутина завершена")
            return
        default:
            fmt.Println("Горутина работает...")
            time.Sleep(1 * time.Second)
        }
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    go Task2(ctx)
    time.Sleep(10 * time.Second)
}
