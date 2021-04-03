package main

import(
	"fmt"
	"errors"
)

func main() {
	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)
	fmt.Println(total)
	result, err := calculate(10, 2, "+")
	if err != nil {
		fmt.Println("terjadi error")
		fmt.Println(err.Error())
	} else {
		fmt.Println(result, err)
	}
	result, err = calculate(10, 2, "-")
	if err != nil {
		fmt.Println("terjadi error")
		fmt.Println(err.Error())
	} else {
		fmt.Println(result, err)
	}
	result, err = calculate(10, 2, "*")
	if err != nil {
		fmt.Println("terjadi error")
		fmt.Println(err.Error())
	} else {
		fmt.Println(result, err)
	}
	result, err = calculate(10, 2, "/")
	if err != nil {
		fmt.Println("terjadi error")
		fmt.Println(err.Error())
	} else {
		fmt.Println(result, err)
	}
	result, err = calculate(10, 2, "x")
	if err != nil {
		fmt.Println("terjadi error")
		fmt.Println(err.Error())
	} else {
		fmt.Println(result, err)
	}
}

func sum(arr []int) int {
	var i int
	i = 0
	for _, value := range arr {
		i = i + value
	}
	return i
}

func calculate(number int, numberTwo int, operator string) (int, error) {
	var hasil int
	var errorResult error

	switch operator {
	case "+":
		hasil = number + numberTwo
	case "-":
		hasil = number - numberTwo
	case "*":
		hasil = number * numberTwo
	case "/":
		hasil = number / numberTwo
	default:
		errorResult = errors.New("error unknown result")
	}

	return hasil, errorResult
}