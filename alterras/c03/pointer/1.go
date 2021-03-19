package main 

import "fmt"

func main() {
	var name string = "Jhon"
	var nameAddress *string = &name
	fmt.Println("name (value)	:", name)					// jhon
	fmt.Println("name (address) :", &name)					// 0xc0001050
	fmt.Println("nameAddress (value)	:", *nameAddress)   // jhon
	fmt.Println("nameAddress (address)  :", nameAddress)	// 0xc0001050
}