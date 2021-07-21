package models

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//This model will contain the properties that represent fields in our database table

type Book struct {
	Id     uint   `json:"id" gorm:"primary_key"`
	Title  string `json: "title"`
	Author string `json: "author"`
}

//To create a book, we need to have a schema that can validate the user’s input
//to prevent us from getting invalid data.
type CreateBookInput struct {
	Title  string `json: "title" binding:"required"`
	Author string `json: "author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json: "title"`
	Author string `json: "author"`
}

//Next, we need to create a utility function called ConnectDatabase that allows us to create a connection
//to the database and migrate our model’s schema.
