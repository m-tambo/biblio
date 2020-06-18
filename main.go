package main

import (
	"biblio/database"
	"biblio/handler"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	db := database.NewDb()
	dao := handler.NewDAO(db)
	r := mux.NewRouter()

	r.Handle("/authors", handler.NewGetAuthorsHandler(dao)).Methods(http.MethodGet)
	r.Handle("/authors", handler.NewCreateAuthorHandler(dao)).Methods(http.MethodPost)
	r.Handle("/authors/{id}", handler.NewDeleteAuthorHandler(dao)).Methods(http.MethodDelete)

	r.Handle("/patrons", handler.NewGetPatronsHandler(dao)).Methods(http.MethodGet)

	http.ListenAndServe(":8008", r)
}