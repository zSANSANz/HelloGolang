package main

import (
	"fmt"
)

func main() {
	var bilangan int
	fmt.Scanf("%d", &bilangan)

	for j:=1; j<=bilangan; j++ {
		for i:=bilangan; i>=j; i-- {
			fmt.Print(" ")
			
		}
		for k:=1; k<=j; k++ {
			fmt.Print("*")
			
		}
		
		for i:=bilangan; i>=j; i-- {
			fmt.Print(" ")
			
		}

		for k:=1; k<=j; k++ {
			fmt.Print(" ")
			
		}

		for k:=1; k<=j; k++ {
			fmt.Print("*")
			
		}
		
			
		fmt.Println("")
	}

	
	fmt.Println("")
	fmt.Println("")

	for j:=1; j<=bilangan; j++ {
		for k:=1; k<=j; k++ {
			fmt.Print(" ")
			
		}
		
		for i:=bilangan; i>=j; i-- {
			fmt.Print("*")
			
		}
		
		for i:=bilangan; i>=j; i-- {
			fmt.Print(" ")
			
		}

		for k:=1; k<=j; k++ {
			fmt.Print(" ")
			
		}

		for i:=bilangan; i>=j; i-- {
			fmt.Print("*")
			
		}
		
		fmt.Println("")
	}
}