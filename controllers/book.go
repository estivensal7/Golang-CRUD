//this file contains our book controllers (i.e all of our CRUD handler functions)

//this tells other files that here we have a package
package controllers

//importing necessary packages
import (
	"Golang-CRUD/models"
	"Golang-CRUD/repository/book"
	"Golang-CRUD/utils"
	"database/sql"
	// "encoding/json"
	"net/http"
	"log"
)

//this struct contains the methods that we will call in this package
type Controller struct {}

//initializing an instance of the Book model
var books []models.Book

// error handling function
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//Our GetBooks method - returns the function that we previously wrote for getBooks()
func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {

	//the returned func takes two parameters ---
	//'w http.ResponseWriter' is used to fill in the HTTP response
	//'r *http.Request' holds the request object
	return func (w http.ResponseWriter, r *http.Request) {

		// creating an instance of the Book struct
		var book models.Book

		// creating variable to hold any errors with a type of 'models.Error'
		var error models.Error
	
		//asign an empty slice to the books variable
		books = []models.Book{}

		//invoking our Book Repo
		bookRepo := bookRepository.BookRepository{}

		//invoking the GetBooks method inside of the book repository and assigning it to two variables - 'books' & 'err'
		books, err := bookRepo.GetBooks(db, book, books)

		//if there is an error returned from the book-repo, we notify the client
		if err != nil {
			error.Message = "Server Error"
			
			//utilize the utils package 'SendError' method - it expects 3 arguments
			// 1 - the ResponseWriter
			// 2 - the status code
			// 3 - the error object
			utils.SendError(w, http.StatusInternalServerError, error)
			
			//return nothing -> exit program
			return
		}

		//if there is no error - we set the content type in the ResponseWriter's Header to JSON
		w.Header().Set("Content-Type", "application/json")

		//utilize the utils package 'SendSuccess' method - passing in the the ResponseWriter & the books slice
		utils.SendSuccess(w, books)
	
	}
}