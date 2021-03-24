package main

import "fmt"

func myMethod() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error Message: ", err)
		}
	}()

	anOddCondition := true
	if anOddCondition {
		panic("I am Panicking")
	}
}

func main() {
	myMethod()
}