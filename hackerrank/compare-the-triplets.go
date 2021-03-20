package main

import "fmt"

func main() {
	var a [3]int
	var b [3]int

	fmt.Scan(&a[0], &a[1], &a[2])
	fmt.Scan(&b[0], &b[1], &b[2])

	alice := 0
	bob := 0

	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {

		} else if a[i] < b[i] {
			alice = alice + 1
		} else if a[i] > b[i] {
			bob = bob + 1
		}
	}

	fmt.Println(bob, alice)
}
