package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krishnatrea/cdp/bootstrap"
	"github.com/krishnatrea/cdp/domain"
)

type LoginController struct {
	Repo   *domain.UserRepository
	Config bootstrap.Config
}

func (sc *LoginController) Login(c *gin.Context) {
	// Login and provide
	var request domain.LoginRequest

	c.ShouldBind(&request)

	// Lets login.


	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
