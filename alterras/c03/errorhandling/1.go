package main

import (
	"fmt"
	"errors"
)

func myFunc(i int) (int, error) {
	if i <= 0 {
		return -1, errors, New("should be greater than zero")
	}
	return i, nil
}

func main() {
	result, err := myFunc(-1)
	fmt.Println(result, err)
}
