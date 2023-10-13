package routers

import (
	"Book-Service/controller"
	"Book-Service/middleware"

	"github.com/gin-gonic/gin"
)

func RatingRouter(router *gin.Engine) {
	router.POST("/rating/:BookId", middleware.Authentication(), controller.Rated())
	router.GET("/rating/:BookId", middleware.Authentication(), controller.GetBookRating())
	router.PUT("/rating/:id", middleware.Authentication(), controller.EditRating())
}
