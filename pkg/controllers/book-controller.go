package controllers

import (
	"bookstore/pkg/models"
	"bookstore/pkg/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	resp, err := json.Marshal(newBooks)

	if err != nil {
		log.Panicf("Error marshalling json: %v", err)
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book_id := params["book_id"]
	book_id_int, err := strconv.ParseInt(book_id, 0, 64)
	if err != nil {
		log.Panicf("Error parsing book_id: %v", err)
	}
	book_details, _ := models.GetBookByID(book_id_int)

	resp, resp_err := json.Marshal(book_details)
	if resp_err != nil {
		log.Panicf("Error marshalling json: %v", resp_err)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	NewBook := CreateBook.CreateBook()
	resp, resp_err := json.Marshal(NewBook)
	if resp_err != nil {
		log.Panicf("Error marshalling json: %v", resp_err)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book_id := params["book_id"]
	book_id_int, err := strconv.ParseInt(book_id, 0, 64)
	if err != nil {
		log.Panicf("Error reading book_id: %v", err)
	}

	res := models.DeleteBook(book_id_int)
	resp, resp_err := json.Marshal(res)
	if resp_err != nil {
		log.Panicf("Error marshalling json: %v", resp_err)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = models.Book{}
	utils.ParseBody(r, &updateBook)
	params := mux.Vars(r)
	book_id := params["book_id"]
	book_id_int, err := strconv.ParseInt(book_id, 0, 64)
	if err != nil {
		log.Panicf("Error reading book_id: %v", err)
	}
	bookDetails, db := models.GetBookByID(book_id_int)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)

	res, res_err := json.Marshal(bookDetails)
	if res_err != nil {
		log.Panicf("Error marshalling json: %v", res_err)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
