package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Data Data `json:"data"`
}

type Data struct {
	kode_pengelolaan string `json:"kode_pengelolaan"`
	nama             string `json:"nama"`
	alamat_jalan     string `json:"alamat_jalan"`
	kecamatan        string `json:"kecamatan"`
	kabupaten_kota   string `json:"kabupaten_kota"`
}

func main() {
	url := "http://jendela.data.kemdikbud.go.id/api/index.php/CcariMuseum/searchGET?nama=museum"

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
	var responseObject Response
	json.Unmarshal(data, &responseObject)

	fmt.Println(responseObject.Data.kode_pengelolaan)
	fmt.Println(responseObject.Data.nama)
	fmt.Println(responseObject.Data.alamat_jalan)
	fmt.Println(responseObject.Data.kecamatan)
	fmt.Println(responseObject.Data.kabupaten_kota)

	fmt.Println(string(data))

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
