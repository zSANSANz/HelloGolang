package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Product holds your product attribute
type Product struct {
	ID       string `json:"id"`
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

func singleProducts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for _, product := range Products {
		if product.ID == id {
			json.NewEncoder(w).Encode(product)
		}
	}
}

func handleRequest() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", home)
	r.HandleFunc("/products", allProducts)
	r.HandleFunc("/products/{id}", singleProducts)

	// pesan kalau aplikasi berjalan
	fmt.Println("Application running")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func main() {
	Products = []Product{
		Product{ID: "1", Title: "First Product", Price: 2000000, Quantity: 5},
		Product{ID: "2", Title: "Second Product", Price: 500000, Quantity: 15},
	}
	handleRequest()
}

/*
	Products is a global variable to hold collection of products
	atau bisa diartikan variable global product berfungsi sebagai pengganti database
*/
var Products []Product
