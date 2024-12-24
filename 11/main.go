package main

import "fmt"

func Task(set1, set2 map[string]struct{}) map[string]struct{} {
    result := make(map[string]struct{})
    for item := range set1 {
        if _, ok := set2[item]; ok {
            result[item] = struct{}{}
        }
    }
    return result
}

func main() {
    set1 := map[string]struct{}{"cat": {}, "dog": {}, "tree": {}}
    set2 := map[string]struct{}{"cat": {}, "god": {}, "tree": {}}
    set := Task(set1, set2)
    for key := range set {
        fmt.Println(key)
    }
}
