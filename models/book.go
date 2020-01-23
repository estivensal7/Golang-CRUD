// This file represents our Book model

// this tells other files that will use the Book model that here we have a package
package models

//Book model
type Book struct {
	ID int `json:id`
	Title string `json:title`
	Author string `json:author`
	Year string `json:year`
}