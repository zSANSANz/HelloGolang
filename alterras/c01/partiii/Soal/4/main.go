package main

import "fmt"

func MeanMedian(arr []float64) {
	n := len(arr)
	tengah := len(arr) / 2

	sum := 0.00

	for i := 0; i < n; i++ {
		sum += (arr[i])
	}

	avg := (float64(sum)) / (float64(n))

	median := arr[tengah] + arr[tengah-1]

	fmt.Println("Mean = ", avg, "Median = ", median)
}

func main() {
	MeanMedian([]float64{1, 2, 3, 4})    // Mean: 2.5, Median: 2.5
	MeanMedian([]float64{1, 2, 3, 4, 5}) // Mean: 3.0, Median: 3.0
	MeanMedian([]float64{3, 5, 7, 5, 3}) // Mean: 4.6, Median: 7.0
}
