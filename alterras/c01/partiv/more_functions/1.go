package main

import (
	"fmt"
)

func sum(numbers ...int) int {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	avg := sum(2, 4, 3, 5)
	fmt.Println(avg)
}