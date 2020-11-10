package main

import (
	"fmt"
)

func main() {
	var names [3]string
	names[0] = "Sandi"
	names[1] = "Permana"
	names[2] = "Ibrahim"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var jeneng = names[0]
	fmt.Println(jeneng)

	var values = [3]int{
		98,
		95,
		90,
	}

	fmt.Println(values)
	fmt.Println("[ ", values[0], ", ", values[0], ", ", values[0], " ]")

	fmt.Println(len(names))

	var lagi [10]string
	fmt.Println(len(lagi))
}
