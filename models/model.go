package model

type Book struct {
	Isbn  string `json:"isbn"`
	Title string `json:"title"`
}

type Books []Book
