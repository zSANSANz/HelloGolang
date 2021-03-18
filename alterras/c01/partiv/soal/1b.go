package main

import (
	"fmt"
)

func main() {
	input_a := "kang"
	input_b := "kangoroo"

	for i := 0; i < len(input_a); i++ {
		for j := 0; j < len(input_b); j++ {
			if string(input_a[i]) == string(input_b[j]) {
				fmt.Print(string(input_a[i]))
			}
		}
	}
}
