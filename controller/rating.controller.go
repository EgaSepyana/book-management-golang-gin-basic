package controller

import (
	"Book-Service/models"
	"Book-Service/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var RatingService = service.NewRatingService()

// GetBooksById		 godoc
// @Summary      Set Rating
// @Description  Takes a book JSON and store in DB. Return saved JSON.
// @Tags         Rating
// @Param 		 BookId 	path	string	true "Book ID"
// @Param request body models.PostRating true "query params"
// @Produce      json
// @Success      201   {object} models.Rating
// @Security     ApiKeyAuth
// @Router       /rating/{BookId} [post]
func Rated() gin.HandlerFunc {
	return func(c *gin.Context) {
		var Rating models.Rating

		book_id := c.Param("BookId")

		user, _ := c.Get("current_user")
		curret_user, _ := user.(models.User)

		if err := c.BindJSON(&Rating); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if validationError := validate.Struct(Rating); validationError != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationError.Error()})
			return
		}

		book, err := BookService.FindByID(book_id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		rating, err := RatingService.AddRating(book.ID, Rating, curret_user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, rating)

	}
}

// PingExample godoc
// GetBooks		 godoc
// @Summary      Get Book Rating
// @Description  Takes a book JSON and store in DB. Return saved JSON.
// @Tags         Rating
// @Param 		 BookId 	path	string	true "Book ID"
// @Produce      json
// @Success      200   {object} interface{}
// @Router       /rating/{BookId} [get]
func GetBookRating() gin.HandlerFunc {
	return func(c *gin.Context) {
		var rating []models.Rating

		bookId := c.Param("BookId")

		book, err := BookService.FindByID(bookId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		Ratings, err := RatingService.GetRating(book.ID, book.Rating_id, &rating)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, Ratings)
	}
}

// GetBooksById		 godoc
// @Summary      Update Rating
// @Description  Takes a book JSON and store in DB. Return saved JSON.
// @Tags         Rating
// @Param 		 id 	path	string	true "Rating  ID"
// @Param request body models.UpdateRating true "query params"
// @Produce      json
// @Success      200   {object} models.Rating
// @Security     ApiKeyAuth
// @Router       /rating/{id} [put]
func EditRating() gin.HandlerFunc {
	return func(c *gin.Context) {
		var rating models.Rating

		rating_id := c.Param("id")

		if err := c.BindJSON(&rating); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		Rating, err := RatingService.Update(rating_id, rating)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, Rating)
	}
}
