package main

import (
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

func handleRequest() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	handleRequest()
}

/*
	Products is a global variable to hold collection of products
	atau bisa diartikan variable global product berfungsi sebagai pengganti database
*/
var Product []Products
