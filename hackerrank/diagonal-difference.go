package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	lsum, rsum := 0, 0
	fmt.Scan(&n)
	a := make([][]int, n)
	for i:=0; i<n; i++ {
		a[i] = make([]int, n)
		for j:=0; j<n; j++ {
			fmt.Scan(&a[i][j])
			if i==j {
				lsum = lsum + a[i][j]
			}
			if i+j == n-1 {
				rsum = rsum + a[i][j] 
			}
		}
	}
	diff := math.Abs(float64(lsum-rsum))
	fmt.Println(diff)
}