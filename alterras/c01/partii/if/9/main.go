package main

import "fmt"

func main() {
	var hour int
	fmt.Scanln(&hour)
	if hour < 12 {
		fmt.Println("Selamat Pagi")
	} else if hour < 18 {
		fmt.Println("Selamat Sore")
	} else {
		fmt.Println("Selamat Malam")
	}
}
