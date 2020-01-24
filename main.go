package main

//imported packages
import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"fmt"

	"Golang-CRUD/driver"
	"Golang-CRUD/models"
	"Golang-CRUD/controllers"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

//the book slice will hold the book record that we are going to create
var books []models.Book

//variable declared to hold all sql.DB functions -- https://golang.org/pkg/database/sql/#DB
var db *sql.DB

func init() {
	//loading all environment variables imported in this file
	gotenv.Load()
}

// error handling function
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	//initializing db variable to call on driver package and connect to db
	db = driver.ConnectDB()

	//invoking our Controllers' methods
	controller := controllers.Controller{}

	//https://github.com/gorilla/mux#install
	//implementing the mux request router
	router := mux.NewRouter()
	
	//creating routes for CRUD capabilities
	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	//starting our server
	fmt.Println("Server is running on PORT 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func getBook(w http.ResponseWriter, r *http.Request) {

	// creating an instance of the Book struct
	var book models.Book

	//invoke this method to grab the value of the params via mux
	params := mux.Vars(r)

	//structuring SQL Query ('$1' is a placeholder value) .. the real value will be passed by 'params["id"]'
	rows := db.QueryRow("SELECT * FROM books WHERE id = $1", params["id"])

	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	logFatal(err)

	json.NewEncoder(w).Encode(books)

}

func addBook(w http.ResponseWriter, r *http.Request) {

	// creating an instance of the Book struct
	var book models.Book

	// holding bookID after the new row is added to the db
	var bookID int

	//decoding the request body, and pointing it to the Book struct instance
	json.NewDecoder(r.Body).Decode(&book)

	//initiating CREATE query for db to create new row, then scan for the new id & point it to the Book struct instance
	err := db.QueryRow(
		"INSERT INTO books (title, author, year) VALUES($1, $2, $3) RETURNING id;", book.Title, book.Author, book.Year).Scan(&bookID)

	//If there's any error - log the error
	logFatal(err)

	json.NewEncoder(w).Encode(bookID)

}

func updateBook(w http.ResponseWriter, r *http.Request) {

	// creating an instance of the Book struct
	var book models.Book

	//decoding the request body, and pointing it to the Book struct instance
	json.NewDecoder(r.Body).Decode(&book)

	//structuring UPDATE query - expecting 2 values
	result, err := db.Exec("UPDATE books SET title=$1, author=$2, year=$3 where id=$4 RETURNING id",
	&book.Title, &book.Author, &book.Year, &book.ID)

	//how many rows have been updated? - any errors?
	rowsUpdated, err := result.RowsAffected()

	//logging any errors
	logFatal(err)

	json.NewEncoder(w).Encode(rowsUpdated)

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

	//invoke this method to grab the value of the params via mux
	params := mux.Vars(r)

	//structuring DELETE query - expecting 2 values
	result, err := db. Exec("DELETE FROM books WHERE id=$1", params["id"])

	//logging any errors
	logFatal(err)

	//how many rows have been deleted? - any errors?
	rowsDeleted, err := result.RowsAffected()

	//logging any errors
	logFatal(err)

	json.NewEncoder(w).Encode(rowsDeleted)

}