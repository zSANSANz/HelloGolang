package main

import (
	"fmt"
)

func main() {
	var bilangan int
	fmt.Scanf("%d", &bilangan)
	for bil:=1; bil<50; bil++ {
		i:=0
		for bag:1; bag<50; bag++ {
			if bil % bag == 0 {
				i++
			}
		}
		for (i==2)&&(bil!=1) {
			fmt.Println(bil)
		}
	}
	
}