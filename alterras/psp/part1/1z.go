package main

import (
	"fmt"
	"math"
)

func main() {
	i, j, k := 0, 0, 0
	fmt.Scanf("%d \t %d \t %d", &i, &j, &k)
	a, b, c := getVal(i, j, k)
	if a == 0 && b == 0 && c == 0 {
		fmt.Println("Not Found")
	} else {
		fmt.Println(a, b, c)
	}
}

func getVal(A, B, C int) (int, int, int) {
	limit := int(math.Sqrt(float64(C)))

	a, b, c, z:= 0, 0, 0, 0
	ketemu := false
	for i:=1; i<=limit; i++ {
		if B % i == 0 {
			for j:=1; j<(A - 1); j++ {
				if B % j == 0 {
					z = A - (i + j)
					if B % z == 0 {
						a = i
						b = j
						c = z
						ketemu = true
						break
					}
				}
			}
		}
		if ketemu {
			break
		}
	}
	if a == 0 {
		return 0, 0, 0
	} else {
		return a, b, c
	}
	
}