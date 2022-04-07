package routes

import (
	"bookstore/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{book_id}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{book_id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{book_id}", controllers.DeleteBook).Methods("DELETE")
}
