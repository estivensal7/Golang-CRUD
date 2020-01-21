package main

//imported packages
import (
	// "encoding/json"
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

}

func getBook(w http.ResponseWriter, r *http.Request) {

}

func addBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}