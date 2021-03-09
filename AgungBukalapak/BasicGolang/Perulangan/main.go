package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Saya sedang belajar Go : ", i)
	}

	a := 1
	for a <= 10 {
		fmt.Println("Dia Belajar Ruby : ", a)
		a++
	}

	title := "Golang the best language"

	for index, letter := range title {
		fmt.Println("index : ", index, " letter ke: ", letter, "berisi huruf : ", string(letter))
	}

	words := "Python is language for data sicience"

	for _, letter := range words {
		fmt.Println(" letter ke: ", letter, "berisi huruf : ", string(letter))
	}
}
