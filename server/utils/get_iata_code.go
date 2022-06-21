package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func GetIataCode(city string) string {
	f, err := os.Open("data/iata_codes.json")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	bytes, _ := ioutil.ReadAll(f)

	var data map[string]string
	json.Unmarshal([]byte(bytes), &data)

	city = strings.ToTitle(city[:1]) + city[1:]
	return data[city]
}
