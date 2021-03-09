package main

import "fmt"

func main() {
	var myAge = 5

	if myAge == 5 {
		fmt.Println("you too young")
	}
	if myAge == 17 {
		fmt.Println("so sweet")
	}
	if myAge >= 17 && myAge <= 30 {
		fmt.Println("My age is between 17 and 30")
	}
	if dadAge := 9; dadAge < myAge {
		fmt.Println(dadAge)
	}

}
