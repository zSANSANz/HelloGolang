package main

import (
	"fmt"
	"time"
)

func Sleep1(ch chan int) {
	time.Sleep(time.Second)
	ch <- 1
}

func Sleep3(ch chan int) {
	time.Sleep(time.Second * 3)
	ch <- 3
}

func Sleep5(ch chan int) {
	time.Sleep(time.Second * 5)
	ch <- 5
}

func main() {
	ch := make(chan int)
	go Sleep1(ch)
	go Sleep3(ch)
	go Sleep5(ch)
	for x := 0; x < 3; x++ {
		fmt.Println(<-ch)
	}
}
