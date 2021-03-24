package main

import "fmt"

func main() {
	fmt.Println(fibonacci(5))
	fmt.Println(fibonacci(6))
	fmt.Println(fibonacci(7))
}

func fibonacci(n int) int {
	if n== 0 || n == 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}