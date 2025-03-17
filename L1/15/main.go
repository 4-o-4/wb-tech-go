package main

var justString string

func someFunc() {
    v := createHugeString(1 << 10)
    justString = string(v[:100]) // Копирование первых 100 символов в новую строку
}

func main() {
    someFunc()
}
