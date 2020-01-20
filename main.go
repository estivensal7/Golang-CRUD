package main

//imported packages
import (
	"encoding/json"
	// "fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
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

func main() {

	//https://github.com/gorilla/mux#install
	//implementing the mux request router
	router := mux.NewRouter()

	//appending data to the books slice
	books = append(books,
		Book{ID: 1, Title: "Golang pointers", Author: "Mr. Golang", Year: "2010"},
		Book{ID: 2, Title: "Goroutines", Author: "Mr. Goroutine", Year: "2011"},
		Book{ID: 3, Title: "Golang routers", Author: "Mr. Router", Year: "2012"},
		Book{ID: 4, Title: "Golang concurrency", Author: "Mr. Currency", Year: "2013"},
		Book{ID: 5, Title: "Golang good parts", Author: "Mr. Good", Year: "2014"})
	
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
	
	//setting the response's content type to json
	w.Header().Set("Content-Type", "application/json")
	log.Println("Get Books is called")
	// An Encoder writes JSON values to an output stream.
	json.NewEncoder(w).Encode(books)

}

func getBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	log.Println("Get Book is called")
	//parameters can be used to create a map of route variables..
	//which can be retrieved calling 'mux.Vars()'
	params := mux.Vars(r)
	log.Println(params)

	// using the 'strconv' package convert the params id from 'str' to 'int'
	i, _ := strconv.Atoi(params["id"])

	//iterating through books to find the matching id numbers
	for _, book := range books {
		if book.ID == i {
			json.NewEncoder(w).Encode(&book)
		}
	}

}

func addBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	log.Println("Add Book is called")

	//create a variable to hold an instance of the 'Book' struct
	var book Book
	// A Decoder reads and decodes JSON values from an input stream. '&book' points the decoded response to the book variable
	json.NewDecoder(r.Body).Decode(&book)

	// setting books = the original books slice + the new book's values
	books = append(books, book)

	//returning a response containing all books
	json.NewEncoder(w).Encode(books)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	log.Println("Update Book is called")

	//create a variable to hold an instance of the 'Book' struct
	var book Book

	// A Decoder reads and decodes JSON values from an input stream. '&book' points the decoded response to the book variable
	json.NewDecoder(r.Body).Decode(&book)

	//iterating through books slice to find the matching book, then updating that book's values based on the UPDATE
	for i, item := range books {
		if item.ID == book.ID {
			books[i] = book
		}
	}

	//returning a response containing all books
	json.NewEncoder(w).Encode(books)

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	log.Println("Delete Book is called")

	//parameters can be used to create a map of route variables..
	//which can be retrieved calling 'mux.Vars()'
	params := mux.Vars(r)

	// using the 'strconv' package convert the params id from 'str' to 'int'
	id, _ := strconv.Atoi(params["id"])

	//iterating through books to find the matching id numbers
	for i, item := range books {
		if item.ID == id {
			books = append(books[:i], books[i+1:]...)
		}
	}
	
	json.NewEncoder(w).Encode(books)
}