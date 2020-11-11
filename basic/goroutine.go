package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go CetakNama("Hai")
	go CetakNama("Hello")
	wg.Wait()
}

func CetakNama(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(time.Second)
		defer wg.Done()
	}
}
