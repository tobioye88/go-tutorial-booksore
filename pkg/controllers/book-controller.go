package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tobioye88/go-tutorial-bookstore/pkg/models"
	"github.com/tobioye88/go-tutorial-bookstore/pkg/utils"
)

var bookModel models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	bookModel := models.GetAllBooks()
	res, _ := json.Marshal(bookModel)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	var bookId = vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book := models.GetBookById(Id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	utils.ParseBody(r, &book)
	b := book.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	utils.ParseBody(r, &book)
	bookId := mux.Vars(r)["bookId"]
	Id, _ := strconv.ParseInt(bookId, 0, 0)
	b := models.UpdateBook(Id, &book)
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)["bookId"]
	Id, _ := strconv.ParseInt(bookId, 0, 0)
	models.DeleteBook(Id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte("{ \"success\": true }"), &jsonMap)
	json.NewEncoder(w).Encode(jsonMap)
}
