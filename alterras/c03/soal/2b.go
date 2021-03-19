package main

import "fmt"

func main() {

	var a = [6]int{1, 2, 3, 9, 7, 8}
	min, max := findMinAndMax(a)
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)

}

func findMinAndMax(a [6]int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}