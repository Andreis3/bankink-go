//go:build local
// +build local

package config

import (
	"os"
)

func init() {
	SERVER_HOST = os.Getenv("SERVER_HOST_DEV")
	SERVER_PORT = os.Getenv("SERVER_PORT_DEV")
	DB_HOST = os.Getenv("DB_HOST_DEV")
	DB_PORT = os.Getenv("DB_PORT_DEV")
	DB_USER = os.Getenv("DB_USER_DEV")
	DB_PASSWORD = os.Getenv("DB_PASSWORD_DEV")
	DB_NAME = os.Getenv("DB_NAME_DEV")
	DB_DRIVER = os.Getenv("DB_DRIVER_DEV")
}
