package main

import "fmt"

func main() {
	var today int = 3
	switch today {
	case 1:
		fmt.Printf("Today is monday")
	case 2:
		fmt.Printf("Today is Tuesday")
	default:
		fmt.Printf("Invalid Date")
	}
}
