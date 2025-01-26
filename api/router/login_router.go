package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/krishnatrea/cdp/api/controller"
	"github.com/krishnatrea/cdp/bootstrap"
	"github.com/krishnatrea/cdp/repository"
	"gorm.io/gorm"
)

func NewLoginRouter(config *bootstrap.Config, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	lc := controller.LoginController{
		Repo:   ur,
		Config: *config,
	}
	group.POST("/login", lc.Login)
}
