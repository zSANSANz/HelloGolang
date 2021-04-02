package main

import (
	"fmt"
)

func main() {
	myMap := map[string]string {
		"ruby": "is beatiful",
		"go": "is super fast",
	}

	for key, value := range myMap {
		fmt.Println("key :", key, " value", value)
	}

	delete(myMap, "ruby")

	for key, value := range myMap {
		fmt.Println("key :", key, " value", value)
	}

	value, err := myMap["p"]
	fmt.Println(value, " ", err)
}