# Sample RESTful API using Golang

This is a simple example of a RESTful API server using the Golang programming language.
The application will track books from my personal library, and allow me to view, update, and delete books from my collection.
I use the gorilla/mux package for HTTP request routing, and provide JSON formatted data in response to those requests.
The goal is to build a flexible API that responds to HTTP GET, PUT, CREATE, and POST requests.


## Structure
```
├── app
│   └── models
│       └── model.go          // Models for our application
│   └── routes
│       └── router.go         // Defines routes and endpoints
│   └── controller
│       └── controller.go     // Defines handlers for endpoints
└── main.go                   // Application entry point    
```

## API (Under construction)

#### /all
* `GET` : Get all books

#### /book/{isbn}
* `GET` : Get a particular book based on its ISBN
* `PUT` : Update a particular book based on its ISBN
* `DELETE` : Delete a particular book based on its ISBN

#### /book/add
* `POST` : Add a book to the library



Todo List
---

- [x] Create Github project
- [x] Write Initial Golang "Hello World" program
- [x] Create a skeleton for handling HTTP requests
- [x] Organize the code with packages
- [ ] Support all CRUD (Create, Retrieve, Update, Delete) requests
- [ ] Connect application to backend database
- [ ] Write tests for the APIs
- [ ] Create a front-end
