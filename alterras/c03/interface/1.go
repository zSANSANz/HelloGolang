package main

import "fmt"

type calculate interface {
	large() int
}

type square struct {
	side int
}

func (s square) large() int {
	return s.side * s.side
}

func main() {
	var dimResult calculate
	dimResult = square{10}
	fmt.Println("large square : ", dimResult.large())
}