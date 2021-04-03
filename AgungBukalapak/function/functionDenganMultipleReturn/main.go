package main

import (
	"fmt"
)

func main() {
	fmt.Println(add(10, 2))

	luas, keliling := add(20, 2)
	fmt.Println(luas, keliling)

	keren, _ := add(20, 2)
	fmt.Println(keren)

	fmt.Println(calculate(3, 2))
}

func add(number, numberTwo int) (int, int) {
	return number * numberTwo, (number + numberTwo) * 2
}

func calculate(panjang int, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)
	return
}