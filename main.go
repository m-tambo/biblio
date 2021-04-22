package main

import (
	"log"
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
	bd := handler.NewBookDAO(db)
	r := mux.NewRouter()
	//use r.exact()?

	r.Handle("/authors", handler.NewGetAuthorsHandler(ad)).Methods(http.MethodGet)
	r.Handle("/authors", handler.NewCreateAuthorHandler(ad)).Methods(http.MethodPost)
	r.Handle("/authors/{id}", handler.NewGetAuthorByIDHandler(ad)).Methods(http.MethodGet)
	r.Handle("/authors/{id}", handler.NewDeleteAuthorHandler(ad)).Methods(http.MethodDelete)

	r.Handle("/patrons", handler.NewGetPatronsHandler(pd)).Methods(http.MethodGet)
	r.Handle("/patrons/{id}", handler.NewGetPatronByIDHandler(pd)).Methods(http.MethodGet)

	r.Handle("/books", handler.NewGetBooksHandler(bd)).Methods(http.MethodGet)
	r.Handle("/books", handler.NewCreateBookHandler(bd)).Methods(http.MethodPost)
	r.Handle("/books/{id}", handler.NewGetBookByIDHandler(bd)).Methods(http.MethodGet)
	r.Handle("/books/{id}", handler.NewDeleteBookHandler(bd)).Methods(http.MethodDelete)

	r.Use(middleware.ExecutionTimer)

	err := http.ListenAndServe(":8008", r)
	if err != nil {
		log.Panic("Could not start server: ", err)
	}
}