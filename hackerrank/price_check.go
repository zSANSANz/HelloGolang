package main

import (
	"fmt"
)

func main() {
	n:=0
	fmt.Scan(&n)
	var products []string
	product := ""
	for i:=1; i<=n; i++ {
		fmt.Scan(&product)
		products = append(products, product)
	}

	n=0
	fmt.Scan(&n)
	var productPrices []float64
	var productPrice float64
	for i:=1; i<=n; i++ {
		fmt.Scan(&productPrice)
		productPrices = append(productPrices, productPrice)
	}

	m:=0
	fmt.Scan(&m)
	var productSolds []string
	var productSold string
	for j:=1; j<=m; j++ {
		fmt.Scan(&productSold)
		productSolds = append(productSolds, productSold)
	}

	m=0
	fmt.Scan(&m)
	var soldPrices []float64
	var soldPrice float64
	for j:=1; j<=m; j++ {
		fmt.Scan(&soldPrice)
		soldPrices = append(soldPrices, soldPrice)
	}

	var kamus []string
	for z:=0; z<n; z++ {
		kamus = append(kamus, products[z]+fmt.Sprintf("%f", productPrices[z]))
	}

	var test []string
	for x:=0; x<m; x++ {
		test = append(test, productSolds[x]+fmt.Sprintf("%f", soldPrices[x]))
	}

	// for _, value := range kamus {
	// 	fmt.Println(value)
	// }

	// for _, value := range test {
	// 	fmt.Println(value)
	// }

	hasil:=m
	for z:=0; z<n; z++ {
		for x:=0; x<m; x++ {
			if kamus[z] == test[x] {
				hasil=hasil-1
			}
		}
	}

	fmt.Println(hasil)
	// for _, value := range products {
	// 	fmt.Println(value)
	// }

	// for _, value := range productPrices {
	// 	fmt.Println(value)
	// }

	// for _, value := range productSolds {
	// 	fmt.Println(value)
	// }

	// for _, value := range soldPrices {
	// 	fmt.Println(value)
	// }

}

