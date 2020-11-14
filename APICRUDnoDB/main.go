package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Product holds your product attribute
type Product struct {
	Title    string `json:"title"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Selamat datang di home page")
}

func allProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Products)
}

func handleRequest() {
	http.HandleFunc("/", home)
	http.HandleFunc("/products", allProducts)
	// pesan kalau aplikasi berjalan
	fmt.Println("Application running")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	Products = []Product{
		Product{Title: "First Product", Price: 2000000, Quantity: 5},
		Product{Title: "Second Product", Price: 500000, Quantity: 15},
	}
	handleRequest()
}

/*
	Products is a global variable to hold collection of products
	atau bisa diartikan variable global product berfungsi sebagai pengganti database
*/
var Products []Product
