package main

import "fmt"

func main() {
	var today int = 1
	switch {
	case today == 1:
		fmt.Println("TOday is Mondey")
	case today == 2:
		fmt.Println("Today is Tuesday")
	default:
		fmt.Printf("Invalid Date")
	}
}
