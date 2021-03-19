package main

import (
	"fmt"
	"strings"
)

func main() {
	// 4. substring
	value := "cat;dog"

	// take substring from index 4 to length of string.
	substring := value[4:len(value)]
	fmt.Println(substring)

	// 5. replace
	s := "this[things]I would like to remove "
	t := strings.Replace(s, "[", "", -1)
	fmt.Printf("%s\n", t)

	// 6. insert
	p := "green"
	index := 2

	q := p[:index] + "H1" + p[index:]
	fmt.Println(p, q)
}