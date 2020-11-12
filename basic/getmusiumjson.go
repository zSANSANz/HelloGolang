package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Nama      string `json:"Name"`
	Alamat    string `json:"alamat_jalan"`
	Kecamatan string `json:"kecamatan"`
	Kota      string `json:"kabupaten_kota"`
}

func main() {
	var jsonString = `{
			"Name": "Art: 1 New Museum",
			"alamat_jalan": "Jl. Rajawali Selatan Raya no. 3 Gn. Sahari Utara, Sawah Besar",
			"kecamatan": "Sawah Besar",
			"kabupaten_kota": "Kota Jakarta Pusat"
		}
		`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Nama Musium :", data.Nama)
	fmt.Println("Alamat Musium  :", data.Alamat)
	fmt.Println("kecamatan Musium :", data.Kecamatan)
	fmt.Println("kota Musium  :", data.Kota)

}
