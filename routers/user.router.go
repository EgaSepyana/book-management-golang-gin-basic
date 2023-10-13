package routers

import (
	"Book-Service/controller"
	"Book-Service/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	router.GET("/user/profile", middleware.Authentication(), controller.Profile())
	router.PUT("/user/setadmin/:email", middleware.Authentication(), controller.Set_Admin())
	router.PUT("/user/changepassword", middleware.Authentication(), controller.ChangePassword())
	router.PUT("/user/disableUser/:email", middleware.Authentication(), controller.DisableUser())
	router.PUT("/user/enableUser/:email", middleware.Authentication(), controller.EnableUser())
	router.PUT("/user/editProfile", middleware.Authentication(), controller.EditProfile())
}
