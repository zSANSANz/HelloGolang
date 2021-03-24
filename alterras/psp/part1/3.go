package main

import(
	"fmt"
)

func startSort (arr[]int) []int {
	if len(arr)==1 {
		return arr
	}

	center := len(arr)/2

	arrA := startSort(arr[:center])
	arrB := startSort(arr[center:])

	return join(arrA, arrB)
}

func main() {
	dragonList := []int{7, 2}
	knightList := []int{2, 1, 8, 5}

	sortedDragon := startSort(dragonList)
	sortedKnight := startSort(knightList)

	fmt.Println(sortedKnight)
	fmt.Println(sortedDragon)

	minTotal := 0
	knightWin := false

	for i:=0; i<len(sortedKnight); i++ {
		for j:=0; j<len(sortedDragon); j++ {
			if sortedKnight[i] >= sortedDragon[j] {
				minTotal += sortedKnight[i]
				if len(sortedDragon) == 1 {
					knightWin = true
				} else {
					sortedDragon =sortedDragon[j+1 :]
				}
				break
			}
		}
		if knightWin {
			break
		}
	}
	if knightWin {
		fmt.Printf("Result: %d", minTotal)
	} else {
		fmt.Println("Knight Fall")
	}
}

func join(a, b[]int)[]int{
	var finalArr []int
		idxa, idxb := 0,0

		lena, lenb := len(a), len(b)

		for i:=0; i<(lena+lenb); i++ {
			if idxb < lenb && idxa < lena {
				if a[idxa] <= b[idxb] {
					finalArr = append(finalArr, a[idxa])
					idxa++
				} else {
					finalArr = append(finalArr, b[idxb])
					idxb++
				}
			} else if idxb == lenb {
				finalArr = append(finalArr, a[idxa])
				idxa++
			} else if idxa == lena {
				finalArr = append(finalArr, b[idxb])
				idxb++
			}
		}
		return finalArr
}

