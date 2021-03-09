package main

import "fmt"

func main() {
	title := "Golang the best language"

	for index, letter := range title {
		if index%2 == 0 {
			fmt.Println(" letter ke: ", index, "berisi huruf : ", string(letter))
		}
	}

	fmt.Println("============================================================")

	for _, letter := range title {
		if string(letter) == "a" ||
			string(letter) == "i" ||
			string(letter) == "u" ||
			string(letter) == "e" ||
			string(letter) == "o" {
			fmt.Println(" letter ke: ", letter, "berisi huruf : ", string(letter))
		}
	}

	fmt.Println("============================================================")

	for _, letter := range title {
		tulisan := string(letter)
		switch tulisan {
		case "a", "i", "u", "e", "o":
			fmt.Println(" letter ke: ", letter, "berisi huruf : ", string(letter))
		}
	}

}
