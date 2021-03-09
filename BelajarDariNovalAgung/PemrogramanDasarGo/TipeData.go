package main

import (
	"fmt"
)

func main() {
	var positiveNumber uint8 = 89
	var negativeNumber = -1234567890

	fmt.Printf("bilangan positif: %d \n", positiveNumber)
	fmt.Printf("bilangan negatif: %d \n", negativeNumber)

	var exist bool = true

	fmt.Printf("exist %t \n ", exist)

	var message string = "halo"
	fmt.Printf("yo bilang %s \n", message)

	message = `nama saya "jhon wick".
				salam kenal.
				mari belajar "golang".`
	fmt.Printf("yo bilang %s \n", message)

}