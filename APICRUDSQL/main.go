package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func handleHome(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(405)
	var status int
	status = 200
	json.NewEncoder(writer).Encode(map[string]interface{}{
		"status":  status,
		"meesage": "Hello World",
	})
}
