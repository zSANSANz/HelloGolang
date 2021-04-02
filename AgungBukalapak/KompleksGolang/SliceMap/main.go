package main

import (
	"fmt"
)

func main() {
	students := []map[string]string{
		{
			"name": "agung",
			"score": "A",
		},
		{
			"name": "link",
			"score": "B",
		},
		{
			"name": "mario",
			"score": "C",
		},
	}

	fmt.Println(students)

	for index, student := range students {
		fmt.Println(index, " :: ", student["name"], " == ", student["score"])
	}

	for _, student := range students {
		for index, s := range student {
			fmt.Println(index, " ||| ", s)
		}
	}
}