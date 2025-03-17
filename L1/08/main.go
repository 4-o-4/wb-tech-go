package main

import "fmt"

func Task(n int64, i int) int64 {
    return n ^ (1 << i)
}

func main() {
    var num int64 = 42
    fmt.Printf("Число: %d\n", num)
    fmt.Printf("Битовый формат: %064b\n", num)

    tmp := Task(num, 1)
    fmt.Printf("Число: %d\n", tmp)
    fmt.Printf("Битовый формат: %064b\n", tmp)

    tmp = Task(num, 2)
    fmt.Printf("Число: %d\n", tmp)
    fmt.Printf("Битовый формат: %064b\n", tmp)
}
