package handler

import (
	"biblio/schema"
	"database/sql"
	"net/http"
)

type BooksDAO interface {
	GetBooks() ([]schema.Book, error)
	GetBookByID() (schema.Book, error)
	CreateBook() error
	DeleteBook(id string) error
}

func NewBookDAO (db *sql.DB) BooksDAO {
	return dao{db: db}
}

type getBooks struct {
	bd BooksDAO
}

type getBookByID struct {
	bd BooksDAO
}

type createBook struct {
	bd BooksDAO
}

type deleteBook struct {
	bd BooksDAO
}

func NewGetBooksHandler(bkd BooksDAO) http.Handler {
	return getBooks{bd: bkd}
}

func NewGetBookByIDHandler(bkd BooksDAO) http.Handler {
	return getBookByID{bd: bkd}
}

func NewCreateBookHandler(bkd BooksDAO) http.Handler {
	return createBook{bd: bkd}
}

func NewDeleteBookHandler(bkd BooksDAO) http.Handler {
	return deleteBook{bd: bkd}
}

func (gb getBooks) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func (gb getBookByID) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func (cb createBook) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func (db deleteBook) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func (d dao) GetBooks() ([]schema.Book, error) {
	return []schema.Book{}, nil
}

func (d dao) GetBookByID() (schema.Book, error) {
	return schema.Book{}, nil
}

func (d dao) CreateBook() error {
	return nil
}

func (d dao) DeleteBook(id string) error {
	return nil
}