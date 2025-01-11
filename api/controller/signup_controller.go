package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krishnatrea/cdp/bootstrap"
	"github.com/krishnatrea/cdp/domain"
)

type SignupController struct {
	Repo   *domain.UserRepository
	Config bootstrap.Config
}

func (sc *SignupController) SignUp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
