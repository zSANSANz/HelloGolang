package main

import (
	"fmt"
)

func main() {
	printMyResult()
	printMyResultInput("snadi")
	sentence := printMyResultOutput("saya")
	fmt.Print(sentence)
}

func printMyResult() {
	fmt.Println("haiii")
}

func printMyResultInput(sentence string) {
	fmt.Println("hallo ", sentence)
}

func printMyResultOutput(sentence string) string {
	newSentence := sentence + " belajar golang"
	return newSentence
}