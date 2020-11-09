package main

import (
	"fmt"
	"sort"
)

func main() {
	ints := []int{1, 4, 5, 6, 8, 2}
	fmt.Println("INPUT: Numerical array")
	fmt.Println("Ints Before Sorting:   ", ints)
	sort.Ints(ints)
	fmt.Println("")
	fmt.Println("OUTPUT: Vertical Barcharts")
	fmt.Println("Ints After Sorting:   ", ints)

}
