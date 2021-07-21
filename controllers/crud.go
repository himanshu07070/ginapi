package controllers

import (
	"fmt"
	"go-gin-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//here we want all books from database
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

//here we insert book into database
func CreateBook(c *gin.Context) {
	var input models.CreateBookInput
	//We first validate the request body by using the ShoulDBindJSON method and pass the schema.
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"something went wrong": err.Error()})
		return
	}
	b := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&b)
	c.JSON(http.StatusOK, gin.H{"book": b})
}

//here we get the desired book
func FindBook(c *gin.Context) {
	var b models.Book // here we get the book models where book structure is defined
	// However, we only get the first book that matches the ID that we got from the request parameter.
	//We also need to check whether the book exists by simply wrapping it inside an if statement.
	if err := models.DB.Where("id=?", c.Param("id")).First(&b).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": b})
}

//now we do updation on book

func UpdateBook(c *gin.Context) {
	var b models.Book
	if err := models.DB.Where("id=?", c.Param("id")).First(&b).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not found"})
		return
	}
	//First, we do to grab a single book and make sure it exists. After we find the book, we need to validate
	// the user input with the UpdateBookInput schema. Finally, we update the book model using the Updates
	// method and return the updated book data to the client.
	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&b).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": b})
}

func DeleteBook(c *gin.Context) {

	var b models.Book
	if err := models.DB.Where("id=?", c.Param("id")).First(&b).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file not found"})
		return
	}
	fmt.Println(b)
	models.DB.Delete(&b)
	//Just like the update controller, we get the book model from the request parameters if it exists and
	//delete it with the Delete method from our database instance, which we get from our middleware. Then,
	// return true as the result since there is no reason to return a deleted book data back to the client.
	c.JSON(http.StatusOK, gin.H{"data removed": true})
}
