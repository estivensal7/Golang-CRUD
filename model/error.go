// This file contains our error model package

// this tells other files that will use the Error model that here we have a package
package models

type Error struct {
	Message string `json:"message"`
}