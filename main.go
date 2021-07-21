//build a bookstore RESTAPI that provides book data and perform some crud operations

//setup a server

package main

import (
	"go-gin-api/controllers"
	"go-gin-api/models"

	"github.com/gin-gonic/gin"
)

func main() {

	//we’ll initialize a new Gin router within the r variable. We’re using the Default router .
	r := gin.Default()

	//Next, we’ll define a GET route to the / endpoint and we need to specify two things: the endpoint and
	//the handler. The endpoint is the path the client wants to fetch and the handler, on the other hand,
	//determines how we provide the data to the client. This is where we put our business logic, such as
	//grabbing the data from the database, validating the user input, and so on.

	models.ConnectionDatabase()                   //here we are connecting the database
	r.GET("/books", controllers.FindBooks)        //here we are getting the request from users of all books
	r.POST("/book", controllers.CreateBook)       // here we are adding the book into database
	r.GET("/book/:id", controllers.FindBook)      //it will find the desired book
	r.PATCH("/book/:id", controllers.UpdateBook)  //it will update the desired book
	r.DELETE("/book/:id", controllers.DeleteBook) //it will just delete the desired book
	//now we run the server
	r.Run(":8000")
}
