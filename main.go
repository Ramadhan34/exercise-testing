package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]interface{} { // Function mengubah data penduduk menjadi lebih informatif
	populationData := []map[string]interface{}{}

	for _, v := range data {
		tokens := strings.Split(v, ";") // Memisahkan data penduduk berdasarkan karakter ";"

		// Membuat map untuk menyimpan data penduduk yang telah diubah
		person := make(map[string]interface{})

		// Menyimpan nama, umur, dan alamat penduduk
		person["name"] = tokens[0]
		age, _ := strconv.Atoi(tokens[1])
		person["age"] = age
		person["address"] = tokens[2]

		if tokens[3] != "" { // Jika ada data tinggi badan
			height, _ := strconv.ParseFloat(tokens[3], 64)
			person["height"] = height
		}

		if tokens[4] != "" { // Jika ada data status pernikahan
			isMarried, _ := strconv.ParseBool(tokens[4])
			person["isMarried"] = isMarried
		}

		// Menambahkan data penduduk yang telah diubah ke dalam slice of map
		populationData = append(populationData, person)
	}
	return populationData
}

func main() {
	data := []string{"Budi;23;Jakarta;;", "Joko;30;Bandung;;true", "Susi;25;Bogor;165.42;"}
	fmt.Println(PopulationData(data))
}
