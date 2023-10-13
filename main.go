package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"Book-Service/routers"

	_ "Book-Service/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Test Air
// @description     A book management service API in Go using Gin framework.
// @description                 Description for what is this security definition being used

//@securityDefinitions.apikey ApiKeyAuth
//@in header
//@name token
//@description                 Description for what is this security definition being used

// @host      localhost:8080

func main() {
	router := gin.Default()
	cors.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "token", "content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatal("Error Loading env File")
	}

	port := os.Getenv("PORT")

	routers.BooksRouter(router)
	routers.AuthRouter(router)
	routers.UserRouter(router)
	routers.RatingRouter(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"Massage": "OK"})
	})

	fmt.Println("swager UI : http://localhost:8080/swagger/any/index.html#/")

	router.Run(":" + port)
}
