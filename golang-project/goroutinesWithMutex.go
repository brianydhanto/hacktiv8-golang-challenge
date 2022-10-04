package main

import (
	"fmt"
	"sync"
)

func main () {
	x := []interface{}{"coba1", "coba2", "coba3"}
	y := []interface{}{"bisa1", "bisa2", "bisa3"}
	var mutex sync.Mutex
	var wg sync.WaitGroup

	for i := 1; i < 5; i++ {

		wg.Add(1)
		go print(i, x, y, &wg, &mutex)
	}

	wg.Wait()
}

func print(index int, s []interface{}, t []interface{}, wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	fmt.Println(s, index)
	mutex.Unlock()

	fmt.Println(t, index)
	wg.Done()
}
