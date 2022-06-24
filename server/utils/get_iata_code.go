package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetIataCode(city string) string {
	absPath, _ := filepath.Abs("../data/iata_codes.json")
	f, err := os.Open(absPath)
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
