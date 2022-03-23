package config

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/santos/banking-go/logger"
)

var (
	SERVER_HOST string
	SERVER_PORT string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_DRIVER   string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.Error(" Error loading .env file " + err.Error())
	}

	SERVER_HOST = os.Getenv("SERVER_HOST")
	SERVER_PORT = os.Getenv("SERVER_PORT")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
	DB_DRIVER = os.Getenv("DB_DRIVER")
}
