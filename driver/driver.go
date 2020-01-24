//This file contains our driver package, connecting our server

//here we tell other files that this is a package named 'driver'
package driver

//importing necessary packages
import (
	"database/sql"
	"log"
	"os"
	"github.com/lib/pq"
)

// error handling function
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//initializing global db variable
var db *sql.DB

func ConnectDB() *sql.DB {
	// Grabbing the ELEPHANTSQL_URL from the .env file then parsing the URL value & setting it equal to the pgURL variable
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	//opening DB connection to pgUrl
	db, err = sql.Open("postgres", pgUrl)
	logFatal(err)

	//db will ping the database - if there are no errors, it won't return anything - if there are any errors, the ping will fill the body of the variable below which we will then pass to the logFatal()
	err = db.Ping()
	logFatal(err)

	//returning the db variable connection instance
	return db
}