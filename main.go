package main

import (
	"fmt"

	"github.com/santos/banking-go/app"
	"github.com/santos/banking-go/config"
	"github.com/santos/banking-go/logger"
)

func main() {
	logger.Info(fmt.Sprintf("Starting the application in port: http://%s:%s", config.SERVER_HOST, config.SERVER_PORT))
	app.Start()
}
