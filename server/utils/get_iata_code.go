package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func GetIataCode(city string) string {
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

	city = strings.ToTitle(city[:1]) + city[1:]
	return data[city]
}
