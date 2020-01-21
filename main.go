package main

//imported packages
import (
	"encoding/json"
	// "fmt"
	"log"
	"os"
	"net/http"
	"github.com/gorilla/mux"
	// "strconv"
	"database/sql"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
)

//book model
type Book struct {
	ID int `json:id`
	Title string `json:title`
	Author string `json:author`
	Year string `json:year`
}

//the book slice will hold the book record that we are going to create
var books []Book

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

	// Grabbing the ELEPHANTSQL_URL from the .env file then parsing the URL value & setting it equal to the pgURL variable
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	//opening DB connection to pgUrl
	db, err = sql.Open("postgres", pgUrl)
	logFatal(err)

	//db will ping the database - if there are no errors, it won't return anything - if there are any errors, the ping will fill the body of the variable below which we will then pass to the logFatal()
	err = db.Ping()
	logFatal(err)


	//https://github.com/gorilla/mux#install
	//implementing the mux request router
	router := mux.NewRouter()
	
	//creating routes for CRUD capabilities
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	//starting our server
	log.Fatal(http.ListenAndServe(":3000", router))
}

//the getBooks func takes two parameters ---
//'w http.ResponseWriter' is used to fill in the HTTP response
//'r *http.Request' holds the request object
func getBooks(w http.ResponseWriter, r *http.Request) {

	// creating an instance of the Book struct
	var book Book

	//asign an empty slice to the books variable
	books = []Book{}

	//invoke the db object Query method - passing in our query statement as well as assigning it to a rows variable. The 'err' body will fill if any errors are returned
	rows, err := db.Query("SELECT * FROM books")
	logFatal(err)
	
	//We are closing the connection after ensuring that the function call is performed
	//Defer is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup.
	defer rows.Close()

	//iterating through the rows to map the values of each row to its corresponding key in the books slice based on the book struct 
	//https://golang.org/pkg/database/sql/#Rows.Next
	for rows.Next() {
		err := rows. Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)

		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)

}

func getBook(w http.ResponseWriter, r *http.Request) {

	// creating an instance of the Book struct
	var book Book

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
	var book Book

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

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}