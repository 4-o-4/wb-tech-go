package main

import (
    "context"
    "fmt"
    "math/rand"
    "os"
    "os/signal"
    "sync"
    "time"
)

func Task(ctx context.Context, id int, dataChan <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("Воркер %d завершает работу\n", id)
            return
        case data, ok := <-dataChan:
            if !ok {
                fmt.Printf("Воркер %d: канал закрыт\n", id)
                return
            }
            fmt.Printf("Воркер %d получил данные: %d\n", id, data)
        }
    }
}

func main() {
    dataChan := make(chan int)
    numWorkers := 4
    ctx, cancel := context.WithCancel(context.Background())
    var wg sync.WaitGroup
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go Task(ctx, i, dataChan, &wg)
    }
    go func() {
        for {
            select {
            case <-ctx.Done():
                close(dataChan)
                return
            default:
                dataChan <- rand.Intn(100)
                time.Sleep(500 * time.Millisecond)
            }
        }
    }()
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt)
    <-stop
    fmt.Println("\nПолучен сигнал завершения...")
    cancel()
    wg.Wait()
    fmt.Println("Программа завершена.")
}
