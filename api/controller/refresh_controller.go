package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krishnatrea/cdp/bootstrap"
	"github.com/krishnatrea/cdp/domain"
	"github.com/krishnatrea/cdp/pkg/tokenutil"
)

type RefreshController struct {
	Repo   domain.UserRepository
	Config *bootstrap.Config
}

func (rc *RefreshController) RefreshToken(c *gin.Context) {
	var request domain.RefreshTokenRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, err := tokenutil.ExtractIdFromToken(request.RefreshToken, rc.Config.RefreshTokenSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Refresh token is not authoristed"})
		return
	}

	user, err := rc.Repo.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found!!"})
		return
	}

	accessToken, err := tokenutil.CreateAccessToken(*user, rc.Config.AccessTokenSecret, rc.Config.AccessTokenTimeOut)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unable to provide access token"})
	}

	refreshToken, err := tokenutil.CreateRefreshToken(*user, rc.Config.RefreshTokenSecret, rc.Config.RefreshTokenTimeOut)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unable to provide access token"})
	}

	response := domain.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, response)
}
