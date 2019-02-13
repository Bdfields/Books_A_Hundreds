package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Isbn  string `json:"isbn"`
	Title string `json:"title"`
}

type Books []Book

func returnAllBooks(w http.ResponseWriter, r *http.Request) {
	books := Books{
		Book{Title: "Cracking the coding interview", Isbn: "1"},
		Book{Title: "Pro Git", Isbn: "2"},
		Book{Title: "Designing Data-Intensive Applications", Isbn: "3"},
		Book{Title: "Amazon Web Services In Action", Isbn: "4"},
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

	fmt.Printf("Hello, Let's go!\n")

	handleRequests()
}
