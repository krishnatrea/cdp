package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/krishnatrea/cdp/api/router"
	"github.com/krishnatrea/cdp/bootstrap"
)

const PORT = "8080"

func main() {

	app := bootstrap.App()

	config := app.Config

	db := app.DB

	timeout := time.Duration(config.ContextTimeout) * time.Second

	gin := gin.Default()

	router.SetUp(config, timeout, db, gin)

	gin.Run(config.ServerAddress)
}
