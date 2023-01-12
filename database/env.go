package database

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

func EnvConnection() string {
	err:= godotenv.Load()
	if err != nil {
		log.Fatal(".env not found")
	}

	return os.Getenv("MONGOURI")
}