package main

import "fmt"

func main() {
	var name = "joko"

	if name == "sandi" {
		fmt.Println("hallo sandi")
	}

	if name == "susanto" {
		fmt.Println("hallo susanto")
	} else {
		fmt.Println("hallo sameone")
	}

	if name == "susanto" {
		fmt.Println("hallo susanto")
	} else if name == "joko" {
		fmt.Println("hello joko")
	} else {
		fmt.Println("hallo sameone")
	}

	var length = len(name)
	if length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}

	if lengthDua := len(name); lengthDua > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}

}
