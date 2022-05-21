package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

type Config struct {
	Env       Env
	Database  DatabaseConfig
	SecretKey string
}

type DatabaseConfig struct {
	Uri    string
	DbName string
}

func GetConfig(env Env) *Config {
	projectName := regexp.MustCompile(`^(.*server)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	envPath := string(rootPath) + "/" + env.GetFileName()

	_, err := os.Stat(envPath)
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error while reading the env file", err)
	}

	cfg := &Config{
		Env: env,
		Database: DatabaseConfig{
			Uri:    os.Getenv("DB_URI"),
			DbName: os.Getenv("DB_NAME"),
		},
		SecretKey: os.Getenv("SECRET_KEY"),
	}

	return cfg
}
