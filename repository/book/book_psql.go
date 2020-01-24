// this file will hold our DB Queries for when we call the controller methods

//this line tells other files that here we have a package
package bookRepository

//importing necessary packages
import (
	"Golang-CRUD/models"
	"database/sql"
)

//creating an empty struct to hold our methods
type BookRepository struct {}

//this method will be called in our controller, and takes in 3 parameters...
// 1 - 'db *sql.DB' - the database instance
// 2 - 'book models.Book' - our 'Book' model/object
// 3 - 'books []model.Book' - the book slice that we will use to append our individual book records to
// will return two values.. a slice of our data, and/or any errors
func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) ([]models.Book, error) {

	//invoke the db object Query method - passing in our query statement as well as assigning it to a rows variable. 
	rows, err := db.Query("SELECT * FROM books")

	//we check if there is an error - if error, return an empty Book model slice, and err message
	if err != nil {
		return []models.Book{}, err
	}

	//iterating through the rows to map the values of each row to its corresponding key in the books slice based on the book struct 
	//https://golang.org/pkg/database/sql/#Rows.Next
	for rows.Next() {
		err = rows. Scan(&book.ID, &book.Title, &book.Author, &book.Year)

		books = append(books, book)
	}

	//we check AGAIN after the scan if there is an error - if error, return an empty Book model slice, and err message
	if err != nil {
		return []models.Book{}, err
	}

	//otherwise we return the book slice, and nil for the error
	return books, nil

}
