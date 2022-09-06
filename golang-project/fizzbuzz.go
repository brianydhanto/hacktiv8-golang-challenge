package main

import "fmt"

func main() {
	var input int

	fmt.Print("Input number:")
	fmt.Scan(&input)

	for i := 1; i <= input; i++ {
        fizzBuzz(i)
    }
}

func fizzBuzz(i int) {
	fizz := "Fizz"
	buzz := "Buzz"

	if i % 3 == 0 && i % 5 == 0 {
		fmt.Println(fizz + buzz)
	} else if i % 3 == 0 {
		fmt.Println(fizz)
	} else if i % 5 == 0 {
		fmt.Println(buzz)
	} else {
		fmt.Println(i)
	}
}
