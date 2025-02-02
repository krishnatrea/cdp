package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krishnatrea/cdp/bootstrap"
	"github.com/krishnatrea/cdp/database/model"
	"github.com/krishnatrea/cdp/domain"
	"github.com/krishnatrea/cdp/pkg/tokenutil"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	Repo   domain.UserRepository
	Config *bootstrap.Config
}

func (sc *SignupController) SignUp(c *gin.Context) {
	var request domain.SignUpRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = sc.Repo.UserExitByEmail(c, request.Email)

	if err == nil {
		c.JSON(http.StatusAlreadyReported, gin.H{
			"message": "user already exist",
		})
		return
	}

	// Create User hash password.

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	request.Password = string(encryptedPassword)

	user := model.User{
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Email:     request.Email,
		Password:  request.Password,
	}

	err = sc.Repo.Create(c, &user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Create Access token
	accessToken, err := tokenutil.CreateAccessToken(user, sc.Config.AccessTokenSecret, sc.Config.AccessTokenTimeOut)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	refreshToken, err := tokenutil.CreateRefreshToken(user, sc.Config.RefreshTokenSecret, sc.Config.RefreshTokenTimeOut)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	signupResponse := domain.SignUpResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signupResponse)
}
