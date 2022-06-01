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
	AppName   string
	Database  DatabaseConfig
	Env       Env
	SecretKey string
}

type DatabaseConfig struct {
	Uri  string
	Name string
}

const (
	APP_NAME   = "APP_NAME"
	DB_URI     = "DB_URI"
	DB_NAME    = "DB_NAME"
	SECRET_KEY = "SECRET_KEY"
)

func GetConfig(env Env) *Config {
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
		APP_NAME:   os.Getenv("APP_NAME"),
		DB_URI:     os.Getenv("DB_URI"),
		DB_NAME:    os.Getenv("DB_NAME"),
		SECRET_KEY: os.Getenv("SECRET_KEY"),
	}

	for k, v := range envVariables {
		if v == "" {
			log.Fatalf("%s is not set", k)
		}
	}

	dbName := fmt.Sprintf("%s-%s", strings.ToLower(string(env)), envVariables[APP_NAME])

	cfg := &Config{
		AppName: envVariables[APP_NAME],
		Database: DatabaseConfig{
			Uri:  envVariables[DB_URI],
			Name: dbName,
		},
		Env:       env,
		SecretKey: envVariables[SECRET_KEY],
	}

	return cfg
}
