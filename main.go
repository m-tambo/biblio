package main

import (
	"net/http"

	"biblio/database"
	"biblio/handler"
	"biblio/middleware"

	"github.com/gorilla/mux"
)

func main() {
	db := database.NewDb()
	ad := handler.NewAuthorDAO(db)
	pd := handler.NewPatronDAO(db)
	r := mux.NewRouter()

	r.Handle("/authors", handler.NewGetAuthorsHandler(ad)).Methods(http.MethodGet)
	r.Handle("/authors", handler.NewCreateAuthorHandler(ad)).Methods(http.MethodPost)
	r.Handle("/authors/{id}", handler.NewDeleteAuthorHandler(ad)).Methods(http.MethodDelete)

	r.Handle("/patrons", handler.NewGetPatronsHandler(pd)).Methods(http.MethodGet)

	r.Use(middleware.ExecutionTimer)

	http.ListenAndServe(":8008", r)
}