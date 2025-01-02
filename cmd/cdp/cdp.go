package main

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/krishnatrea/cdp/database"
)

const PORT = "8080"

func main() {
	database.ConnectDatabase()
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "CDP by krishnatrea")
	})
	err := router.Run(fmt.Sprintf(":%s", PORT))
	if err != nil {
		slog.Error("error while serving the server", "error", err)
	}
}
