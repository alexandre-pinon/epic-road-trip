package config

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName       string
	Database      DatabaseConfig
	Env           Env
	SecretKey     string
	GoogleKey     string
	BaseUrlGoogle string
	AmadeusKey    string
	AmadeusSecret string
}

type DatabaseConfig struct {
	Uri  string
	Name string
}

const (
	APP_NAME        = "APP_NAME"
	DB_URI          = "DB_URI"
	DB_NAME         = "DB_NAME"
	GO_MODE         = "GO_MODE"
	SECRET_KEY      = "SECRET_KEY"
	GOOGLE_KEY      = "GOOGLE_KEY"
	BASE_URL_GOOGLE = "BASE_URL_GOOGLE"
	AMADEUS_KEY     = "AMADEUS_KEY"
	AMADEUS_SECRET  = "AMADEUS_SECRET"
)

func GetConfig() *Config {
	projectName := regexp.MustCompile(`^(.*server)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	envPath := string(rootPath) + "/.env"

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
		APP_NAME:        os.Getenv(APP_NAME),
		DB_URI:          os.Getenv(DB_URI),
		DB_NAME:         os.Getenv(DB_NAME),
		GO_MODE:         os.Getenv(GO_MODE),
		SECRET_KEY:      os.Getenv(SECRET_KEY),
		GOOGLE_KEY:      os.Getenv(GOOGLE_KEY),
		BASE_URL_GOOGLE: os.Getenv(BASE_URL_GOOGLE),
		AMADEUS_KEY:     os.Getenv(AMADEUS_KEY),
		AMADEUS_SECRET:  os.Getenv(AMADEUS_SECRET),
	}

	for k, v := range envVariables {
		if v == "" {
			log.Fatalf("%s is not set", k)
		}
	}

	if err := Env(envVariables[GO_MODE]).IsValid(); err != nil {
		log.Fatal(err)
	}

	dbName := fmt.Sprintf("%s-%s", strings.ToLower(envVariables[GO_MODE]), envVariables[DB_NAME])

	cfg := &Config{
		AppName: envVariables[APP_NAME],
		Database: DatabaseConfig{
			Uri:  envVariables[DB_URI],
			Name: dbName,
		},
		Env:           Env(envVariables[GO_MODE]),
		SecretKey:     envVariables[SECRET_KEY],
		GoogleKey:     envVariables[GOOGLE_KEY],
		BaseUrlGoogle: envVariables[BASE_URL_GOOGLE],
		AmadeusKey:    envVariables[AMADEUS_KEY],
		AmadeusSecret: envVariables[AMADEUS_SECRET],
	}

	return cfg
}
