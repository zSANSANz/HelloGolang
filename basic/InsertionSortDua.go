package main

import (
	"fmt"
)

func swap(arrayzor []int, i, j int) {
	tmp := arrayzor[j]
	arrayzor[j] = arrayzor[i]
	arrayzor[i] = tmp
}
func insertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			swap(array, j, j-1)
		}
	}
}
func main() {
	array := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("Unsorted array: ", array)
	insertionSort(array)
	fmt.Println("Sorted array: ", array)
}
