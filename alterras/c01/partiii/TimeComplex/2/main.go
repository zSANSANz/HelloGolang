package main

import (
	"fmt"
)

func main() {
	var A[10] int
	fmt.Println(linier(10, A))
}

func linier(n int, A[]int) int {
	for i := 0; i < n; i++ {
		if A[i] == 0 {
			return 0
		}
	}
	return 1
}