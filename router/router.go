package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"data": "Hello to the CDP"})
		})
	}
}
