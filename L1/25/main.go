package main

import "time"

func Task(duration time.Duration) {
    start := time.Now()
    for {
        if time.Since(start) >= duration {
            break
        }
    }
}

func main() {
    Task(2 * time.Second)
}
