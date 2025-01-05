package main

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/krishnatrea/cdp/model"
)

const PORT = "8080"

func main() {
	if err := model.ConnectDatabase(); err != nil {
		slog.Error("Failed to migrate database:", "Error", err)
	}
	if err := model.RunMigrations(); err != nil {
		slog.Error("Failed to migrate database:", "Error", err)
	}
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "CDP by krishnatrea")
	})
	err := router.Run(fmt.Sprintf(":%s", PORT))
	if err != nil {
		slog.Error("error while serving the server", "error", err)
	}
}
