package controller

import (
	"Book-Service/helper"
	"Book-Service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// GetBooks		 godoc
// @Summary      Profile
// @Description  Takes a book JSON and store in DB. Return saved JSON.
// @Tags         User
// @Produce      json
// @Success      200   {object} models.User
// @Security     ApiKeyAuth
// @Router       /user/profile [get]
func Profile() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("current_user")
		curret_user, ok := user.(models.User)

		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"Massage": "No Authentication"})
			return
		}

		c.JSON(http.StatusOK, curret_user)
	}
}

// PingExample godoc
// GetBooks		 godoc
// @Summary      Set Admin
// @Description  Takes a book JSON and store in DB. Return saved JSON.
// @Tags         User
// @Param 		 email 	path	string	true "Email"
// @Produce      json
// @Success      200   {object} map[string]interface{}
// @Security     ApiKeyAuth
// @Router       /user/setadmin/{email} [put]
func Set_Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Param("email")

		userActive, _ := c.Get("current_user")
		curret_user, ok := userActive.(models.User)

		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"Massage": "No Authentication"})
			return
		}

		if curret_user.Role != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"Massage": "No Authentication"})
			return
		}

		user, err := userService.FindUser(email)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"err": err.Error()})
			return
		}

		if user.Role == "admin" {
			c.JSON(http.StatusForbidden, gin.H{"err": "This User Is Admin"})
			return
		}

		user.Role = "admin"
		err = userService.Update(user)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"err": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Massage": "Success"})
	}
}

// PingExample godoc
// GetBooks		 godoc
// @Summary      Disable User
// @Description  Takes a book JSON and store in DB. Return saved JSON.
// @Tags         User
// @Param 		 email 	path	string	true "Email"
// @Produce      json
// @Success      200   {object} map[string]interface{}
// @Security     ApiKeyAuth
// @Router       /user/disableUser/{email} [put]
func DisableUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Param("email")

		userActive, _ := c.Get("current_user")
		curret_user, ok := userActive.(models.User)

		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"Massage": "No Authentication"})
			return
		}

		if curret_user.Role != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"Massage": "No Authentication"})
			return
		}

		user, err := userService.FindUser(email)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"err": err.Error()})
			return
		}

		if user.Disable == true {
			c.JSON(http.StatusForbidden, gin.H{"err": "This User Already Disable"})
			return
		}

		if user.Role == "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"Massage": "Can't Disable Admin"})
			return
		}

		user.Disable = true
		err = userService.Update(user)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"err": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Massage": "Success"})

	}
}

// PingExample godoc
// GetBooks		 godoc
// @Summary      Enable user
// @Description  Takes a book JSON and store in DB. Return saved JSON.
// @Tags         User
// @Param 		 email 	path	string	true "Email"
// @Produce      json
// @Success      200   {object} map[string]interface{}
// @Security     ApiKeyAuth
// @Router       /user/enableUser/{email} [put]
func EnableUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Param("email")

		userActive, _ := c.Get("current_user")
		curret_user, ok := userActive.(models.User)

		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"Massage": "No Authentication"})
			return
		}

		if curret_user.Role != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"Massage": "No Authentication"})
			return
		}

		user, err := userService.FindUser(email)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"err": err.Error()})
			return
		}

		if user.Disable == false {
			c.JSON(http.StatusForbidden, gin.H{"err": "This User Already Enable"})
			return
		}

		user.Disable = false
		err = userService.Update(user)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"err": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Massage": "Success"})

	}
}

// PingExample godoc
// GetBooks		 godoc
// @Summary      Change Password
// @Description  Takes a book JSON and store in DB. Return saved JSON.
// @Tags         User
// @Produce      json
// @Param request body models.ChangePassword true "query params"
// @Success      200   {object} map[string]interface{}
// @Security     ApiKeyAuth
// @Router       /user/changepassword [put]
func ChangePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		var password models.ChangePassword

		user, _ := c.Get("current_user")
		curret_user, ok := user.(models.User)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"Massage": "No Authentication"})
			return
		}

		if err := c.BindJSON(&password); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		validateErr := validate.Struct(password)
		if validateErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": validateErr.Error()})
			return
		}

		validationPassword := helper.VerifyPassword(password.OldPassword, curret_user.Password)
		if !validationPassword {
			c.JSON(http.StatusBadRequest, gin.H{"err": "Old Password Is Invalid"})
			return
		}

		if password.OldPassword == password.NewPassword {
			c.JSON(http.StatusBadRequest, gin.H{"err": "New Password and Old Password is Same"})
			return
		}

		curret_user.Password = helper.HashPassword(password.NewPassword)
		err := userService.Update(curret_user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Massage": "Password Updated"})

	}
}

// PingExample godoc
// GetBooks		 godoc
// @Summary      Change Password
// @Description  Takes a book JSON and store in DB. Return saved JSON.
// @Tags         User
// @Produce      json
// @Param request body models.Profile true "query params"
// @Success      200   {object} map[string]interface{}
// @Security     ApiKeyAuth
// @Router       /user/editProfile [put]
func EditProfile() gin.HandlerFunc {
	return func(c *gin.Context) {

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		userActive, _ := c.Get("current_user")
		curret_user, ok := userActive.(models.User)

		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"Massage": "No Authentication"})
			return
		}

		if user.Address == "" {
			user.Address = curret_user.Address
		}

		if user.First_Name == "" {
			user.First_Name = curret_user.First_Name
		}

		if user.Last_Name == "" {
			user.Last_Name = curret_user.Last_Name
		}

		NewProfile, err := userService.UpdateProfile(user, curret_user.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, NewProfile)

	}
}
