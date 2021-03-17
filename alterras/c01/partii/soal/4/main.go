package main 

import "fmt"

func main() {

	var bil1, bil2, hasil int = 1, 1, 1
	fmt.Scanf("%d %d", &bil1, &bil2)

	for i:=1; i<=bil2; i++ {
		hasil = hasil*bil1
	}
	
	fmt.Println(hasil)
}