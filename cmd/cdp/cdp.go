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
	r := gin.Default()
	err := r.Run(fmt.Sprintf(":%s", PORT))
	if err != nil {
		slog.Error("error while serving the server", "error", err)
	}
}
