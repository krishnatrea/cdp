package main

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/krishnatrea/cdp/model"
	"github.com/krishnatrea/cdp/router"
)

const PORT = "8080"

func main() {
	if err := model.ConnectDatabase(); err != nil {
		slog.Error("Failed to migrate database:", "Error", err)
	}
	// if err := model.RunMigrations(); err != nil {
	// 	slog.Error("Failed to migrate database:", "Error", err)
	// }
	r := gin.Default()

	router.SetUpRoutes(r)

	err := r.Run(fmt.Sprintf(":%s", PORT))
	if err != nil {
		slog.Error("error while serving the server", "error", err)
	}
}
