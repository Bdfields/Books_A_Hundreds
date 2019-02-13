package main

import (
	"Books_A_Hundreds/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var c = &controller.Controller{}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", c.HomePageFunc)
	myRouter.HandleFunc("/all", c.ReturnAllBooks)
	myRouter.HandleFunc("/book/{id}", c.ReturnSingleBook)

	log.Fatal(http.ListenAndServe(":1906", myRouter))
}
func main() {

	fmt.Printf("Running Books A Hundreds Server!\n")

	handleRequests()
}
