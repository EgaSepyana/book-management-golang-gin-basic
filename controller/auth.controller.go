package controller

import (
	"Book-Service/helper"
	"Book-Service/models"
	"Book-Service/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userService = service.NewUserService()

// PingExample godoc
// GetBooks		 godoc
// @Summary      SignIn
// @Description  Takes a book JSON and store in DB. Return saved JSON.
// @Tags         Auth
// @Param request body models.SignIn true "query params"
// @Produce      json
// @Success      201   {object} map[string]interface{}
// @Router       /signin [post]
func SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var User models.User

		if err := c.BindJSON(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(User)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		errMassage := userService.SingIn(User)
		if errMassage != "" {
			c.JSON(http.StatusBadRequest, gin.H{"massage": errMassage})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"massage": "Successfull"})
	}
}

// PingExample godoc
// GetBooks		 godoc
// @Summary      Login
// @Description  Takes a book JSON and store in DB. Return saved JSON.
// @Tags         Auth
// @Param request body models.Login true "query params"
// @Produce      json
// @Success      200   {object} map[string]interface{}
// @Router       /login [post]
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var logs models.Login

		if err := c.BindJSON(&logs); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}

		user, err := userService.FindUser(logs.Email)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"Massage": "Email Not Exist"})
			return
		}

		passwordValid := helper.VerifyPassword(logs.Password, user.Password)
		if !passwordValid {
			c.JSON(http.StatusForbidden, gin.H{"Massage": "Wrong Password"})
			return
		}

		token, err := userService.Token(user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, token)

	}
}
