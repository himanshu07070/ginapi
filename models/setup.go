package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

/*Inside this function, we create a new connection with the gorm.Open method. Here, we specify which kind
of database we plan to use and how to access it ,here we’ll use SQLite and store our data inside the test.db file. To connect our server
 to the database, we need to import the database’s driver, which is located inside the github.com/jinzhu/
 gorm/dialects module.*/

//here we create the global variable that is going to be used by controller to get acces to our database

var DB *gorm.DB

func ConnectionDatabase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("failed to connect database")
	}
	database.AutoMigrate(&Book{})
	// Next, we migrate the database schema using AutoMigrate. Make sure to call this method on each model
	//you have created.

	DB = database
	//Lastly, we populate the the DB variable with our database instance and it is one of the way to pass
	// database connection to controllers and another way is "Create Struct to hold DB Connection" but
	//that is lengthy procedure.
}

//now we need to build the controllers
