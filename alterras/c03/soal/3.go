package main

import (
	"entities"
	"fmt"
)

func main() {
	var products = []entities.Product{
		entities.Product{
			Id:       "p01",
			Name:     "name 1",
			Price:    5,
			Quantity: 9,
			Status:   false,
		},
		entities.Product{
			Id:       "p02",
			Name:     "name 2",
			Price:    2,
			Quantity: 8,
			Status:   true,
		},
		entities.Product{
			Id:       "p03",
			Name:     "name 3",
			Price:    11,
			Quantity: 7,
			Status:   false,
		},
	}

	min, max := FindMaxAndMin(products)

	fmt.Println("Product Min")
	Display(min)

	fmt.Println("Product Max")
	Display(max)
}

func FindMaxAndMin(products []entities.Product) (min entities.Product, max entities.Product) {
	min = products[0]
	max = products[0]
	for _, product := range products {
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
	fmt.Println("id: ", product.Id)
	fmt.Println("name: ", product.Name)
	fmt.Println("price: ", product.Price)
	fmt.Println("quantity: ", product.Quantity)
	fmt.Println("status: ", product.Status)
	fmt.Println("total: ", (product.Price * float64(product.Quantity)))