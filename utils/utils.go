package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvVars struct {
	DB_USER     string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

func LoadEnv() EnvVars {
	// load the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading environment variables\n")
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	return EnvVars{
		DB_HOST:     dbHost,
		DB_USER:     dbUser,
		DB_PASSWORD: dbPassword,
		DB_PORT:     dbPort,
		DB_NAME:     dbName,
	}
}

// GetValue returns configuration value based on a given key from environment variables
func GetValue(key string) string {
	// load the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading environment variable\n")
	}

	// return the value based on a given key
	return os.Getenv(key)
}
