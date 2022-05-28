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
	Env       Env
	Database  DatabaseConfig
	SecretKey string
}

type DatabaseConfig struct {
	Uri  string
	Name string
}

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

	dbUri := os.Getenv("DB_URI")
	if dbUri == "" {
		log.Fatal("DB_URI is not set")
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("DB_NAME is not set")
	}

	dbName = fmt.Sprintf("%s-%s", strings.ToLower(string(env)), dbName)

	cfg := &Config{
		Env: env,
		Database: DatabaseConfig{
			Uri:  dbUri,
			Name: dbName,
		},
	}

	return cfg
}
