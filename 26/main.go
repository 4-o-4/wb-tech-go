package main

func Task(str string) bool {
    set := make(map[int32]bool)
    for _, c := range str {
        if set[c] {
            return false
        }
        set[c] = true
    }
    return true
}
