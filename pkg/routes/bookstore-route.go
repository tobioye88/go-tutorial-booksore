package routes

import (
	"github.com/gorilla/mux"
	"github.com/tobioye88/go-tutorial-bookstore/pkg/controllers"
)

var RegisterBookstoreRoutes = func(route *mux.Router) {
	route.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	route.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	route.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods("GET")
	route.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	route.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
