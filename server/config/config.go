package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alexandre-pinon/epic-road-trip/utils"
	"github.com/joho/godotenv"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Google   GoogleConfig
	Amadeus  AmadeusConfig
}

type AppConfig struct {
	Env    Env
	Name   string
	Secret string
}

type AmadeusConfig struct {
	BaseUrl string
	Key     string
	Secret  string
}

type GoogleConfig struct {
	BaseUrl string
	Key     string
}

type DatabaseConfig struct {
	Uri  string
	Name string
}

const (
	APP_ENV          = "APP_ENV"
	APP_NAME         = "APP_NAME"
	APP_SECRET       = "APP_SECRET"
	DB_URI           = "DB_URI"
	DB_NAME          = "DB_NAME"
	GOOGLE_BASE_URL  = "GOOGLE_BASE_URL"
	GOOGLE_KEY       = "GOOGLE_KEY"
	AMADEUS_BASE_URL = "AMADEUS_BASE_URL"
	AMADEUS_KEY      = "AMADEUS_KEY"
	AMADEUS_SECRET   = "AMADEUS_SECRET"
)

func GetConfig() *Config {
	rootPath := utils.GetRootPath()
	envPath := rootPath + "/.env"

	_, err := os.Stat(envPath)
	if err != nil {
		log.Print("[WARNING]: .env file was not found.\nAttempting to retrieve env variables directly...")
	} else {
		err = godotenv.Load(envPath)
		if err != nil {
			log.Fatal("Error while reading the env file", err)
		}
	}

	envVariables := map[string]string{
		APP_ENV:          os.Getenv(APP_ENV),
		APP_NAME:         os.Getenv(APP_NAME),
		APP_SECRET:       os.Getenv(APP_SECRET),
		DB_URI:           os.Getenv(DB_URI),
		DB_NAME:          os.Getenv(DB_NAME),
		GOOGLE_BASE_URL:  os.Getenv(GOOGLE_BASE_URL),
		GOOGLE_KEY:       os.Getenv(GOOGLE_KEY),
		AMADEUS_BASE_URL: os.Getenv(AMADEUS_BASE_URL),
		AMADEUS_KEY:      os.Getenv(AMADEUS_KEY),
		AMADEUS_SECRET:   os.Getenv(AMADEUS_SECRET),
	}

	for k, v := range envVariables {
		if v == "" {
			log.Fatalf("%s is not set", k)
		}
	}

	if err := Env(envVariables[APP_ENV]).IsValid(); err != nil {
		log.Fatal(err)
	}

	dbName := fmt.Sprintf("%s-%s", strings.ToLower(envVariables[APP_ENV]), envVariables[DB_NAME])

	cfg := &Config{
		App: AppConfig{
			Env:    Env(envVariables[APP_ENV]),
			Name:   envVariables[APP_NAME],
			Secret: envVariables[APP_SECRET],
		},
		Database: DatabaseConfig{
			Uri:  envVariables[DB_URI],
			Name: dbName,
		},
		Google: GoogleConfig{
			BaseUrl: envVariables[GOOGLE_BASE_URL],
			Key:     envVariables[GOOGLE_KEY],
		},
		Amadeus: AmadeusConfig{
			BaseUrl: envVariables[AMADEUS_BASE_URL],
			Key:     envVariables[AMADEUS_KEY],
			Secret:  envVariables[AMADEUS_SECRET],
		},
	}

	return cfg
}
