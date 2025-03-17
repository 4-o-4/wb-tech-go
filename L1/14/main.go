package main

import "fmt"

func Task(v interface{}) {
    switch v.(type) {
    case int:
        fmt.Println("Тип: int")
    case string:
        fmt.Println("Тип: string")
    case bool:
        fmt.Println("Тип: bool")
    case chan int:
        fmt.Println("Тип: channel (chan int)")
    default:
        fmt.Println("Тип: неизвестный")
    }
}

func main() {
    values := []interface{}{
        42,
        "Hello, world!",
        true,
        make(chan int),
        3.14,
    }

    for _, v := range values {
        fmt.Printf("Значение: %v -> ", v)
        Task(v)
    }
}
