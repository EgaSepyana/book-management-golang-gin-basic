package routers

import (
	"Book-Service/controller"
	"Book-Service/middleware"

	"github.com/gin-gonic/gin"
)

func BooksRouter(router *gin.Engine) {
	router.POST("/books", middleware.Authentication(), controller.CreateBook())
	router.GET("/books", controller.GetBooks())
	router.GET("/books/:id", controller.GetBookById())
	router.PUT("/books/:id", middleware.Authentication(), controller.UpdateBook())
	router.DELETE("/books/:id", middleware.Authentication(), controller.DeleteBook())
}
