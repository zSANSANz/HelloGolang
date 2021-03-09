package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// User holds your product attribute
type User struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Nama_Lengkap string `json:"nama_lengkap"`
	Foto         string `json:"foto"`
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Selamat datang di home page")
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(Users)
}

func singleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	for _, user := range Users {
		if user.ID == id {
			json.NewEncoder(w).Encode(user)
		}
	}
}

func createUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Users = append(Users, user)
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, p := range Users {
		if p.ID == id {
			Users[i].Username = user.Username
			Users[i].Password = user.Password
			Users[i].Nama_Lengkap = user.Nama_Lengkap
			Users[i].Foto = user.Foto
			json.NewEncoder(w).Encode(Users[i])
			return
		}
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	for i, p := range Users {
		if p.ID == id {
			Users = append(Users[:i], Users[i+1:]...)
			json.NewEncoder(w).Encode(p)
			return
		}
	}
}

func handleRequest() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", home)
	r.HandleFunc("/users", allUsers).Methods("GET")
	r.HandleFunc("/users/{id}", singleUser).Methods("GET")
	r.HandleFunc("/users", createUsers).Methods("POST")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	// pesan kalau aplikasi berjalan
	fmt.Println("Application running")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func main() {
	Users = []User{
		User{ID: "1", Username: "sandi", Password: "123", Nama_Lengkap: "sandi permana", Foto: "abc.img"},
		User{ID: "2", Username: "joko", Password: "123", Nama_Lengkap: "joko samudra", Foto: "def.img"},
	}
	handleRequest()
}

/*
	Users is a global variable to hold collection of users
	atau bisa diartikan variable global product berfungsi sebagai pengganti database
*/
var Users []User
