package main

import (
    "fmt"
    "sync"
)

func main () {
    x := []interface{}{"coba1", "coba2", "coba3"}
    y := []interface{}{"bisa1", "bisa2", "bisa3"}
    var wg sync.WaitGroup

    for i := 1; i < 5; i++ {

        wg.Add(1)
        go printSecondInterface(i, x, &wg)

        wg.Add(1)
        go printFirstInterface(i, y, &wg)
    }

    wg.Wait()
}

func printFirstInterface(index int, s []interface{}, wg *sync.WaitGroup) {
    fmt.Println(s, index)
    wg.Done()
}

func printSecondInterface(index int, s []interface{}, wg *sync.WaitGroup) {
    fmt.Println(s, index)
    wg.Done()
}
