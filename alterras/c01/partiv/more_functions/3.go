package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Later")
	}()

	fmt.Println("first")
}