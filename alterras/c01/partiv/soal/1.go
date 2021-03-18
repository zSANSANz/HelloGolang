package main

import (
	"fmt"
	"strings"
)

func main() {

	input_a := strings.Split("akashi", "")
	input_b := strings.Split("akashi", "")

	fmt.Println(len(input_a))

	for i := 0; i <= (len(input_a) + 1); i++ {
		for j := 1; j <= 5; j++ {
			if input_a[i] == input_b[j] {
				fmt.Print(input_a[i])
			}
		}
	}
}
