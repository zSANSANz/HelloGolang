package main

import (
	"entities"
	"fmt"
)

func main() {
	var students = []entities.Student{
		entities.Student{
			Name:     "rizky",
			Price:    80,
		},
		entities.Student{
			Name:     "kobar",
			Price:    75,
		},
	}

	min, max := FindMaxAndMin(students)

	fmt.Println("Product Min")
	Display(min)

	fmt.Println("Product Max")
	Display(max)
}

func FindMaxAndMin(students []entities.Product) (min entities.Product, max entities.Product) {
	min = students[0]
	max = students[0]
	for _, product := range students {
		if product.Price > max.Price {
			max = product
		}
		if product.Price < min.Price {
			min = product
		}
	}
	return min, max
}

func Display(product entities.Product) {
	fmt.Println("name: ", product.Name)
	fmt.Println("price: ", product.Price)
	fmt.Println("quantity: ", product.Quantity)
	fmt.Println("status: ", product.Status)
	fmt.Println("total: ", (product.Price * float64(product.Quantity)))