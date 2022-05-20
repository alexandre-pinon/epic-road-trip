package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database  DatabaseConfig
	SecretKey string
}

type DatabaseConfig struct {
	Uri    string
	DbName string
}

func GetConfig() *Config {
	_, err := os.Stat(".env")

	if !os.IsNotExist(err) {
		err := godotenv.Load(".env")

		if err != nil {
			log.Println("Error while reading the env file", err)
			panic(err)
		}
	}

	config := &Config{
		Database: DatabaseConfig{
			Uri:    os.Getenv("DB_URI"),
			DbName: os.Getenv("DB_NAME"),
		},
		SecretKey: os.Getenv("SECRET_KEY"),
	}

	return config
}
