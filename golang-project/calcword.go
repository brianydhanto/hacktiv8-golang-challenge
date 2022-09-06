package main

import (
    "fmt"
)

func main() {
    text := `selamat malam`

    var length = len([]rune(text))
    var collect = make(map[string]int)

    for i := 0; i < length; i++ {
        var txt = string(text[i])
        fmt.Println(string(text[i]))
        collect[txt] += 1
    }

    fmt.Print(collect)
}


