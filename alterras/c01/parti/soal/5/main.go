package main

import (
	"fmt"
)

func main() {
	T := 20.0
	r := 4.0
	const Pi float64 = 22.0/7.0

	L := (2 * Pi * (r*r)) + (2 * Pi * r * T)
	fmt.Println(L)
}