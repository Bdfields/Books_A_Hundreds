package main

import (
	model "Books_A_Hundreds/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func returnAllBooks(w http.ResponseWriter, r *http.Request) {
	books := model.Books{
		model.Book{Title: "Cracking the Coding Interview", Isbn: "1"},
		model.Book{Title: "Pro Git", Isbn: "2"},
		model.Book{Title: "Designing Data-Intensive Applications", Isbn: "3"},
		model.Book{Title: "Amazon Web Services In Action", Isbn: "4"},
	}

	fmt.Println("Endpoint: returnAllBooks")
	json.NewEncoder(w).Encode(books)
}
func homePageFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In handler")
}

func returnSingleBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w, "Key: "+key)

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePageFunc)
	myRouter.HandleFunc("/all", returnAllBooks)
	myRouter.HandleFunc("/book/{id}", returnSingleBook)

	log.Fatal(http.ListenAndServe(":1906", myRouter))
}
func main() {

	fmt.Printf("Running Books A Hundreds Server!\n")

	handleRequests()
}
