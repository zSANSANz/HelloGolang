package main

import  (
	"fmt"
)

func main() {
	var myMap map[string]int
	myMap = map[string]int{}
	myMap["ruby"] = 9
	myMap["ytohn"] = 1
	myMap["golang"] = 2
	myMap["java"] = 3
	myMap["ruby"] = 4
	fmt.Println(myMap)

	myLanguage := map[string]string{
		"satu" : "one",
		"dua" : "two",
		"tiga" : "three",
		"empat" : "four",
		"lima" : "five",
	}
	fmt.Println(myLanguage)
}