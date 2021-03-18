package main

import (
	"fmt"
)

func appendCategory(a []string) []string {

	check := make(map[string]int)
	d := append(a)
	res := make([]string, 0)
	for _, val := range d {
		check[val] = 1
	}

	for letter, _ := range check {
		res = append(res, letter)
	}

	return res
}

func main() {
	fmt.Println(appendCategory([]string{"2", "3", "3", "3", "6", "9", "9"}))

	fmt.Println(appendCategory([]string{"2", "2", "2", "11"}))

}
