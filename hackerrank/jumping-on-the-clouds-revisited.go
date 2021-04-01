package main

import (
	"fmt"
)

func main() {
	n, k := 0, 0
	var c []int
	fmt.Scanf("%d\t%d", &n, &k)
	j := 0
	for i:=0; i<n; i++ {
		fmt.Scan(&j)
		c = append(c, j)
	}

	energy := 100
	for i:=0; i<n; i=i+k {
		fmt.Printf("energy = %d - %d - %d", energy, k, c[i])
		energy = energy - k - c[i]
		fmt.Println()
	}
	fmt.Println(energy)
}