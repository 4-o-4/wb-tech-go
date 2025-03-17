package main

import (
    "strings"
    "unicode/utf8"
)

func Task(str string) string {
    var sb strings.Builder
    for i := len(str); i > 0; {
        r, size := utf8.DecodeLastRuneInString(str[:i])
        sb.WriteRune(r)
        i -= size
    }
    return sb.String()
}
