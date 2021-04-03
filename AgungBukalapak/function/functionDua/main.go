package main

import (
	"fmt"
)

func add(number, numberTwo int) int {
	result := number + numberTwo
	return result
}

func main() {
	result := add(10, 20)
	fmt.Println(result)
}