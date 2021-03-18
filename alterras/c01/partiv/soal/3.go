package main

import "fmt"

func main() {
	var input_a = []int{3, 6, 10, 12, 15}
	var input_b = []int{1, 3, 5, 10, 16}
	for i := 0; i < len(input_a); i++ {
		for j := 0; j < len(input_b); j++ {
			if input_a[i] == input_b[j] {
				fmt.Println(input_a[i])
			}
		}
	}
}
