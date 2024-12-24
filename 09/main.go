package main

import "fmt"

func Task(inputChan <-chan int, outputChan chan<- int) {
    for x := range inputChan {
        outputChan <- x * x
    }
    close(outputChan)
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    inputChan := make(chan int)
    outputChan := make(chan int)

    go Task(inputChan, outputChan)
    go func() {
        for _, num := range numbers {
            inputChan <- num
        }
        close(inputChan)
    }()
    for result := range outputChan {
        fmt.Println(result)
    }
}
