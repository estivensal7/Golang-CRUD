//this file will contain error messages for the client - handler functions

//this tells other files that here we have a package
package utils

//importing necessary packages
import (
	"encoding/json"
	"Golang-CRUD/models"
	"net/http"
)

//here we take 3 parameters
// -the ResponseWriter handles our response that's associated with the request
// -the status will be of type - integer - ex. 304, 404, 501 status codes
// -the err will be of type 'Error' from the models package
func SendError(w http.ResponseWriter, status int, err models.Error) {

	//sending an HTTP response header with the provided status code
	w.WriteHeader(status)

	//encoding error message
	json.NewEncoder(w).Encode(err)
}

//here we take 2 parameters
// -the ResponseWriter handles our response that's associated with the request
// -the data parameter will be of type - interface - because we don't want to be strict on the type of data that we will expect to be returned --- https://medium.com/rungo/interfaces-in-go-ab1601159b3a
func SendSuccess() {
	
	//encoding response data
	json.NewEncoder(w).Encode(data)
}