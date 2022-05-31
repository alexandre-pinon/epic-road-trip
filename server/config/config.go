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

	appName := os.Getenv("APP_NAME")
	dbUri := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	secretKey := os.Getenv("SECRET_KEY")

	for _, s := range []string{appName, dbUri, dbName, secretKey} {
		if s == "" {
			log.Fatalf("%s is not set", s)
		}
	}

	dbName = fmt.Sprintf("%s-%s", strings.ToLower(string(env)), dbName)

	cfg := &Config{
		AppName: appName,
		Database: DatabaseConfig{
			Uri:  dbUri,
			Name: dbName,
		},
		Env:       env,
		SecretKey: secretKey,
	}

	return cfg
}
