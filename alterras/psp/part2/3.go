package main

import "fmt"

func main() {
	fmt.Println(Solution(10, 30, 40))
}

func Solution(X int, Y int, D int) int {
    result := (Y-X)/D
    if (Y-X)%D > 0 {
        result++
    }
    return result
}