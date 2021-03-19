package main

import "fmt"

type Student struct {
	name 	[]string
	score	[]int
}

func (s Student) average() float64 {
	return 20.5
}

func (s Student) Min() (min int, nama string) {
	return 20, "joko"
}

func main() {
	s := Student{
		name: {"Rizky","ismael"},
		score: 80,
	}
	fmt.Println(s.average())

	
}