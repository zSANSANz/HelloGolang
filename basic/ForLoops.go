package main

import (
	"fmt"
)

func main() {
	counter := 1
	for counter <= 13 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	for counter2 := 1; counter2 <= 10; counter2++ {
		fmt.Println("Perulangan ke", counter2)
	}

	slice := []string{"sandi", "peramna", "ibrahim"}
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])

	for i := 0; i < len(slice); i++ {
		fmt.Println(i, " ", slice[i])
	}

	for i, value := range slice {
		fmt.Println("index", i, "=", value)
	}

	for _, value := range slice {
		fmt.Println("=", value)
	}

	person := make(map[string]string)
	person["name"] = "sanji"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
