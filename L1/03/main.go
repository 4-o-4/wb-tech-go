package main

import (
    "fmt"
    "sync"
)

func Task(start, end int, numbers []int, resultChan chan<- int) {
    sum := 0
    for i := start; i < end; i++ {
        sum += numbers[i] * numbers[i]
    }
    resultChan <- sum
}

func main() {
    numbers := []int{2, 4, 6, 8, 10}
    numG := 2
    resultChan := make(chan int, numG)
    var wg sync.WaitGroup
    chunkSize := len(numbers) / numG
    for i := 0; i < numG; i++ {
        wg.Add(1)
        start := i * chunkSize
        end := (i + 1) * chunkSize
        if i == numG-1 {
            end = len(numbers)
        }
        go func(start, end int) {
            defer wg.Done()
            Task(start, end, numbers, resultChan)
        }(start, end)
    }
    wg.Wait()
    close(resultChan)
    totalSum := 0
    for sum := range resultChan {
        totalSum += sum
    }
    fmt.Println("Сумма квадратов чисел:", totalSum)
}
