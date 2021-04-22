package main

import (
	"fmt"
)

func main() {
	x:=3

	for i:=0; i<=7; i++ {
		if i%x == 1 {
			fmt.Print("x ")
		} else {
			fmt.Print("y ")
		}
	}
}