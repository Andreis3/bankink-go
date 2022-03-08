package main

import (
	"github.com/santos/banking-go/app"
	"github.com/santos/banking-go/logger"
)

func main() {
	//log.Println("Starting our application...")
	logger.Info("Starting the application in port: http://localhost:3000")
	app.Start()
}
