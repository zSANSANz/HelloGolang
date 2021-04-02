package main

import (
	"fmt"
)

func main() {
	scores := []int{100, 80, 75, 92, 70 ,93, 88,67}
	jumlah := 0
	for _, score := range scores {
		jumlah = jumlah + score
	}
	rata := 0.00
	rata = float64(jumlah) / float64(len(scores))
	fmt.Println(rata)

	scores2 := []int{100, 80, 75, 92, 70 ,93, 88,67}

	var goodScores []int
	for _, score2 := range scores2 {
		if score2 >= 90 {
			goodScores = append(goodScores, score2)
		}
	}

	for _, goodScore := range goodScores {
		fmt.Print(goodScore, " ")
	}
}