package main

import (
	"fmt"
)

func main() {
	i:=0
	for (i<9) {
		i += 1

		if (i%2 == 0) {
			fmt.Print("masuk ")
		} else {
			fmt.Print("keluar ")
		}
	}
}