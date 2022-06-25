package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func CityToIata(city string) string {
	data := getCitiesIataCodes()
	city = strings.ToTitle(city[:1]) + city[1:]
	return data[city]
}

func IataToCity(iata string) string {
	data := getCitiesIataCodes()
	for city, code := range data {
		if code == iata {
			return city
		}
	}
	return ""
}

func getCitiesIataCodes() map[string]string {
	rootPath := GetRootPath()
	envPath := rootPath + "/data/iata_codes.json"

	_, err := os.Stat(envPath)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(envPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bytes, _ := ioutil.ReadAll(f)

	var data map[string]string
	json.Unmarshal([]byte(bytes), &data)

	return data
}
