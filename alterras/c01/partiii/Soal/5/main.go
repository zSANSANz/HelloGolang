package main

import (
	"fmt"
)

func PairSum(arr []int, target int) {
	satu := 0
	dua := 0

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				satu = i
				dua = j
			}
		}
	}

	fmt.Println(dua, satu)
}

func main() {
	PairSum([]int{1, 2, 3, 4, 6}, 6)
	PairSum([]int{2, 5, 9, 11}, 11)
	fmt.Println(" ")
}
