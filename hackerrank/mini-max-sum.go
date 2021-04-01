package main

import (
	"fmt"
)

func main() {
	var positive [5]int
	
	for i:=0; i<=4; i++ {
		fmt.Scan(&positive[i])
	}

	min := 9223372036854775807
	for i:=0; i<=4; i++ {
		if min >= positive[i] {
			min = positive[i]
		}
	}
	jumlahBilMax := 0
	doubleMax := 0
	for i:=0; i<=4; i++ {
		if positive[i] == min {
			if doubleMax == positive[i] {
				jumlahBilMax = jumlahBilMax + positive[i]
			} else {
				doubleMax = positive[i]
			}
		}
		if positive[i] != min {
			jumlahBilMax = jumlahBilMax + positive[i]
		}
	}

	max := 0
	for i:=0; i<=4; i++ {
		if min <= positive[i] {
			max = positive[i]
		}
	}
	jumlahBilMin := 0
	doubleMin := 0
	for i:=0; i<=4; i++ {
		if positive[i] == max {
			if doubleMin == positive[i] {
				jumlahBilMin = jumlahBilMin + positive[i]
			} else {
				doubleMin = positive[i]
			}
		}
		if positive[i] != max {
			jumlahBilMin = jumlahBilMin + positive[i]
		}
	}
	fmt.Printf("%d %d", jumlahBilMin, jumlahBilMax)
}