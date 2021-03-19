package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func greeting(hour int) {
	if hour < 12 {
		fmt.Println("Selamat Pagi")
	} else if hour < 18 {
		fmt.Println("Selamat Sore")
	} else {
		fmt.Println("Selamat malam")
	}
}

func main() {
	hour := 15
	greeting(hour)
	sayHello()
}