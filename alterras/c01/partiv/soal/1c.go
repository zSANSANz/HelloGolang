package main

import "fmt"

func compare(input_a, input_b string) {
	for i := 0; i < len(input_a); i++ {
		for j := 0; j < len(input_b); j++ {
			if string(input_a[i]) == string(input_b[j]) {
				fmt.Print(string(input_a[i]))
			}
		}
	}
	fmt.Println("")
}

func main() {
	compare("akashi", "aka")
	compare("kangooro", "kang")
}
