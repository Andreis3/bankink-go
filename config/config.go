package config

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/santos/banking-go/logger"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.Error(" Error loading .env file " + err.Error())
	}
}

var SERVER_HOST = os.Getenv("SERVER_HOST")
var SERVER_PORT = os.Getenv("SERVER_PORT")
var DB_HOST = os.Getenv("DB_HOST")
var DB_PORT = os.Getenv("DB_PORT")
var DB_USER = os.Getenv("DB_USER")
var DB_PASSWORD = os.Getenv("DB_PASSWORD")
var DB_NAME = os.Getenv("DB_NAME")
var DB_DRIVER = os.Getenv("DB_DRIVER")
