package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/krishnatrea/cdp/bootstrap"
	"gorm.io/gorm"
)

func NewSignupRouter(config *bootstrap.Config, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	// TODO: setup signup route.
}
