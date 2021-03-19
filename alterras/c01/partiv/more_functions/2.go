package main

import "fmt"

func main() {
	// Anonymous function
	func() {
		fmt.Println("Welcome!, to GeeksforGeeks")
	}()

	// Assigning an anonymous function to a variable
	value := func() {
		fmt.Println("Welcome! to GeeksforGeeks")
	}
	value()

	// Passing arguments in anonymous function
	func(sentence string) {
		fmt.Println(sentence)
	}("GeeksforGeeks")
}