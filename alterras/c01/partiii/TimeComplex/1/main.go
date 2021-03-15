package main

import (
	"fmt"
)

func main () {
	fmt.Println("woooee")
	fmt.Println(dominant(1000))
}

func dominant (n int) int {
	var result int = 0
	for i := 0; i < n; i++ {
		result += 1
	}
	return result
}