package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/krishnatrea/cdp/pkg/tokenutil"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")

		if len(t) == 2 {
			authToken := t[1]
			authorized, err := tokenutil.IsAuthorized(authToken, secret)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
				return
			}
			if authorized {
				userID, err := tokenutil.ExtractIdFromToken(authToken, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Status Unauthorized"})
					c.Abort()
					return
				}
				c.Set("x-user-id", userID)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Status Unauthorized",
			})
			c.Abort()
		}
	}
}
