package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

type Museum struct {
	Data      []string `json:"data"`
	Kode      string   `json:"data:kode_pengelolaan"`
	Nama      string   `json:"nama"`
	Alamat    string   `json:"alamat_jalan"`
	Kecamatan string   `json:"kecamatan"`
	Kabupaten string   `json:"kabupaten_kota"`
}

func main() {
	url := "http://jendela.data.kemdikbud.go.id/api/index.php/CcariMuseum/searchGET?nama=museum"

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		response, err := http.Get(url)

		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
			log.Fatal(err)
		}

		fmt.Println(string(data))

		var museum Museum
		var jsonData = []byte(data)
		json.Unmarshal(jsonData, &museum)
		fmt.Printf("kode pengelolaan: %s, nama: %s, alamat: %s, kecamatan: %s, kabupaten atau kota: %s", museum.Kode, museum.Nama, museum.Alamat, museum.Kecamatan, museum.Kabupaten)
		wg.Done()
	}()
	wg.Wait()

	var musium = [][]string{{"nama", "Art: 1 New Museum"}, {"alamat_jalan", "Jl. Rajawali Selatan Raya no. 3 Gn. Sahari Utara, Sawah Besar"}}

	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range musium {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}

}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
