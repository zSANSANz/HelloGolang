package main

import (
	"fmt"
)

func main() {
	counter := 1

	for counter <= 13 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	for counter2 := 1; counter2 <= 10; counter2++ {
		fmt.Println("Perulangan ke", counter2)
	}
}
