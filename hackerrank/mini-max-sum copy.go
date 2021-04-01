package main

import (
	"fmt"
	"sort"
)

func main() {
	positive := make([]int, 5)
	
	for i:=0; i<=4; i++ {
		fmt.Scan(&positive[i])
	}

	jumlahBilMin, jumlahBilMax := 0, 0

	sort.Ints(positive)
	for i:=0; i<4; i++ {
		jumlahBilMin = jumlahBilMin + positive[i]
	}

	for i:=1; i<=4; i++ {
		jumlahBilMax = jumlahBilMax + positive[i]
	}
	
	fmt.Printf("%d %d", jumlahBilMin, jumlahBilMax)
}