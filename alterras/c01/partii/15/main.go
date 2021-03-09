package main

import "fmt"

func main() {
	sum := 1
	for sum < 9 {
		sum += sum
	}
	fmt.Println(sum)
}
