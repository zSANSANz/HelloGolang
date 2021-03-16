package main

import (
	"fmt"
)

func main() {
	input := 17
	if input % 1 == 0 {
		if input % input == 0 {
			if input % 2 == 0 {
				fmt.Println("Tidak")
			} else {
				fmt.Println("YA")
			}
		}
	}
}