package main

import "fmt"

func main() {
	var languages [5]string
	languages[0] = "Go"
	languages[1] = "Ruby"
	languages[2] = "JavaScript"
	languages[3] = "C#"
	languages[4] = "Python"

	fmt.Println(languages)
	fmt.Println(len(languages))

	frameworks := [5]string{"Gin", "Spear", "NodeJS", "MVC5", "Djanggo"}

	fmt.Println(frameworks)
	length := len(frameworks)
	fmt.Println(length)

	javascriptFrameworks := [...]string{
		"NodeJS",
		"ReactJS",
		"VueJS",
		"AngularJS",
		"ReactNative",
		"ExpressJS",
	}

	fmt.Println(javascriptFrameworks)
	fmt.Println(len(javascriptFrameworks))

	for index, lang := range javascriptFrameworks {
		fmt.Println("index ke : ", index, " isinya : ", lang)
	}
}
