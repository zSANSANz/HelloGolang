package main

import (
	"fmt"
)

func main() {
	fmt.Println(" ")
	fmt.Println(" ")
	ints := []int{1, 4, 5, 6, 8, 2}
	fmt.Println("INPUT: Numerical array")
	fmt.Println("")
	fmt.Println(ints)
	fmt.Println("")
	fmt.Println("OUTPUT: Vertical Barcharts")
	fmt.Println("")
	VerticalBarcharts(ints)
	fmt.Println("INPUT: Numerical array")
	fmt.Println("")
	fmt.Println(ints)
	fmt.Println("")
	fmt.Println("OUTPUT:")
	fmt.Println("")
	fmt.Println("- Sorted array (ascending)")
	fmt.Println("- Steps visualization")
	fmt.Println("")
	VerticalBarcharts(ints)
	insertionSort(ints)

}

func VerticalBarcharts(values []int) {
	cols := len(values)
	max := 8

	for r := 0; r < max; r = r + 1 {
		for c := 0; c < cols; c = c + 1 {
			if r+values[c] >= max {
				fmt.Print("| ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println("")

	}

	for counter := 0; counter <= 5; counter++ {
		fmt.Print(values[counter], " ")
	}
	fmt.Println(" ")
	fmt.Println(" ")
}

func insertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			swap(array, j, j-1)
			VerticalBarcharts(array)
		}
	}
}

func swap(arrayzor []int, i, j int) {
	tmp := arrayzor[j]
	arrayzor[j] = arrayzor[i]
	arrayzor[i] = tmp
}
