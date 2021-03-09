package main

import ("fmt")

func main() {
	var firstName string = "jhon"

	var lastName string 
	lastName = "witch"
	fmt.Printf("hallo %s \n %s", firstName, lastName)
	fmt.Println("hallo", firstName, lastName,  "!")

	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	fmt.Println(first, second, third)

	binatang := "sapi"

	fmt.Println(binatang)

	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hadi"

	fmt.Println(one, isFriday, twoPointTwo, say)

	fmt.Printf("isi variable one adalah %s \n sedangkan isFriday adalah %s \n dan twoPointTwo adalah %s \n dan say adalah %s \n", one, isFriday, twoPointTwo, say)

	name := "apa sih"
	// _ := "bisa gag"

	juragan, _ := "wew", "what"
	fmt.Println(name)

	alamat := new(string)
	*alamat = "terlalu"

	var p *int

	i := 42
	p = &i

	fmt.Println(*p)
	*p = 31 

	fmt.Println(alamat, juragan)
	fmt.Println(*alamat, juragan)

	fmt.Println(*p)

	var positiveNumber uint8 = 89
	var negativeNumber = -12345678349

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)
}