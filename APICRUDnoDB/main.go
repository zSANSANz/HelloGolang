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
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(Products)
}

func singleProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	for _, product := range Products {
		if product.ID == id {
			json.NewEncoder(w).Encode(product)
		}
	}
}

func createProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var product Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Products = append(Products, product)
	json.NewEncoder(w).Encode(product)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	var product Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, p := range Products {
		if p.ID == id {
			Products[i].Title = product.Title
			Products[i].Price = product.Price
			Products[i].Quantity = product.Quantity
			json.NewEncoder(w).Encode(Products[i])
			return
		}
	}
}

func handleRequest() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", home)
	r.HandleFunc("/products", allProducts).Methods("GET")
	r.HandleFunc("/products/{id}", singleProduct).Methods("GET")
	r.HandleFunc("/products", createProducts).Methods("POST")
	r.HandleFunc("/products/{id}", updateProduct).Methods("PUT")

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
