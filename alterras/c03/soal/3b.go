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
	s1 := Student{
		name: "Rizky",
		score: 80,
	}
	fmt.Println(s1.average())
	s2 := Student{
		name: "Kobar",
		score: 75,
	}
	fmt.Println(s2.average())

	
}