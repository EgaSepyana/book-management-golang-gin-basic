package middleware

import (
	"Book-Service/helper"
	"Book-Service/models"
	"Book-Service/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userService = service.NewUserService()

func get_current_user(email string) models.User {
	user, _ := userService.FindUser(email)
	return user
}

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authentication"})
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}

		current_user := get_current_user(claims.Email)

		c.Set("current_user", current_user)
		c.Next()
	}
}
