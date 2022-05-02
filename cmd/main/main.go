package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tobioye88/go-tutorial-bookstore/pkg/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookstoreRoutes(router)
	http.Handle("/", router)
	fmt.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
