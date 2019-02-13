package controller

import (
	"Books_A_Hundreds/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Controller ...
type Controller struct {
}

// ReturnAllBooks GET
func (c *Controller) ReturnAllBooks(w http.ResponseWriter, r *http.Request) {
	books := models.Books{
		models.Book{Title: "Cracking the Coding Interview", Isbn: "1"},
		models.Book{Title: "Pro Git", Isbn: "2"},
		models.Book{Title: "Designing Data-Intensive Applications", Isbn: "3"},
		models.Book{Title: "Amazon Web Services In Action", Isbn: "4"},
		models.Book{Title: "Becoming Michelle Obama", Isbn: "5"},
	}

	fmt.Println("Endpoint: returnAllBooks")
	json.NewEncoder(w).Encode(books)
}

// HomePageFunc GET
func (c *Controller) HomePageFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In handler")
}

// ReturnSingleBook GET
func (c *Controller) ReturnSingleBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w, "Key: "+key)

}
