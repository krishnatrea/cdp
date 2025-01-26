package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krishnatrea/cdp/bootstrap"
	"github.com/krishnatrea/cdp/domain"
	"github.com/krishnatrea/cdp/internal/tokenutil"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	Repo   domain.UserRepository
	Config bootstrap.Config
}

func (lc *LoginController) Login(c *gin.Context) {
	// Login and provide
	var request domain.LoginRequest

	c.ShouldBind(&request)
	// Fetch User by email
	user, err := lc.Repo.FetchByEmail(c, request.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "username or password is not correct.",
		})
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "username or password is not correct.",
		})
		return
	}

	accessToken, err := tokenutil.CreateAccessToken(*user, lc.Config.AccessTokenSecret, lc.Config.AccessTokenTimeOut)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	refreshToken, err := tokenutil.CreateRefreshToken(*user, lc.Config.RefreshTokenSecret, lc.Config.RefreshTokenTimeOut)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// generate access token and refresh token

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
}
