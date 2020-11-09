package main

import (
	"fmt"
)

func main() {
	sayHelloTo("Eko", "khannedy")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
