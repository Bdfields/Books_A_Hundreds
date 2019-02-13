package main

import (
	router "Books_A_Hundreds/routes"
	"fmt"
	"log"
	"net/http"
)

func handleRequests() {
	myRouter := router.NewRouter()
	log.Fatal(http.ListenAndServe(":1906", myRouter))
}
func main() {
	fmt.Printf("Running Books A Hundreds Server!\n")
	handleRequests()
}
