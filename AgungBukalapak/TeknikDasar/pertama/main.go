package main

import (
	"fmt"
	"pertama/calculation"
)

func main() {
	fmt.Println("haiii")
	sentence := TestAja()

	result := calculation.Add(8, 9)

	perkalian := calculation.Multiply(2, 2)

	fmt.Println(sentence, result, perkalian)

}
