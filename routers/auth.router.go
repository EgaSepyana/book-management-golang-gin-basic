package routers

import (
	"Book-Service/controller"

	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.Engine) {
	router.POST("/signin", controller.SignIn())
	router.POST("/login", controller.Login())
}
