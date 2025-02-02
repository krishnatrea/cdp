package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/krishnatrea/cdp/api/controller"
	"github.com/krishnatrea/cdp/bootstrap"
	"github.com/krishnatrea/cdp/repository"
	"gorm.io/gorm"
)

func NewRefreshTokenRouter(config *bootstrap.Config, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)

	rc := &controller.RefreshController{
		Repo:   ur,
		Config: config,
	}

	group.POST("/refresh", rc.RefreshToken)
}
