package main

import "fmt"

func main() {
	fmt.Println("hai")
	fmt.Println("pppppppp")

	var nama string = "Sandi Permana"
	fmt.Printf("nama saya %s sudah tau yaa", nama)
	fmt.Println(" nama saya ", nama, " sudah tau yaa")

	var umur int = 10
	fmt.Println("aku punya ayam sebanyak ", umur)

	var a bool = true
	fmt.Println("apasih ", a)

	const tempat_tinggal string = "jalan merkurius"
	fmt.Println(" tempat tinggalku di ", tempat_tinggal)

	/*
		tidak bisa mengubah isi dari sebuah const
		tempat_tinggal = "jalan yang berbeda"
		fmt.Println(" tempat tinggalku di ", tempat_tinggal)
		namun jika variabel bisa di rubaah
	*/

	nama = "Susanto Oni"
	fmt.Println(" nama saya ganti menjadi ", nama)
}
