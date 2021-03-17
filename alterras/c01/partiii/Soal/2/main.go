package main

import (
	"fmt"
)

func pow(x, n int) int {
	switch n {
	case 1:
		return x
	default:
		return x*pow(x,(n-1));
	}
}

func main() {
	fmt.Println(pow(2, 3)) // 8
	fmt.Println(pow(7, 2)) // 49
}