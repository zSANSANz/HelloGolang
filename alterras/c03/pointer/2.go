package main

import "fmt"

func main() {
	var name string = "jhon"
	var nameAddress *string = &name
	fmt.Println("name (value)	:", name)  // jhon
	fmt.Println("name (address) :", &name) // 0xc20000a220
	fmt.Println("nameAddress (value)	:", *nameAddress) //jhon
	fmt.Println("nameAddress (value)	:", nameAddress)  // 0xc2000a220

	name = "doe"

	fmt.Println("")
	fmt.Println("name (value)	:", name) // Doe
	fmt.Println("name (address)	:", &name) // 0xc2000a220
	fmt.Println("nameAddress (value)	:", *nameAddress) // doe
	fmt.Println("nameAddress (address)	:", nameAddress) // 0x20000a220

	*nameAddress = "deni"

	fmt.Println("")
	fmt.Println("name (value)	:", name) // Doe
	fmt.Println("name (address)	:", &name) // 0xc2000a220
	fmt.Println("nameAddress (value)	:", *nameAddress) // doe
	fmt.Println("nameAddress (address)	:", nameAddress) // 0x20000a220

	
}