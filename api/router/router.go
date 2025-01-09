package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/krishnatrea/cdp/api/middleware"
	"github.com/krishnatrea/cdp/bootstrap"
	"gorm.io/gorm"
)

func SetUp(config *bootstrap.Config, timeout time.Duration, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")

	// All Public APIs
	NewLoginRouter(config, timeout, db, publicRouter)
	NewLoginRouter(config, timeout, db, publicRouter)
	// NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(config.AccessTokenSecret))
	// All Private APIs
	// NewProfileRouter(env, timeout, db, protectedRouter)
	// NewTaskRouter(env, timeout, db, protectedRouter)
}
