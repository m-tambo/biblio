package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"biblio/schema"
	"database/sql"
	"github.com/gorilla/mux"
)

type dao struct {
	db *sql.DB
}

type AuthorsDAO interface {
	GetAuthors() ([]schema.Author, error)
	CreateAuthor()
	DeleteAuthor(id string)
}

func NewAuthorDAO(db *sql.DB) AuthorsDAO {
	return dao{db: db}
}

type getAuthors struct {
	adao AuthorsDAO
}

type createAuthor struct {
	adao AuthorsDAO
}

type deleteAuthor struct {
	adao AuthorsDAO
}

func NewGetAuthorsHandler(adao AuthorsDAO) http.Handler {
	return getAuthors{adao: adao}
}

func NewCreateAuthorHandler(adao AuthorsDAO) http.Handler {
	return createAuthor{adao: adao}
}

func NewDeleteAuthorHandler(adao AuthorsDAO) http.Handler {
	return deleteAuthor{adao: adao}
}

// the handlers form responses for the incoming requests,
// hold the business logic of the user registration process,
// and write responses and break remaining steps when one step fails.
func (ga getAuthors) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	authors, err := ga.adao.GetAuthors()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

func (ca createAuthor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var a schema.Author
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ca.adao.CreateAuthor()
	w.WriteHeader(http.StatusNoContent)
	w.Header().Set("Content-Type", "application/json")
}

func (da deleteAuthor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var id string
	if val, ok := mux.Vars(r)["id"]; ok {
		id = val
	}

	da.adao.DeleteAuthor(id)
	w.WriteHeader(http.StatusNoContent)
	w.Header().Set("Content-Type", "application/json")
}

func (d dao) GetAuthors() ([]schema.Author, error) {
	var got []schema.Author

	qry := "SELECT id, first_name, last_name, dob, created FROM author"

	rows, err := d.db.Query(qry)
	if err != nil {
		log.Printf("Error getting Authors: %v", err)
		return nil, err
	}

	for rows.Next() {
		var a schema.Author
		err = rows.Scan(&a.ID, &a.FirstName, &a.LastName, &a.Dob, &a.Created)
		if err != nil {
			log.Printf("Error reading Authors response: %v", err)
			return nil, err
		}

		got = append(got, a)
	}

	return got, nil
}

func (d dao) CreateAuthor() {
	var a schema.Author

	qry := `
		INSERT INTO author (first_name, last_name, dob, created)
		VALUES ($1, $2, $3, $4)
	`
	_, err := d.db.Exec(qry, a.FirstName, a.LastName, a.Dob, time.Now())
	if err != nil {
		log.Printf("Error creating Author: %v", err)
	}
}

func (d dao) DeleteAuthor(id string) {
	qry := "DELETE FROM author WHERE id = " + id

	_, err := d.db.Exec(qry)
	if err != nil {
		log.Printf("Error deleting Author: %v", err)
	}
}