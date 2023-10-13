package controller

import (
	"Book-Service/models"
	"Book-Service/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var BookService = service.NewBookService()
var validate = validator.New()

// PingExample godoc
// GetBooks		 godoc
// @Summary      post book
// @Description  Takes a boo	k JSON and store in DB. Return saved JSON.
// @Tags         booksss
// @Param request body models.PostBook true "query params"
// @Produce      json
// @Success      201   {object} map[string]interface{}
// @Security     ApiKeyAuth
// @Router       /books [post]
func CreateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var book models.Books

		fmt.Println("test")

		userActive, _ := c.Get("current_user")
		curret_user, ok := userActive.(models.User)

		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"Massage": "No Authentication"})
			return
		}

		if curret_user.Disable {
			c.JSON(http.StatusUnauthorized, gin.H{"Massage": "Your Account Is Disable"})
		}

		if err := c.BindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": http.StatusBadRequest})
			return
		}

		validatorError := validate.Struct(book)
		if validatorError != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validatorError.Error(), "status": http.StatusBadRequest, "data": ""})
			return
		}

		newbook, err := BookService.Create(book, curret_user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": http.StatusBadRequest, "data": ""})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": newbook})
	}
}

// PingExample godoc
// GetBooks		 godoc
// @Summary      get book
// @Description  Takes a book JSON and store in DB. Return saved JSON.
// @Tags         books
// @Produce      json
// @Success      200   {object} map[string]interface{}
// @Router       /books [get]
func GetBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		var books []models.Books
		allBooks, err := BookService.FindAll(&books)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, allBooks)
	}
}

// PingExample godoc
// GetBooks		 godoc
// @Summary      get book
// @Description  Takes a book JSON and store in DB. Return saved JSON.
// @Tags         books
// @Param 		 id 	path	string	true "Book ID"
// @Produce      json
// @Success      200   {object} models.Books
// @Router       /books/{id} [get]
func GetBookById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		book, err := BookService.FindByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, book)
	}
}

// GetBooksById		 godoc
// @Summary      Update Book
// @Description  Takes a book JSON and store in DB. Return saved JSON.
// @Tags         books
// @Param 		 id 	path	string	true "Book ID"
// @Param request body models.UpdatedBook true "query params"
// @Produce      json
// @Success      201   {object} models.Books
// @Security     ApiKeyAuth
// @Router       /books/{id} [put]
func UpdateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var newbook models.Books

		if err := c.BindJSON(&newbook); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		userActive, _ := c.Get("current_user")
		curret_user, ok := userActive.(models.User)

		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"Massage": "No Authentication"})
			return
		}

		if curret_user.Disable {
			c.JSON(http.StatusUnauthorized, gin.H{"Massage": "Your Account Is Disable"})
		}

		oldBook, err := BookService.FindByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if oldBook.Author != curret_user.Email {
			c.JSON(http.StatusUnauthorized, gin.H{"Massage": "No Authentication"})
		}

		if newbook.Content == "" {
			newbook.Content = oldBook.Content
		}

		if newbook.Description == "" {
			newbook.Description = oldBook.Description
		}
		if newbook.Title == "" {
			newbook.Title = oldBook.Title
		}

		book, err := BookService.Update(newbook, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, book)
	}
}

// GetBooksById		 godoc
// @Summary      Delete Book
// @Description  Takes a book JSON and store in DB. Return saved JSON.
// @Tags         books
// @Param 		 id 	path	string	true "Book ID"
// @Produce      json
// @Success      204              {string}  string    "Data Deleted"
// @Security     ApiKeyAuth
// @Router       /books/{id} [delete]
func DeleteBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		userActive, _ := c.Get("current_user")
		curret_user, ok := userActive.(models.User)

		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"Massage": "No Authentication"})
			return
		}

		if curret_user.Disable {
			c.JSON(http.StatusUnauthorized, gin.H{"Massage": "Your Account Is Disable"})
		}

		book, err := BookService.FindByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if book.Author != curret_user.Email {
			c.JSON(http.StatusUnauthorized, gin.H{"Massage": "No Authentication"})
		}

		err = BookService.Delete(id, book.Rating_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	}
}
