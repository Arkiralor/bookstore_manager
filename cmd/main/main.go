package main

import (
	"bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	new_router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(new_router)
	http.Handle("/", new_router)
	log.Fatal(http.ListenAndServe("localhost:8000", new_router))
}
