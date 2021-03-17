package main

import (
	"fmt"
)

func ArrayMerge(arrayA, arrayB []string) []string {
	return append(arrayA, arrayB...)
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
}
