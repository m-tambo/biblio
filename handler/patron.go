package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"biblio/schema"
	"database/sql"
)

type PatronsDAO interface {
	GetPatrons() []schema.Patron
	GetPatronByID(id string) schema.Patron
}

func NewPatronDAO(db *sql.DB) PatronsDAO {
	return dao{db: db}
}

type getPatrons struct {
	pdao PatronsDAO
}

type getPatronByID struct {
	pdao PatronsDAO
}

func NewGetPatronsHandler(pdao PatronsDAO) http.Handler {
	return getPatrons{pdao: pdao}
}

func NewGetPatronByIDHandler(pdao PatronsDAO) http.Handler {
	return getPatronByID{pdao: pdao}
}

func (gp getPatrons) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	patrons := gp.pdao.GetPatrons()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patrons)
}

func (gp getPatronByID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var id string
	if val, ok := mux.Vars(r)["id"]; ok {
		id = val
	}

	patron := gp.pdao.GetPatronByID(id)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patron)
}

func (d dao) GetPatrons() []schema.Patron {
	var patrons []schema.Patron
	qry := "SELECT id, first_name, last_name, street, city, state, zip, dob, created FROM patron;"

	rows, err := d.db.Query(qry)
	if err != nil {
		log.Printf("Error retrieving patrons %v", err)
	}

	for rows.Next() {
		var p schema.Patron
		err = rows.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Street, &p.City, &p.State, &p.Zip, &p.Dob, &p.Created)
		if err != nil {
			log.Printf("Error reading Patrons response: %v", err)
		}

		patrons = append(patrons, p)
	}

	return patrons
}

func (d dao) GetPatronByID(id string) schema.Patron {
	qryPatron := "SELECT * FROM patron WHERE id = " + id + ";"

	qryBooks := "SELECT b.id, title FROM patron AS p " +
		"JOIN patron_book AS pb ON p.id = pb.patron_id " +
		"JOIN book AS b ON pb.book_id = b.id WHERE p.id =" + id + ";"

	var p schema.Patron
	row := d.db.QueryRow(qryPatron)
	row.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Street, &p.City, &p.State, &p.Zip, &p.Dob, &p.Created)

	var books []schema.Book
	rows, err := d.db.Query(qryBooks)
	if err != nil {
		log.Printf("Error getting Patron Books from db: %v", err)
	}
	for rows.Next() {
		var b schema.Book
		err := rows.Scan(&b.ID, &b.Title)
		if err != nil {
			log.Printf("Error reading book row: %v", err)
		}

		books = append(books, b)
	}

	p.Books = books
	return p
}