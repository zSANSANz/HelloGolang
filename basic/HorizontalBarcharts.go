package main

import (
	"fmt"
)

func main() {
	HorizontalBarcharts()
}

func HorizontalBarcharts() {
	ints := []int{1, 4, 5, 6, 8, 2}

	i := 0

	for i <= 5 {
		fmt.Print(ints[i], " ")
		j := 1
		for j <= ints[i] {
			fmt.Print("-")
			j++
		}
		fmt.Println("")
		i++
	}
}
