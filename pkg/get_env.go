package pkg

import (
	"os"
	"github.com/joho/godotenv"
)

// LoadEnv loads the environment variables from the .env file.
func LoadEnv() error {
	return godotenv.Load()
}

// GetEnv gets the value of the environment variable specified by key.
func GetEnv(key string) string {
	return os.Getenv(key)
}