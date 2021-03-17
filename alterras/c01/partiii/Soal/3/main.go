package main

import (
	"fmt"
)

func ArrayMerge(arrayA, arrayB []string) []string {
	return append([]string{"1","2"}, []string{"3","4"})
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
}