package main

import  (
	"fmt"
)

func main() {
	var erson [6]string
	erson[0] = "nol"
	fmt.Println(erson[0])

	terson := [8]string {
		"kosong",
		"satu",
		"dua",
		"tiga",
	}

	fmt.Println(terson[0])


	// array
	gerson := [...]string {
		"enol",
	}
	gerson[0] = "sji"
	fmt.Println(gerson[0])

	// slice
	ganok := []string {
		"ganok",
	}
	ganok = append(ganok, "satu")
	fmt.Println(ganok[0])
	fmt.Println(ganok[1])


	var nelson [1]string
	nelson[0] =	"emty"
	fmt.Println(nelson[0])


	// slices
	var gamingConsoles []string
	gaming := append(gamingConsoles, "zero")
	fmt.Println(gaming[0])

	ganok = append(ganok, "dau")
	ganok = append(ganok, "tiga")
	ganok = append(ganok, "emnat")
	ganok = append(ganok, "lima")

	for index, gnk := range ganok  {
		fmt.Println(index, "isinya ", gnk)
	}
}