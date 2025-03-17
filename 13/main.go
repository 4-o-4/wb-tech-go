package main

func Task(a, b *int) {
    *b, *a = *a, *b
}
