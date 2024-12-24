package main

import "strings"

func Task(str string) string {
    var sb strings.Builder
    split := strings.Split(str, " ")
    for i := len(split); i != 0; i-- {
        sb.WriteString(split[i-1])
        if i != 1 {
            sb.WriteByte(' ')
        }
    }
    return sb.String()
}
