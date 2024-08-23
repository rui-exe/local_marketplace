package middleware

import (
	"backend/helpers"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			c.Abort()
			return
		}

		//Strip the Bearer string from the token
		tokenParts := strings.Split(clientToken, " ")
		if len(tokenParts) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid/Malformed auth token format"})
			c.Abort()
			return
		}

		clientToken = tokenParts[1]

		claims, err := helpers.ValidateToken(clientToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "message": err.Error()})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.User_type)
		c.Set("username", claims.Username)
		fmt.Println(claims)
		c.Next()
	}
}
